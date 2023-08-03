package npt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNPTFromTime(t *testing.T) {
	tm := time.Now()

	time.Sleep(time.Second)
	nt := Now()

	nt.FromTime(tm)

	assert.Equal(t, tm.Unix(), nt.Time().Unix()) //NPT has precison up to a second
}

func TestNPTAdd(t *testing.T) {
	tm := time.Now()

	nt := Now()
	nt.FromTime(tm)

	nt.Add(time.Minute)
	tm = tm.Add(time.Minute)
	assert.Equal(t, tm.Unix(), nt.Time().Unix()) //NPT has precison up to a second

	nt.Add(time.Minute * -3)
	tm = tm.Add(time.Minute * -3)
	assert.Equal(t, tm.Unix(), nt.Time().Unix()) //NPT has precison up to a second
}
