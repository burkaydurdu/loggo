# Learn Go
![](./img/golang.svg)

**Go** is a statically typed, compiled programing language.

### Variables

``var``, ``const`` declares one or more variables. </br>
``var`` is mutable.
``const`` is immutable.

```go
var name string
var animal, school string

var (
 count, order int
 time float32
)

const message = "Hello"
const animal, school string

const (
 offset int = 12
)
```

### Private and Public

An identifier is exported if it begins with a **capital** letter </br>
For example: </br>

```go
 type Address struct {
  name string // private
  Nickname string // public
 }

 func GetPublicName() {} // public

 func setPrivateName() {} // private
```