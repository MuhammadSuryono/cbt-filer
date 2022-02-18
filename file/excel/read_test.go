package excel

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	excel := NewExcelHandler("Template Soal.xlsx")
	rows, _ := excel.ReadExcel()

	fmt.Println(rows)
}
