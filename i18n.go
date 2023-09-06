package datetime

import (
	"fmt"

	"github.com/go-zoox/core-utils/strings"
	"github.com/spf13/cast"
)

// I18nTranslations is the translations.
type I18nTranslations map[string]string

// I18nLocales is the locales.
type I18nLocales map[string]I18nTranslations

// SetLocales is the week days.
func SetLocales(locales I18nLocales) {
	Locales = locales
}

// SetLanguage sets the language.
func SetLanguage(language string) {
	Language = language
}

// translate translates the given key with the given arguments.
func translate(locales I18nLocales, locale string, key string, data ...map[string]any) (string, error) {
	if locales == nil {
		return "", fmt.Errorf("locales not loaded")
	}

	// Get the translation for the given key.
	translations, ok := locales[locale]
	if !ok {
		return "", fmt.Errorf("locale(%s) not found", locale)
	}

	translation, ok := translations[key]
	if !ok {
		return "", fmt.Errorf("translation key(%s) not found", key)
	}

	// If no data was given, return the translation.
	if len(data) == 0 || data[0] == nil {
		return cast.ToString(translation), nil
	}

	return strings.Format(translation, data[0]), nil
}

// t is an alias for translate.
func t(key string, data ...map[string]any) string {
	translation, err := translate(Locales, Language, key, data...)
	if err != nil {
		return key
	}

	return translation
}
