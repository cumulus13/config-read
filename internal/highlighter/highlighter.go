package highlighter

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

// Highlighter handles syntax highlighting for config files
type Highlighter struct {
	theme    *chroma.Style
	lexer    chroma.Lexer
	fileType string
}

// NewHighlighter creates a new Highlighter with the specified theme and file type
func NewHighlighter(themeName string, fileType string) (*Highlighter, error) {
	// Register custom themes
	RegisterThemes()
	
	// Get theme
	theme := styles.Get(themeName)
	if theme == nil {
		// Try loading custom Fruity theme
		theme = GetFruityTheme()
		if theme == nil {
			return nil, fmt.Errorf("theme '%s' not found", themeName)
		}
	}

	// Get lexer for file type
	lexer := lexers.Get(fileType)
	if lexer == nil {
		// Try to match by filename
		lexer = lexers.Match(fileType)
		if lexer == nil {
			// Fallback to plain text
			lexer = lexers.Fallback
		}
	}

	// Configure lexer for better config file handling
	lexer = chroma.Coalesce(lexer)

	return &Highlighter{
		theme:    theme,
		lexer:    lexer,
		fileType: fileType,
	}, nil
}

// Highlight applies syntax highlighting to the content
func (h *Highlighter) Highlight(content string) string {
	// Create formatter for terminal output with 256 colors
	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	// Tokenize the content
	iterator, err := h.lexer.Tokenise(nil, content)
	if err != nil {
		// If tokenization fails, return original content
		return content
	}

	// Format the tokens with the theme
	var buf strings.Builder
	err = formatter.Format(&buf, h.theme, iterator)
	if err != nil {
		// If formatting fails, return original content
		return content
	}

	return buf.String()
}

// DetectFileType detects the file type based on extension and filename
func DetectFileType(filePath string) string {
	ext := strings.ToLower(filepath.Ext(filePath))
	baseName := strings.ToLower(filepath.Base(filePath))
	
	// Map common config file extensions
	extMap := map[string]string{
		".yaml":      "yaml",
		".yml":       "yaml",
		".json":      "json",
		".toml":      "toml",
		".ini":       "ini",
		".cfg":       "ini",
		".conf":      "ini",
		".env":       "dotenv",
		".xml":       "xml",
		".hcl":       "hcl",
		".tf":        "hcl",
		".properties": "java-properties",
		".editorconfig": "ini",
		".cnf":       "ini",
		".config":    "xml",
		".plist":     "xml",
		".lock":      "yaml",
		".template":  "yaml",
	}

	if lexer, ok := extMap[ext]; ok {
		return lexer
	}

	// Try to detect by filename (no extension or special names)
	nameMap := map[string]string{
		"dockerfile":        "docker",
		"makefile":          "makefile",
		"jenkinsfile":       "groovy",
		"vagrantfile":       "ruby",
		".env":              "dotenv",
		".env.example":      "dotenv",
		".env.local":        "dotenv",
		".env.development":  "dotenv",
		".env.production":   "dotenv",
		".env.staging":      "dotenv",
		"gemfile":           "ruby",
		"rakefile":          "ruby",
		"procfile":          "yaml",
		"docker-compose.yml": "yaml",
		"docker-compose.yaml": "yaml",
	}

	if lexer, ok := nameMap[baseName]; ok {
		return lexer
	}

	// Try to detect by content type (basic heuristics)
	if strings.Contains(baseName, "docker") {
		return "docker"
	}
	if strings.Contains(baseName, "make") {
		return "makefile"
	}
	if strings.HasPrefix(baseName, ".") && strings.Contains(baseName, "env") {
		return "dotenv"
	}

	// Default fallback
	return "plaintext"
}

// GetFileTypeDescription returns a human-readable description of the file type
func GetFileTypeDescription(fileType string) string {
	descriptions := map[string]string{
		"yaml":             "YAML Configuration",
		"json":             "JSON Configuration",
		"toml":             "TOML Configuration",
		"ini":              "INI Configuration",
		"dotenv":           "Environment Variables",
		"xml":              "XML Configuration",
		"hcl":              "HCL Configuration",
		"docker":           "Dockerfile",
		"makefile":         "Makefile",
		"groovy":           "Groovy Script",
		"ruby":             "Ruby Script",
		"java-properties":  "Java Properties",
		"plaintext":        "Plain Text",
	}

	if desc, ok := descriptions[fileType]; ok {
		return desc
	}
	return fmt.Sprintf("%s File", strings.ToUpper(fileType))
}

// IsBinaryFileType checks if the file type is binary
func IsBinaryFileType(fileType string) bool {
	binaryTypes := []string{
		"binary",
		"image",
		"video",
		"audio",
		"compressed",
		"executable",
	}
	
	for _, bt := range binaryTypes {
		if strings.Contains(fileType, bt) {
			return true
		}
	}
	return false
}

// GetSupportedExtensions returns all supported file extensions
func GetSupportedExtensions() []string {
	return []string{
		".yaml", ".yml",
		".json",
		".toml",
		".ini", ".cfg", ".conf", ".cnf",
		".env",
		".xml", ".config", ".plist",
		".hcl", ".tf",
		".properties",
		".editorconfig",
	}
}
