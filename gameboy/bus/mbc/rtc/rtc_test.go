package rtc

import (
	"fmt"
	"testing"
)

func TestBasicGet(t *testing.T) {
	clock := NewVirtualClock()
	rtc := NewRTC(clock)

	clock.forward(5, 8, 12, 2)
	if !assertClockEquals(5, 8, 12, 2, rtc) {
		t.Fatal("Not equal")
	}
}

func TestLatch(t *testing.T) {
	clock := NewVirtualClock()
	rtc := NewRTC(clock)

	clock.forward(5, 8, 12, 2)

	rtc.Latch()
	clock.forward(10, 5, 19, 4)
	if !assertClockEquals(5, 8, 12, 2, rtc) {
		t.Fatal("Not equal")
	}
	rtc.Unlatch()
	assertClockEquals(5+10, 8+5, 12+19, 2+4, rtc)
}

func TestCounterOverflow(t *testing.T) {
	clock := NewVirtualClock()
	rtc := NewRTC(clock)

	clock.forward(511, 23, 59, 59)
	if rtc.IsCounterOverflow() {
		t.Fatal("Should be false")
	}

	clock.forward(0, 0, 0, 1)
	if !assertClockEquals(0, 0, 0, 0, rtc) {
		t.Fatal("should be 0,0,0,0")
	}
	if !rtc.IsCounterOverflow() {
		t.Fatal("Should be true")
	}

	clock.forward(10, 5, 19, 4)
	if !assertClockEquals(10, 5, 19, 4, rtc) {
		t.Fatal("should be 10,5,19,4")
	}
	if !rtc.IsCounterOverflow() {
		t.Fatal("Should be true")
	}

	rtc.ClearCounterOverflow()
	if !assertClockEquals(10, 5, 19, 4, rtc) {
		t.Fatal("should be 10,5,19,4")
	}
	if rtc.IsCounterOverflow() {
		t.Fatal("Should be false")
	}
}

func TestSetClock(t *testing.T) {
	clock := NewVirtualClock()
	rtc := NewRTC(clock)

	clock.forward(10, 5, 19, 4)
	if !assertClockEquals(10, 5, 19, 4, rtc) {
		t.Fatal("should be 10,5,19,4")
	}

	rtc.SetHalt(true)
	if !rtc.IsHalt() {
		t.Fatal("halt should be true")
	}

	rtc.SetDays(10)
	rtc.SetHours(16)
	rtc.SetMinutes(21)
	rtc.SetSeconds(32)
	clock.forward(1, 1, 1, 1) // should be ignored after unhalt
	rtc.SetHalt(false)

	if rtc.IsHalt() {
		t.Fatal("halt should be false")
	}
	if !assertClockEquals(10, 16, 21, 32, rtc) {
		t.Fatal("should be 10,16,21,32")
	}

	clock.forward(2, 2, 2, 2)
	if !assertClockEquals(12, 18, 23, 34, rtc) {
		t.Fatal("should be 12,18,23,34")
	}
}

func assertClockEquals(days int64, hours int64, minutes int64, seconds int64, rtc *RTC) bool {
	if days != rtc.Days() {
		fmt.Printf("Days not equal. Got %v, want: %v", rtc.Days(), days)
		return false
	}
	if hours != rtc.Hours() {
		fmt.Printf("Hours not equal. Got %v, want: %v", rtc.Hours(), hours)
		return false
	}
	if minutes != rtc.Minutes() {
		fmt.Printf("Minutes not equal. Got %v, want: %v", rtc.Minutes(), minutes)
		return false
	}
	if seconds != rtc.Seconds() {
		fmt.Printf("Seconds not equal. Got %v, want: %v", rtc.Seconds(), seconds)
		return false
	}
	return true
}
