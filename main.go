package main

import (
	"net/http"
	"net/http/cookiejar"

	"belly/data"
)


func main() {
	jar, _ := cookiejar.New(nil)

	app := App{
		Client: &http.Client{Jar: jar},
	}

	app.Login()
	showQuery := GetQueries()
	showList, idList := app.GetProjects(showQuery)
	for i, show := range showList {
		TicketData := app.GetTickets(idList[i])
		data.DataTest(show, TicketData)
	}
	data.Display()
}