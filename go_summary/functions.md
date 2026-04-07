# Go Functions

A function is a block of statements that can be used repeatedly in a program.

A function will not execute automatically when a page loads.

A function will be executed by a call to the function.
Create a Function

To create (often referred to as declare) a function, do the following:

    Use the func keyword.
    Specify a name for the function, followed by parentheses ().
    Finally, add code that defines what the function should do, inside curly braces {}.

Syntax
```go
func FunctionName() {
  // code to be executed
}
```
Call a Function

Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

In the example below, we create a function named "myMessage()". The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function. The function outputs "I just got executed!". To call the function, just write its name followed by two parentheses ():
Example
```go
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}
```
Result:
I just got executed!

A function can be called multiple times.
Example
```go
package main
import ("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage()
  myMessage()
  myMessage()
}
```
Result:
I just got executed!
I just got executed!
I just got executed!


### Naming Rules for Go Functions

    > A function name must start with a letter
    > A function name can only contain alpha-numeric characters and underscores . (A-z, 0-9, and _ )
    > Function names are case-sensitive
    > A function name cannot contain spaces
    If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used

Tip: Give the function a name that reflects what the function does!
## Go Output Functions

Go has three functions to output text:

    Print()
    Println()
    Printf()

### The Print() Function

The Print() function prints its arguments with their default format.
Example

Print the values of i and j:
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i)
  fmt.Print(j)
}

Result:
HelloWorld
Example

If we want to print the arguments in new lines, we need to use \n.
```go
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n")
  fmt.Print(j, "\n")
}
```
Result:
Hello
World

Tip: \n creates new lines.
Example

It is also possible to only use one Print() for printing multiple variables.
```go
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, "\n",j)
}
```
Result:
Hello
World
Example

If we want to add a space between string arguments, we need to use " ":
```go
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Print(i, " ", j)
}
```
Result:
Hello World
Example

* Print() inserts a space between the arguments if neither are strings:
```go
package main
import ("fmt")

func main() {
  var i,j = 10,20

  fmt.Print(i,j)
}

Result:
10 20
```

### The Println() Function

The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:
Example
```go
package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
}
```
Result:
Hello World
### The Printf() Function

The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

    %v is used to print the value of the arguments
    %T is used to print the type of the arguments

Example
```go
package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}
```
Result:
i has value: Hello and type: string
j has value: 15 and type: int