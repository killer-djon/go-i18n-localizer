package go_i18n_localizer

import (
	"fmt"
	"github.com/killer-djon/go-i18n-localizer/helpers"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Translation struct {
	config *helpers.LocalizerConfig
	texts  map[string]interface{}
	bindParams []string
	contexts []string
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
			var newMap = make(map[string]interface{})
			trans.FlattenMap(mapText, "", newMap)

			texts[match[1]] = newMap
		}
	}

	trans.texts = texts
}

// FlattenMap - method for flatten multidimensional map
// to single level map values
func (trans *Translation) FlattenMap(value interface{}, prefix string, m map[string]interface{}) {
	base := ""
	if prefix != "" {
		base = prefix+"."
	}

	for key, text := range value.(map[interface{}]interface{}) {
		switch reflect.ValueOf(text).Kind() {
		case reflect.String:
			m[base + key.(string)] = text
			break
		case reflect.Map:
			trans.FlattenMap(text, base + key.(string), m)
			break
		}
	}
}

// SetBindParams Simple setter for bind params for replacement
func (trans *Translation) BindParams(params map[string]interface{}) {
	if params != nil {
		for key, item := range params {
			var valueItem string
			switch reflect.ValueOf(item).Kind() {
			case reflect.String:
				valueItem = item.(string)
				break
			case reflect.Int:
				valueItem = strconv.Itoa(item.(int))
				break
			case reflect.Float64:
			case reflect.Float32:
				valueItem = fmt.Sprintf("%f", item)
				break
			}
			trans.bindParams = append(trans.bindParams, key, valueItem)
		}
	}
}

// SetUsedContexts instantiate used contexts for display messages by key
func (trans *Translation) SetContexts(context ...string) {
	var defaultContext = []string{trans.config.DefaultContext}
	if len(context) > 0 {
		defaultContext = context
	}

	trans.contexts = defaultContext
}

// ParseString return string text for requests by key
// before search text by key must parse incoming key
func (trans Translation) ParseString(keyString string) string {
	var translateText string
	for _, contextItem := range trans.contexts {
		if trans.texts[contextItem] != nil {
			textMap := trans.texts[contextItem].(map[string]interface{})
			if textMap[keyString] != nil {
				parsedText := textMap[keyString].(string)
				if trans.bindParams !=nil {
					//log.Println("Bind params", trans.bindParams)
					re := strings.NewReplacer(trans.bindParams...)
					parsedText = re.Replace(parsedText)
				}

				translateText += parsedText
			}
		}
	}

	return translateText
}
