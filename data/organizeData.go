package data

import "fmt"

type TicketTypeSales struct {
	TicketType string
	Price      string
	Count      string
	Total      string
}


type indShowData struct {
	show string
	ticketInfo []TicketTypeSales
}

var fullPacket []indShowData


func DataTest(show string, ticketData []TicketTypeSales) {
	singleShowLoad := indShowData {
		show: show,
		ticketInfo: ticketData,
	}
	singleShowLoad.populatePacket()
}

func (singleShowLoad *indShowData) populatePacket() {
	fullPacket = append(fullPacket, *singleShowLoad)
}

func Display() {
	fmt.Println(fullPacket[4].ticketInfo[1].Count)
}