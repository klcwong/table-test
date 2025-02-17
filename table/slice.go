package table

import (
	"reflect"
	"strconv"
)

func getSliceHeader(data any) []string {
	value := reflect.ValueOf(data)
	elementType := value.Type().Elem().Kind()
	header := []string{}
	header = append(header, "Index")
	switch elementType {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)
			for j := 0; j < element.Len(); j++ {
				for len(header) <= element.Len() {
					header = append(header, strconv.Itoa(len(header)-1))
				}
			}
		}
	default:
		header = append(header, "Value")
	}
	return header
}

func getSliceBody(data any) [][]string {
	value := reflect.ValueOf(data)
	elementType := value.Type().Elem().Kind()
	body := [][]string{}
	switch elementType {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			element := value.Index(i)
			newValues := []string{}
			newValues = append(newValues, strconv.Itoa(i))
			for j := 0; j < element.Len(); j++ {
				newValues = append(newValues, getValueString(value.Index(i).Index(j)))
			}
			body = append(body, newValues)
		}
	default:
		for i := 0; i < value.Len(); i++ {
			newValues := []string{}
			newValues = append(newValues, strconv.Itoa(i))
			newValues = append(newValues, getValueString(value.Index(i)))
			body = append(body, newValues)
		}
	}
	return body
}

func printSlice(slice any) {
	header := getSliceHeader(slice)
	body := getSliceBody(slice)
	widths := getWidths(header, body)
	printTable(header, body, widths)
}
