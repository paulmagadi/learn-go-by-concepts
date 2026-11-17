Create a small logging system using an interface.

1. Define an interface called `Logger` with one method:

   ```
   Log(message string)
   ```

2. Create two types that satisfy this interface:

   * `ConsoleLogger` which prints the message to the terminal.

   * `PrefixedLogger` which also prints to the terminal but adds a fixed prefix, for example `[INFO]`.

3. Write a function called `runLog` that accepts a `Logger` and a message, then calls the `Log` method.

4. In `main()`, create both loggers and pass them to `runLog` with different messages.

Run it using:

```
go run main.go
```