package main

import "fmt"

var s = NewServer()
var w = make(map[string]string)

func setupHandles() {

	s.HandleFunc("GET", "/articles", func(c *Context) {

		var a []string
		for k, _ := range w {
			a = append(a, k)
		}
		fmt.Println(a)
		c.ListArticles(a)
	})

	s.HandleFunc("GET", "/articles/:name", func(c *Context) {

		if val, ok := w[c.Params["name"].(string)]; ok {
			c.ReadAnArticle(val, true)
		} else {
			c.ReadAnArticle(val, false)
		}
	})

	s.HandleFunc("PUT", "/articles/:name", func(c *Context) {

		c.Request.ParseForm()
		var PutVal string

		for k, _ := range c.Request.Form {
			PutVal = k
		}

		if _, ok := w[c.Params["name"].(string)]; ok {
			c.StoreAnArticle(true)
			w[c.Params["name"].(string)] = PutVal
		} else {
			c.StoreAnArticle(false)
			w[c.Params["name"].(string)] = PutVal
		}
	})

}

func main() {

	setupHandles()

	s.Run(":9090")
}
