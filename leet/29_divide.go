package leet

func divide(dividend int, divisor int) int {
	MAX := 1<<31 - 1
	if divisor == 0 {
		return MAX
	}

	sign := 1
	if dividend < 0 {
		sign *= -1
		dividend *= -1
	}
	if divisor < 0 {
		sign *= -1
		divisor *= -1
	}
	res := 0
	tempDivisor := divisor
	mul := 1
	for dividend >= divisor {
		for dividend < tempDivisor && tempDivisor > divisor {
			tempDivisor >>= 1
			mul >>= 1
		}
		dividend -= tempDivisor
		res += mul
		tempDivisor <<= 1
		mul <<= 1
	}

	res *= sign
	if res > MAX {
		return MAX
	}

	return res
}
