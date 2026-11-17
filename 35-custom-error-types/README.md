Create a custom error type named `NotFoundError` with a field called `Resource`.

1. Implement the `Error() string` method so it returns a message describing which resource was not found.

2. Write a function `findUser(id int)` that always returns a `NotFoundError` for this exercise.

3. In `main()`, call `findUser`, check the error, and print it.

Run it using:

```
go run main.go
```