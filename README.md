go-yandex-translate
===================

Go (golang) Yandex Translate API wrapper

Usage:

```
package main

import (
  "fmt"
  "github.com/icrowley/go-yandex-translate"
)

func main() {
  tr := yandex_translate.New(YOUR_API_KEY)

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
