Create a new package called `messages`.

Inside it, define a function named `WelcomeMessage` that returns a custom greeting string.

Follow these steps:

1. Create a folder called `messages`.

2. Inside it, create a file called `text.go`.

3. Define an exported function named `WelcomeMessage` that takes a `name` and returns:

   ```
   "Welcome, <name>!"
   ```

4. Import the `messages` package into `main.go`.

5. Call `messages.WelcomeMessage()` and print the returned string.

Run your program with:

```
go run .
```