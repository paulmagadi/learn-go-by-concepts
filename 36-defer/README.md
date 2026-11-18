Write a function `saveReport()` that:

1. Prints `"Saving report"` at the start.
2. Defers two cleanup messages: `"Closing resources"` and `"Cleaning temporary data"`.
3. Prints `"Writing data"`.

Then call `saveReport()` inside `main()` and confirm the deferred messages run in reverse order.

Run it using:

```
go run main.go
```