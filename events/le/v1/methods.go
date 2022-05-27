//LazyEvent v1 is a simple event log package that can work with files & CLI simultaneously.
//It unifies event struct and helps in creating event log & event map for any app.
//
//LE can prove useful in apps that require fast but easy-readable logging or event stacking.
package v1

import (
	"github.com/lazybark/go-helpers/cli/clf"
)

//GetText returns current text value
func (e Event) GetText() string {
	return e.text
}

//Text sets new text value for event
func (e Event) Text(s string) Event {
	e.text = s
	return e
}

//Append adds string to event text
func (e Event) Append(s string) Event {
	e.text += s
	return e
}

//Src sets different source for the event
func (e Event) Src(s Evsource) Event {
	e.source = s
	return e
}

//ConsoleOnly marks event to be logged only in CLI
func (e Event) ConsoleOnly() Event {
	e.loggerType = console
	return e
}

//LogfileOnly marks event to be logged only in logfile
func (e Event) LogfileOnly() Event {
	e.loggerType = file
	return e
}

//EscapeAnsi marks event as EscapeAnsi = true, which means that in case
//passing to a file, all ANSI escape sequences will be deleted from event text.
//It's useful in cases when same text must be logged to file and CLI,
//but CLI must have colored output.
//
//Method may work relatively slow on big texts
func (e Event) EscapeAnsi() Event {
	e.escapeAnsi = true
	return e
}

//Log logs the event via event processor that it's associated to
func (e Event) Log() {
	e.proc.log(e)
}

//Inf sets event type to Info
func (e Event) Inf() Event {
	e.etype = TInfo
	return e
}

//Err sets event type to Error
func (e Event) Err() Event {
	e.etype = TError
	return e
}

//Note sets event type to Note
func (e Event) Note() Event {
	e.etype = TNote
	return e
}

//Warn sets event type to Warning
func (e Event) Warn() Event {
	e.etype = TWarning
	return e
}

//Panic sets event type to Panic
func (e Event) Panic() Event {
	e.etype = TPanic
	return e
}

//Crit sets event type to critical
func (e Event) Crit() Event {
	e.etype = TCritical
	return e
}

//Fatal sets event type to fatal
func (e Event) Fatal() Event {
	e.etype = TFatal
	return e
}

//Red sets event CLI text color to Red
func (e Event) Red() Event {
	e.format = clf.CRed
	return e
}

//Red sets event CLI text color to Green
func (e Event) Green() Event {
	e.format = clf.CGreen
	return e
}

//Red sets event CLI text color to Yellow
func (e Event) Yellow() Event {
	e.format = clf.CYellow
	return e
}

//Red sets event CLI text color to Blue
func (e Event) Blue() Event {
	e.format = clf.CBlue
	return e
}

//Red sets event CLI text color to Magenta
func (e Event) Magenta() Event {
	e.format = clf.CMagenta
	return e
}

//Red sets event CLI text color to Cyan
func (e Event) Cyan() Event {
	e.format = clf.CCyan
	return e
}

//Red sets event CLI text color to Gray
func (e Event) Gray() Event {
	e.format = clf.CGray
	return e
}

//Red sets event CLI text background to Black
func (e Event) BgBlack() Event {
	e.format = clf.BgBlack
	return e
}

//Red sets event CLI text background to Red
func (e Event) BgRed() Event {
	e.format = clf.BgRed
	return e
}

//Red sets event CLI text background to Green
func (e Event) BgGreen() Event {
	e.format = clf.BgGreen
	return e
}

//Red sets event CLI text background to Yellow
func (e Event) BgYellow() Event {
	e.format = clf.BgYellow
	return e
}

//Red sets event CLI text background to Magenta
func (e Event) BgMagenta() Event {
	e.format = clf.BgMagenta
	return e
}

//Red sets event CLI text background to Cyan
func (e Event) BgCyan() Event {
	e.format = clf.BgCyan
	return e
}

//Red sets event CLI text background to White
func (e Event) BgWhite() Event {
	e.format = clf.BgWhite
	return e
}

//Red sets event CLI text background to Underline
func (e Event) Underline() Event {
	e.format = clf.Underline
	return e
}
