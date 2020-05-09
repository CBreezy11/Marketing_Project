package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	showSearchURL = "https://secure.independenttickets.com/backstage/home/events.php?v=s&stat=any"
)



func (app *App) GetProjects(searchQuery string) ([]string, []string) {
	var showList []string
	var idList []string

	data := url.Values{
		"s": {searchQuery},
	}

	client := app.Client
	response, err := client.PostForm(showSearchURL, data)
	if err != nil {
		log.Fatalln("Error fetching response. ", err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".singleEvent").Each(func(i int, s *goquery.Selection) {
		show := s.Find(".title")
		if strings.Contains(strings.ToLower(show.Text()), searchQuery) {
			showList = append(showList, show.Text())
			linktag := s.Find(".nav a:first-child")
			link, _ := linktag.Attr("href")
			idList = append(idList, strings.Split(link, "=")[1])
		} else {
			fmt.Println("I cannot find that headliner, please check your spelling")
			os.Exit(1)
		}
	})
	return showList, idList
}




