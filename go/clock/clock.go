package clock

import (
	"strconv"
)

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	clock := Clock{hour, minute}

	clock.hour = clock.hour % 24
	clock.FormatMinutesLessThanZero()
	clock.FormatMinutesGreaterThanSixty()
	clock.FormatHoursLessThanZero()
	clock.minute = clock.minute % 60

	return clock
}

func (clock Clock) Add(minutes int) Clock {
	clock.minute = clock.minute + minutes
	clock.FormatMinutesLessThanZero()
	clock.FormatMinutesGreaterThanSixty()
	clock.FormatHoursLessThanZero()
	clock.minute = clock.minute % 60

	return clock
}

func (clock *Clock) FormatMinutesGreaterThanSixty() {
	if clock.minute >= 60 {
		clock.hour = (clock.hour + clock.minute/60) % 24
	}
}

func (clock *Clock) FormatMinutesLessThanZero() {
	for clock.minute < 0 {
		clock.minute = clock.minute + 60
		clock.hour = clock.hour - 1
	}
}

func (clock *Clock) FormatHoursLessThanZero() {
	for clock.hour < 0 {
		clock.hour = clock.hour + 24
	}
}

func (clock Clock) String() string {
	hour := strconv.Itoa(clock.hour)
	minute := strconv.Itoa(clock.minute)

	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(minute) == 1 {
		minute = "0" + minute
	}

	return hour + ":" + minute
}
