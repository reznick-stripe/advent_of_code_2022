package cursor

import "errors"

func Scan(input string, length int) (int, error) {
	r := []rune(input)

	if length > len(r) {
		return 0, errors.New("length is longer than input")
	}

LOOP_OVER_INPUT:
	for i := 0; i < len(r)-length; i++ {
		lookback := input[i : length+i]
		if len(lookback) != length {
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
		return i + length, nil
	}

	return 0, errors.New("not found")
}
