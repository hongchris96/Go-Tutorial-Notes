package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) //[[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [0,0,0]
}

func threeSum(nums []int) [][]int {
	uniqueTriplets := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		sets := twoSum(nums[i+1:], -nums[i])
		for _, set := range sets {
			duplicate := false
			for _, triplet := range uniqueTriplets {
				if arrEqual(triplet, set) {
					duplicate = true
					break
				}
			}
			if !duplicate {
				uniqueTriplets = append(uniqueTriplets, set)
			}
		}
	}
	return uniqueTriplets
}

func twoSum(nums []int, target int) [][]int {
	triplets := [][]int{}
	numLocation := map[int]int{}
	for i, num := range nums {
		if _, ok := numLocation[target-num]; ok {
			triplet := []int{-target, target - num, num}
			sort.Ints(triplet)
			triplets = append(triplets, triplet)
		}
		numLocation[num] = i
	}
	return triplets
}

func arrEqual(nums1, nums2 []int) bool {
	for i := 0; i < 3; i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
