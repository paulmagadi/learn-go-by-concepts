Write a function called `squareRoot` that takes an integer and returns:

1. The square root (use `float64` for the result)
2. An error if the input is negative

Hint: You'll need to import the `math` package and use `math.Sqrt()` to calculate the square root. You'll also need to convert the integer input to `float64` before passing it to `math.Sqrt()`.

In `main()`, call the function with both valid and invalid inputs to test your code. Print the error when one occurs. Print the result when everything is fine.

Run it using:
```
go run main.go
```