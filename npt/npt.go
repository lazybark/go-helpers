package npt

import "time"

//NPT is a time-holding object that does not have timezone pointer.
//It's more memory-effective than default Go time.Time and can be used in apps that store time
//in memory only and do not serialize it.
//
//For example: some action buffer that use time only for it's internal purposes of starting tasks
//at specific moment.
type NPT struct {
	sec  int64
	nsec uint32
}

//Now creates NPT from current moment in time
func Now() NPT {
	var t NPT
	t.sec = time.Now().Unix()
	t.nsec = uint32(time.Now().UnixNano())
	return t
}

//ToNow sets internal NPT values to current moment
func (t *NPT) ToNow() {
	*t = Now()
}

//FromTime sets internal NPT values to specified time value
func (t *NPT) FromTime(tm time.Time) {
	t.sec = tm.Unix()
	t.nsec = uint32(tm.UnixNano())
}

//ToNow sets internal NPT values to current moment
func (t *NPT) Add(d time.Duration) {
	gt := t.Time()
	gt.Add(d)
	t.FromTime(gt)
}

//Time cunverts NPT into Time object
func (t *NPT) Time() time.Time {
	return time.Unix(t.sec, int64(t.nsec))
}
