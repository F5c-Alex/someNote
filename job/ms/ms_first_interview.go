//普通二叉树转平衡二叉树
package main

import "fmt"

type  TreeNode struct {
	left *TreeNode
	right *TreeNode
	val int
}

type ListNode struct {
	next *ListNode
	val int
}
func treeToList(treeRoot *TreeNode)*ListNode{
	list := &ListNode{}
	p := list
	var dfs func(node*TreeNode)
	dfs = func(node *TreeNode){
		if node == nil{
			return
		}
		dfs(node.left)
		newListNode := &ListNode{
			next: nil,
			val:  node.val,
		}
		p.next = newListNode
		p = p.next
		dfs(node.right)
	}
	dfs(treeRoot)
	return list.next
}

func listToBST(listRoot *ListNode)*TreeNode{
	treeRoot := &TreeNode{}
	var div func(*ListNode)*TreeNode
	div = func(head *ListNode)*TreeNode{
		if head == nil{
			return nil
		}
		if head.next == nil{
			return &TreeNode{
				left:  nil,
				right: nil,
				val:   head.val,
			}
		}
		fastPointer := head
		slowPointer := head
		lastPointer := slowPointer
		for(fastPointer != nil && fastPointer.next != nil){
			fastPointer = fastPointer.next
			fastPointer = fastPointer.next
			lastPointer = slowPointer
			slowPointer = slowPointer.next
		}
		treeRoot := &TreeNode{
			left:  nil,
			right: nil,
			val:   slowPointer.val,
		}
		//fmt.Println("root",slowPointer.val)
		lastPointer.next = nil
		treeRoot.left = div(head)
		treeRoot.right = div(slowPointer.next)
		//fmt.Println("left",treeRoot.left.val)
		//fmt.Println("right",treeRoot.right.val)
		return treeRoot
	}
	treeRoot = div(listRoot)
	//fmt.Println("root",treeRoot.val)

	return treeRoot
}
func main(){
	treeNode1 := &TreeNode{
		left:  nil,
		right: nil,
		val:   1,
	}

	treeNode3 := &TreeNode{
		left:  nil,
		right: nil,
		val:   3,
	}
	treeNode2 := &TreeNode{
		left:  treeNode1,
		right: treeNode3,
		val:   2,
	}


	treeNode6 :=&TreeNode{
		left:  nil,
		right: nil,
		val:   6,
	}
	treeNode8 :=&TreeNode{
		left:  nil,
		right: nil,
		val:   8,
	}
	treeNode7 :=&TreeNode{
		left:  treeNode6,
		right: treeNode8,
		val:   7,
	}
	treeNode5 :=&TreeNode{
		left:  nil,
		right: treeNode7,
		val:   5,
	}

	treeNode4 :=&TreeNode{
		left:  treeNode2,
		right: treeNode5,
		val:   4,
	}


	l := treeToList(treeNode4)
	printList(l)
	fmt.Println("-----------")
	treeResult := listToBST(l)

	fmt.Println(treeResult.val)

	fmt.Println(treeResult.left.val)
	fmt.Println(treeResult.right.val)

	fmt.Println(treeResult.left.left.val)
	fmt.Println(treeResult.left.right.val)

	fmt.Println(treeResult.right.left.val)
	fmt.Println(treeResult.right.right.val)
}

func printList(list *ListNode){
	for list!=nil{
		fmt.Println(list.val)
		list = list.next
	}
}



