package arraysslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(arraysToSum ...[]int) []int {
	lenOfNumbers := len(arraysToSum)
	sums := make([]int, lenOfNumbers)

	for _, arr := range arraysToSum {
		sums = append(sums, Sum(arr))
	}
	return sums
}

func SumAllTails(arraysToSum ...[]int) []int {
	var tailsums []int

	for _, arr := range arraysToSum {
		if len(arr) == 0 {
			tailsums = append(tailsums, 0)
		} else {
			tail := arr[1:]
			tailsums = append(tailsums, Sum(tail))
		}
	}
	return tailsums
}
