package data

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
)

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

type IndShowData struct {
	Show string
	TicketInfo []TicketTypeSales
	TicketDaySalesInfo []TicketDaySales
	ZipCodeSalesInfo []ZipCodeSales
}

var fullPacket []IndShowData


func DataTest(show string, ticketData []TicketTypeSales, ticketSalesData []TicketDaySales, zipCodeSalesData []ZipCodeSales) {
	singleShowLoad := IndShowData {
		Show: show,
		TicketInfo: ticketData,
		TicketDaySalesInfo: ticketSalesData,
		ZipCodeSalesInfo: zipCodeSalesData,
	}
	fullPacket = append(fullPacket, singleShowLoad)
}

func Display() {
	fmt.Printf("%#v", fullPacket)
}

func StartExcel(showQuery string) {

	fmt.Println("We have the data now preparing the Excel Sheet")

	xlsx, err := excelize.OpenFile("./ShowData.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	
	populate(fullPacket, xlsx)
	fileQuery := strings.Split(showQuery, " ")
	filename := strings.Join(fileQuery, "")

	err = xlsx.SaveAs("./" + filename + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}	
}