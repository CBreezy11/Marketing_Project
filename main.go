package main

import (
	"net/http"
	"net/http/cookiejar"

	"belly/data"
	"belly/scrape"
)


func main() {
	jar, _ := cookiejar.New(nil)

	App := scrape.App{
		Client: &http.Client{Jar: jar},
	}

	App.Login()
	showQuery := scrape.GetQueries()
	showList, idList := App.GetProjects(showQuery)
	for i, show := range showList {
		TicketData := App.GetTickets(idList[i])
		DaySalesData := App.GetDaySales(idList[i])
		data.DataTest(show, TicketData, DaySalesData)
	}
	data.Display()
}