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
# U9
# U0
