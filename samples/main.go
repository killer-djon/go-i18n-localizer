package main

import (
	. "github.com/killer-djon/go-i18n-localizer"
	"github.com/killer-djon/go-i18n-localizer/helpers"
	"log"
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

	translation.SetContexts("messages")
	translation.BindParams(map[string]interface{}{
		":replacement": "Иванов",
		":last_name": "Петров",
		":count": 2,
	})
	msg := translation.ParseString("message.by_replace")

	log.Println("Localized text", msg)
}
