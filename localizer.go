package main

import (
	"go-i18n-localizer/helpers"
	"log"
)


func main() {
	//localizerConfig := helpers.ParseJson("./config.json")
	//localizerConfig := &helpers.LocalizerConfig{
	//	Locale:          "en",
	//	TranslationPath: "translation",
	//	DefaultContext:  "user",
	//}

	localizerConfig := helpers.NewTranslation("./config.json")

	log.Println("Start localizer", localizerConfig)
}