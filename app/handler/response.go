package handler

type (
	// Response is the http json response schema
	Response struct {
		Data interface{} `json:"data"`
	}

	// PaginatedResponse is the paginated response json schema
	// we not use it yet
	PaginatedResponse struct {
		Count    int         `json:"count"`
		Next     string      `json:"next"`
		Previous string      `json:"previous"`
		Results  interface{} `json:"results"`
	}

	// Ers is the root node of a JSON:API error response
	Ers struct {
		Errors interface{} `json:"errors"`
	}
	// Source is the source of the error in the code with a pointer to the line/function
	Source struct {
		Pointer string `json:"pointer"`
	}
	// Error is the actual error message contained by the root errors node
	Error struct {
		Status int    `json:"status"`
		Source Source `json:"source"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
	}
)

// NewResponse is the Response struct factory function
func NewJAPIResponse(content interface{}) *Response {
	return &Response{
		Data: content,
	}
}

// NewPaginatedResponse will created http paginated response
func NewPaginatedResponse(count int, next, prev string, results interface{}) *Response {
	return &Response{
		Data: &PaginatedResponse{
			Count:    count,
			Next:     next,
			Previous: prev,
			Results:  results,
		},
	}
}

// NewErrorResponse is the Error Response struct factory function
func NewErrorResponse(status int, pointer string, title string, detail string) *Ers {
	e := Error{
		Status: status,
		Source: Source{pointer},
		Title:  title,
		Detail: detail,
	}
	return &Ers{
		Errors: []Error{
			e,
		},
	}
}
