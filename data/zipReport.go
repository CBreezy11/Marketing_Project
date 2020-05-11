package data

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func zipReport(test []IndShowData, xlsx *excelize.File) {
	inputStartIndex := 2
	var cell string
	var formCell string
	for s := range(test) {
		for i := range(test[s].ZipCodeSalesInfo) {
			entryRow := convertIntToString(inputStartIndex)
			switch test[s].ZipCodeSalesInfo[i].City {
			case "Total":
			default:
				cell = "B" + entryRow
				xlsx.SetCellValue("zip dump", cell, test[s].ZipCodeSalesInfo[i].City)
				cell = "C" + entryRow
				xlsx.SetCellValue("zip dump", cell, test[s].ZipCodeSalesInfo[i].State)
				cell = "D" + entryRow
				z, err := strconv.Atoi(test[s].ZipCodeSalesInfo[i].Zip)
				if err != nil {
					xlsx.SetCellValue("zip dump", cell, test[s].ZipCodeSalesInfo[i].Zip)
				}else {
					xlsx.SetCellValue("zip dump", cell, z)
				}
				cell = "E" + entryRow
				xlsx.SetCellValue("zip dump", cell, convertToNumber(test[s].ZipCodeSalesInfo[i].Quantity))
				cell = "F" + entryRow
				xlsx.SetCellValue("zip dump", cell, splitDollarString(test[s].ZipCodeSalesInfo[i].Sales))
				cell = "A" + entryRow
				formCell = "D" +entryRow
				_ = xlsx.SetCellFormula("zip dump", cell, "IFERROR(VLOOKUP("+formCell+",ZipMaster!A:D,4,0),\"Out of State\")")
				inputStartIndex ++
			}
		}
	}
}