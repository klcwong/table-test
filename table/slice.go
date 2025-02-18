package table

import (
	"reflect"
	"strconv"
)

func getSliceHeader(data any) []string {
	value := reflect.ValueOf(data)
	elementType := value.Type().Elem().Kind()
	header := []string{}
	header = append(header, getColorStr("Index", colors.yellow))
	switch elementType {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)
			for j := 0; j < element.Len(); j++ {
				for len(header) <= element.Len() {
					title := getColorStr(strconv.Itoa(len(header)-1), colors.green)
					header = append(header, title)
				}
			}
		}
	default:
		header = append(header, getColorStr("Value", colors.green))
	}
	return header
}

func getSliceBody(data any) [][]string {
	value := reflect.ValueOf(data)
	elementType := value.Type().Elem().Kind()
	body := [][]string{}
	for i := 0; i < value.Len(); i++ {
		newValues := []string{}
		newValues = append(newValues, getColorStr(strconv.Itoa(i), colors.yellow))
		switch elementType {
		case reflect.Slice, reflect.Array:
			element := value.Index(i)
			for j := 0; j < element.Len(); j++ {
				newValues = append(newValues, getValueString(value.Index(i).Index(j)))
			}
		default:
			newValues = append(newValues, getValueString(value.Index(i)))
		}
		body = append(body, newValues)
	}
	return body
}

func printSlice(slice any) {
	header := getSliceHeader(slice)
	body := getSliceBody(slice)
	widths := getWidths(header, body)
	printTable(header, body, widths)
}
