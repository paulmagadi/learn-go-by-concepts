Create a struct type called `Book` with these fields:

* `Title` (string)

* `Pages` (int)

Write a function called `addPages` that takes a pointer to a `Book` as its first parameter and an integer amount as its second parameter. The function should add the amount to the book's page count.

In `main()`, create a book, call `addPages`, and print the updated page count to confirm the original struct changed.

Run it using:

```
go run main.go
```