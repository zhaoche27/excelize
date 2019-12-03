package excelize

type SheetCap struct {
	name           string
	rowCap, colCap int
}

func NewSheetCap(name string, rowCap, colCap int) *SheetCap {
	return &SheetCap{name: name, rowCap: rowCap, colCap: colCap}
}
