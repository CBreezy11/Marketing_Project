package data

import "fmt"

type TicketTypeSales struct {
	TicketType string
	Price      string
	Count      string
	Total      string
}

type TicketDaySales struct {
	Day string
	Quantity string
	Orders string
	Sales string
}


type indShowData struct {
	show string
	ticketInfo []TicketTypeSales
	ticketDaySalesInfo []TicketDaySales
}

var fullPacket []indShowData


func DataTest(show string, ticketData []TicketTypeSales, ticketSalesData []TicketDaySales) {
	singleShowLoad := indShowData {
		show: show,
		ticketInfo: ticketData,
		ticketDaySalesInfo: ticketSalesData,
	}
	singleShowLoad.populatePacket()
}

func (singleShowLoad *indShowData) populatePacket() {
	fullPacket = append(fullPacket, *singleShowLoad)
}

func Display() {
	for i := range fullPacket {
		fmt.Println(fullPacket[i],"\n\n\n\n")
	}
}