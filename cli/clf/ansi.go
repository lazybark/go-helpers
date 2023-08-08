// clf if a command line format package to add colors and styles to CLI-apps
package clf

import "fmt"

type Format int

const (
	formats_start Format = iota

	//FNone exists for compatibility with any external methods that require clf formatting
	FNone

	//FReset resets the formatting
	FReset

	//CRed marks start of red text color
	CRed

	//CGreen marks start of green text color
	CGreen

	//CYellow marks start of yellow text color
	CYellow

	//CBlue marks start of blue text color
	CBlue

	//CMagenta marks start of magenta text color
	CMagenta

	//CCyan marks start of cyan text color
	CCyan

	//CGray marks start of gray text color
	CGray

	//CWhite marks start of white text color
	CWhite

	//BgBlack marks start of black background
	BgBlack

	//BgRed marks start of red background
	BgRed

	//BgGreen marks start of green background
	BgGreen

	//BgYellow marks start of yellow background
	BgYellow

	//BgBlue marks start of blue background
	BgBlue

	//BgMagenta marks start of magenta background
	BgMagenta

	//BgCyan marks start of cyan background
	BgCyan

	//BgWhite marks start of white background
	BgWhite

	Bold
	Underline
	Blink
	Reverse
	UnReverse

	formats_end
)

var FormatStrings = [...]string{
	"",
	"\033[0m",
	"\033[31m",
	"\033[32m",
	"\033[33m",
	"\033[34m",
	"\033[35m",
	"\033[36m",
	"\033[37m",
	"\033[97m",
	"\033[40m",
	"\033[41m",
	"\033[42m",
	"\033[43m",
	"\033[44m",
	"\033[45m",
	"\033[46m",
	"\033[47m",
	"\033[1m",
	"\033[4m",
	"\033[5m",
	"\033[7m",
	"\033[27m",
}

//CheckType returns true if c is a legit Format
func (c Format) CheckType() bool {
	return formats_start < c && c < formats_end
}

//String returns ansi-formatting of associated Format value
func (c Format) String() string {
	if !c.CheckType() {
		return ""
	}

	return FormatStrings[c-1]
}

//ansi returns ansi-formatting, represented by c
func ansi(c Format) string {
	if !c.CheckType() {
		return ""
	}

	return FormatStrings[c-1]
}

//Red marks text as Red
func Red(args ...interface{}) string {
	return ansi(CRed) + fmt.Sprint(args...) + ansi(FReset)
}

//Green marks text as Green
func Green(args ...interface{}) string {
	return ansi(CGreen) + fmt.Sprint(args...) + ansi(FReset)
}

//Yellow marks text as Yellow
func Yellow(args ...interface{}) string {
	return ansi(CYellow) + fmt.Sprint(args...) + ansi(FReset)
}

//Blue marks text as Blue
func Blue(args ...interface{}) string {
	return ansi(CBlue) + fmt.Sprint(args...) + ansi(FReset)
}

//Magenta marks text as Magenta
func Magenta(args ...interface{}) string {
	return ansi(CMagenta) + fmt.Sprint(args...) + ansi(FReset)
}

//Cyan marks text as Cyan
func Cyan(args ...interface{}) string {
	return ansi(CCyan) + fmt.Sprint(args...) + ansi(FReset)
}

//Gray marks text as Gray
func Gray(args ...interface{}) string {
	return ansi(CGray) + fmt.Sprint(args...) + ansi(FReset)
}

//White marks text as White
func White(args ...interface{}) string {
	return ansi(CWhite) + fmt.Sprint(args...) + ansi(FReset)
}

//BBlack marks text as Black background
func BBlack(args ...interface{}) string {
	return ansi(BgBlack) + fmt.Sprint(args...) + ansi(FReset)
}

//BRed marks text as Red background
func BRed(args ...interface{}) string {
	return ansi(BgRed) + fmt.Sprint(args...) + ansi(FReset)
}

//BGreen marks text as Green background
func BGreen(args ...interface{}) string {
	return ansi(BgGreen) + fmt.Sprint(args...) + ansi(FReset)
}

//BYellow marks text as Yellow background
func BYellow(args ...interface{}) string {
	return ansi(BgYellow) + fmt.Sprint(args...) + ansi(FReset)
}

//BBlue marks text as Blue background
func BBlue(args ...interface{}) string {
	return ansi(BgBlue) + fmt.Sprint(args...) + ansi(FReset)
}

//BMagenta marks text as Magenta background
func BMagenta(args ...interface{}) string {
	return ansi(BgMagenta) + fmt.Sprint(args...) + ansi(FReset)
}

//BCyan marks text as Cyan background
func BCyan(args ...interface{}) string {
	return ansi(BgCyan) + fmt.Sprint(args...) + ansi(FReset)
}

//BWhite marks text as White background
func BWhite(args ...interface{}) string {
	return ansi(BgWhite) + fmt.Sprint(args...) + ansi(FReset)
}

//FBold marks text as bold-styled
func FBold(args ...interface{}) string {
	return ansi(Bold) + fmt.Sprint(args...) + ansi(FReset)
}

//FUnderline marks text as underline-styled
func FUnderline(args ...interface{}) string {
	return ansi(Underline) + fmt.Sprint(args...) + ansi(FReset)
}

//FBlink marks text as blinking
func FBlink(args ...interface{}) string {
	return ansi(Blink) + fmt.Sprint(args...) + ansi(FReset)
}

//FReverse marks text as reversed style
func FReverse(args ...interface{}) string {
	return ansi(Reverse) + fmt.Sprint(args...) + ansi(FReset)
}

//FUnReverse marks text as unreversed style
func FUnReverse(args ...interface{}) string {
	return ansi(UnReverse) + fmt.Sprint(args...) + ansi(FReset)
}

//Reset adds formatting reset directive
func Reset() string {
	return ansi(FReset)
}
