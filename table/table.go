package table

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type cell []string

type row []cell

func Print(data any) {
	value := reflect.ValueOf(data)
	valueType := value.Type().Kind()
	switch valueType {
	case reflect.Slice, reflect.Array:
		printSlice(data)
	default:
		fmt.Println("NOT SUPPORT!")
	}
}

func getWidths(header row, boby []row) []int {
	const spaces = 3
	rows := []row{header}
	rows = append(rows, boby...)
	widths := make([]int, len(header))
	for _, row := range rows {
		for i, cell := range row {
			width := 0
			for _, item := range cell {
				width = max(width, getVisibleLen(item)+spaces)
			}
			widths[i] = max(widths[i], width)
		}
	}
	return widths
}

func printTable(header row, body []row, widths []int) {
	printTop(widths)
	printRow(header, widths)
	for _, row := range body {
		printLine(widths)
		printRow(row, widths)
	}
	printBottom(widths)
}

func printTop(widths []int) {
	fmt.Printf("┌")
	for i := range widths {
		for j := 0; j < widths[i]; j++ {
			fmt.Printf("─")
		}
		if i != len(widths)-1 {
			fmt.Printf("┬")
		} else {
			fmt.Printf("┐\n")
		}
	}
}

func printRow(row row, widths []int) {
	rowNum := 0
	for _, cell := range row {
		rowNum = max(rowNum, len(cell))
	}

	for i := 0; i < rowNum; i++ {
		fmt.Printf("│")
		for j, cell := range row {
			width := widths[j] - 1
			if len(cell) > i {
				diff := len(cell[i]) - getVisibleLen(cell[i])
				width = width + diff
			}
			printFormat := " %-" + strconv.Itoa(width) + "s"
			if len(row) > j && len(cell) > i {
				fmt.Printf(printFormat, cell[i])
			} else {
				fmt.Printf(printFormat, "")
			}
			fmt.Printf("│")
		}
		fmt.Printf("\n")
	}
}

func printLine(widths []int) {
	fmt.Printf("├")
	for i := range widths {
		for j := 0; j < widths[i]; j++ {
			fmt.Printf("─")
		}
		if i != len(widths)-1 {
			fmt.Printf("┼")
		} else {
			fmt.Printf("┤\n")
		}
	}
}

func printBottom(widths []int) {
	fmt.Printf("└")
	for i := range widths {
		for j := 0; j < widths[i]; j++ {
			fmt.Printf("─")
		}
		if i != len(widths)-1 {
			fmt.Printf("┴")
		} else {
			fmt.Printf("┘\n")
		}
	}
}

func getValueStr(value reflect.Value) string {
	val := value
	valType := val.Kind()
	if valType == reflect.Interface {
		valType = val.Elem().Kind()
		val = val.Elem()
	}
	switch valType {
	case reflect.Slice, reflect.Array:
		str := "["
		for i := 0; i < val.Len(); i++ {
			str += " " + getValueStr(val.Index(i))
			if i != val.Len()-1 {
				str += ","
			}
		}
		str += " ]"
		return str
	case reflect.String:
		return getRefreshedStr(val.String())
	default:
		return fmt.Sprintf("%v", val)
	}
}

func getRefreshedStr(str string) string {
	newStr := str
	for len(newStr) > 0 && newStr[0] == '\b' {
		newStr = newStr[1:]
	}
	tabSpaces := strings.Repeat(" ", 8)
	newStr = strings.ReplaceAll(newStr, "\t", tabSpaces)
	return newStr
}

func getVisibleLen(str string) int {
	re := regexp.MustCompile(`\x1b\[[0-9;]*[mGKH]`)
	tempStr := re.ReplaceAllString(str, "")

	visible := []rune{}
	for _, char := range tempStr {
		if char == '\b' {
			if len(visible) > 0 {
				visible = visible[:len(visible)-1]
			}
		} else {
			visible = append(visible, char)
		}
	}

	return len(visible)
}
