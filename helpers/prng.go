package helpers

func PRNG(input string) int64 {
	// PRNG But Not Really Random hehe :)
	// Generate random number base on input string
	var result int64
	for _, v := range input {
		result += int64(v)
	}

	return result
}