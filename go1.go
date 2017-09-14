package main

import (
	"fmt"
	"math"
)

// Given two equal length string of numbers contains only 4's and 7's
// write a function to return the minimum operations needed using only
// swap and replace operations to make the 2nd string same as 1st one.
// 
// Example 1:
// str1 := "4747"
// str2 := "7474"
// minOps(str1, str2) would print "2 swaps, 0 replaces"
// (this can also be 4 replaces or 1 swap 2 replaces, but neither is optimal)
//
// Example 2
// str1 := "44444"
// str2 := "77777"
// minOps(str1, str2) would print "0 swaps, 5 replaces"
//
// Example 3
// str1 := "44447"
// str2 := "77774"
// minOps(str1, str2) would print "1 swaps, 3 replaces"

func main() {
	testCases := []map[string]string{
		{
			"str1": "4747774747",
			"str2": "7474474477",
		},
		{
			"str1": "7774747",
			"str2": "4474477",
		},
		{
			"str1": "4747",
			"str2": "4477",
		},
		{
			"str1": "44444",
			"str2": "77777",
		},
		{
			"str1": "4",
			"str2": "7",
		},
		{
			"str1": "",
			"str2": "",
		},
	}
	
	for _, c := range testCases {
		minReplaces, minSwaps := byMath(c["str1"], c["str2"])
		minReplaces2, minSwaps2 := byBrute(c["str1"], c["str2"])
	
		fmt.Printf("Use math: min replaces = %d, min swaps = %d\n", minReplaces, minSwaps)
		fmt.Printf("Use brute: min replaces = %d, min swaps = %d\n\n", minReplaces2, minSwaps2)
	}
}

func toNum(num rune) int {
	return int(num - '0')
}

// method 1 => brute force
func byBrute(str1, str2 string) (replaces, swaps int) {
	if len(str1) != len(str2) {
		panic("Must take equal length strings as input.")
	}
	b1, b2 := []byte(str1), []byte(str2)
	i := 0
	for i < len(b1) {
		c1 := b1[i]
		c2 := b2[i]
				
		if c1 != c2 {
			// exausted the list, still couldn't find the swaps, do a replace
			if i == len(b1) - 1 {
				replaces++
			}
			
			// use replace unless no swap avaiable
			k := i + 1
			done := false
			for k < len(b1) && !done {
				if b1[k] == c2 {
					if  b1[k] != b2[k] {
						b1[k] = c1
						swaps++
					} else {
						replaces++
					}
					done = true
				} else if k == len(b1) - 1 {
					replaces++
				}
				k++
			}
		}
		i++
	}
	return replaces, swaps
}

// method 2 => math
func byMath(str1, str2 string) (replaces, swaps int) {
	if len(str1) != len(str2) {
		panic("Must take equal length strings as input.")
	}
	
	var diffSum, diffCount int
	for i, c := range str1 {
		n1 := toNum(c)
		n2 := toNum(rune(str2[i]))
		
		if n1 != n2 {
			diffCount += 1
		}
		diffSum += n1 - n2
	}
	replaces = int(math.Abs(float64(diffSum / 3)))
	swaps = (diffCount - replaces) / 2
	return replaces, swaps
}
