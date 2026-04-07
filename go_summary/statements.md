## Go Statements

fmt.Println("Hello World!") is a statement.

In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".

Hitting the Enter key adds ";" to the end of the line implicitly (does not show up in the source code).

The left curly bracket { cannot come at the start of a line.

Run the following code and see what happens:
Example
```go
package main
import ("fmt")

func main()
{
  fmt.Println("Hello World!")
} 
```
## Go Comments
Go Comments

A comment is a text that is ignored upon execution.

Comments can be used to explain the code, and to make it more readable.

Comments can also be used to prevent code execution when testing an alternative code.

Go supports single-line or multi-line comments.
Go Single-line Comments

Single-line comments start with two forward slashes (//).

Any text between // and the end of the line is ignored by the compiler (will not be executed).
Example
```go
// This is a comment
package main
import ("fmt")

func main() {
  // This is a comment
  fmt.Println("Hello World!")
}
```

The following example uses a single-line comment at the end of a code line:
Example
```go
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!") // This is a comment
}
```
Go Multi-line Comments

Multi-line comments start with /* and ends with */.

Any text between /* and */ will be ignored by the compiler:
Example
```go
package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen, and it is amazing */
  fmt.Println("Hello World!")
} 
```

Tip: It is up to you which you want to use. Normally, we use // for short comments, and /* */ for longer comments.

REMOVE ADS
Comment to Prevent Code Execution

You can also use comments to prevent the code from being executed.

The commented code can be saved for later reference and troubleshooting.
Example
```go
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
  // fmt.Println("This line does not execute")
}
```