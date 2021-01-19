package go_i18n_localizer

import "go-i18n-localizer/helpers"

type Translation struct {
	config *helpers.LocalizerConfig
	texts interface{}
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

}