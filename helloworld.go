package main

import "fmt"
func main(){
	var i int
	fmt.Println("Hello World !!!")

	arr := [6]int{2,9,4,6,3,4}
	fmt.Printf("Arr = %v \n",arr)
	var s1[] int = arr[1:3]
	fmt.Printf("Slices S1 = %v \n",s1)
	// creating slices

	s2 := make([]int,6)
	for i=0;i<len(s2);i++{
		s2[i] = i
	}
	s2 = append(s2,6)
    fmt.Println(len(s2))
	fmt.Printf("Slices S2 = %v \n",s2)

	s3 := []int{}
	s3 = append(s3,1)
	fmt.Printf("Slices S3 = %v \n",s3)
}