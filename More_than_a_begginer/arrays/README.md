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

## Using `fmt.Scanf` for fixed array size: Input is taken using loops

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

## Using `fmt.Scanf()` for variable length array - slices

```go
var numbers []int

for {
    var num int
    fmt.Print("Enter a number (or 0 to stop): ")
    fmt.Scanf("%d", &num)

    if num == 0 {
        break
    }

    numbers = append(numbers, num) // Append to the slice
}

fmt.Println("You entered:", numbers) // Checking your input
```

# What is len and cap?

The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself.\
The capacity of a slice, accessible by the built-in function `cap()`, reports the maximum length the slice may assume.

Here is a function to append data to a slice. If the data exceeds the capacity, the slice is reallocated.
```go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
```

# Variadic

Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the "..." syntax in the function signature.

A variadic function receives the variadic arguments as a slice.
```go
func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for _, str := range strs {
        final += str
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}
```