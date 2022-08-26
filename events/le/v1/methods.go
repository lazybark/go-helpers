// PACKAGE IS DEPRECATED: use github.com/lazybark/lazyevent instead
// LazyEvent v1 is a simple event log package that can work with files & CLI simultaneously.
// It unifies event struct and helps in creating event log & event map for any app.
//
// LE can prove useful in apps that require fast but easy-readable logging or event stacking.
package v1

import (
	"time"

	"github.com/lazybark/go-helpers/cli/clf"
)

// GetText returns current text value
func (e Event) GetText() string {
	return e.text
}

// Text sets new text value for event and returns new Event
func (e Event) Text(s string) Event {
	e.text = s
	return e
}

// TextSet sets new text value for event
func (e *Event) TextSet(s string) {
	e.text = s
}

// Append adds string to event text
func (e *Event) Append(s string) {
	e.text += s
}

// Time sets different time for the event and returns new Event
func (e Event) Time(t time.Time) Event {
	e.time.FromTime(t)
	return e
}

// Time sets different time for the event
func (e *Event) TimeSet(t time.Time) {
	e.time.FromTime(t)
}

// Src sets different source for the event and returns new Event
func (e Event) Src(s Evsource) Event {
	e.source = s
	return e
}

// SrcSet sets different source for the event
func (e *Event) SrcSet(s Evsource) {
	e.source = s
}

// ConsoleOnly marks event to be logged only in CLI and returns new Event
func (e Event) ConsoleOnly() Event {
	e.loggerType = console
	return e
}

// ConsoleOnlySet marks event to be logged only in CLI
func (e *Event) ConsoleOnlySet() {
	e.loggerType = console
}

// LogfileOnly marks event to be logged only in logfile and returns new Event
func (e Event) LogfileOnly() Event {
	e.loggerType = file
	return e
}

// LogfileOnlySet marks event to be logged only in logfile
func (e *Event) LogfileOnlySet() {
	e.loggerType = file
}

// EscapeAnsi marks event as EscapeAnsi = true, which means that in case
// passing to a file, all ANSI escape sequences will be deleted from event text.
// It's useful in cases when same text must be logged to file and CLI,
// but CLI must have colored output.
//
// Method may work relatively slow on big texts
func (e Event) EscapeAnsi() Event {
	e.escapeAnsi = true
	return e
}

// EscapeAnsiSet marks event as EscapeAnsi = true, which means that in case
// passing to a file, all ANSI escape sequences will be deleted from event text.
// It's useful in cases when same text must be logged to file and CLI,
// but CLI must have colored output.
//
// Method may work relatively slow on big texts
func (e *Event) EscapeAnsiSet() {
	e.escapeAnsi = true
}

// Log logs the event via event processor that it's associated to
func (e Event) Log() {
	e.proc.log(e)
}

// Info sets event type to Info and returns new Event
func (e Event) Info() Event {
	e.etype = TInfo
	return e
}

// InfoSet sets event type to Info
func (e *Event) InfoSet() {
	e.etype = TInfo
}

// Error sets event type to Error and returns new Event
func (e Event) Error() Event {
	e.etype = TError
	return e
}

// ErrorSet sets event type to Error
func (e *Event) ErrorSet() {
	e.etype = TError
}

// Note sets event type to Note and returns new Event
func (e Event) Note() Event {
	e.etype = TNote
	return e
}

// NoteSet sets event type to Note and returns new Event
func (e *Event) NoteSet() {
	e.etype = TNote
}

// Warning sets event type to Warning and returns new Event
func (e Event) Warning() Event {
	e.etype = TWarning
	return e
}

// WarningSet sets event type to Warning
func (e *Event) WarningSet() {
	e.etype = TWarning
}

// Panic sets event type to Panic and returns new Event
func (e Event) Panic() Event {
	e.etype = TPanic
	return e
}

// PanicSet sets event type to Panic
func (e *Event) PanicSet() {
	e.etype = TPanic
}

