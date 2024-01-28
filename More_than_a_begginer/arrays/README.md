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

# Taking input as arrays

Using `fmt.Scanf` for fixed array size: Input is taken using loops

```go
var numbers [5]int

fmt.Println("Enter 5 numbers:")

for i := 0; i < 5; i++ {
    fmt.Scanf("%d", &numbers[i]) // Read each element into the array
}

fmt.Println("You entered:", numbers)
```

In GO we can't create arrays of variable length due to reasons like memory efficiency and type safety.\
But in GO we have slices as a very powerful tool.\
Slices to Rescue:
- Don't own their memory, and reference a section of an underlying array. 
- While it may seem like you're creating slices independently, they always have an underlying array, even if you don't explicitly declare one.
- Can grow or shrink as needed using the **append** function.