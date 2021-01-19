package go_i18n_localizer

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

	localizerConfig := helpers.NewTranslationConfig("./config.json")
	translation := NewTranslation(localizerConfig)

	log.Println("Start localizer", translation)
}
