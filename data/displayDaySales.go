package data

import "github.com/360EntSecGroup-Skylar/excelize"

func displayDaySales(test []IndShowData, xlsx *excelize.File) {
	for s := range(test) {
		var last5Ref int
		var cell string
		var first5Quant int 
		sheetIndex := convertIntToString(7 + s)
		for i := range(test[s].TicketDaySalesInfo) {
			var t bool = true
			switch t {
			case i < 5:
				first5Quant += convertToNumber(test[s].TicketDaySalesInfo[i].Quantity)
				first5Index := convertIntToString(21 + i)
				cell = "A" + first5Index
				xlsx.SetCellValue(sheetIndex, cell, test[s].TicketDaySalesInfo[i].Day)
				cell = "B" + first5Index
				xlsx.SetCellValue(sheetIndex, cell, convertToNumber(test[s].TicketDaySalesInfo[i].Quantity))
				cell = "C" + first5Index
				xlsx.SetCellValue(sheetIndex, cell, convertToNumber(test[s].TicketDaySalesInfo[i].Orders))
				cell = "D" + first5Index
				xlsx.SetCellValue(sheetIndex, cell, splitDollarString(test[s].TicketDaySalesInfo[i].Sales))
				cell = "E" +first5Index
				xlsx.SetCellValue(sheetIndex, cell, runningPercent(first5Quant))
			case i > 5 && len(test[s].TicketDaySalesInfo) < 10:
				last5Index := convertIntToString(27 + last5Ref)
				cell = "A" + last5Index
				xlsx.SetCellValue(sheetIndex, cell, test[s].TicketDaySalesInfo[i].Day)
				cell = "B" + last5Index
				xlsx.SetCellValue(sheetIndex, cell, convertToNumber(test[s].TicketDaySalesInfo[i].Quantity))
				cell = "C" + last5Index
				xlsx.SetCellValue(sheetIndex, cell, convertToNumber(test[s].TicketDaySalesInfo[i].Orders))
				cell = "D" + last5Index
				xlsx.SetCellValue(sheetIndex, cell, splitDollarString(test[s].TicketDaySalesInfo[i].Sales))
				last5Ref ++
			case i > 10:
				first := len(test[s].TicketDaySalesInfo) - 6
				if i >= first &&  i < len(test[s].TicketDaySalesInfo) - 1 {
					last5Index := convertIntToString(27 + last5Ref)
					cell = "A" + last5Index
					xlsx.SetCellValue(sheetIndex, cell, test[s].TicketDaySalesInfo[i].Day)
					cell = "B" + last5Index
					xlsx.SetCellValue(sheetIndex, cell, convertToNumber(test[s].TicketDaySalesInfo[i].Quantity))
					cell = "C" + last5Index
					xlsx.SetCellValue(sheetIndex, cell, convertToNumber(test[s].TicketDaySalesInfo[i].Orders))
					cell = "D" + last5Index
					xlsx.SetCellValue(sheetIndex, cell, splitDollarString(test[s].TicketDaySalesInfo[i].Sales))
					last5Ref ++					
				}
			}
		}
	}
}

func runningPercent(quant int) float64 {
	currentQuantPercent := float64(quant) / 450
	return currentQuantPercent
}