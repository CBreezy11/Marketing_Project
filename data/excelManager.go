package data

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)

func populate(test []IndShowData, xlsx *excelize.File) {

	fmt.Println("Writing the summary sheet")
	displayTypes(test, xlsx)
	fmt.Println("Writing the daily sales data")
	displayDaySales(test, xlsx)
	fmt.Println("Preparing the zip Report")
	zipReport(test, xlsx)
	fmt.Println("Adding all the needed formulas to the Excel sheet")
	setFormulas(xlsx)
}