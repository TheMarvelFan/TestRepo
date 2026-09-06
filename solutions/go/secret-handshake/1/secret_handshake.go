package secrethandshake

import (
    "fmt"
    "slices"
)

var actions = []string{"", "jump", "close your eyes", "double blink", "wink"}
const reverseBit = 0

func Handshake(code uint) []string {
	if code > 31 {
        return nil
    }

    ret := []string{}

    binary := fmt.Sprintf("%05b", code)
    rev := false

    for pos, bit := range binary {
        if bit - '0' == 1 {
            if pos > 0 {
                ret = append(ret, actions[pos])
            } else {
                rev = true
            }
        }
    }

    if !rev {
        slices.Reverse(ret)
    }

    return ret
}
