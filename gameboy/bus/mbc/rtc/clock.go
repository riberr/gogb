package rtc

import "time"

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func NewRealClock() *RealClock {
	return &RealClock{}
}

func (rc *RealClock) Now() time.Time {
	return time.Now()
}

type VirtualClock struct {
	clock time.Time
}

func NewVirtualClock() *VirtualClock {
	return &VirtualClock{time.Now()}
}

func (vc *VirtualClock) Now() time.Time {
	return vc.clock
}

func (vc *VirtualClock) forward(days int, hours int, minutes int, seconds int) {
	vc.clock = vc.clock.Add(
		time.Duration(days*24)*time.Hour +
			time.Duration(hours)*time.Hour +
			time.Duration(minutes)*time.Minute +
			time.Duration(seconds)*time.Second,
	)
}
