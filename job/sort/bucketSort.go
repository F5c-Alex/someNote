package main

import "fmt"
func bucket_sort(arr []int){
	if len(arr) <= 1{
		return
	}
	var max_num = arr[0]
	for _,v := range arr{
		if max_num < v{
			max_num = v
		} 
	}
	max_num++
	bucket := make([]int,max_num)
	for _,v := range arr{
		bucket[v]++
	}
	j := 0
	for k,v := range bucket{
		for v != 0{
			arr[j] = k
			j++
			v--
		}
	}
	fmt.Println(arr)
}

func main(){
	arr := []int{92,3,1,4,123,51,3,45}
	bucket_sort(arr)
}