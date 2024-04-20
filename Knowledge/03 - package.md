# Go Package

> A package is a collection of source files/code in the same directory that are compiled together. 
> Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
> - https://go.dev/doc/code#Organization

Package คือไฟล์ของ Go ทุกไฟล์่ทีอยู่ใน**โฟลเดอร์เดียวกันที่จะถูก compile** ไปด้วยกัน โดยทุก ๆ ฟังก์ชั่น ตัวแปร หรือค่าคงที่ในไฟล์ ๆ หนึ่งใน package จะใช้ร่วมกันได้ในทุก ๆ ไฟล์ ใน package เดียวกัน

Package is simply a Go file, dependency, and/or a open-sourced go files available for use.

Think in 2 side; 
1. Local Go package: All Go files in the local project.
2. Online Open-Source Go package: All go libraries, packages available on the internet.

Every Go files in the project need a keyword `package` declared at the top; `package <package-name>`.

For the main(primary) package, `package` needed to be declared as `main` and it reuquires `main()` function in the package. Otherwise Go will raise error that `function main is undeclared in the main package`.


**IMPORTANT:**

1. Public Function, which available for all modules, must have signature start with captial letter.
      - `func HelloWorldPublic()`
2. Private Function, only available to just that package. Will have 1st non-capital letter in the signature
      - `func helloWorldPrivate()`

## Create customized Go Package in the Project

Reference to `00 - basic, syntax, modules, packaging/01 - experiment`

The directory `greetings` contains a new package called `greetings`.

However, this package is contained within the same module. Therefore, package `main` can call the code from that package.

`main` package can import the sub-package by using the same module URI with additional path to the target. package

Here, the `main` package import `greetings` package with `import "example.com/yourusername/go-example/greetings"`


### **Installing Go Package**

```sh
go get github.com/google/uuid
```

The above snippet add a dependency for Go project into `go.mod` and `go.sum` files. 

The package can now be used in the project.

**View all the added packages**

```sh
go list -m all 
```

**Import/Use package within Go file**

