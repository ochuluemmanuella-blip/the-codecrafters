# The Go programming Language
Go (also called Golang) is a programming language created by Google that is designed to be simple, fast, and efficient. It is used to build things like websites, servers, and large systems that need to handle many tasks at the same time. Unlike some languages that are slower or more complex, Go is compiled (which makes it run quickly) and statically typed (which helps catch errors early). One of its strongest features is built-in support for concurrency, meaning it can run multiple tasks at once without much difficulty. Go also manages memory automatically, so developers don’t have to worry as much about cleaning up unused data. Overall, it’s a great language for beginners who want something powerful but easy to understand.

### Summary (5 Key Points):

Go is a simple, fast, and efficient programming language created by Google
It is used for web servers, networking, and large-scale applications
It runs quickly because it is a compiled language
It can handle multiple tasks at once using built-in concurrency features
It automatically manages memory, making it easier for beginners to use
## Go Get Started

To start using Go, you need two things:

    A text editor, like VS Code, to write Go code
    A compiler, like GCC, to translate the Go code into a language that the computer will understand
## Go Quickstart

Let's create our first Go program.

    * Launch the VS Code editor
    * Open the extension manager or alternatively, press Ctrl + Shift + x
    * In the search box, type "go" and hit enter
    * Find the Go extension by the GO team at Google and install the extension
    * After the installation is complete, open the command palette by pressing Ctrl + Shift + p
    * Run the Go: Install/Update Tools command
    * Select all the provided tools and click OK

VS Code is now configured to use Go.

Open up a terminal window and type:
go mod init example.com/hello

Do not worry if you do not understand why we type the above command. Just think of it as something that you always do, and that you will learn more about in a later chapter.

Create a new file (File > New File). Copy and paste the following code and save the file as helloworld.go (File > Save As):
helloworld.go
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}

Now, run the code: Open a terminal in VS Code and type:
* go run .\helloworld.go
> Hello World!

Congratulations! You have now written and executed your first Go program.

If you want to save the program as an executable, type and run:
* go build .\helloworld.go


#### helloworld.go

Code:
package main
import ("fmt")

func main() {
 fmt.Println("Hello World!")
}

Result:
Hello World! 
# Go Syntax
A Go file consists of the following parts:

    Package declaration
    Import packages
    Functions
    Statements and expressions

Look at the following code, to understand it better:
Example
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
Example explained

Line 1: In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.

Line 2: import ("fmt") lets us import files included in the fmt package.

Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more readable.

Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.

Line 5: fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".

Note: In Go, any executable code belongs to the main package.
