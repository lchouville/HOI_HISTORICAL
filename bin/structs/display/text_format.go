package hoiMod

import "strconv"

type Text string
type Color string

const (
	// Text
	Text_Bold          Text = "\033[1m"
	Text_Italic        Text = "\033[3m"
	Text_Underline     Text = "\033[4m"
	Text_Blink         Text = "\033[5m"
	Text_Reversed      Text = "\033[7m"
	Text_Hiddens       Text = "\033[8m"
	Text_Strikethrough Text = "\033[9m"
	// Color
	Color_Black   Text = "\033[38;2;0;0;0m" // ANSI CODE with RGB
	Color_Red     Text = "\033[38;2;255;0;0m"
	Color_Green   Text = "\033[38;2;0;255;0m"
	Color_Yellow  Text = "\033[38;2;255;255;0m"
	Color_Orange  Text = "\033[38;2;255;136;0m"
	Color_Blue    Text = "\033[38;2;0;0;255m"
	Color_Magenta Text = "\033[38;2;120;0;76m"
	Color_Cyan    Text = "\033[38;2;0;162;211m"
	Color_White   Text = "\033[38;2;255;255;255m"
	// Reset
	ResetAll           Text = "\033[m"
	ResetBold          Text = "\033[22m"
	ResetItalic        Text = "\033[23m"
	ResetUnderline     Text = "\033[24m"
	ResetReversed      Text = "\033[27m"
	ResetHiddens       Text = "\033[28m"
	ResetStrikethrough Text = "\033[29m"
)

func ColorFontRGB(r int, g int, b int) string {
	res := ("\033[38;2;" +
		strconv.Itoa(r) + ";" +
		strconv.Itoa(g) + ";" +
		strconv.Itoa(b) + "m")
	return res
}
func ColorBackGroundRGB(r int, g int, b int) string {
	res := ("\033[48;2;" +
		strconv.Itoa(r) + ";" +
		strconv.Itoa(g) + ";" +
		strconv.Itoa(b) + "m")
	return res
}
