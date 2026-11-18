Create a type called `Account` with two fields:

* `Owner` (string)

* `Balance` (float64)

Write two methods:

1. `Deposit(amount float64)`

   * Uses a pointer receiver

   * Adds the amount to the balance

2. `BalanceInfo() string`

   * Must use a value receiver (not a pointer receiver)

   * Returns a string describing the account owner and balance

In `main()`, create an account, call `Deposit` a few times, then print the result of `BalanceInfo()`.

Run it using:

```
go run main.go
```