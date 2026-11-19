Create a Go file named **messy.go** with intentionally inconsistent formatting. Include at least a function and a print statement.

Then run:

```bash
gofmt -w messy.go
```
Open the file again and observe how the tool has cleaned up indentation, spacing, and layout.

Extra credit: If you want to go further, run:

```bash
go fmt
```
inside your module to format every file.