package switchdemo

// nolint:unused
func switchDemo(code string) string {
	var name string

	switch code {
	case "1", "2":
		name = "answer"
	case "3": // want `Duplicate case body found for case "3": and case "1","2": Is it a copy and paste?`
		name = "answer"
	default:
		name = "answer"
	}
	return name
}
