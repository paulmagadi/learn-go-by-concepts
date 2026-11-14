1. Create variables for a product:
   * Name: `"Notebook"`
   * Price: `9.99`

2. Use `fmt.Printf` to print: `Notebook costs $9.99`

3. Use `fmt.Sprintf` to build the same string and store it in a variable.

4. Print the result using `fmt.Println`.


## Common Format Verbs

| Verb  | Description               | Example                        |
| ----- | ------------------------- | ------------------------------ |
| `%s`  | String                    | `"Go"` → `Go`                  |
| `%d`  | Integer (decimal)         | `42` → `42`                    |
| `%f`  | Float (decimal)          | `3.14` → `3.140000`            |
| `%v`  | Default format (any type) | `true` → `true`                |
| `%T`  | Print type of value       | `"text"` → `string`            |
| `%+v` | Struct with field names   | `{Name: "Alice"}` → `Name:Alice` |