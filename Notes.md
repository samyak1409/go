## Task

Madhu sir:

> Core Language & Framework
>
> - Go 1.24 - Primary programming language
> - Gin - HTTP web framework for building REST APIs
> - Google Wire - Dependency injection framework
>
> Database & ORM
>
> - Ent - Entity framework for Go (similar to GORM but more type-safe)
> - Goose - Database migration tool

## Notes

- Go is a high-level, general-purpose programming language that is statically-typed and compiled.
- It is known for the simplicity of its syntax and the efficiency of development that it enables through the inclusion of a large standard library supplying many needs for common projects.
- It was designed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson, and publicly announced in November 2009.
- It is syntactically similar to C, but also has garbage collection, structural typing, and CSP-style concurrency.
- It is often referred to as Golang to avoid ambiguity and because of its former domain name, golang.org, but its proper name is Go.

- Based on C

- Code file -> main func -> entry point

- Go doesn't require virtual environments like Python. Instead, it uses a single installation and manages dependencies at the project level.

- ```sh
  go mod init <module-path>
  ```

  Creates a `go.mod`.

  `go.mod` is like `requirements.txt` -> tracks dependencies

- ```sh
  go run .
  ```

- A package is a way to group functions, and it's made up of all the files in the same directory.

- [Standard library](https://pkg.go.dev/std) packages

- Import any external package in the code, and run:

  ```
  go mod tidy
  ```

  Go will add its module as a requirement to `go.mod`.

- [Go Playground](https://go.dev/play)

- In Go, a name is exported if it begins with a capital letter. For example, Pi is an exported name which is exported from the math package.

  Any "unexported" names are not accessible from outside the \*package\*.

- A function can return any number of results (values).

- Naked return statements should be used only in short functions. They can harm readability in longer functions.

- [Basic types](https://go.dev/tour/basics/11)

- `if`, `for`: `()` for expression not required, but `{}` for block code required.

- Static typing / type hints are added after the var name not before.

- No try/except exists, intentionally, to make the code more robust, as you would need to think more on what could go wrong.
