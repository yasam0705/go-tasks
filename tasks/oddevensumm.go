package tasks

func OddEven(arr []int) (int, int) {
	var oddSum, evenSum int

	for _, v := range arr {
		if v%2 == 0 {
			evenSum += v
		} else {
			oddSum += v
		}
	}
	return oddSum, evenSum
}
