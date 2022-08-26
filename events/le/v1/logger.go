// PACKAGE IS DEPRECATED: use github.com/lazybark/lazyevent instead
// LazyEvent v1 is a simple event log package that can work with files & CLI simultaneously.
// It unifies event struct and helps in creating event log & event map for any app.
//
// LE can prove useful in apps that require fast but easy-readable logging or event stacking.
package v1

import (
	"fmt"
	"os"

	"github.com/lazybark/go-helpers/cli/clf"
)

// logger makes log records for events according to specified parameters
type logger struct {
	delimeter string
	types     []Etype
	file      *os.File
	lt        loggerType
}

// loggerType represents basic logger type to control log info flow
type loggerType uint8

const (
	//all makes logger suitable to any log message
	all loggerType = iota

	//file makes logger suitable only to messages that are meant to be recorded into file
	file

	//console makes logger suitable only to messages for CLI-output
	console
)

// log logs event according to its parameters
func (l *logger) log(e Event) error {
	sc, sf := l.getSource(e.source)
	if l.file != nil {
		if e.escapeAnsi {
			e.AvoidANSI()
		}
		//No need for mutex in this case as we call loggers one by one, waiting until last one finishes recording
		_, err := l.file.WriteString(fmt.Sprintf("%s%s%s%s%s%s\n", e.time.Time(), l.delimeter, e.etype, l.delimeter, sf, e.text))
		if err != nil {
			return err
		}
	} else {
		if e.format != clf.FNone {
			fmt.Printf("%s%s%s%s%s%s%s%s\n", e.format.String(), e.time.Time(), l.delimeter, e.etype, l.delimeter, sf, e.text, clf.FReset.String())
		} else {
			fmt.Printf("%s%s%s%s%s%s\n", e.time.Time(), l.delimeter, e.etype, l.delimeter, sc, e.text)
		}
	}
	return nil
}

// getSource returns formatted source for logger. First return is source for CLI,
// second is for file (will be different only in case ANSI-escapes were added)
func (l *logger) getSource(source Evsource) (string, string) {
	var s string
	if source.Text != "" {
		s = fmt.Sprintf("%s%s%s%s", source.Open, source.Text, source.Close, l.delimeter)
	}
	//Add some spices
	if source.Format == "red" {
		return clf.Red(s), s
	} else if source.Format == "cyan" {
		return clf.Cyan(s), s
	} else if source.Format == "green" {
		return clf.Green(s), s
	} else if source.Format == "yellow" {
		return clf.Yellow(s), s
	} else if source.Format == "blue" {
		return clf.Blue(s), s
	} else if source.Format == "magenta" {
		return clf.Magenta(s), s
	} else if source.Format == "gray" {
		return clf.Gray(s), s
	} else if source.Format == "black_back" {
		return clf.BBlack(s), s
	} else if source.Format == "red_black" {
		return clf.BRed(s), s
	} else if source.Format == "green_black" {
		return clf.BGreen(s), s
	} else if source.Format == "yellow_back" {
		return clf.BYellow(s), s
	} else if source.Format == "blue_back" {
		return clf.BBlue(s), s
	} else if source.Format == "magenta_back" {
		return clf.BMagenta(s), s
	} else if source.Format == "cyan_back" {
		return clf.BCyan(s), s
	} else if source.Format == "white_back" {
		return clf.BWhite(s), s
	} else if source.Format == "underline" {
		return clf.FUnderline(s), s
	} else if source.Format == "bold" {
		return clf.FBold(s), s
	}

	return s, s
}

// suits checks if logger has suitable type and matches event type
func (l *logger) suits(et Etype, lt loggerType) bool {
	for _, etype := range l.types {
		if etype == TAll || etype == et {
			if lt == all || lt == l.lt {
				return true
			}
		}
	}
	return false
}
