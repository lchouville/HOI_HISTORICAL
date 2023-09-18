package hoiMod

type Color string

const (
	ColorBlack  Color = "\u001b[30m" // ANSI CODE
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\033[93m"
	// ColorYellow  Color = "\u001b[33m"
	ColorOrange  Color = "\033[38;5;208m"
	ColorBlue    Color = "\u001b[34m"
	ColorMagenta Color = "\u001b[35m"
	ColorCyan    Color = "\u001b[36m"
	ColorWhite   Color = "\u001b[37m"
	ColorReset   Color = "\u001b[0m"
)
