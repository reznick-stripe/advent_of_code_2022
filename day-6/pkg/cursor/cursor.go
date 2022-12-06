package cursor

func Scan(input string) (bool, int) {
	r := []rune(input)

LOOP_OVER_INPUT:
	for i := 0; i < len(r); i++ {
		lookback := input[i : 4+i]
		if len(lookback) != 4 {
			continue LOOP_OVER_INPUT
		}
		counts := make(map[rune]int)

		for _, r := range lookback {
			counts[r]++
			if counts[r] > 1 {
				continue LOOP_OVER_INPUT
			}
		}

		for _, v := range counts {
			if v > 1 {
				continue LOOP_OVER_INPUT
			}
		}
		return true, i + 4
	}

	return false, 0
}
