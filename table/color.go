package table

type color string

var colors = struct {
	green  color
	yellow color
}{
	green:  "\033[32m",
	yellow: "\033[33m",
}

func getColorStr(str string, color color) string {
	return string(color) + str + "\033[0m"
}
