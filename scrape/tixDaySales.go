package scrape

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"belly/data"
)

const (
	ticketDaySalesGetURL = "https://secure.independenttickets.com/backstage/event/report_daily.php?id="
)



func (app *App) GetDaySales(showId string) []data.TicketDaySales {
	var DayTicketSales [] data.TicketDaySales
	
	ticketURL := ticketDaySalesGetURL + showId
	client := app.Client
	response, err := client.Get(ticketURL)
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
				ticketEntry := data.TicketDaySales{}
				row.Find("td").Each(func(dataIndex int, entry *goquery.Selection) {
					switch dataIndex {
					case 0:
						ticketEntry.Day = strings.TrimSpace(entry.Text())
					case 1:
						ticketEntry.Quantity = strings.TrimSpace(entry.Text())
					case 2:
						ticketEntry.Orders = strings.TrimSpace(entry.Text())
					case 3:
						ticketEntry.Sales = strings.TrimSpace(entry.Text())
					}
				})
				DayTicketSales = append(DayTicketSales, ticketEntry)
			})
			totals := s.Find("tfoot")
			totals.Find("tr").Each(func(footIndex int, footData *goquery.Selection) {
				ticketEntry := data.TicketDaySales{}
				footData.Find("td").Each(func(footDataIndex int, footEntry *goquery.Selection) {
					switch footDataIndex {
					case 0:
						ticketEntry.Day = strings.TrimSpace(footEntry.Text())
					case 1:
						ticketEntry.Quantity = strings.TrimSpace(footEntry.Text())
					case 2:
						ticketEntry.Orders = strings.TrimSpace(footEntry.Text())
					case 3:
						ticketEntry.Sales = strings.TrimSpace(footEntry.Text())
					}
				})
				DayTicketSales = append(DayTicketSales, ticketEntry)
			})
		}
	})
	return DayTicketSales
}