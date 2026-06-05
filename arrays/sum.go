package arrays

func Sum(numbers []int) int {

	//sum := 0
	//
	//n := len(numbers)
	//
	//for i := 0; i < n; i++ {
	//
	//	sum += numbers[i]
	//
	//}
	//
	//return sum

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum

}
