# Variables

## Basic variables
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

## Short Variable Declaration

```go
// Normal variable declaration
var anotherStr string = "Another random piece of sh**"

// short variable declaration
randomString := "This is a random string value"

```

## Creating and computed constants

```go
package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour// ?

	// don't edit below this line
	fmt.Println("number of seconds in an hour:", secondsInHour)
}

```