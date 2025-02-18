package table

import (
	"fmt"
	"reflect"
	"regexp"
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
		widths = append(widths, getVisibleLen(item)+spaces)
	}
	for _, row := range boby {
		for i, item := range row {
			widths[i] = max(widths[i], getVisibleLen(item)+spaces)
		}
	}
	return widths
}

func printTable(header []string, body [][]string, widths []int) {
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

func printRow(row []string, widths []int) {
	fmt.Printf("│")
	for i := range widths {
		diff := len(row[i]) - getVisibleLen(row[i])
		width := widths[i] - 1 + diff
		printFormat := " %-" + strconv.Itoa(width) + "s"
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
		str := "["
		for i := 0; i < value.Len(); i++ {
			str += " " + getValueString(value.Index(i))
			if i != value.Len()-1 {
				str += ","
			}
		}
		str += " ]"
		return str
	default:
		return fmt.Sprintf("%v", value)
	}
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
