package data

import "github.com/360EntSecGroup-Skylar/excelize"

func displayTypes(test []IndShowData, xlsx *excelize.File) {
	var index int
	var cell string
	totalTicketSold := 0
	totalRevenue := 0
	for s := range(test) {
		numTixType := 0
		summaryCellStartIndex := 7 + s
		row := convertIntToString(summaryCellStartIndex)
		cell = "A" + row
		xlsx.SetCellValue("Summary", cell, test[s].Show)
		index = xlsx.NewSheet(row)
		_ = xlsx.CopySheet(4, index)
		xlsx.SetCellValue(row, "A2", test[s].Show)
		for i := range(test[s].TicketInfo) {
			ticketCellStartIndex := 4 + i
			ticketRow := convertIntToString(ticketCellStartIndex)
			switch test[s].TicketInfo[i].TicketType {
			case "COMPS":
			case "PRINTONLYS":				
			case "":
			default:
				numTixType ++
				cell = "A" + ticketRow
				xlsx.SetCellValue(row, cell, test[s].TicketInfo[i].TicketType)
				cell = "B" + ticketRow
				xlsx.SetCellValue(row, cell, splitDollarString(test[s].TicketInfo[i].Price))
				cell = "C" + ticketRow
				xlsx.SetCellValue(row, cell, convertToNumber(test[s].TicketInfo[i].Count))
				cell = "D" + ticketRow
				xlsx.SetCellValue(row, cell, splitDollarString(test[s].TicketInfo[i].Total))
			case "TOTAL":
				totalTicketSold += convertToNumber(test[s].TicketInfo[i].Count)
				totalRevenue += splitDollarString(test[s].TicketInfo[i].Total)
				cell = "D" + row
				xlsx.SetCellValue("Summary", cell, convertToNumber(test[s].TicketInfo[i].Count))
				cell = "E" + row
				xlsx.SetCellValue("Summary", cell, splitDollarString(test[s].TicketInfo[i].Total))
				cell = "A" + ticketRow
				xlsx.SetCellValue(row, cell, test[s].TicketInfo[i].TicketType)
				cell = "B" + ticketRow
				xlsx.SetCellValue(row, cell, test[s].TicketInfo[i].Price)
				cell = "C" + ticketRow
				xlsx.SetCellValue(row, cell, convertToNumber(test[s].TicketInfo[i].Count))
				cell = "D" + ticketRow
				xlsx.SetCellValue(row, cell, splitDollarString(test[s].TicketInfo[i].Total))
			}
		}
		cell = "C" + row
		xlsx.SetCellValue("Summary", cell, numTixType) 
	}
	xlsx.SetCellValue("Summary", "B1", len(test))
	xlsx.SetCellValue("Summary", "B2", getAverages(totalTicketSold, len(test)))
	xlsx.SetCellValue("Summary", "B3", totalTicketSold)
	xlsx.SetCellValue("Summary", "B4", getAverages(totalRevenue, len(test)))
	xlsx.SetCellValue("Summary", "B5", totalRevenue)
}

