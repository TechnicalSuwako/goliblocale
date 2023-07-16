# liblocale

自分のプロジェクトで使うには：

```sh
go get gitler.moe/suwako/goliblocale
```

## 例
```json
{
  "hello": "こんにちは"
}
```

```go
package main

import (
  "fmt"
  "gitler.moe/suwako/goliblocale"
)

func main () {
  i18n, err := goliblocale.GetLocale("ja")
  if err != nil {
    fmt.Println("liblocaleエラー：%v", err)
    return
  }

  fmt.Println(i18n["hello"])
}
```
