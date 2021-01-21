# Golang translation
Golang translation package inspired by symfony translation component

[![License MIT](https://img.shields.io/apm/l/vim-mode.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Build Status](https://travis-ci.com/killer-djon/go-i18n-localizer.svg?branch=master)](https://travis-ci.com/killer-djon/go-i18n-localizer)
[![Go Report Card](https://goreportcard.com/badge/github.com/killer-djon/go-i18n-localizer)](https://goreportcard.com/report/github.com/killer-djon/go-i18n-localizer)

### How to user
```go
package main

import (
    "github.com/killer-djon/go-i18n-localizer/helpers"
    . "github.com/killer-djon/go-i18n-localizer"
)     
//First step you must create config structure

// 1. From config json file
localizerConfig := helpers.ParseJson("./config.json")

//OR

// 2. Make config structure point
localizerConfig := &helpers.LocalizerConfig{
	Locale:          "en",
	TranslationPath: "translation",
	DefaultContext:  "user",
}

// Then create translate interface instance
translation := NewTranslation(localizerConfig)
// and paste localized files direct to translate directory
// all files must be names like:
// <context_name>.<local>.yaml
// as example users.en.yaml
// and after that point you can call translate method

translation.SetContexts("users") // context in this case is prefix for file name
// for get first level text you can call
msg := translation.ParseString("First user")
// for second level and another level by dimension you can call
msg := translation.ParseString("user.first")
// You can create pattern for messages 
// and when call them can replace
translation.BindParams(map[string]interface{}{
    ":count_children": 2,
})
msg := translation.ParseString("profile.first")
```

##### Your file can have some keys like this
```yaml
// FIle name users.en.yaml
First user: 'First user'
second_user: 'Second user'
user:
  first: 'First user 2'
  second: 'Second user 2'
profile:
  first: 'Mrs. Smith have :count_children children'
```