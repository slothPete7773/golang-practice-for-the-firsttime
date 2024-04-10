# 1. Introduction of Backend

- Middle man(Server) for handling communication between Frontend, Client, Application with the Server, API, Database.
- Go is **aiming to provide a better backend system**.

## What is Go/GoLang

- Started inventing in 2007, by Robert Griesemer, Rob Pike และ Ken Thompson
- Release in 2009 as Open-source by Google.
- Aim to provide simple, simplified, efficient, easy to read, and easy to code.

## Hightlight of Go

- Static Type
- Compilation first.
  - Required Compilation for machine-readable for best performance.
- Garbage Collection
  - Automatically managed and collected the left-over memory and unused variables.
- Concurrency Support

# 2. Try with Go!

Running the Go file without compiling the binary executable file.
```sh
$ go run main.go
```

Compile

# 3. Go Module

A collection of GO packages that are related and dependent in the project.

## Go Package

> A package is a collection of source files in the same directory that are compiled together. 
> Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
> - https://go.dev/doc/code#Organization

Package คือไฟล์ของ Go ทุกไฟล์่ทีอยู่ใน**โฟลเดอร์เดียวกันที่จะถูก compile** ไปด้วยกัน โดยทุก ๆ ฟังก์ชั่น ตัวแปร หรือค่าคงที่ในไฟล์ ๆ หนึ่งใน package จะใช้ร่วมกันได้ในทุก ๆ ไฟล์ ใน package เดียวกัน