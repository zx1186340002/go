// package main

// import "fmt"

// type Animal interface {
// 	eat()
// 	sleep()
// }

// type Dog interface {
// 	Animal
// 	goOut()
// 	eat()
// }
// type balabala struct {
// }
// type Fly struct{}

// func (balabala) eat() {
// 	fmt.Println("吃了")
// }
// func (balabala) sleep() {
// 	fmt.Println("吃了12")
// }
// func (Fly) eat() {
// 	fmt.Println("吃了1234")
// }

// func (Fly) sleep() {
// 	fmt.Println("吃4了")
// }
// func (Fly) goOut() {
// 	fmt.Println("as")
// }

// func main() {
// 	var A Animal = balabala{}
// 	var B Animal = Fly{}
// 	A.eat()
// 	B.eat()

// }
