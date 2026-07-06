package masker

import (
	"regexp"
	"strings"
)

type Masker struct {
	patterns []string
}

var defaultSensitivePatterns = []string{
	// Common sensitive key names
	`password`,
	`passwd`,
	`secret`,
	`token`,
	`api[_-]?key`,
	`apikey`,
	`auth`,
	`credential`,
	`private[_-]?key`,
	`access[_-]?key`,
	`access[_-]?token`,
	`refresh[_-]?token`,
	`jwt`,
	`ssn`,
	`credit[_-]?card`,
	`app_password`,
	
	// AWS specific
	`aws[_-]?access[_-]?key[_-]?id`,
	`aws[_-]?secret[_-]?access[_-]?key`,
	`aws[_-]?session[_-]?token`,
	
	// Database
	`connection[_-]?string`,
	`db[_-]?password`,
	`db[_-]?pass`,
	
	// Generic
	`pwd`,
	`passphrase`,
	`encryption[_-]?key`,
	`signing[_-]?key`,
	`oauth`,
	`client[_-]?secret`,
}

func NewMasker(additionalPatterns []string) *Masker {
	patterns := append(defaultSensitivePatterns, additionalPatterns...)
	return &Masker{
		patterns: patterns,
	}
}

func (m *Masker) MaskContent(content string, fileType string) string {
	maskedContent := content

	switch fileType {
	case "yaml", "json", "toml", "dotenv", "ini":
		maskedContent = m.maskKeyValueContent(maskedContent, fileType)
	default:
		// For unknown types, try generic masking
		maskedContent = m.maskGenericContent(maskedContent)
	}

	return maskedContent
}

func (m *Masker) maskKeyValueContent(content string, fileType string) string {
	for _, pattern := range m.patterns {
		switch fileType {
		case "yaml":
			content = m.maskYAML(content, pattern)
		case "json":
			content = m.maskJSON(content, pattern)
		case "toml":
			content = m.maskTOML(content, pattern)
		case "dotenv":
			content = m.maskDotEnv(content, pattern)
		case "ini":
			content = m.maskINI(content, pattern)
		}
	}
	return content
}

func (m *Masker) maskYAML(content string, pattern string) string {
	// Match YAML key-value pairs: key: value
	re := regexp.MustCompile(`(?im)^(\s*)` + pattern + `(\s*:\s*)(.+)$`)
	return re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 4 {
			indent := parts[1]
			key := parts[2]
			value := strings.TrimSpace(parts[3])
			maskedValue := strings.Repeat("*", len(value))
			return indent + pattern + key + maskedValue
		}
		return match
	})
}

func (m *Masker) maskJSON(content string, pattern string) string {
	// Match JSON key-value pairs: "key": "value" or "key": value
	re := regexp.MustCompile(`(?i)"(` + pattern + `)"\s*:\s*"([^"]*)"`)
	content = re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 3 {
			key := parts[1]
			value := parts[2]
			maskedValue := strings.Repeat("*", len(value))
			return `"` + key + `": "` + maskedValue + `"`
		}
		return match
	})

	// Also match non-string values (numbers, booleans, etc.)
	re2 := regexp.MustCompile(`(?i)"(` + pattern + `)"\s*:\s*([^,}\s]+)`)
	content = re2.ReplaceAllStringFunc(content, func(match string) string {
		parts := re2.FindStringSubmatch(match)
		if len(parts) == 3 {
			key := parts[1]
			value := parts[2]
			maskedValue := strings.Repeat("*", len(value))
			return `"` + key + `": ` + maskedValue
		}
		return match
	})

	return content
}

func (m *Masker) maskTOML(content string, pattern string) string {
	// Match TOML key-value pairs: key = "value" or key = value
	re := regexp.MustCompile(`(?im)^(\s*)` + pattern + `(\s*=\s*)"([^"]*)"`)
	content = re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 4 {
			indent := parts[1]
			key := parts[2]
			value := parts[3]
			maskedValue := strings.Repeat("*", len(value))
			return indent + pattern + key + `"` + maskedValue + `"`
		}
		return match
	})

	// Match non-quoted values
	re2 := regexp.MustCompile(`(?im)^(\s*)` + pattern + `(\s*=\s*)([^\s#]+)`)
	content = re2.ReplaceAllStringFunc(content, func(match string) string {
		parts := re2.FindStringSubmatch(match)
		if len(parts) == 4 {
			indent := parts[1]
			key := parts[2]
			value := parts[3]
			maskedValue := strings.Repeat("*", len(value))
			return indent + pattern + key + maskedValue
		}
		return match
	})

	return content
}

func (m *Masker) maskDotEnv(content string, pattern string) string {
	// Match .env key-value pairs: KEY=value or KEY="value"
	re := regexp.MustCompile(`(?im)^(` + pattern + `)\s*=\s*"([^"]*)"`)
	content = re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 3 {
			key := parts[1]
			value := parts[2]
			maskedValue := strings.Repeat("*", len(value))
			return key + `="` + maskedValue + `"`
		}
		return match
	})

	re2 := regexp.MustCompile(`(?im)^(` + pattern + `)\s*=\s*([^\s#]+)`)
	content = re2.ReplaceAllStringFunc(content, func(match string) string {
		parts := re2.FindStringSubmatch(match)
		if len(parts) == 3 {
			key := parts[1]
			value := parts[2]
			maskedValue := strings.Repeat("*", len(value))
			return key + "=" + maskedValue
		}
		return match
	})

	return content
}

func (m *Masker) maskINI(content string, pattern string) string {
	// Match INI key-value pairs: key = value or key=value
	re := regexp.MustCompile(`(?im)^(\s*)` + pattern + `(\s*=\s*)(.+)$`)
	return re.ReplaceAllStringFunc(content, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 4 {
			indent := parts[1]
			key := parts[2]
			value := strings.TrimSpace(parts[3])
			maskedValue := strings.Repeat("*", len(value))
			return indent + pattern + key + maskedValue
		}
		return match
	})
}

func (m *Masker) maskGenericContent(content string) string {
	// Generic key=value masking
	for _, pattern := range m.patterns {
		re := regexp.MustCompile(`(?im)` + pattern + `\s*[:=]\s*([^\s,;]+)`)
		content = re.ReplaceAllStringFunc(content, func(match string) string {
			re2 := regexp.MustCompile(`(?im)(` + pattern + `\s*[:=]\s*)([^\s,;]+)`)
			parts := re2.FindStringSubmatch(match)
			if len(parts) == 3 {
				prefix := parts[1]
				value := parts[2]
				maskedValue := strings.Repeat("*", len(value))
				return prefix + maskedValue
			}
			return match
		})
	}
	return content
}