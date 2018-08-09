[![Build Status](https://travis-ci.org/dafanasev/go-yandex-translate.svg?branch=master)](https://travis-ci.org/dafanasev/go-yandex-translate)
[![GoDoc](https://godoc.org/github.com/dafanasev/go-yandex-translate?status.svg)](https://godoc.org/github.com/dafanasev/go-yandex-translate)
[![Go Report Card](https://goreportcard.com/badge/github.com/dafanasev/go-yandex-translate)](https://goreportcard.com/report/github.com/dafanasev/go-yandex-translate)
[![Coverage Status](https://coveralls.io/repos/github/dafanasev/go-yandex-translate/badge.svg)](https://coveralls.io/github/dafanasev/go-yandex-translate)

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
    fmt.Println(translation.Result())
  }
}
```
