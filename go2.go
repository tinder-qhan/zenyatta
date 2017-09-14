package main

import (
	"fmt"
	"strings"
	"bytes"
)

// Given a valide JSON string, write a prettyPrint functino to
// print it with proper indentation (2 spaces)
//
// e.g.
// prettyPrint(`{"car":{"made":"toyota","year":2013}}`) should output
//
// {
//   "car":{
//     "made":"toyota",
//     "year":2013
//   }
// }

const (	
	LCB = byte(123) // }
	RCB = byte(125) // }
	LSB = byte(91)  // [
	RSB = byte(93)  // ]
	CM  = byte(44)  // ,
)

var testCases = []string{
	`{"foo":"bar"}`,
	`[{"foo":"bar"},{"loo":123},{"buk":{"soo":10}}]`,
	`{"foo":"bar","bee":{"boo":123,"fee":{"baz":"bla"}},"bal":{"buz":"bya"}}`,
}

func main() {
	for i, t := range testCases {
		b := prettyPrint(t)
		fmt.Printf("test %d:\n%s\n\n", i + 1, b.String())
	}
}

func prettyPrint(str string) bytes.Buffer {
	var (
		b bytes.Buffer
		lev = 0
	)
	
	for _, c := range []byte(str) {
		if c == LCB || c == LSB {
			fmt.Fprint(&b, string(c) + "\n")
			lev++
			fmt.Fprint(&b, strings.Repeat("  ",  lev))
		} else if c == RCB || c == RSB {
			lev--
			fmt.Fprint(&b, "\n" + strings.Repeat("  ", lev) + string(c))
		} else if c == CM {
			fmt.Fprint(&b, ",\n" + strings.Repeat("  ", lev))
		} else {
			fmt.Fprint(&b, string(c))
		}
	}
	return b
}

