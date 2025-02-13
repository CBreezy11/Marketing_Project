package scrape

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"belly/data"
)

const (
	ticketGetURL = "https://secure.independenttickets.com/backstage/event/report_ticket_count.php?id="
)



func (app *App) GetTickets(showId string) []data.TicketTypeSales {
	var TicketSales []data.TicketTypeSales
	ticketURL := ticketGetURL + showId
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
		if (i) == 1 {
			daydata := s.Find("tbody")
			daydata.Find("tr").Each(func(rowIndex int, row *goquery.Selection) {
				ticketEntry := data.TicketTypeSales{}
				if rowIndex > 0 {
					row.Find("td").Each(func(dataIndex int, entry *goquery.Selection) {
						switch dataIndex {
						case 0:
							ticketEntry.TicketType = strings.TrimSpace(entry.Text())
						case 2:
							ticketEntry.Price = strings.TrimSpace(entry.Text())
						case 3:
							ticketEntry.Count = strings.TrimSpace(entry.Text())
						case 4:
							ticketEntry.Total = strings.TrimSpace(entry.Text())
						}
					})
				}
				TicketSales = append(TicketSales, ticketEntry)
			})
			totals := s.Find("tfoot")
			totals.Find("tr").Each(func(footIndex int, footData *goquery.Selection) {
				ticketEntry := data.TicketTypeSales{}
				footData.Find("td").Each(func(footDataIndex int, footEntry *goquery.Selection) {
					switch footDataIndex {
					case 0:
						ticketEntry.TicketType = strings.TrimSpace(footEntry.Text())
					case 1:
						ticketEntry.Count = strings.TrimSpace(footEntry.Text())
					case 2:
						ticketEntry.Total = strings.TrimSpace(footEntry.Text())
					}
				})
				TicketSales = append(TicketSales, ticketEntry)
			})
		}
	})
	return TicketSales
}
