package main

import (
	"fmt"
	"html/template"
	htemplate "html/template"
	"net/http"
)

/*
  Color Options:
    "#54B399" = Green
    "#6092C0" = Blue
    "#D36086" = Red
    "#9170B8" = Purple
    "#CA8EAE" = Pink
    "#D6BF57" = Yellow
    "#B9A888" = Gold
    "#DA8B45" = Orange
    "#AA6556" = Brown
    "#E7664C" = Blood Orange
*/

// Home ...
type Home struct {
	Title     string
	Bookshelf []Article
}

var (
	shelf = []Article{
		{
			// article metadata
			ID:    0,
			Title: "How TCP Works in Outer Space",
			Tags:  []string{"networking", "tcp", "space"},
			// article look
			Color:      "#54B399",
			Image:      "pelican",
			LeftMargin: 650,
			TopMargin:  -300,
			// shelf look
			BookPositionHTML: htemplate.HTMLAttr(getBookPosition("#54B399", 100)),
		},
		{
			// article metadata
			ID:    1,
			Title: "What does Netflix's OpenConnect do?",
			Tags:  []string{"networking", "tcp", "space"},
			// article look
			Color:      "#D36086",
			Image:      "butterfly",
			LeftMargin: 650,
			TopMargin:  -300,
			// shelf look
			BookPositionHTML: htemplate.HTMLAttr(getBookPosition("#D36086", -100)),
		},
		{
			// article metadata
			ID:    2,
			Title: "Starting to use Go's Generics",
			Tags:  []string{"networking", "tcp", "space"},
			// article look
			Color:      "#D6BF57",
			Image:      "lobster",
			LeftMargin: 650,
			TopMargin:  -300,
			// shelf look
			BookPositionHTML: htemplate.HTMLAttr(getBookPosition("#D6BF57", 50)),
		},
		{
			// article metadata
			ID:    3,
			Title: "Explaining EUV to a 5 year old",
			Tags:  []string{"networking", "tcp", "space"},
			// article look
			Color:      "#9170B8",
			Image:      "chameleon",
			LeftMargin: 650,
			TopMargin:  -300,
			// shelf look
			BookPositionHTML: htemplate.HTMLAttr(getBookPosition("#9170B8", -20)),
		},
		{
			// article metadata
			ID:    4,
			Title: "Understanding Formula 1 Telemetry",
			Tags:  []string{"networking", "tcp", "space"},
			// article look
			Color:      "#E7664C",
			Image:      "ostrich",
			LeftMargin: 650,
			TopMargin:  -300,
			// shelf look
			BookPositionHTML: htemplate.HTMLAttr(getBookPosition("#E7664C", -130)),
		},
	}
)

func getHTMLID(articleID string) string {
	return fmt.Sprintf("src=\"../static/%s.png\"", articleID)
}

func getArticle(leftMargin, topMargin int) string {
	return fmt.Sprintf("style=\"position: relative;left: 100px;padding: 30px;font-size: 40px;\"", leftMargin, topMargin, topMargin)
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate, _ := template.ParseFiles(HomeTemplatePath)

	home := Home{
		Title:     "Karan Gugle",
		Bookshelf: shelf,
	}

	homeTemplate.Execute(w, home)
}
