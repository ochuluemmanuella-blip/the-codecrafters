## Go Variables

Variables are containers for storing data values.
Go Variable Types

In Go, there are different types of variables, for example:

    * int- stores integers (whole numbers), such as 123 or -123
    float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
    * string - stores text, such as "Hello World". String values are surrounded by double quotes
    * bool- stores values with two states: true or false

More about different variable types, will be explained in the Go Data Types chapter.
Declaring (Creating) Variables

In Go, there are two ways to declare a variable:
1. With the var keyword:

Use the var keyword, followed by variable name and type:
Syntax
var variablename type = value

Note: You always have to specify either type or value (or both).
2. With the := sign:

Use the := sign, followed by the variable value:
Syntax
variablename := value

Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

Note: It is not possible to declare a variable using :=, without assigning a value to it.
Variable Declaration With Initial Value

If the value of a variable is known from the start, you can declare the variable and assign a value to it on one line:
Example
```go
package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}
```
Note: The variable types of student2 and x is inferred from their values.

REMOVE ADS
Variable Declaration Without Initial Value

In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:
Example
```go
package main
import ("fmt")

func main() {
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```
Example explained

In this example there are 3 variables:

    a
    b
    c

These variables are declared but they have not been assigned initial values.

By running the code, we can see that they already have the default values of their respective types:

    a is ""
    b is 0
    c is false


### Value Assignment After Declaration

It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.
Example
```go
package main
import ("fmt")

func main() {
  var student1 string
  student1 = "John"
  fmt.Println(student1)
}
```
Note: It is not possible to declare a variable using ":=" without assigning a value to it.
Difference Between var and :=

There are some small differences between the var var :=:
var 	:=
Can be used inside and outside of functions 	Can only be used inside functions
Variable declaration and value assignment can be done separately 	Variable declaration and value assignment cannot be done separately (must be done in the same line)
Example

This example shows declaring variables outside of a function, with the var keyword:
```go
package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```
Example

Since := is used outside of a function, running the program results in an error.
```go
package main
import ("fmt")

a := 1

func main() {
  fmt.Println(a)
}
```
Result:
./prog.go:5:1: syntax error: non-declaration statement outside function body

## Go Multiple Variable Declaration
Go Multiple Variable Declaration

In Go, it is possible to declare multiple variables on the same line.
Example

This example shows how to declare multiple variables on the same line:
```go
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```
Note: If you use the type keyword, it is only possible to declare one type of variable per line.

If the type keyword is not specified, you can declare different types of variables on the same line:
Example
```go
package main
import ("fmt")

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```
### Go Variable Declaration in a Block

Multiple variable declarations can also be grouped together into a block for greater readability:
Example
```go
package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```
### Go Variable Naming Rules
Go Variable Naming Rules

A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

Go variable naming rules:

    * A variable name must start with a letter or an underscore character (_)
    * A variable name cannot start with a digit
    * A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
    * Variable names are case-sensitive (age, Age and AGE are three different variables)
    * There is no limit on the length of the variable name
    * A variable name cannot contain spaces
    * The variable name cannot be any Go keywords

Multi-Word Variable Names

Variable names with more than one word can be difficult to read.

There are several techniques you can use to make them more readable:
#### Camel Case

Each word, except the first, starts with a capital letter:
myVariableName = "John"
#### Pascal Case

Each word starts with a capital letter:
MyVariableName = "John"
#### Snake Case

Each word is separated by an underscore character:
my_variable_name = "John"


## Go Constants
Go Constants

If a variable should have a fixed value that cannot be changed, you can use the const keyword.

The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
Syntax
const CONSTNAME type = value

Note: The value of a constant must be assigned when you declare it.
Declaring a Constant

Here is an example of declaring a constant in Go:
Example
```go
package main
import ("fmt")

const PI = 3.14

func main() {
  fmt.Println(PI)
}
```
Constant Rules

    Constant names follow the same naming rules as variables
    Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
    Constants can be declared both inside and outside of a function

## Constant Types

There are two types of constants:

    Typed constants
    Untyped constants

## Typed Constants

Typed constants are declared with a defined type:
Example
```go
package main
import ("fmt")

const A int = 1

func main() {
  fmt.Println(A)
}
```

## Untyped Constants

Untyped constants are declared without a type:
Example
```go
package main
import ("fmt")

const A = 1

func main() {
  fmt.Println(A)
}
```
Note: In this case, the type of the constant is inferred from the value (means the compiler decides the type of the constant, based on the value).
Constants: Unchangeable and Read-only

When a constant is declared, it is not possible to change the value later:
Example
```go
package main
import ("fmt")

func main() {
  const A = 1
  A = 2
  fmt.Println(A)
}
```
Result:
./prog.go:8:7: cannot assign to A
Multiple Constants Declaration

Multiple constants can be grouped together into a block for readability:
Example
```go
package main
import ("fmt")

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
```
