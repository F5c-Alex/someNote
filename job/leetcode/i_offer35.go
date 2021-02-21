package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//使用哈希进行结点映射
func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	n := head
	for n != nil {
		node := new(Node)
		m[n] = node
		n = n.Next
	}
	n = head
	for n != nil {
		node := m[n]
		node.Val = n.Val
		node.Next = m[n.Next]
		node.Random = m[n.Random]
		n = n.Next
	}
	return m[head]
}

//解二
// 我们也可以不使用哈希表的额外空间来保存已经拷贝过的结点，而是将链表进行拓展，在每个链表结点的旁边拷贝，比如 A->B->C 变成 A->A'->B->B'->C->C'，然后将拷贝的结点分离出来变成 A->B->C和A'->B'->C'，最后返回 A'->B'->C'。

// 作者：z1m
// 链接：https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/solution/lian-biao-de-shen-kao-bei-by-z1m/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
