//普通解法：
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		k := num % 2
		if k == 1 {
			res++
		}
		num = num / 2
	}
	return res
}

//进阶解法
func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}
	var res int = 0
	for num > 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

// 作者：huanglian1219
// 链接：https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/solution/goti-jie-yi-wei-cao-zuo-by-huanglian1219/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

//进阶解法2
//利用n&n-1将n最右侧的1转为0
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		res++
		num &= num - 1
	}
	return res
}
