# goapi
Simple REST Api with MongoDB and gorilla/mux

.env file
```
server_host=api:8080
mongo_user=root
mongo_password=password
mongo_host=db
mongo_port=27017
```

mongo.evn file

```
server_host=api:8080
mongo_user=root
mongo_password=password
mongo_host=db
mongo_port=27017
```
The server actually runs on localhost:8080 even though the term Printf shows api:8080.

## Note about Docker
For some reason docker keeps a cache of main.out that doesn't show code changes every time. To get around this I'm currently running the following script to keep it fresh:

`docker-compose build --force-rm --no-cache && docker-compose up --detach && docker-compose logs -f`