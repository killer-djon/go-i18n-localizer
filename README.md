# Golang translation
Golang translation package inspired by symfony translation component

[![License MIT](https://img.shields.io/apm/l/vim-mode.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Build Status](https://travis-ci.com/killer-djon/gimdownloader.svg?branch=master)](https://travis-ci.com/killer-djon/gimdownloader)
[![Go Report Card](https://goreportcard.com/badge/github.com/killer-djon/gimdownloader)](https://github.com/killer-djon/go-i18n-localizer)

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

// 2. Make config structure point
```