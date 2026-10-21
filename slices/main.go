package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic array
// array + useful methods

func main() {
	// unitialized slice is nil
	var nums []int 
	fmt.Println("nums:", nums) // []
	fmt.Println("nums is nil:", nums == nil) // true
	fmt.Println("length of nums:", len(nums)) // 0
	var nums2 = make([]int,5) // [0 0 0 0 0]
	fmt.Println("nums2:", nums2)
	fmt.Println(cap(nums2)) // 5 // -> maximum no of elements can fit
	fmt.Println("length of nums2:", len(nums2)) // 5
	var nums3 = make([]int, 2,5) // [0 0 0 0 0]
	fmt.Println("nums3:", nums3)
	fmt.Println(cap(nums3)) // 5 // -> maximum no of elements can fit
	fmt.Println("length of nums3:", len(nums3)) // 2
	nums = append(nums, 1) // [1]
	fmt.Println("nums:", nums)
	fmt.Println("length of nums:", len(nums)) // 1
	fmt.Println(cap(nums)) // 1 // -> maximum no of elements can fit

	// double the capacity when the slice is full
	nums2 = append(nums2, 1) // [1]
	fmt.Println("nums2:", nums2)
	fmt.Println("length of nums2:", len(nums2)) // 5
	fmt.Println(cap(nums2)) // 5 // -> maximum no of elements can fit

	 nums4:=[]int{}
	 nums4 = append(nums4, 1) // [1]
	 fmt.Println("nums4:", nums4)
	 fmt.Println("length of nums4:", len(nums4)) // 1
	 nums4[0] = 2 // [2]

	 // copy slice
	 var nums5 = make([]int, 2, 5) // [0 0]	
	 nums5 = make([]int,len(nums4))
	 copy(nums5, nums4)
	 fmt.Println("nums5:", nums5) // [2 0]
	 // agar length of nums5 is less than nums4, then only copy the elements which can fit in nums5
	 // agar length = 0 nothing will be copied

	 // slice operator
	 nums = []int{1,2,3}
	 fmt.Println("nums[0:2]:",nums[0:2]) // exclude value at index 2
	 fmt.Println("nums[1:3]:",nums[1:3]) // exclude value at index 3

	 // slice
	 var nums6 = []int{2,3,4}
	 var nums7 = []int{2,3,4} // [2,3,4]
	 fmt.Println(slices.Equal(nums6,nums7))

	 var nums2d = [][]int{
		 {1, 2, 3},
		 {4, 5, 6},
	 }
	 fmt.Println("nums2d:", nums2d)

}