// Critical sets event type to critical and returns new Event
func (e Event) Critical() Event {
	e.etype = TCritical
	return e
}

// Critical sets event type to critical
func (e *Event) CriticalSet() {
	e.etype = TCritical
}

// Fatal sets event type to fatal and returns new Event
func (e Event) Fatal() Event {
	e.etype = TFatal
	return e
}

// Fatal sets event type to fatal
func (e *Event) FatalSet() {
	e.etype = TFatal
}

// Red sets event CLI text color to Red and returns new Event
func (e Event) Red() Event {
	e.format = clf.CRed
	return e
}

// RedSet sets event CLI text color to Red
func (e *Event) RedSet() {
	e.format = clf.CRed
}

// Green sets event CLI text color to Green and returns new Event
func (e Event) Green() Event {
	e.format = clf.CGreen
	return e
}

// GreenSet sets event CLI text color to Green
func (e *Event) GreenSet() {
	e.format = clf.CGreen
}

// Red sets event CLI text color to Yellow and returns new Event
func (e Event) Yellow() Event {
	e.format = clf.CYellow
	return e
}

// Red sets event CLI text color to Blue and returns new Event
func (e Event) Blue() Event {
	e.format = clf.CBlue
	return e
}

// Red sets event CLI text color to Magenta and returns new Event
func (e Event) Magenta() Event {
	e.format = clf.CMagenta
	return e
}

// Red sets event CLI text color to Cyan and returns new Event
func (e Event) Cyan() Event {
	e.format = clf.CCyan
	return e
}

// Red sets event CLI text color to Gray and returns new Event
func (e Event) Gray() Event {
	e.format = clf.CGray
	return e
}

// Red sets event CLI text color to Gray
func (e *Event) GraySet() {
	e.format = clf.CGray
}

// Red sets event CLI text background to Black and returns new Event
func (e Event) BgBlack() Event {
	e.format = clf.BgBlack
	return e
}

// Red sets event CLI text background to Black
func (e *Event) BgBlackSet() {
	e.format = clf.BgBlack
}

// Red sets event CLI text background to Red and returns new Event
func (e Event) BgRed() Event {
	e.format = clf.BgRed
	return e
}

// Red sets event CLI text background to Red
func (e *Event) BgRedSet() {
	e.format = clf.BgRed
}

// Red sets event CLI text background to Green and returns new Event
func (e Event) BgGreen() Event {
	e.format = clf.BgGreen
	return e
}

// Red sets event CLI text background to Green
func (e *Event) BgGreenSet() {
	e.format = clf.BgGreen
}

// Red sets event CLI text background to Yellow and returns new Event
func (e Event) BgYellow() Event {
	e.format = clf.BgYellow
	return e
}

// BgYellowSet sets event CLI text background to Yellow
func (e *Event) BgYellowSet() {
	e.format = clf.BgYellow
}

// Red sets event CLI text background to Magenta and returns new Event
func (e Event) BgMagenta() Event {
	e.format = clf.BgMagenta
	return e
}

// Red sets event CLI text background to Magenta
func (e *Event) BgMagentaSet() {
	e.format = clf.BgMagenta
}

// Red sets event CLI text background to Cyan and returns new Event
func (e Event) BgCyan() Event {
	e.format = clf.BgCyan
	return e
}

// BgCyanSet sets event CLI text background to Cyan
func (e *Event) BgCyanSet() {
	e.format = clf.BgCyan
}

// Red sets event CLI text background to White and returns new Event
func (e Event) BgWhite() Event {
	e.format = clf.BgWhite
	return e
}

// BgWhiteSet sets event CLI text background to White
func (e *Event) BgWhiteSet() {
	e.format = clf.BgWhite
}

// Red sets event CLI text style to Underline and returns new Event
func (e Event) Underline() Event {
	e.format = clf.Underline
	return e
}

// UnderlineSet sets event CLI text background to Underline
func (e *Event) UnderlineSet() {
	e.format = clf.Underline
}
