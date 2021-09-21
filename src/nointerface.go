package main

import "fmt"

type dog struct{}

type fish struct{}

type bird struct{}

func (dog) caminar() string {
	return "I am a dog and I walk"
}

func (fish) caminar() string {
	return "I am a fish and I swim"
}

func (bird) caminar() string {
	return "I am a bird and I fly"
}

func moveDog(d dog) {
	fmt.Println(d.caminar())
}

func moveFish(f fish) {
	fmt.Println(f.caminar())
}

func moveBird(b bird) {
	fmt.Println(b.caminar())
}

func main() {
	d := dog{}
	moveDog(d)
	f := fish{}
	moveFish(f)
	b := bird{}
	moveBird(b)
}
