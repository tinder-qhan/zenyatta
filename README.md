### Practice Go

Question 1 (go1.go)

```
Given two strings of numbers contains only 4's and 7's,
write a function to return the minimum operations needed using only
swap and replace operations to make the 2nd string same as 1st one.

Example 1:
str1 := "4747"
str2 := "7474"
minOps(str1, str2) => "2 swaps, 0 replaces"
(this can also be 4 replaces or 1 swap 2 replaces, but neither is optimal)

Example 2
str1 := "44444"
str2 := "77777"
minOps(str1, str2) => "0 swaps, 5 replaces"

Example 3
str1 := "44447"
str2 := "77774"
minOps(str1, str2) => "1 swaps, 3 replaces"
```

Question 2 (go2.go)

```
Given a valide JSON string, write a prettyPrint function to
print it with proper indentation (2 spaces)

e.g.
prettyPrint(`{"car":{"made":"toyota","year":2013}}`) =>

{
  "car":{
    "made":"toyota",
    "year":2013
  }
}
```