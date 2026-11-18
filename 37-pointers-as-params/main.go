package main

import "fmt"

func addTen(x *int) {
	*x = *x + 10
}

//
type Book struct {
	Title string
	Pages int
}

func addPages(b *Book, amount int) {
	b.Pages += amount
}

//

type User struct {
	Name string
	Score int
}

func increaseScoreCopy(u User) {
	u.Score++
}

func increaseScorePtr(u *User) {
	u.Score++
}

func main() {
	n := 5
	addTen(&n)
	fmt.Println(n)

	//
	novel := Book{Title: "The Explorer", Pages: 200}
	addPages(&novel, 50)
	fmt.Println(novel.Pages)

	//
	userA := User{Name: "Jane", Score: 10}
	userB := User{Name: "Bob", Score: 10}

	increaseScoreCopy(userA)
	increaseScorePtr(&userB)

	fmt.Println(userA.Score)
	fmt.Println(userB.Score)
}