package framework

func SpacingIndentation(i int, len int) string {
	if len > i {
		return ""
	}
	// return string of spaces
	i = i - len
	var str string
	for i > 0 {
		str += " "
		i--
	}
	return str
}
