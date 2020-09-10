package leetcode

/** Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.
  * Return the quotient after dividing dividend by divisor.
  */
func divide(dividend, divisor int) int {
	res := 0

	absdd := dividend
	if dividend < 0 {
		absdd = -dividend
	}
	absdi := divisor
	if divisor < 0 {
		absdi = -divisor
	}

	for absdd >= absdi {
		shiftAbsDi := absdi
		shift := 1
		for absdd >= shiftAbsDi {
			shift = shift << 1
			shiftAbsDi = shiftAbsDi << 1
		}
		res += (shift >> 1)
		absdd -= (shiftAbsDi >> 1)
	}
	if (dividend <= 0 && divisor > 0) || (dividend >= 0 && divisor < 0) {
		res = -res
	}
	if res < -(1 << 31) || res > (1 << 31 -1) {
		res = 1 << 31 - 1
	}
	return res
}