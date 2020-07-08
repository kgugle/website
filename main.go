package main

import (
	"fmt"

	"io/ioutil"
	"net/http"

	htemplate "html/template"
	"text/template"
)

// Article ...
type Article struct {
	Title         string
	Content       string
	ImagePosition htemplate.HTMLAttr
}

func main() {
	http.HandleFunc("/", article)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		panic(err)
	}
}

func getImagePosition(leftMargin, topMargin int) string {
	return fmt.Sprintf("style=\"position: relative;left: %dpx; top: %dpx; margin-bottom: %dpx;\"", leftMargin, topMargin, topMargin)
}

func article(w http.ResponseWriter, r *http.Request) {
	articleTemplate, _ := template.ParseFiles("./template/article.html")
	articleContent, _ := ioutil.ReadFile("./article/1.html")

	article := Article{
		Title:         "How TCP Works in Outer Space",
		Content:       string(articleContent),
		ImagePosition: htemplate.HTMLAttr(getImagePosition(375, -225)),
	}

	articleTemplate.Execute(w, article)
}
