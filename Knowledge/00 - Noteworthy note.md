# Note about Package
1. Go only has concept of Function, Package and Module
   - It does not have any OOP concept, no Class & Object
2. Only 1 Package per 1 Directory
3. Package has Internal directory feature to limit accessibility from outside of package
    - Let's say, `Greetings` package has sub-directory called `internal` and a package called `speaking` inside that sub-directory.
    - `main` package that call upon `Greetings` package will not able to call `speaking` package inside `internal` directory.