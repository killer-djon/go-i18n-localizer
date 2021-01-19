package main

import (
	. "github.com/killer-djon/go-i18n-localizer"
	"github.com/killer-djon/go-i18n-localizer/helpers"
)

func main() {
	//localizerConfig := helpers.ParseJson("./config.json")
	//localizerConfig := &helpers.LocalizerConfig{
	//	Locale:          "en",
	//	TranslationPath: "translation",
	//	DefaultContext:  "user",
	//}

	localizerConfig := helpers.NewTranslationConfig("./config.json")
	translation := NewTranslation(localizerConfig)

	translation.ParseString("simple message", nil, nil)
}
