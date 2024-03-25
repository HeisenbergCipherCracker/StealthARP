package lib

type Ether struct {
	domain int
}

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
	ASCII string
}

func colorizeString(clr string, s string) string {
	switch {
	case clr == "red":
		return s + "\x1b[38;2;255;0;0m█\x1b[0m"

	case clr == "green":
		return s + "\x1b[38;2;0;255;0m█\x1b[0m"

	case clr == "blue":
		return s + "\x1b[38;2;0;0;255m█\x1b[0m"
	default:
		return s
	}
}

func main() {
}
