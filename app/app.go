package app

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rakmaster/goapi/app/controller"
	"github.com/rakmaster/goapi/app/db"
	"github.com/rakmaster/goapi/config"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"golang.org/x/net/context"
)

// App has the mongo database and router instances
type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

// ConfigAndRunApp will create and initialize App structure. App factory function.
func ConfigAndRunApp(config *config.Config) {
	app := new(App)
	app.Initialize(config)
	app.Run(config.ServerHost)
}

// Initialize initialize the app with
func (app *App) Initialize(config *config.Config) {
	app.DB = db.InitialConnection("golang", config.MongoURI())
	app.createIndexes()

	app.Router = mux.NewRouter()
	app.setRouters()
}

// SetupRouters will register routes in router
func (app *App) setRouters() {
	app.Post("/person", app.handleDbRequest(controller.CreatePerson))
	app.Patch("/person/{id}", app.handleDbRequest(controller.UpdatePerson))
	app.Put("/person/{id}", app.handleDbRequest(controller.UpdatePerson))
	app.Get("/person/{id}", app.handleDbRequest(controller.GetPerson))
	app.Get("/person", app.handleDbRequest(controller.GetPersons))
	app.Get("/person", app.handleDbRequest(controller.GetPersons), "page", "{page}")
	app.Get("/", app.handleHtRequest(controller.ShowDefault))
}

// createIndexes will create unique and index fields.
func (app *App) createIndexes() {
	// username and email will be unique.
	keys := bsonx.Doc{
		{Key: "username", Value: bsonx.Int32(1)},
		{Key: "email", Value: bsonx.Int32(1)},
	}
	people := app.DB.Collection("people")
	db.SetIndexes(people, keys)
}

// Get will register Get method for an endpoint
func (app *App) Get(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("GET").Queries(queries...)
}

// Post will register Post method for an endpoint
func (app *App) Post(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("POST").Queries(queries...)
}

// Put will register Put method for an endpoint
func (app *App) Put(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("PUT").Queries(queries...)
}

// Patch will register Patch method for an endpoint
func (app *App) Patch(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("PATCH").Queries(queries...)
}

// Delete will register Delete method for an endpoint
func (app *App) Delete(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("DELETE").Queries(queries...)
}

// Run will start the http server on host that you pass in. host:<ip:port>
func (app *App) Run(host string) {
	// use signals for shutdown server gracefully.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	go func() {
		http.Handle("/", app.Router)
		log.Fatal(http.ListenAndServe(":8080", app.Router))
	}()
	log.Printf("Server is listening on http://%s\n", host)
	sig := <-sigs
	log.Println("Signal: ", sig)

	log.Println("Stoping MongoDB Connection...")
	app.DB.Client().Disconnect(context.Background())
}

// RequestDbHandlerFunction is a custome type that help us to pass db arg to all endpoints
type RequestDbHandlerFunction func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

// handleRequest is a middleware we create for pass in db connection to endpoints.
func (app *App) handleDbRequest(handler RequestDbHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		handler(app.DB, w, r)
	}
}

// RequestHtHandlerFunction handles a regular http request
type RequestHtHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (app *App) handleHtRequest(handler RequestHtHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		handler(w, r)
	}
}
