package scrape

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"belly/data"
)

const (
	zipCodeURL = "https://secure.independenttickets.com/backstage/event/report_market.php?id="
)



func (app *App) GetZipCode(showId string) []data.ZipCodeSales {
	var ZipCodeData []data.ZipCodeSales
	
	zipURL := zipCodeURL + showId
	client := app.Client
	response, err := client.Get(zipURL)
	if err != nil {
		log.Fatalln("Error fetching response. ", err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("TABLE.report").Each(func(i int, s *goquery.Selection) {
		if (i) == 2 {
			daydata := s.Find("tbody")
			daydata.Find("tr").Each(func(rowIndex int, row *goquery.Selection) {
				zipEntry := data.ZipCodeSales{}
					row.Find("td").Each(func(dataIndex int, entry *goquery.Selection) {
						switch dataIndex {
						case 0:
							zipEntry.City = strings.TrimSpace(entry.Text())
						case 1:
							zipEntry.State = strings.TrimSpace(entry.Text())
						case 2:
							zipEntry.Zip = strings.TrimSpace(entry.Text())
						case 3:
							zipEntry.Quantity = strings.TrimSpace(entry.Text())
						case 4:
							zipEntry.Sales = strings.TrimSpace(entry.Text())
						}
					})
				ZipCodeData = append(ZipCodeData, zipEntry)
			})
			totals := s.Find("tfoot")
			totals.Find("tr").Each(func(footIndex int, footData *goquery.Selection) {
				zipEntry := data.ZipCodeSales{}
				footData.Find("td").Each(func(footDataIndex int, footEntry *goquery.Selection) {
					switch footDataIndex {
					case 0:
						zipEntry.City = strings.TrimSpace(footEntry.Text())
					case 1:
						zipEntry.State = strings.TrimSpace(footEntry.Text())
					case 2:
						zipEntry.Sales = strings.TrimSpace(footEntry.Text())
					}
				})
				ZipCodeData = append(ZipCodeData, zipEntry)
			})
		}
	})
	return ZipCodeData
}
