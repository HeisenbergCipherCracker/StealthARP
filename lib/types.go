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

func colorizeString(c Color, s string) string {
	switch {
	case c.Red > 0:
		return s + "\x1b[38;2;255;0;0m█\x1b[0m"

	case c.Green > 0:
		return s + "\x1b[38;2;0;255;0m█\x1b[0m"

	default:
		return s
	}
}

func main() {
}
