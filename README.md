# khmer-number

Lib for convert Khmer number into Khmer words.


```sh
  fmt.Println(Num2Word("10005.55"))
package main

import (
  "fmt"
  kn "github.com/khmerlang/khmer-number"
)

func main() {
  fmt.Println(kn.Num2Word("1"))
  fmt.Println(kn.Num2Word("2"))
  fmt.Println(kn.Num2Word("10005.55"))
  fmt.Println(kn.Num2Word("12.05"))
}
```
