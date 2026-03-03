package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
    Hour int
    Minute int
}

func New(h, m int) Clock {
	c := Clock {
        Hour: h,
        Minute: m,
    }

    c.Hour %= 24

    for c.Hour < 0 {
        c.Hour += 24
    }

    for c.Minute < 0 {
        c.Minute += 60
        c.Hour -= 1

        if (c.Hour == -1) {
            c.Hour = 23
        }
    }

    for c.Minute >= 60 {
        c.Minute -= 60
        c.Hour = (c.Hour + 1) % 24
    }

    return c
}

func (c Clock) Add(m int) Clock {
	c.Minute += m

    for c.Minute >= 60 {
        c.Minute -= 60
        c.Hour = (c.Hour + 1) % 24
    }

    return c
}

func (c Clock) Subtract(m int) Clock {
	c.Minute -= m

    for c.Minute < 0 {
        c.Minute += 60
        c.Hour -= 1

        if (c.Hour == -1) {
            c.Hour = 23
        }
    }

    return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
