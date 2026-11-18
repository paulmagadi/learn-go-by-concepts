package main

import "fmt"

type Account struct {
    Owner   string
    Balance float64
}

func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}

func (a Account) BalanceInfo() string {
    return fmt.Sprintf("%s has a balance of %.2f", a.Owner, a.Balance)
}

//
type Number struct {
    n int
}

func (x Number) DoubleValue() {
    x.n = x.n * 2
}

func (x *Number) DoublePointer() {
    x.n = x.n * 2
}


func main() {
    acc := Account{Owner: "Alice", Balance: 100}
    acc.Deposit(50)
    acc.Deposit(25)
    fmt.Println(acc.BalanceInfo())


	//
	a := Number{n: 10}
    a.DoubleValue()
    fmt.Println("After DoubleValue:", a.n)

    b := Number{n: 10}
    b.DoublePointer()
    fmt.Println("After DoublePointer:", b.n)
}