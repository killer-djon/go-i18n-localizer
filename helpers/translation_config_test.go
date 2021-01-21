package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestMakeDefaultConfig(t *testing.T) {
	var translateConfig = NewTranslationConfig()

	if translateConfig != &DefaultLocaleConfig {
		t.Errorf("%v=%v default config for translation instance is not equal", translateConfig, &DefaultLocaleConfig)
	}
}

func TestMakeFileConfig(t *testing.T) {
	var testFileName = "test.json"
	type testingFileConfig struct {
		Locale          string `json:"locale"`
		TranslationPath string `json:"translation_path"`
		DefaultContext  string `json:"default_context"`
	}

	var newTestingFile = testingFileConfig{
		Locale:          "en",
		TranslationPath: "translations",
		DefaultContext:  "messages",
	}

	jsonTestData, _ := json.MarshalIndent(newTestingFile, "", "")
	_ = ioutil.WriteFile(testFileName, jsonTestData, 0644)

	defer os.Remove(testFileName)
	translateConfig := NewTranslationConfig(testFileName)
	if translateConfig.Locale != newTestingFile.Locale {
		t.Errorf("Locale loaded from config file is not identical with struct config")
	}

	if translateConfig.TranslationPath != newTestingFile.TranslationPath {
		t.Errorf("TranslationPath loaded from config file is not identical with struct config")
	}

	if translateConfig.DefaultContext != newTestingFile.DefaultContext {
		t.Errorf("DefaultContext loaded from config file is not identical with struct config")
	}


}
