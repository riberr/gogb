package rtc

import (
	"fmt"
	"time"
)

type RTC struct {
	clock      Clock
	clockStart time.Time
	latchStart time.Time

	offsetSec int64

	halt        bool
	haltSeconds int64
	haltMinutes int64
	haltHours   int64
	haltDays    int64
}

func NewRTC(clock Clock) *RTC {
	return &RTC{
		clock:      clock,
		clockStart: clock.Now(),
	}
}

func (rtc *RTC) Latch() {
	rtc.latchStart = rtc.clock.Now() //time.Now()
}

func (rtc *RTC) Unlatch() {
	rtc.latchStart = time.Time{}
}

func (rtc *RTC) Seconds() int64 {
	return rtc.timeInSeconds() % 60
}
func (rtc *RTC) Minutes() int64 {
	return (rtc.timeInSeconds() % (60 * 60)) / 60
}
func (rtc *RTC) Hours() int64 {
	return (rtc.timeInSeconds() % (60 * 60 * 24)) / (60 * 60)
}
func (rtc *RTC) Days() int64 {
	return rtc.timeInSeconds() % (60 * 60 * 24 * 512) / (60 * 60 * 24)
}

func (rtc *RTC) IsHalt() bool {
	fmt.Printf("halt should be %v\n", rtc.halt)
	return rtc.halt
}

func (rtc *RTC) IsCounterOverflow() bool {
	return rtc.timeInSeconds() >= 60*60*24*512
}

func (rtc *RTC) SetSeconds(seconds int64) {
	if !rtc.halt {
		return
	}
	rtc.haltSeconds = seconds
}
func (rtc *RTC) SetMinutes(minutes int64) {
	if !rtc.halt {
		return
	}
	rtc.haltMinutes = minutes
}
func (rtc *RTC) SetHours(hours int64) {
	if !rtc.halt {
		return
	}
	rtc.haltHours = hours
}
func (rtc *RTC) SetDays(days int64) {
	if !rtc.halt {
		return
	}
	rtc.haltDays = days
}

func (rtc *RTC) SetHalt(halt bool) {
	if halt && !rtc.halt {
		rtc.Latch()
		rtc.haltSeconds = rtc.Seconds()
		rtc.haltMinutes = rtc.Minutes()
		rtc.haltHours = rtc.Hours()
		rtc.haltDays = rtc.Days()
		rtc.Unlatch()
	} else if !halt && rtc.halt {
		rtc.offsetSec = rtc.haltSeconds + rtc.haltMinutes*60 + rtc.haltHours*60*60 + rtc.haltDays*60*60*24
		rtc.clockStart = rtc.clock.Now()
	}
	rtc.halt = halt
	fmt.Printf("halt is now %v\n", rtc.halt)
}

func (rtc *RTC) ClearCounterOverflow() {
	for rtc.IsCounterOverflow() {
		rtc.offsetSec -= 60 * 60 * 24 * 512
	}
}

func (rtc *RTC) timeInSeconds() int64 {
	var now time.Time
	if rtc.latchStart.IsZero() {
		now = rtc.clock.Now()
	} else {
		now = rtc.latchStart
	}
	return (now.Sub(rtc.clockStart).Milliseconds() / 1000) + rtc.offsetSec
}
