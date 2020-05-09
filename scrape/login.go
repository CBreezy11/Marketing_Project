package scrape

import (
	"io/ioutil"
	"log"
	"net/url"
	"net/http"
)

const (
	loginURL      = "http://secure.independenttickets.com/backstage/index.php"
)

var (
	username = "manager@bellyupaspen.com"
	password = "Brandon15"
)

type App struct {
	Client *http.Client
}


func (app *App) Login() {
	client := app.Client

	data := url.Values{
		"user": {username},
		"pass": {password},
	}

	response, err := client.PostForm(loginURL, data)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
}