//微软三道笔试
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := []int{5,5,5,5,5,5}
// 	fmt.Println(solution(a))
// }

// func sumIsEven(a,b int) bool {
// 	if (a+b)%2 == 0 {
// 		return true
// 	}else{
// 		return false
// 	}
// }
// func solution(a []int)int{
// 	if len(a) == 0 || len(a) == 1 {
// 		return 0
// 	}
// 	var i int
// 	for i = 0; i < len(a)-1; i++ {
// 		if !sumIsEven(a[i],a[i+1]) {
// 			break
// 		}
// 	}
// 	if i == len(a)-1 {
// 		return int(len(a)/2)
// 	}
// 	i++
// 	b := append(a[i:],a[:i]...)
// 	j := 0
// 	res := 0
// 	for j < len(a)-1{
// 		if sumIsEven(b[j],b[j+1]){
// 			res++
// 			j+=2
// 		}else{
// 			j++
// 		}
// 	}
// 	return res
// }
package main

import (
	"fmt"
)

func main() {
	u :=2
	l:=2
	c :=[]int{2,1,1,0,1}
	k := Solution(u,l,c)
	fmt.Println(k)
}

func Solution(U int, L int, C []int) string {
	l := len(C)
	uSlice := make([]int, l)
	// lSlice := make([]int, l)
	uSum := U
	lSum := L
	for i:=0;i<l;i++{
		uSlice[i] = -1
		// lSlice[i] = -1
	}
	for j:=0;j<l;j++{
		if C[j] == 0{
			uSlice[j] = 0
			// lSlice[j] = 0
		}else if C[j] == 2{
			uSlice[j] = 1
			// lSlice[j] = 1
			uSum--
			lSum--
		}
	}
	for k:=0;k<l;k++{
		if uSlice[k] == -1{
			if uSum >0{
				uSlice[k] = 1
				uSum--
			}else if lSum >0{
				uSlice[k] = 0
				lSum--
			}else{
				fmt.Println("Error")
				return "IMPOSSIBLE"
			}
			fmt.Println(uSum,lSum,k)
		}
	}
	if uSum != 0 || lSum != 0{
		return "IMPOSSIBLE"
	}
	first := ""
	second := ""
	for k,v := range uSlice{
		first += fmt.Sprint(v)
		second += fmt.Sprint(C[k]-v)
	}

	return first + "," + second

}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := []int{1,1}
// 	fmt.Println(Solution(a))
// }
// func calSlope(a,b int) (slope int){
// 	k := b - a 
// 	if k == 0{
// 		slope = 0
// 	}else if k>0{
// 		slope = 1
// 	}else{
// 		slope = -1
// 	}
// 	return slope
// }
// func Solution(blocks []int) int {
// 	l := len(blocks)
// 	if l == 0{
// 		return 0
// 	}
// 	if l == 1{
// 		return 1
// 	}
	
// 	record := []int{0}
// 	slope := calSlope(blocks[0],blocks[1])
	

// 	for i:=1;i<len(blocks)-1;i++{
// 		curSlope := calSlope(blocks[i],blocks[i+1])
// 		if curSlope != slope{
// 			slope = curSlope
// 			record = append(record,i)
// 		}
// 	}
// 	record = append(record,l-1)

// 	res := 0
// 	//calculate the first point
// 	// cur := record[0]
// 	// if blocks[cur] < blocks[cur++]{

// 	// }
// 	fmt.Println(record)
// 	for j := 0;j <len(record);j++{
// 		x := j
// 		for x+1 <len(record) && blocks[record[x+1]] >= blocks[record[x]]{
// 			x++
// 		}
// 		y := j
// 		for y-1>=0 && blocks[record[y-1]] >= blocks[record[y]]{
// 			y--
// 		}
		
// 		if res < record[x] - record[y] +1{
// 			fmt.Println(x,",",y)
// 			res = record[x] - record[y] +1
// 		}
// 	}
// 	return res

// }