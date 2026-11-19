Write a program that:

1. Logs a startup message using log.Println.
2. Calls a function called loadUser that always returns an error.
3. If the error is not nil, use log.Fatal to log the error and exit.
4. Otherwise, log a success message.

Run it using:

```bash
go run main.go
```