package main

import "fmt"

/*

Contains Duplicate
==================

Link: https://leetcode.com/problems/contains-duplicate/
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:
Input: nums = [1,2,3,1]
Output: true

Example 2:
Input: nums = [1,2,3,4]
Output: false

Example 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

*/

func main() {
	output := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	fmt.Println("output", output)
}

func containsDuplicate(nums []int) bool {

	evaluatedNumbers := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, isPresent := evaluatedNumbers[nums[i]]
		if isPresent {
			return true
		}
		evaluatedNumbers[nums[i]] = nums[i]
	}
	return false
}
