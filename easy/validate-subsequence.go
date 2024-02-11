package main

func IsValidSubsequence(array []int, sequence []int) bool {
	if len(sequence) > len(array) {
		return false
	}

	var pointer int
	for _, num := range array {
		if pointer == len(sequence) {
			break
		}
		
		if num == sequence[pointer] {
			pointer += 1
		}
	}
	
	return pointer == len(sequence)
}
