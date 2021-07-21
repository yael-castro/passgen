package password

func GenerateWordSet(length int, letters []rune) (words [][]rune) {
	counters := make([]int, length)
	maxValue := len(letters) - 1

	for i := range counters {
		counters[i] = maxValue
	}

	for counters[0] > 0 {
		word := make([]rune, length)

		for i := range counters {
			index := counters[i]
			word[i] = letters[index]
		}

		words = append(words, word)

		index := length - 1
		counters[index]--

		for index != 0 && counters[index] == -1 {
			counters[index] = maxValue
			index--
			counters[index]--
		}
	}

	return
}
