package main

import (
	"fmt"
	"strings"
)

/*

Valid Anagram
=============

Link: https://leetcode.com/problems/valid-anagram/
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.


Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

*/

func main() {

	result := isAnagram("anagram", "nagaram")
	fmt.Println("result", result)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArray := []rune(s)
	for i := 0; i < len(sArray); i++ {
		t = strings.Replace(t, fmt.Sprintf("%c", sArray[i]), "", 1)
	}
	return t == ""

}
