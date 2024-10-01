package arrays

func Sum(numbers []int) int {

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(collections ...[]int) []int {
	var total []int

	for i := 0; i < len(collections); i++ {
		total = append(total, Sum(collections[i]))
	}

	return total
}
func SumAllTails(collections ...[]int) []int {
	var total []int

	for _, numbers := range collections {
		if len(numbers) == 0 {
			continue
		}
		tail := numbers[1:]
		total = append(total, Sum(tail))
	}

	return total
}
