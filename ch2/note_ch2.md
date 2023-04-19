# U7 Boring Ol' Hello World
本單元僅編寫以下代碼，其意義於下個章節解說。
```Go
package main

import "fmt"

func main () {
	fmt.Println("Hi there!")
}
```
# U8 Five Important Questions
1. How do we run the code in our project?
2. What does 'package main' mean?
3. What does 'import "fmt"' mean?
4. What's the 'func' thing?
5. How is the main.go file organized?

本單元僅回答第一個問題。
執行 Go 代碼：在終端機輸入以下指令 $ go run 檔案路徑

## Go CLI
- go build
編譯並產生可執行檔，但不執行。例如：
 $ go build main.go
以上指令會產出名為 main 的檔案
然後在檔案所在位置輸入以下指令即可執行檔案
 $ ./main

- go run
編譯並執行
# U9 Go Packages
Key Words
package
package name
Executable Package
Reusable Package

## Package
Package 相當於 Project 或 Workspace
一個 Package 當中可能包含多個 .go 檔，每個 .go 檔首行代碼都註明其所屬的 Package。
例：Package Main 包含 main.go, support.go, helper.go...


## Types of Packages
- Executable Package
可藉由 $ go build fileName.go 生成執行檔
舉凡首行代碼為 package main 者，即是 Executable Package。
Executable Package 之中必須有個名為 main 的函式。

- Reusable Package
類似套件，非執行檔。
首行代碼非 package main 者，即是 Reusable Package。
對 Reusable Package 執行 $ go build 並不會產生執行檔。

# U10 Import Statements
import "fmt" 的 fmt 是 format 的縮寫
Go 內建的函式庫：
crypto, debug, encoding, fmt, io, math...

main package 可以 import fmt(Standard Lib), calculator(Reusable Package), uploader(Reusable, Package)...

[Go 內建函式庫官方文件](https://pkg.go.dev/std)

## U11 File Orgnization
- What's the 'func' thing?
```Go
func main() {
    fmt.Println("Hello")
}
```
- How is the main.go file organized?
```Go
package main // Package declaration

import "fmt"
 
 // Declare functions, tell Go to do things.
func main() {
    fmt.Println("Hello")
}
```

