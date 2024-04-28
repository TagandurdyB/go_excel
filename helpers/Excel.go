package helpers

import (
	"fmt"
	"strconv"

	_ "image/gif" // Import necessary image formats (add more as needed)
	_ "image/jpeg"
	_ "image/png"

	// "io/ioutil"

	"github.com/xuri/excelize/v2"
)

type ExcelModel struct {
	Value   interface{}
	Column  string
	Row     int
	Width   *float64
	Height  *float64
	IsImage *bool
}

func ExcelModelsfromArray(values [][]interface{}) ([]ExcelModel, error) {
	// Initialize empty slice for models
	models := make([]ExcelModel, 0, len(values))

	for _, v := range values {
		// Check for expected length of inner slice
		if len(v) != 6 {
			return nil, fmt.Errorf("invalid data slice length: expected 6, got %d", len(v))
		}

		var width, height *float64

		if val, ok := v[3].(int); ok {
			f := float64(val) // Convert int to float64 if necessary
			width = &f
		} else if v[3] != nil {
			f, ok := (v[3].(float64))
			if !ok {
				return nil, fmt.Errorf("invalid type for width: expected float64, got %T", v[3])
			}
			width = &f
		}

		if val, ok := v[4].(int); ok {
			f := float64(val) // Convert int to float64 if necessary
			height = &f
		} else if v[4] != nil {
			f, ok := v[4].(float64)
			if !ok {
				return nil, fmt.Errorf("invalid type for height: expected float64, got %T", v[4])
			}
			height = &f
		}

		// Handle potential type conversions and errors
		column, ok := v[1].(string)
		if !ok {
			return nil, fmt.Errorf("invalid type for column: expected string, got %T", v[1])
		}
		row, ok := v[2].(int)
		if !ok {
			return nil, fmt.Errorf("invalid type for row: expected int, got %T", v[2])
		}
		var isImage *bool
		if v[5] != nil {
			b, ok := v[5].(bool)
			if !ok {
				return nil, fmt.Errorf("invalid type for isImage: expected bool, got %T", v[5])
			}
			isImage = &b
		}

		// Append model with safe conversions
		models = append(models, ExcelModel{
			Value:   v[0],
			Column:  column,
			Row:     row,
			Width:   width,
			Height:  height,
			IsImage: isImage,
		})
	}

	return models, nil
}

func ExcelCreate(filePath string, data []ExcelModel) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	sheetName := "Talyplar"
	f.NewSheet(sheetName)
	f.DeleteSheet("Sheet1")
	// Set active sheet of the workbook.
	// f.SetActiveSheet(index)

	// imagePath := ".base/uploads/empty.jpeg"
	// excelPath := ".base/image_export.xlsx"
	// exportImageToExcel(imagePath, excelPath)

	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center", // Set horizontal alignment to center
			Vertical:   "center", // Set vertical alignment to middle
			WrapText:   true,     // Wrap text if it overflows the cell
		},
		// Border: []excelize.Border{
		// 	{Type: "left", Color: "0000FF", Style: 3},
		// 	{Type: "top", Color: "00FF00", Style: 4},
		// 	{Type: "bottom", Color: "FFFF00", Style: 5},
		// 	{Type: "right", Color: "FF0000", Style: 6},
		// 	{Type: "diagonalDown", Color: "A020F0", Style: 7},
		// 	{Type: "diagonalUp", Color: "A020F0", Style: 8},
		// },
	})
	ErrH("Error in excel style", err)
	last := data[len(data)-1]
	key := last.Column + (strconv.Itoa(last.Row))
	f.SetCellStyle(sheetName, "A1", key, style)

	for _, v := range data {
		key := v.Column + (strconv.Itoa(v.Row))
		if v.Width != nil {
			f.SetColWidth(sheetName, v.Column, v.Column, *v.Width)
		}
		if v.Height != nil {
			f.SetRowHeight(sheetName, v.Row, *v.Height)
		}
		if v.IsImage != nil && *v.IsImage {
			imagePath := v.Value.(string)
			err := f.AddPicture(sheetName, key, imagePath,
				&excelize.GraphicOptions{
					// OffsetX: 15,
					// OffsetY: 10,
					// Hyperlink:       "https://github.com/xuri/excelize",
					// HyperlinkType:   "External",
					// PrintObject:     &enable,
					// LockAspectRatio: false,
					// Locked:          &disable,
					// Positioning:     "oneCell",
					AutoFit:         true,
					LockAspectRatio: true,
				})
			ErrH("Error adding image to Excel file:", err)
		} else {
			f.SetCellValue(sheetName, key, v.Value)
		}
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs(filePath); err != nil {
		fmt.Println(err)
	}
	
	
}

func ExcelExportImgToXlsx(imagePath string, excelPath string) {
	// Create a new Excel file
	f := excelize.NewFile()

	// Define sheet name
	sheetName := "Sheet1"
	f.NewSheet(sheetName)

	// Adjust column width for image size
	// desiredWidth := 300.0 // Replace with your desired image width in pixels
	f.SetColWidth(sheetName, "A", "A", 20)

	// (Optional) Adjust row height for image size
	// desiredHeight := 200.0 // Replace with your desired image height in pixels
	f.SetRowHeight(sheetName, 1, 120)

	//   // Open the image file
	//   img, err := f.addImage(imagePath)
	//   if err != nil {
	//     fmt.Println("Error opening image file:", err)
	//     return
	//   }

	// Insert the image into the sheet at a specific cell (adjust as needed)
	cell := "A1"
	err := f.AddPicture(sheetName, cell, imagePath,
		&excelize.GraphicOptions{
			// OffsetX: 15,
			// OffsetY: 10,
			// Hyperlink:       "https://github.com/xuri/excelize",
			// HyperlinkType:   "External",
			// PrintObject:     &enable,
			// LockAspectRatio: false,
			// Locked:          &disable,
			// Positioning:     "oneCell",
			AutoFit:         true,
			LockAspectRatio: true,
		})
	if err != nil {
		fmt.Println("Error adding image to Excel file:", err)
		return
	}

	// (Optional) Set image dimensions (adjust as needed)
	// width := 200
	// height := 150
	// if err := f.SetSheetPrProperty(sheetName, excelize.SheetPrWidth, width); err != nil {
	// 	fmt.Println("Error setting sheet width:", err)
	// }
	// if err := f.SetSheetPrProperty(sheetName, excelize.SheetPrHeight, height); err != nil {
	// 	fmt.Println("Error setting sheet height:", err)
	// }

	// Save the Excel file
	if err := f.SaveAs(excelPath); err != nil {
		fmt.Println("Error saving Excel file:", err)
	} else {
		fmt.Println("Image exported successfully to Excel file:", excelPath)
	}
}
