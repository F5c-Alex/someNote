package leetcode

//分组异或
func singleNumbers(nums []int) []int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	bit := 1
	for bit&a == 0 {
		bit = bit << 1
	}
	b, c := 0, 0
	for _, v := range nums {
		if bit&v == 0 {
			b ^= v
		} else {
			c ^= v
		}
	}
	return []int{b, c}
}
