# What packages are we working with?

## `fmt`
Provides functions for formatted input and output operations like:

- `fmt.Println()`: printing values to console, followed by a newline
- `fmt.Printf()`: printing formatted outputs, this gives us more
control over output layout
    ```go
    name := "John"
    fmt.Printf("Hello, %s!\n", name)
    ```
- `fmt.Scan()` and `fmt.Scanf()`: taking input from console
- `fmt.Sprintf()`: this is one of the strange ones, this is used for formating strings without printing then, this maybe useful in creating formatted strings.
    ```go
    number := 10
    message := fmt.Sprintf("%d is a magic number", number)

    //This will create a string "10 is a magic number" without printing it.
    ```

## `log`
