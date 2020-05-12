package main

import (
	"net/http"
	"net/http/cookiejar"
	"fmt"

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
	showList, idList = scrape.RemoveShows(showList, idList)
	fmt.Println("\nGetting ticket info for all the shows")
	fmt.Println("Getting the daily sales info for all the shows")
	fmt.Println("Getting the zip reports for all the shows")
	for i, show := range showList {
		TicketData := App.GetTickets(idList[i])
		DaySalesData := App.GetDaySales(idList[i])
		ZipSalesData := App.GetZipCode(idList[i])
		data.DataTest(show, TicketData, DaySalesData, ZipSalesData)
	}
	data.StartExcel(showQuery)
<<<<<<< HEAD
	fmt.Println("Done!")
=======
	fmt.Println("Done")
>>>>>>> 6397941ff6d267591a3c8f6a381a9a12f347ac38
}


