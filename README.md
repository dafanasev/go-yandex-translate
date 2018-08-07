go-yandex-translate
===================

Go Yandex Translate API wrapper

Usage:

```
package main

import (
  "fmt"
  "github.com/dafanasev/go-yandex-translate"
)

func main() {
  tr := translate.New(YOUR_API_KEY)

  response, err := tr.GetLangs("en")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.Langs)
    fmt.Println(response.Dirs)
  }

  translation, err := tr.Translate("ru", "A lazy dog")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(translation)
    fmt.Println(translation.Result())
  }
}
```
