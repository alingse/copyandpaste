package switchdemo

// nolint:unused
func switchDemo(code string) string {
	var name string

	switch code { // want `Duplicate case body found for case "3": and case "1","2": Is it a copy and paste?`
	case "1", "2":
		name = "answer"
	case "3":
		name = "answer"
	}
	return name
}
