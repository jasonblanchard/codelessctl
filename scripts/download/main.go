package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gocolly/colly/v2"
)

type story struct {
	CaseNumber string
	Title      string
	Text       string
}

func main() {
	stories := []*story{}

	c := colly.NewCollector(
		colly.AllowedDomains("thecodelesscode.com"),
		colly.Async(false),
	)

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		fmt.Println(fmt.Sprintf("%v: Parsing...", time.Now()))
		caseNumber := e.ChildText(".supertitle")
		title := e.ChildText(".title")
		text := e.ChildText(".story")

		story := &story{
			CaseNumber: caseNumber,
			Title:      title,
			Text:       text,
		}

		stories = append(stories, story)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// c.Visit("http://thecodelesscode.com/case/234")
	for i := 1; i < 235; i++ {
		url := fmt.Sprintf("http://thecodelesscode.com/case/%v", i)
		c.Visit(url)
	}

	c.Wait()

	storiesJson, err := json.MarshalIndent(stories, "", " ")

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("stories.json", storiesJson, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println(len(stories))
}
