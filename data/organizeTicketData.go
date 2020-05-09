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

type ZipCodeSales struct {
	City, State, Zip, Quantity, Sales string
}

type indShowData struct {
	show string
	ticketInfo []TicketTypeSales
	ticketDaySalesInfo []TicketDaySales
	zipCodeSalesInfo []ZipCodeSales
}

var fullPacket []indShowData


func DataTest(show string, ticketData []TicketTypeSales, ticketSalesData []TicketDaySales, zipCodeSalesData []ZipCodeSales) {
	singleShowLoad := indShowData {
		show: show,
		ticketInfo: ticketData,
		ticketDaySalesInfo: ticketSalesData,
		zipCodeSalesInfo: zipCodeSalesData,
	}
	singleShowLoad.populatePacket()
}

func (singleShowLoad *indShowData) populatePacket() {
	fullPacket = append(fullPacket, *singleShowLoad)
}

func Display() {
	fmt.Printf("%#v", fullPacket)
}