package clf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnsi(t *testing.T) {
	//Start with empty cases
	s := ansi(formats_start)
	assert.Empty(t, s)
	s = ansi(formats_start - 1)
	assert.Empty(t, s)
	s = ansi(formats_end)
	assert.Empty(t, s)
	s = ansi(formats_end + 1)
	assert.Empty(t, s)
	s = ansi(FNone)
	assert.Empty(t, s)

	//Randoms
	s = ansi(FReset)
	assert.Equal(t, "\033[0m", s)
	s = ansi(UnReverse)
	assert.Equal(t, "\033[27m", s)

	//Whole list
	for i, f := range FormatStrings {
		s = ansi(Format(i + 1))
		assert.Equal(t, f, s)
	}
}

func TestAnsiFormatting(t *testing.T) {
	text := "some text for CLI"

	s := Red(text)
	assert.Equal(t, ansi(CRed)+text+ansi(FReset), s)

	s = Red(text)
	assert.Equal(t, ansi(CRed)+text+ansi(FReset), s)
	s = Green(text)
	assert.Equal(t, ansi(CGreen)+text+ansi(FReset), s)
	s = Yellow(text)
	assert.Equal(t, ansi(CYellow)+text+ansi(FReset), s)
	s = Blue(text)
	assert.Equal(t, ansi(CBlue)+text+ansi(FReset), s)
	s = Magenta(text)
	assert.Equal(t, ansi(CMagenta)+text+ansi(FReset), s)
	s = Cyan(text)
	assert.Equal(t, ansi(CCyan)+text+ansi(FReset), s)
	s = Gray(text)
	assert.Equal(t, ansi(CGray)+text+ansi(FReset), s)
	s = White(text)
	assert.Equal(t, ansi(CWhite)+text+ansi(FReset), s)
	s = BBlack(text)
	assert.Equal(t, ansi(BgBlack)+text+ansi(FReset), s)
	s = BRed(text)
	assert.Equal(t, ansi(BgRed)+text+ansi(FReset), s)
	s = BGreen(text)
	assert.Equal(t, ansi(BgGreen)+text+ansi(FReset), s)
	s = BYellow(text)
	assert.Equal(t, ansi(BgYellow)+text+ansi(FReset), s)
	s = BBlue(text)
	assert.Equal(t, ansi(BgBlue)+text+ansi(FReset), s)
	s = BMagenta(text)
	assert.Equal(t, ansi(BgMagenta)+text+ansi(FReset), s)
	s = BCyan(text)
	assert.Equal(t, ansi(BgCyan)+text+ansi(FReset), s)
	s = BWhite(text)
	assert.Equal(t, ansi(BgWhite)+text+ansi(FReset), s)
	s = FBold(text)
	assert.Equal(t, ansi(Bold)+text+ansi(FReset), s)
	s = FUnderline(text)
	assert.Equal(t, ansi(Underline)+text+ansi(FReset), s)
	s = FBlink(text)
	assert.Equal(t, ansi(Blink)+text+ansi(FReset), s)
	s = FReverse(text)
	assert.Equal(t, ansi(Reverse)+text+ansi(FReset), s)
	s = FUnReverse(text)
	assert.Equal(t, ansi(UnReverse)+text+ansi(FReset), s)
	s = Reset()
	assert.Equal(t, ansi(FReset), s)
}
