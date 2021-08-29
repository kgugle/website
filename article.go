package main

import (
	"fmt"
	"html/template"
	htemplate "html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ArticleInfo ...
type ArticleInfo struct {
	Title string
}

// Article ...
type Article struct {
	Title string
	ID    int
	Tags  []string

	Color      string
	Image      string
	LeftMargin int
	TopMargin  int

	BookPositionHTML  htemplate.HTMLAttr
	ColorHTML         htemplate.HTMLAttr
	ImageHTML         htemplate.HTMLAttr
	ImagePositionHTML htemplate.HTMLAttr
	ContentHTML       htemplate.HTML
}

func getBookPosition(color string, left int) string {
	return fmt.Sprintf("style=\"background-color: %s;cursor: pointer;position: relative;left: %dpx;padding: 30px;font-size: 40px; border-radius: 10px;\"", color, left)
}

func getImage(articleID string) string {
	return fmt.Sprintf("src=\"../static/%s.png\"", articleID)
}

func getImagePosition(leftMargin, topMargin int) string {
	return fmt.Sprintf("style=\"position: relative;left: %dpx; top: %dpx; margin-bottom: %dpx;\"", leftMargin, topMargin, topMargin)
}

func getBackgroundColor(color string) string {
	return fmt.Sprintf("style=\"background-color: %s;\"", color)
}

func article(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["article"]
	ID, err := strconv.Atoi(articleID)
	if err != nil {
		return
	}

	articleInfo := shelf[ID]

	articleTemplate, _ := template.ParseFiles(ArticleTemplatePath)
	articleContent, _ := ioutil.ReadFile(fmt.Sprintf(ArticleContentPath, articleID))

	article := Article{
		ID:    articleInfo.ID,
		Title: articleInfo.Title,
		Image: articleInfo.Image,

		ColorHTML:         htemplate.HTMLAttr(getBackgroundColor(articleInfo.Color)),
		ImageHTML:         htemplate.HTMLAttr(getImage(articleInfo.Image)),
		ImagePositionHTML: htemplate.HTMLAttr(getImagePosition(articleInfo.LeftMargin, articleInfo.TopMargin)),
		ContentHTML:       htemplate.HTML(articleContent),
	}

	articleTemplate.Execute(w, article)
}
