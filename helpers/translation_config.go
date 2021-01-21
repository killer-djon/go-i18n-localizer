package helpers

import (
	"encoding/json"
	"log"
	"os"
)

const (
	CONFIG_FILE              = "/go/bin/config.json"
	DEFAULT_LOCALE           = "ru"
	DEFAULT_TRANSLATION_PATH = "translations"
	DEFAULT_CONTEXT          = "messages"
)

// LocalizerConfig structure for configuration
type LocalizerConfig struct {
	Locale          string `json:"locale"`
	TranslationPath string `json:"translation_path"`
	DefaultContext  string `json:"default_context"`
}

var DefaultLocaleConfig = LocalizerConfig{
	Locale:          DEFAULT_LOCALE,
	TranslationPath: DEFAULT_TRANSLATION_PATH,
	DefaultContext:  DEFAULT_CONTEXT,
}

// NewTranslation create localizer config instance
// of the structure to implement
func NewTranslationConfig(configFile ...string) *LocalizerConfig {
	return ParseJson(configFile...)
}

// ParseJson try to parse json config file
// for implements configuration options
func ParseJson(configFile ...string) *LocalizerConfig {
	var cFile *string
	if len(configFile) > 0 {
		cFile = &configFile[0]
	}

	if cFile != nil {
		jsonFile, err := os.Open(*cFile)

		if err != nil {
			log.Printf("Error when try to open json config file %v", err)
			return nil
		}

		defer jsonFile.Close()


		decoder := json.NewDecoder(jsonFile)
		err = decoder.Decode(&DefaultLocaleConfig)

		if err != nil {
			log.Println("Can't decode json config file, check json config and try again", err)
			return nil
		}
	}

	return &DefaultLocaleConfig
}

