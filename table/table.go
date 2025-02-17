package table

import (
	"fmt"
	"reflect"
	"strconv"
)

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

func getWidths(header []string, boby [][]string) []int {
	const spaces = 3
	widths := []int{}
	for _, item := range header {
		widths = append(widths, len(item)+spaces)
	}
	for _, row := range boby {
		for i, item := range row {
			widths[i] = max(widths[i], len(item)+spaces)
		}
	}
	return widths
}

func printTable(header []string, body [][]string, widths []int) {
	printTop(widths)
	printRow(header, widths, true)
	for _, row := range body {
		printLine(widths)
		printRow(row, widths, false)
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

func printRow(row []string, widths []int, isHeader bool) {
	fmt.Printf("│")
	for i := range widths {
		printFormat := " %-" + strconv.Itoa(widths[i]-1) + "s"
		if i == 0 {
			printFormat = " \033[38;5;214m%-" + strconv.Itoa(widths[i]-1) + "s\033[0m"
		}
		if isHeader {
			printFormat = " \033[32m%-" + strconv.Itoa(widths[i]-1) + "s\033[0m"
		}
		if len(row) > i {
			fmt.Printf(printFormat, row[i])
		} else {
			fmt.Printf(printFormat, "")
		}
		fmt.Printf("│")
	}
	fmt.Printf("\n")
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

func getValueString(value reflect.Value) string {
	valueType := value.Kind()
	switch valueType {
	case reflect.Slice, reflect.Array:
		str := "[abc"
		for i := 0; i < value.Len(); i++ {

		}
		str += "]"
		return str
	case reflect.String:
		return fmt.Sprintf("\"%v\"", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}
