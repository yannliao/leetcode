package leetcode

func reverse(x int) int {
	const intMax = 2147483647
	const intMin = -2147483648
	var result int32
	if x > intMax || x < intMin {
		return 0
	}
	tmp := int32(x)
	for tmp != 0 {
		tail := tmp % 10
		tmpResult := result*10 + tail
		if (tmpResult-tail)/10 != result {
			return 0
		}
		result = tmpResult
		tmp = tmp / 10
	}
	return int(result)
}
