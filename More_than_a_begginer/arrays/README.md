# Arrays in GO
Arrays are fixed-size groups of variables of the same type.
The type `[n]T` is an array of n values of type T

For Example:
```go
// To declare an array of 10 integers
var myInts [10]int

// declaring an initialized literal
primes := [6]int{2, 3, 5, 7, 11, 13}
```

# Slices in GO
99 times out of 100 you will use a slice instead of an array when working with ordered lists. Arrays are fixed in size. Once you make an array like `[10]int` you can't add an 11th element.

Syntax for using a sliced array:
```go
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}

// The Syntax
// arrayname[lowIndex:highIndex]
// arrayname[lowIndex:]
// arrayname[:highIndex]
// arrayname[:]
```