package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/s4kibs4mi/empleo/core"
)

type EuroTechJobs struct {
	CurrentPage int
	EndPage     int
	URL         string
	RootURL     string
}

func (i *EuroTechJobs) Init() error {
	i.CurrentPage = 1
	i.EndPage = 1
	i.URL = "https://www.eurotechjobs.com/jobs/java_developer"
	i.RootURL = "https://www.eurotechjobs.com"
	return nil
}

func (i *EuroTechJobs) Fetch() ([]core.Empleo, bool) {
	var empleos []core.Empleo
	c := colly.NewCollector()
	c.OnHTML(".jobinfo", func(e *colly.HTMLElement) {
		emp := core.Empleo{}
		emp.Title = e.DOM.Find(".col-lg-12").Find("a").Text()
		emp.Link = fmt.Sprintf("%s/%s", i.RootURL, e.DOM.Find(".col-lg-12").Find("a").AttrOr("href", ""))
		emp.Company = e.DOM.Find(".companyName").Text()
		emp.Location = e.DOM.Find(".location").Text()
		emp.Time = "none"
		empleos = append(empleos, emp)
	})
	c.Visit(i.URL)
	return empleos, false
}

var EUROTECHJOBS EuroTechJobs
