//LazyEvent v1 is a simple event log package that can work with files & CLI simultaneously.
//It unifies event struct and helps in creating event log & event map for any app.
//
//LE can prove useful in apps that require fast but easy-readable logging or event stacking.
package v1

import (
	"fmt"
	"regexp"
	"time"

	"github.com/lazybark/go-helpers/cli/clf"
	"github.com/lazybark/go-helpers/npt"
)

//ansiEscaper is a regexp to find ANSI escape characters in text
var ansiEscaper = regexp.MustCompile(`\033\[\d*m`)

//Delimeter is the default delimeter for log parts
var Delimeter = "	"

//Event is a basic event struct that holds all necessary info to build log record
type Event struct {
	etype      Etype
	source     Evsource
	time       npt.NPT
	text       string
	loggerType loggerType

	//EscapeAnsi makes logger to remove ANSI escape sequences in log text
	//in case it's being written to file
	escapeAnsi bool

	format clf.Format

	proc *EvProc
}

//ChainEvent is the struct to keep event in events chain (in memory)
//of event processor.
type chainEvent struct {
	etype  Etype
	source Evsource
	time   npt.NPT
	text   string
}

//EventConverted represents serialized (except Time) values of event fields that are meant to be exported.
//It should be used to return event chain to external routines.
type EventConverted struct {
	Type   string
	Source string
	Time   time.Time
	Text   string
}

//Esource represents source and should be created by external user
type Evsource struct {
	Text   string
	Format string
	Open   string
	Close  string
}

var (
	//EvsEmpty is an empty source to create log records with no sources
	EvsEmpty = Evsource{}

	//EvsDebug is a default event source to mark debug messages
	EvsDebug = Evsource{
		Text:   "DEBUG",
		Format: "",
		Open:   "[",
		Close:  "]",
	}

	//EvsMain represents main() function as an event source
	EvsMain = Evsource{
		Text:   "MAIN",
		Format: "",
		Open:   "[",
		Close:  "]",
	}
)

//Event returns new Event instance with default values and args serialized into string,
//but source is set to s
func (s Evsource) Event(p *EvProc, args ...interface{}) Event {
	return p.newEvent(TAll, s, fmt.Sprint(args...), all, false, clf.FNone)
}

//Info returns new Event instance with Info type, default values and args serialized into string,
//but source is set to s
func (s Evsource) Info(p *EvProc, args ...interface{}) Event {
	return p.newEvent(TInfo, s, fmt.Sprint(args...), all, false, clf.FNone)
}

//Note returns new Event instance with Note type, default values and args serialized into string,
//but source is set to s
func (s Evsource) Note(p *EvProc, args ...interface{}) Event {
	return p.newEvent(TNote, s, fmt.Sprint(args...), all, false, clf.FNone)
}

//Warning returns new Event instance with Warning type, default values and args serialized into string,
//but source is set to s
func (s Evsource) Warning(p *EvProc, args ...interface{}) Event {
	return p.newEvent(TWarning, s, fmt.Sprint(args...), all, false, clf.FNone)
}

//Error returns new Event instance Error Error type, default values and args serialized into string,
//but source is set to s
func (s Evsource) Error(p *EvProc, args ...interface{}) Event {
	return p.newEvent(TError, s, fmt.Sprint(args...), all, false, clf.FNone)
}

//Panic returns new Event instance with Panic type, default values and args serialized into string,
//but source is set to s
func (s Evsource) Panic(p *EvProc, args ...interface{}) Event {
	return p.newEvent(TPanic, s, fmt.Sprint(args...), all, false, clf.FNone)
}

//Critical returns new Event instance with Critical type, default values and args serialized into string,
//but source is set to s
func (s Evsource) Critical(p *EvProc, args ...interface{}) Event {
	return p.newEvent(TCritical, s, fmt.Sprint(args...), all, false, clf.FNone)
}

//Fatal returns new Event instance with Fatal type, default values and args serialized into string,
//but source is set to s
func (s Evsource) Fatal(p *EvProc, args ...interface{}) Event {
	return p.newEvent(TFatal, s, fmt.Sprint(args...), all, false, clf.FNone)
}

//Etype is the type of event to correctly treat it.
type Etype uint8

//eTypeNames is an array with text-represented Etypes
var eTypeNames = [...]string{
	"",
	"Info",
	"Note",
	"Warning",
	"Error",
	"Panic",
	"Critical",
	"Fatal",
}

const (
	e_types_start Etype = iota

	//All represents all levels. It will be empty in logs and passed to every logger
	//according to loggertype only
	TAll

	//Info informs human about any event in app
	TInfo

	//Note is like info, but with higher priority (e.g. for human-reader to make notes)
	TNote

	//Warning warns human about potentially dangreous situation
	TWarning

	//Error reports that something bad has happened
	TError

	//Critical reports that something really bad has happened
	TCritical

	//Panic makes app panic after event is logged
	TPanic

	//Fatal makes app exit after event is logged
	TFatal

	e_types_end
)

//String returns event type text representation
func (l Etype) String() string {
	if l <= e_types_start || l >= e_types_end {
		return ""
	}
	return eTypeNames[l-1]
}

//Check returns false in case event type in unknown and should not be used
func (t Etype) Check() bool {
	if t <= e_types_start || t >= e_types_end {
		return false
	}
	return true
}

//IsErr returns true in case event type represents error. It can be used
//by external app to treat events according to its internal logic.
func (e Etype) IsErr() bool {
	//All events above Warning are errors, but with different priority and consequences
	return e > TWarning
}

//IsErr returns true in case event type represents error. It can be used
//by external app to treat events according to its internal logic.
func (e Event) IsErr() bool {
	return e.etype.IsErr()
}

//AvoidANSI uses regex to drop all ANSI escape sequences in event text.
func (e *Event) AvoidANSI() {
	e.text = ansiEscaper.ReplaceAllString(e.text, "")
}
