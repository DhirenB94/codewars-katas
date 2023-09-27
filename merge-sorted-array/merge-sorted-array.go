package main

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
// To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored.
// nums2 has a length of n.

func merge(nums1 []int, m int, nums2 []int, n int) {
	endNum1Ints := m - 1
	endNum2Ints := n - 1 
	endNum1FullArray := m + n - 1

	for endNum1Ints >= 0 && endNum2Ints >= 0 {
		if nums2[endNum2Ints] > nums1[endNum1Ints] {
			nums1[endNum1FullArray] = nums2[endNum2Ints]
			endNum2Ints--
			endNum1FullArray--
		} else {
			nums1[endNum1FullArray] = nums1[endNum1Ints]
			endNum1Ints--
			endNum1FullArray--
		}
	}

	//If there are remaining elements in nums2, copy them into nums1
	//if m = 0 check no longer needed now
	for endNum2Ints >= 0 { //0
		nums1[endNum1FullArray] = nums2[endNum2Ints]
		endNum2Ints --
		endNum1FullArray --
	}
}

//Sudocode
//check the end on nums1 arrray in terms of ints that arent 0 against the end of nums 2 array
//if end nums 2 > end nums1 ints, end num 1 all = end num 2 AND end nums2 -1 end nums 1 full -1
//this will check 3 v 6 = 1, 2, 3, 0, 0, 6
// then 3 v 5 = 1, 2, 3, 0, 5, 6
// then 3 v 2 --> else add end num1 to end num full and miunus
