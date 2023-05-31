package numeric

import "strconv"

func GetShortestRepresentationByThousands(num int64) (string, int) {
	s1 := strconv.FormatInt(num, 10)
	counter := 0
	for i := len(s1) - 1; i >= 0; i -= 1 {
		if s1[i] == '0' {
			counter += 1
			continue
		}
		break
	}

	thousands := counter / 3
	zerosToRemove := thousands * 3

	shortNum := s1[:len(s1)-zerosToRemove]
	return shortNum, thousands
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
