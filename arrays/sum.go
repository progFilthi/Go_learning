package arrays

func Sum(nums []int) int {

	sum := 0

	//n := len(nums)
	//
	//for i := 0; i < n; i++ {
	//	//sum = sum + nums[i]
	//	sum += nums[i]
	//
	//}

	for _, num := range nums {
		sum += num
	}

	return sum
}
