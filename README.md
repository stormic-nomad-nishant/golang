
# Golang

A self paced learning project for teaching Golang


## Installing go

You can install golang by following the official doc https://go.dev/learn/


## Running a Go program


There are 2 ways to run a go program, the first one is long way:

1.  In order to run a go program the first thing you need to do is compile it :

```bash
$ go build file_name.go
```

This will create a binary file that you can run as follows:

```bash
$ ./file_name
```

2. The second one is the short way of doing the same thing above using the go run command which abstracts the compilation step :

```bash
$ go run file_name.go
```

## Background of Golang

-   Go was originally created by Robert Griesemar, Rob Pike & Ken Thompson at Google.
-   Why Have GO ?
    -   Python : Easy to use but slow as its interprted language
    -   Jave : Quick but the type system is very complex
    -   C/C++ : Quick , but type system is complex and takes a lot of time to compile code.
    -   Concurrancy is later patched in the above language.
-   Go is Strongly typed language.
    -   A strongly typed language makes sure that the a variable cannot be changed over time. i.e if you declare a variable to hold `a` as an int then you cannot later choose to store a string value to it
    -   Strongly typed is a concept used to refer to a programming language that enforces strict restrictions on intermixing of values with differing data types. When such restrictions are violated and error (exception) occurs.
    -   <https://stackoverflow.com/questions/2690544/what-is-the-difference-between-a-strongly-typed-language-and-a-statically-typed>
-   Go is also a statically typed language.
    -   A language is statically typed if the type of a variable is known at compile time.
    -   Statically typed programming languages do type checking (i.e. the process of verifying and enforcing the constraints of types) at compile-time as opposed to run-time. Ex: Java, C, C++
    -   Dynamically typed programming languages do type checking at run-time as opposed to compile-time. Ex: Perl, Ruby, Python, PHP, JavaScript
-   GO Key Features:
    -   Simplicity
    -   Extreamly fast compile time
    -   Automatically Garbage Collection : <https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8>
    -   Built-in Concurrancy : <https://www.golang-book.com/books/intro/10>
    -   Compile to standalone binaries: Everything once your program is compiled is wrapped up in one binary.
