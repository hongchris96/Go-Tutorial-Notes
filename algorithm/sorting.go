package main

import "fmt"

func main() {
	// Merge sort
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	sorted := mergeSort(unsorted) // [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	fmt.Println(sorted)

	// Quick sort
	qsorted := quickSort(unsorted)
	fmt.Println(qsorted)

	// Bucket sort
	bucketSorted := bucketSort([]int{328, 834, 329, 983, 249, 832, 564, 895, 329, 493, 721, 980, 469, 890, 342})
	fmt.Println(bucketSorted)
	// Heap sort
	// Counting sort
	// Selection sort
}

func mergeSort(list []int) []int {
	// base case
	if len(list) <= 1 {
		return list
	}

	// recursive case
	midIdx := len(list) / 2
	leftHalf := append([]int{}, list[:midIdx]...)
	rightHalf := append([]int{}, list[midIdx:]...)

	return mergeHelper(mergeSort(leftHalf), mergeSort(rightHalf))
}

func mergeHelper(list1, list2 []int) []int {
	newList := []int{}

	for len(list1) > 0 && len(list2) > 0 {
		if list1[0] <= list2[0] {
			newList = append(newList, list1[0])
			list1 = list1[1:]
		} else {
			newList = append(newList, list2[0])
			list2 = list2[1:]
		}
	}

	if len(list1) > 0 {
		newList = append(newList, list1...)
	} else if len(list2) > 0 {
		newList = append(newList, list2...)
	}

	return newList
}

func quickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	pivot := list[0]
	listLength := len(list)
	smaller := filter(list[1:listLength], func(num int) bool { return num <= pivot })
	greater := filter(list[1:listLength], func(num int) bool { return num >= pivot })

	newSlice := append(quickSort(smaller), pivot)
	return append(newSlice, quickSort(greater)...)
}

func filter(list []int, condition func(num int) bool) []int {
	filteredList := []int{}
	for _, num := range list {
		if condition(num) {
			filteredList = append(filteredList, num)
		}
	}
	return filteredList
}

func bucketSort(nums []int) []int {
	answer := []int{}
	buckets := [10][]int{}
	max := nums[0]
	maxDig := 1
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	max /= 10
	for max > 0 {
		maxDig *= 10
		max /= 10
	}

	for _, num := range nums {
		targetBucket := num / maxDig
		buckets[targetBucket] = append(buckets[targetBucket], num)
	}
	for _, bucket := range buckets {
		quickSort(bucket)
		answer = append(answer, bucket...)
	}
	return answer
}
