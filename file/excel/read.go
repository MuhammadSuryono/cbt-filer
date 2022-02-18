package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func (e *ExcelHandler) ReadExcel() ([]TemplateQuestion, error) {
	xlsx, err := excelize.OpenFile(e.Filename)
	if err != nil {
		return nil, err
	}

	var dataRows []TemplateQuestion
	rows := getRows(xlsx, getSheetName(xlsx))
	for iRow, row := range rows {
		if iRow == 0 {
			continue
		}

		if row[0] == "" {
			continue
		}

		var dataRow TemplateQuestion
		for iCol, colCell := range row {
			if iCol > 10 {
				continue
			}

			switch iCol {
			case Question:
				dataRow.Question = colCell
			case Option1:
				dataRow.Option1 = colCell
			case Option2:
				dataRow.Option2 = colCell
			case Option3:
				dataRow.Option3 = colCell
			case Option4:
				dataRow.Option4 = colCell
			case Option5:
				dataRow.Option5 = colCell
			case Answer:
				dataRow.Answer = convertToAlphabetOption(colCell)
			case GroupName:
				dataRow.GroupName = colCell
			case GroupDescription:
				dataRow.GroupDescription = colCell
			case FilenameGroup:
				dataRow.FilenameGroup = colCell
			case QuestionGroup:
				dataRow.QuestionGroup = colCell
			default:
				fmt.Println("Unknown column:", iCol)
			}
		}
		dataRows = append(dataRows, dataRow)
	}

	return dataRows, nil
}

func getSheetName(xlsx *excelize.File) string {
	return xlsx.GetSheetMap()[1]
}

func getRows(xlsx *excelize.File, sheetName string) [][]string {
	return xlsx.GetRows(sheetName)
}

func convertToAlphabetOption(option string) string {
	fmt.Printf("t1: %T\n", option)
	options := map[string]string{
		"1":   "a",
		"2":   "b",
		"3":   "c",
		"4":   "d",
		"5":   "e",
		"1.0": "a",
		"2.0": "b",
		"3.0": "c",
		"4.0": "d",
		"5.0": "e",
	}
	return options[option]
}
