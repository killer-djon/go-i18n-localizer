package go_i18n_localizer

import (
	"fmt"
	"github.com/killer-djon/go-i18n-localizer/helpers"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type Translation struct {
	config *helpers.LocalizerConfig
	texts map[string]interface{}
}

func NewTranslation(config *helpers.LocalizerConfig) *Translation {
	var trans = Translation{
		config: config,
		texts:  nil,
	}

	trans.ParseTextLocalized()
	return &trans
}

func (trans *Translation) ParseTextLocalized() {
	transDir := trans.config.TranslationPath
	files, err := ioutil.ReadDir(transDir)
	if err != nil {
		log.Fatal(err)
	}

	var texts = make(map[string]interface{})
	for _, f := range files {
		re := regexp.MustCompile("([a-z]+)\\.([a-z]{2,4})")
		match := re.FindStringSubmatch(f.Name())

		if len(match) > 2 && match[2] == trans.config.Locale {
			yamlFile, err := os.Open(fmt.Sprintf("%s/%s", transDir, f.Name()))

			if err != nil {
				log.Println("Error to open transliteration file", err)
				continue
			}

			defer yamlFile.Close()

			var mapText interface{}
			yamlByte, _ := ioutil.ReadAll(yamlFile)
			yaml.Unmarshal(yamlByte, &mapText)

			texts[match[1]] = mapText
		}
	}

	trans.texts = texts
}

func (trans Translation) ParseString(keyString string, params map[string]interface{}, context *string) {
	log.Println(keyString, params, context)
}