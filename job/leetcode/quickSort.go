func quickSort(arr []int) []int {
	copy := append(int{}, arr...)

	var sort func([]int)
	sort = func(ori []int) {
		if len(ori) < 2 {
			return
		}
		var i, j int
		var loopj = true
		ref := ori[0]
		for i, j = 0, len(ori); i != j; {
			if loopj {
				if ref > ori[j] {
					ori[i] = ori[j]
					i++
					loopj = !loopj
				} else {
					j--
				}
			} else {
				if ref < ori[i] {
					ori[j] = ori[i]
					j--
					loopj = !loopj
				} else {
					i++
				}
			}
		}
		ori[i] = ref
		sort(ori[:i])
		sort(ori[i+1:])
	}
	sort(copy)
	return copy
}