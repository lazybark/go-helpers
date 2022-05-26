//LazyEvent is a simple event log package that can work with files & CLI simultaneously.
//It unifies event struct and helps in creating event log & event map for any app.
//It's easy to add custom event processing methods and wrap LE-events around.
//
//LE can prove useful in apps that require fast but easy-readable logging or event stacking.
package le

import (
	v1 "github.com/lazybark/go-helpers/events/le/v1"
)

//NewV1 returns pointer to new EventProcessor of v1
func NewV1(chainLength int) *v1.EvProc {
	return v1.New(chainLength)
}
