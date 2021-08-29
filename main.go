package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	StaticDir = "/static/"

	HomeTemplatePath = "./template/home.html"

	ArticleTemplatePath = "./template/article.html"
	ArticleContentPath  = "./article/%s.html"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/{article}", article)

	router.
		PathPrefix(StaticDir).
		Handler(http.StripPrefix(StaticDir, http.FileServer(http.Dir("."+StaticDir))))

	if err := http.ListenAndServe(":1234", router); err != nil {
		panic(err)
	}
}
