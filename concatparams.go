package piscine

func ConcatParams(args []string) string {
	str := ""
	for index, data := range args {
		str += data
		if index != len(args)-1 {
			str += "\n"
		}
	}
	return str
}
