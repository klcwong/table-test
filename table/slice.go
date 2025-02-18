package table

import (
	"reflect"
	"strconv"
	"strings"
)

func getSliceHeader(data any) row {
	value := reflect.ValueOf(data)
	elementType := value.Type().Elem().Kind()
	header := row{}
	header = append(header, cell{getColorStr("Index", colors.yellow)})
	switch elementType {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)
			for j := 0; j < element.Len(); j++ {
				for len(header) <= element.Len() {
					index := len(header) - 1
					title := getColorStr(strconv.Itoa(index), colors.green)
					header = append(header, cell{title})
				}
			}
		}
	case reflect.Struct:
		for i := 0; i < value.Type().Elem().NumField(); i++ {
			field := value.Type().Elem().Field(i)
			header = append(header, cell{getColorStr(field.Name, colors.green)})
		}
	case reflect.Map:
		// To-do
	default:
		header = append(header, cell{getColorStr("Value", colors.green)})
	}
	return header
}

func getSliceBody(data any) []row {
	value := reflect.ValueOf(data)
	body := []row{}
	elementType := value.Type().Elem().Kind()
	for i := 0; i < value.Len(); i++ {
		newValues := row{}
		newValues = append(newValues, cell{getColorStr(strconv.Itoa(i), colors.yellow)})
		element := value.Index(i)
		switch elementType {
		case reflect.Slice, reflect.Array:
			for j := 0; j < element.Len(); j++ {
				newValues = append(newValues, cell{getValueStr(element.Index(j))})
			}
		case reflect.Struct:
			for j := 0; j < element.NumField(); j++ {
				newValues = append(newValues, cell{getValueStr(element.Field(j))})
			}
		case reflect.Map:
			//
		case reflect.String:
			cell := strings.Split(getValueStr(value.Index(i)), "\n")
			newValues = append(newValues, cell)
		default:
			newValues = append(newValues, cell{getValueStr(value.Index(i))})
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
