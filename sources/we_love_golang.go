package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/pathao-eng/empleo/core"
	"strings"
)

type WeLoveGolang struct {
	CurrentPage int
	EndPage     int
	URL         string
}

func (w *WeLoveGolang) Init() error {
	w.CurrentPage = 1
	w.EndPage = 1
	w.URL = "https://www.welovegolang.com"
	return nil
}

func (w *WeLoveGolang) Fetch() ([]core.Empleo, bool) {
	var empleos []core.Empleo
	c := colly.NewCollector()
	c.OnHTML(".stream-item", func(e *colly.HTMLElement) {
		emp := core.Empleo{}
		emp.Title = e.DOM.Find(".media-body").Find(".media-heading").Find("span").Text()
		emp.Company = e.DOM.Find(".media-body").Find(".company").Find("strong").Text()
		if strings.TrimSpace(emp.Company) == "" {
			emp.Company = e.DOM.Find(".media-body").Find(".company").Find("span").Text()
		}

		l, ok := e.DOM.Find(".media-body").Find(".media-heading").Find("a").Attr("href")
		if ok {
			emp.Link = fmt.Sprintf("%s%s", w.URL, l)
		}

		emp.Location = e.DOM.Find(".media-body").Find(".location").Find("span").Text()

		e.DOM.Find(".media-body").Find(".job-tag").Each(func(i int, s *goquery.Selection) {
			emp.Tags = append(emp.Tags, s.Text())
		})

		emp.Time = e.DOM.Find(".pull-right").Find("time").AttrOr("datetime", "none")
		empleos = append(empleos, emp)
	})
	c.Visit(w.URL)
	return empleos, false
}

var INSTANCEWELOVEGOLANG WeLoveGolang
