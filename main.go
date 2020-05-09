package main

import (
	"fmt"
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
	fmt.Println(showList)
	fmt.Println(idList)
	datatest := app.GetTickets(idList[0])
	fmt.Println(datatest)
}