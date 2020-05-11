package data

import "github.com/360EntSecGroup-Skylar/excelize"

func setFormulas(xlsx *excelize.File) {
	var cell string
	var formCell string
	for rowIndex := 3; rowIndex < 16; rowIndex ++ {
		cell = "B" + convertIntToString(rowIndex)
		formCell = "A" + convertIntToString(rowIndex)
		_ = xlsx.SetCellFormula("ZipSummary", cell, "=SUMIF('zip dump'!A:A," + formCell + ",'zip dump'!E:E)")
		cell = "C" + convertIntToString(rowIndex)
		_ = xlsx.SetCellFormula("ZipSummary", cell, "=SUMIF('zip dump'!A:A,'ZipSummary'!" + formCell + ",'zip dump'!F:F)")
		cell = "D" + convertIntToString(rowIndex)
		formCell = "B" + convertIntToString(rowIndex)
		_ = xlsx.SetCellFormula("ZipSummary", cell, "=+" + formCell + "/B$16")
	}
	cell = "B16"
	_ = xlsx.SetCellFormula("ZipSummary", cell, "=SUM(B3:B15)")
	cell = "C16"
	_ = xlsx.SetCellFormula("ZipSummary", cell, "=SUM(C3:C15)")
	cell = "D16"
	_ = xlsx.SetCellFormula("ZipSummary", cell, "=SUM(D3:D15)")
	cell = "B19"
	_ = xlsx.SetCellFormula("ZipSummary", cell, "=SUM(B3:B9)")
	cell = "C19"
	_ = xlsx.SetCellFormula("ZipSummary", cell, "=+B19/$B$16")
	cell = "B20"
	_ = xlsx.SetCellFormula("ZipSummary", cell, "=SUM(B3:B7)")
	cell = "C20"
	_ = xlsx.SetCellFormula("ZipSummary", cell, "=+B20/$B$19")
	cell = "B21"
	_ = xlsx.SetCellFormula("ZipSummary", cell, "=SUM(B8:B10)")
	cell = "C21"
	_ = xlsx.SetCellFormula("ZipSummary", cell, "=+B21/$B$19")
}
