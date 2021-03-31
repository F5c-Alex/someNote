package main

import (
	"fmt"
	"strconv"
)

func quickSort(arr []int)[]int{
	copy := append([]int{},arr...)
	var sort func([]int)
	sort = func(arr []int){
		if len(arr) < 2{
			return
		}
		ref := arr[0]
		loopj := true
		var i,j int
		for i,j = 0,len(arr)-1;i!=j;{
			if loopj{
				if ref>arr[j]{
					arr[i]=arr[j]
					i++
					loopj=false
				}else{
					j--
				}
			}else{
				if ref<arr[i]{
					arr[j]=arr[i]
					j--
					loopj=true
				}else{
					i++
				}
			}
		}
		arr[i]=ref
		sort(arr[0:i])
		sort(arr[i+1:])
	}
	sort(copy)
	return copy
}

func main(){
	//arr := []int{2,1,5,2,1,4,6,8,10}

	//fmt.Println(quickSort(arr))
}