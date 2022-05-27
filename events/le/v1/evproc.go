//LazyEvent v1 is a simple event log package that can work with files & CLI simultaneously.
//It unifies event struct and helps in creating event log & event map for any app.
//
//LE can prove useful in apps that require fast but easy-readable logging or event stacking.
package v1

import (
	"fmt"
	"os"
	"sync"

	"github.com/lazybark/go-helpers/cli/clf"
	"github.com/lazybark/go-helpers/fsw"
	"github.com/lazybark/go-helpers/npt"
)

//EvProc is the controlling mechanism for event logging
type EvProc struct {
	delimeter   string
	chainLength int
	sourceLengh int
	loggers     []*logger
	evChain     []*chainEvent
	lcMutex     sync.Mutex
	ec          chan (Event)
}

//New returns pointer no new EvProc
func New(chainLength int) *EvProc {
	p := &EvProc{delimeter: Delimeter, chainLength: chainLength, sourceLengh: len(EvsDebug.Text), loggers: make([]*logger, 0), evChain: make([]*chainEvent, 0), lcMutex: sync.Mutex{}, ec: make(chan Event, 10)}
	go p.start()
	return p
}

//SetDelimeter changes delimeter beteen log parts for all loggers.
//Some delimeters can make log entries hard to read or just bad looking
func (p *EvProc) SetDelimeter(d string) {
	p.delimeter = d
	for _, l := range p.loggers {
		l.delimeter = d
	}
}

//Source creates new source for events that will be printed out next to type.
//It's better to keep source names about the same length, otherwise log lines
//may become hard to read.
func (p *EvProc) Source(t string, f string, o string, c string) Evsource {
	return Evsource{
		Text:   t,
		Format: f,
		Open:   o,
		Close:  c,
	}
}

//start launches routine to log events
func (p *EvProc) start() {
	var err error
	for e := range p.ec {
		for _, l := range p.loggers {
			if !l.suits(e.etype, e.loggerType) {
				continue
			}
			err = l.log(e)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error making log entry: %s", err)
			}
		}
		if e.etype == TPanic {
			panic(e.text)
		}
		if e.etype == TFatal {
			os.Exit(2)
		}
	}
}

//EventChain returns pointer to slice of events with length specified during EvProc creation
func (p *EvProc) EventChain() *[]*EventConverted {
	var ec []*EventConverted
	for _, e := range p.evChain {
		ec = append(ec, &EventConverted{Type: e.etype.String(), Source: e.source.Text, Time: e.time.Time(), Text: e.text})
	}
	return &ec
}

//eventToChain adds new event to chain and removes oldest event in case new one does not fit
func (p *EvProc) eventToChain(e Event) {
	len := len(p.evChain)
	min := 0
	p.lcMutex.Lock()
	if len >= p.chainLength {
		min = len - p.chainLength + 1
		p.evChain = p.evChain[min:]
	}
	p.evChain = append(p.evChain, &chainEvent{etype: e.etype, source: e.source, time: e.time, text: e.text})
	p.lcMutex.Unlock()
}

//newEvent creates new Event and puts it to chain
func (p *EvProc) newEvent(t Etype, s Evsource, text string, lt loggerType, ea bool, format clf.Format) Event {
	e := Event{etype: t, source: s, time: npt.Now(), text: text, loggerType: lt, escapeAnsi: ea, format: format, proc: p}
	if p.chainLength > 0 {
		p.eventToChain(e)
	}
	return e
}

func (p *EvProc) log(e Event) {
	p.eventToChain(e)
	p.ec <- e
}

//NewFile creates new logger to write into file and the file itself in case it does not exist.
//NewFile will create all directories in specified path.
func (p *EvProc) NewFile(path string, truncate bool, types ...Etype) error {
	if len(types) == 0 {
		types = append(types, TAll)
	} else {
		for _, t := range types {
			if !t.Check() {
				return fmt.Errorf("unknown event type: %d", t)
			}
		}
	}

	f, err := fsw.MakePathToFile(path, truncate)
	if err != nil {
		return fmt.Errorf("[EvProc] error making log path: %w", err)
	}
	var l logger
	l.delimeter = p.delimeter
	l.file = f
	l.lt = file
	l.types = make([]Etype, len(types))
	l.types = append(l.types, types...)
	p.loggers = append(p.loggers, &l)

	return nil
}

//NewConsole creates new logger to log messages in CLI
func (p *EvProc) NewConsole(types ...Etype) error {
	if len(types) == 0 {
		types = append(types, TAll)
	} else {
		for _, t := range types {
			fmt.Println(t)
			if !t.Check() {
				return fmt.Errorf("unknown event type: %d", t)
			}
		}
	}

	var l logger
	l.delimeter = p.delimeter
	l.lt = console
	l.types = make([]Etype, len(types))
	l.types = append(l.types, types...)
	p.loggers = append(p.loggers, &l)

	return nil
}

//Event returns new Event instance with default values and args serialized into string
func (p *EvProc) Event(args ...interface{}) Event {
	return p.newEvent(TAll, EvsEmpty, fmt.Sprint(args...), all, false, clf.FNone)
}

//Inf returns new Event instance with Info type, default values and args serialized into string
func (p *EvProc) Inf(args ...interface{}) Event {
	return p.newEvent(TInfo, EvsEmpty, fmt.Sprint(args...), all, false, clf.FNone)
}
