//clf if a command line format package to add colors and styles to CLI-apps
package clf

import "fmt"

type Format int

const (
	formats_start Format = iota

	FReset
	CRed
	CGreen
	CYellow
	CBlue
	CMagenta
	CCyan
	CGray
	CWhite
	BgBlack
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite

	Bold
	Underline
	Blink
	Reverse
	UnReverse

	formats_end
)

var FormatStrings = [...]string{
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

//String returns ansi-formatting of associated Format value
func (c Format) String() string {
	if c <= formats_start || c >= formats_end {
		return ""
	}
	return FormatStrings[c-1]
}

//ansi returns ansi-formatting, represented by c
func ansi(c Format) string {
	if c <= formats_start || c >= formats_end {
		return ""
	}
	return FormatStrings[c-1]
}

func Red(args ...interface{}) string {
	return ansi(CRed) + fmt.Sprint(args...) + ansi(FReset)
}

func Green(args ...interface{}) string {
	return ansi(CGreen) + fmt.Sprint(args...) + ansi(FReset)
}

func Yellow(args ...interface{}) string {
	return ansi(CYellow) + fmt.Sprint(args...) + ansi(FReset)
}

func Blue(args ...interface{}) string {
	return ansi(CBlue) + fmt.Sprint(args...) + ansi(FReset)
}

func Magenta(args ...interface{}) string {
	return ansi(CMagenta) + fmt.Sprint(args...) + ansi(FReset)
}

func Cyan(args ...interface{}) string {
	return ansi(CCyan) + fmt.Sprint(args...) + ansi(FReset)
}
func Gray(args ...interface{}) string {
	return ansi(CGray) + fmt.Sprint(args...) + ansi(FReset)
}
func White(args ...interface{}) string {
	return ansi(CWhite) + fmt.Sprint(args...) + ansi(FReset)
}

func BBlack(args ...interface{}) string {
	return ansi(BgBlack) + fmt.Sprint(args...) + ansi(FReset)
}

func BRed(args ...interface{}) string {
	return ansi(BgRed) + fmt.Sprint(args...) + ansi(FReset)
}

func BGreen(args ...interface{}) string {
	return ansi(BgGreen) + fmt.Sprint(args...) + ansi(FReset)
}

func BYellow(args ...interface{}) string {
	return ansi(BgYellow) + fmt.Sprint(args...) + ansi(FReset)
}

func BBlue(args ...interface{}) string {
	return ansi(BgBlue) + fmt.Sprint(args...) + ansi(FReset)
}

func BMagenta(args ...interface{}) string {
	return ansi(BgMagenta) + fmt.Sprint(args...) + ansi(FReset)
}

func BCyan(args ...interface{}) string {
	return ansi(BgCyan) + fmt.Sprint(args...) + ansi(FReset)
}

func BWhite(args ...interface{}) string {
	return ansi(BgWhite) + fmt.Sprint(args...) + ansi(FReset)
}

func FBold(args ...interface{}) string {
	return ansi(Bold) + fmt.Sprint(args...) + ansi(FReset)
}

func FUnderline(args ...interface{}) string {
	return ansi(Underline) + fmt.Sprint(args...) + ansi(FReset)
}

func FBlink(args ...interface{}) string {
	return ansi(Blink) + fmt.Sprint(args...) + ansi(FReset)
}

func FReverse(args ...interface{}) string {
	return ansi(Reverse) + fmt.Sprint(args...) + ansi(FReset)
}

func FUnReverse(args ...interface{}) string {
	return ansi(UnReverse) + fmt.Sprint(args...) + ansi(FReset)
}

func Reset() string {
	return ansi(FReset)
}
