package go_i18n_localizer

import (
	"github.com/killer-djon/go-i18n-localizer/helpers"
	"reflect"
	"testing"
)

func TestTranslation_FlattenMap(t *testing.T) {
	translate := NewTranslation(&helpers.LocalizerConfig{
		Locale:          "en",
		TranslationPath: helpers.DEFAULT_TRANSLATION_PATH,
		DefaultContext:  helpers.DEFAULT_CONTEXT,
	})

	var texts = make(map[interface{}]interface{})
	texts = map[interface{}]interface{}{
		"simple_text": "text",
		"second_level": map[interface{}]interface{}{
			"simple": "text",
			"another": "text",
		},
	}

	var newMap = make(map[string]interface{})
	translate.FlattenMap(texts, "", newMap)

	for _, text := range newMap {
		if reflect.ValueOf(text).Kind() == reflect.Map {
			t.Errorf("Incoming map is not flattened")
			break
		}
	}
}

func TestTranslation_ParseString(t *testing.T) {
	translate := NewTranslation(&helpers.LocalizerConfig{
		Locale:          "en",
		TranslationPath: helpers.DEFAULT_TRANSLATION_PATH,
		DefaultContext:  helpers.DEFAULT_CONTEXT,
	})
	var texts = make(map[interface{}]interface{})
	texts = map[interface{}]interface{}{
		"simple_text": "text",
		"second_level": map[interface{}]interface{}{
			"simple": "text",
			"another": "text",
		},
	}
	var newMap = make(map[string]interface{})
	translate.FlattenMap(texts, "", newMap)

	translate.texts[helpers.DEFAULT_CONTEXT] = newMap
	mapValues := reflect.ValueOf(translate.texts)

	if len(mapValues.MapKeys()) != 2 {
		t.Errorf("Count keys after flatten is not correct, must be 3 keys")
	}

	translate.SetContexts(helpers.DEFAULT_CONTEXT)
	if translate.ParseString("simple_text") != "text" {
		t.Errorf("Incorrect parse string from translation file by top level")
	}

	if translate.ParseString("second_level.simple") != "text" {
		t.Errorf("Incorrect parse string from translation file by second level")
	}
}