// package main

// import (
// 	"errors"
// 	"fmt"
// 	"reflect"
// )

// func abc(int) (string, error) {
// 	return "", errors.New("报错了")
// }

// func main() {
// 	fmt.Println("nihao ")
// 	defer func() {
// 		fmt.Println(reflect.TypeOf(recover()))
// 		ab := 10
// 		fmt.Println(ab)
// 	}()
// 	defer func() {
// 		fmt.Println(reflect.TypeOf(recover()))
// 		ab := 100000
// 		fmt.Println(ab)
// 		panic("wqwq")
// 	}()

// 	panic("212")
// 	a := 1
// 	fmt.Println(a)
// }
