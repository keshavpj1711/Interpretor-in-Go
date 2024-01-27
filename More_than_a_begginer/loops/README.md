# Introduction to loops
In go we only have a for loop and not while loop. 
Syntax for writing loops is similar to that of C-like.
For example:
```go
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Initializing; Condition; After

// Prints 0 through 9
```

Now one problem which arises with while loop not present is sometimes,
we need to run a loop infinitely in that case what can we do.
In that case we can take the use of **omitting conditions** and **break** as well as **continue** statements
