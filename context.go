package main

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Params map[string]interface{}

	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

type HandlerFunc func(*Context)

func (c *Context) ReadAnArticle(str string, found bool) {

	// Content-Type: text/html
	c.ResponseWriter.Header().Set("Content-Type", "text/html")

	if found {
		c.ResponseWriter.WriteHeader(http.StatusOK)
		c.ResponseWriter.Write([]byte(str))
	} else {
		c.ResponseWriter.WriteHeader(http.StatusNotFound)
	}
}

func (c *Context) StoreAnArticle(found bool) {

	c.ResponseWriter.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	if found {
		c.ResponseWriter.WriteHeader(http.StatusOK)
	} else {
		c.ResponseWriter.WriteHeader(http.StatusCreated)
	}
}

func (c *Context) ListArticles(v interface{}) {
	// Content-Type: application/json
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	// HTTP Status -> Status OK
	c.ResponseWriter.WriteHeader(http.StatusOK)

	// v -> json
	if err := json.NewEncoder(c.ResponseWriter).Encode(v); err != nil {
		// If there's a err then RenderErr
		c.RenderErr(http.StatusInternalServerError, err)
	}
}

func (c *Context) RenderErr(code int, err error) {

	if err != nil {

		if code > 0 {
			// Normal code -> HTTP Status code
			http.Error(c.ResponseWriter, http.StatusText(code), code)
		} else {
			// Abnormal code -> HTTP Status StatusInternalServerError
			defaultErr := http.StatusInternalServerError
			http.Error(c.ResponseWriter, http.StatusText(defaultErr), defaultErr)
		}
	}
}
