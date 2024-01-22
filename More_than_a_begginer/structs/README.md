# Intro to structs
We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:
```go
type car struct {
  Make string
  Model string
  Height int
  Width int
}
```
Structs has a similar usecase to dictionaries.\
Now as we already know how to define a struct lets actually use this created variable:
```go
var myCar car
myCar.Make = //Enter your value
myCar.Model = //Enter your value 
myCar.Height = //Enter your value
myCar.Width = //Enter your value
```

# Nested Structs
```go
type car struct {
  Make string
  Model string
  Height int
  Width int
  FrontWheel Wheel
  BackWheel Wheel
}

type Wheel struct {
  Radius int
  Material string
}
```
How to use these vars:
```go
myCar := car{}
myCar.FrontWheel.Radius = 5
```