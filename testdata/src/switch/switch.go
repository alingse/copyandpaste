package switchdemo

// nolint:unused
func switchDemo(code string) string {
	var name string

	switch code {
	case "1", "2":
		name = "answer"
	case "3":
		name = "answer"
	}
	return name
}
