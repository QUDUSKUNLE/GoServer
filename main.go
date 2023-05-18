package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := httprouter.New()
	router.GET("/", HomeHandler)

	// Post collections
	router.GET("/posts", PostsIndexHandler)
	router.POST("/posts", PostsCreateHandler)

	// Posts singular
	router.GET("/posts/:id", PostsShowHandler)
	router.PUT("/posts/:id", PostsUpdateHandler)
	router.GET("/posts/:id/edit", PostsEditHandler)

	fmt.Println("Serving running on port: "+port)
	http.ListenAndServe("localhost:"+port, router)
}

func HomeHandler(res http.ResponseWriter, req *http.Request, router httprouter.Params) {
	fmt.Fprintln(res, "Home")
}

func PostsIndexHandler(res http.ResponseWriter, req *http.Request, router httprouter.Params) {
	fmt.Fprintln(res, "Posts index")
}

func PostsCreateHandler(res http.ResponseWriter, req *http.Request, router httprouter.Params) {
	fmt.Fprintln(res, "Creates index")
}

func PostsShowHandler(res http.ResponseWriter, req *http.Request, router httprouter.Params) {
	id := router.ByName("id")
	fmt.Fprintln(res, "showing post", id)
}

func PostsUpdateHandler(res http.ResponseWriter, req *http.Request, router httprouter.Params) {
	fmt.Fprintln(res, "Post update")
}

func PostsDeleteHandler(res http.ResponseWriter, req *http.Request, router httprouter.Params) {
	fmt.Fprintln(res, "Post delete")
}

func PostsEditHandler(res http.ResponseWriter, req *http.Request, router httprouter.Params) {
	fmt.Fprintln(res, "Post edit")
}
