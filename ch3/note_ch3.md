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

# U17 Slice and For Loops
Go 有兩種陣列
- Array: Fixed length list of things
- Slice: An array that can grow or shrink

Slice 的值必須是相同的型別

## 宣告變數並賦予 Slice 為其值

    變數名稱 := []型別{值1, 值2, 值3,...}
	cards := []string{"Ace of Diamonds", newCard()}

## Slice 插入新值

    變數名稱 = append(變數名稱, 插入的值)
	cards = append(cards, "Six of Spades")

以上操作不會對原本的 Slice 造成變動，而是產生一個新的 Slice 給變數 cards

## Slice For Loop 迴圈

    for 陣列索引值, 目前迭代標的 := range 陣列{
        欲執行的代碼
    }

	for i, card := range cards {
		fmt.Println(i, card)
	}

# U18
# U19 Custom Type Declarations
# U20 Receiver Functions
在 func 關鍵詞與函式名稱之間加入圓括弧，
其中依序填入 type 實體的簡稱（慣例是一或二個英文字母），以及 type 名稱。
經此設定後，綁定 Receiver 的 每個 type 實體即可使用 Receiver Function 
```Go
/*
func (type實體 type名稱) 函式名稱(引數) {
	for i, value := range type實體 {
		fmt.Println(i, value)
	}
}
*/ 

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

測驗5
以下代碼合法。
```Go
package main
import "fmt"
type book string
func (b book) printTitle() {
    fmt.Println(b)
}
func main() {
    var b book = "Harry Potter"
    b.printTitle()
}
```
# U21 Creating a New Deck

# U22
# U23
# U24
# U25
# U26
# U27
# U28
# U29
# U3
