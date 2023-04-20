# U13 Project Overview

# U14 New Project Folder

在 package main 當中一定要宣告 func main(){}。
main 函式會在執行檔案時自動執行。

# U15 Variable Declarations
變數宣告
	var card string = "Ace of Spades" // 等號左邊是宣告（初始化）變數，右邊是變數賦值
    var 變數名稱 資料型別 = 值

    card := "Ace of Spades"
    變數名稱 := 值

重新賦值
    card = "Five of Diamonds"
    變數名稱 = 值


資料型別
bool, string, int, float64...

測驗2
```Go
// 不合法的代碼，在函式外不能為變數賦值。
package main
import "fmt"
deckSize := 20
func main(){
    fmt.Println(deckSize)
}
```

```Go
// 合法的代碼，在函式外宣告（初始化）變數符合 Go 的規定。
package main
import "fmt"
var deckSize int
func main(){
    deckSize = 50
    fmt.Println(deckSize)
}
```

# U16 Functions and Return Types
函式宣告：
func 函式名稱 () 回傳型別 {
    return 回傳值
}

func newCard() string {
	return "Five of Diamonds"
}

測驗3
函式宣告有提升作用，宣告及調用順序可以顛倒。
同一 package 的不同檔案可共用彼此的函式，且不需編寫表示引用的代碼。

# U17

# U18

# U19

# U20
