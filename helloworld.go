package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbConnection() {
	fmt.Println("Drivers :", sql.Drivers())
	fmt.Println(" ******** DB connection process begins ********* ")
	db, err := sql.Open("mysql", "somnath@somnath@tcp(localhost:3306)/booking_driver_flow")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println(" ******** DB connection was successfully done ********* ")
}

func main() {

	dbConnection()
	//	helloWorld()
}

func helloWorld() {
	var i int
	fmt.Println("Hello World !!!")

	arr := [6]int{2, 9, 4, 6, 3, 4}
	fmt.Printf("Arr = %v \n", arr)
	var s1 []int = arr[1:3]
	fmt.Printf("Slices S1 = %v \n", s1)
	// creating slices

	s2 := make([]int, 6)
	for i = 0; i < len(s2); i++ {
		s2[i] = i
	}
	s2 = append(s2, 6)
	fmt.Println(len(s2))
	fmt.Printf("Slices S2 = %v \n", s2)

	s3 := []int{}
	s3 = append(s3, 1)
	fmt.Printf("Slices S3 = %v \n", s3)

	m1 := map[string]string{"name": "somnath", "email": "gmail"}
	for k, element := range m1 {
		fmt.Printf("Key : %v , Value : %v \n", k, element)
	}
	fmt.Printf("Map m1 %v", m1)
	delete(m1, "name")
	fmt.Printf("Map m1 %v", m1)
}
