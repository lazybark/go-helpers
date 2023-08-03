package npt

import (
	"time"
)

// NPT is a time-holding object that does not have timezone pointer.
// It's more memory-effective than default Go time.Time and can be used in apps that store time
// in memory only and do not serialize it.
//
// For example: some action buffer that use time only for it's internal purposes of starting tasks
// at specific moment.
//
// Important: NPT does not provide exact precision up to nano. Max precision is up to a second.
type NPT struct {
	epoch int64
}

// Now creates NPT from current moment in time
func Now() NPT {
	t := NPT{
		epoch: time.Now().Unix(),
	}

	return t
}

// ToNow sets internal NPT values to current moment
func (t *NPT) ToNow() {
	*t = Now()
}

// FromTime sets internal NPT values to specified time value
func (t *NPT) FromTime(tm time.Time) {
	t.epoch = tm.Unix()
}

// ToNow sets internal NPT values to current moment
func (t *NPT) Add(d time.Duration) {
	gt := t.Time()
	gt = gt.Add(d)
	t.FromTime(gt)
}

// Time cunverts NPT into Time object
func (t *NPT) Time() time.Time {
	return time.Unix(t.epoch, 0)
}
