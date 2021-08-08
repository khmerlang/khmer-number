# khmer-number

Lib for convert Khmer number into Khmer words.

```sh
go get github.com/khmerlang/khmer-number

```

```sh
package main

import (
  "fmt"
  kn "github.com/khmerlang/khmer-number"
)

func main() {

  fmt.Println(kn.Num2Word("1"))
  // មួយ
  fmt.Println(kn.Num2Word("2"))
  // ពីរ
  fmt.Println(kn.Num2Word("10005.55"))
  // មួយម៉ឺនប្រាំចុចហាសិបប្រាំ
  fmt.Println(kn.Num2Word("12.05"))
  // ដប់ពីរចុចសូន្យប្រាំ
}
```

### Special thank to @invisal  for converting logic
