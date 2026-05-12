package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//go:embed locales/*
var localesFS embed.FS

var translations map[string]string
var currentLocale string

func Init() {
	lang := detectLanguage()
	currentLocale = lang

	data, err := localesFS.ReadFile(fmt.Sprintf("locales/%s.json", lang))
	if err != nil {
		// Fallback a inglés si el idioma detectado no existe
		data, _ = localesFS.ReadFile("locales/en.json")
		currentLocale = "en"
	}

	if err := json.Unmarshal(data, &translations); err != nil {
		fmt.Printf("Error loading translations: %v\n", err)
	}
}

func T(key string) string {
	val, ok := translations[key]
	if !ok {
		return key
	}
	return val
}

func GetLocale() string {
	return currentLocale
}

func detectLanguage() string {
	// 1. Prioridad: Variable de entorno VTX_LANG
	if env := os.Getenv("VTX_LANG"); env != "" {
		return strings.ToLower(env[:2])
	}

	// 2. Variable LANG en Linux/macOS
	if lang := os.Getenv("LANG"); lang != "" {
		return strings.ToLower(lang[:2])
	}

	// 3. TODO: Windows registry para mayor precisión si fuera necesario
	
	return "en" // Default
}
