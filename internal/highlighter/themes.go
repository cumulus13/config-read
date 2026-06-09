package highlighter

import (
	"sync"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/styles"
)

var (
	themeRegistryOnce sync.Once
)

// init automatically registers themes when package is imported
func init() {
	RegisterThemes()
}

// RegisterThemes registers all custom themes with Chroma
func RegisterThemes() {
	themeRegistryOnce.Do(func() {
		// Register Fruity theme and its variants
		fruityTheme := GetFruityTheme()
		if fruityTheme != nil {
			styles.Register(fruityTheme)
		}
		
		// Register Fruity variants
		variants := GetAdditionalFruityVariants()
		for _, variant := range variants {
			if variant != nil {
				styles.Register(variant)
			}
		}
		
		// Register custom Monokai Pro theme
		monokaiPro := chroma.MustNewStyle("monokai-pro", chroma.StyleEntries{
			chroma.Background:    "#2d2a2e bg:#2d2a2e",
			chroma.Text:          "#fcfcfa",
			chroma.Comment:       "#727072 italic",
			chroma.Keyword:       "#ff6188 bold",
			chroma.KeywordType:   "#ff6188",
			chroma.LiteralString: "#a9dc76",
			chroma.LiteralNumber: "#ab9df2",
			chroma.NameFunction:  "#78dce8 bold",
			chroma.NameClass:     "#78dce8 bold",
			chroma.NameTag:       "#ff6188",
			chroma.Operator:      "#ff6188",
		})
		styles.Register(monokaiPro)
		
		// Register custom Dracula Pro theme
		draculaPro := chroma.MustNewStyle("dracula-pro", chroma.StyleEntries{
			chroma.Background:    "#22212c bg:#22212c",
			chroma.Text:          "#f8f8f2",
			chroma.Comment:       "#7970a9 italic",
			chroma.Keyword:       "#ff79c6 bold",
			chroma.KeywordType:   "#8be9fd",
			chroma.LiteralString: "#f1fa8c",
			chroma.LiteralNumber: "#bd93f9",
			chroma.NameFunction:  "#50fa7b bold",
			chroma.NameClass:     "#50fa7b bold",
			chroma.NameTag:       "#ff79c6",
			chroma.Operator:      "#ff79c6",
		})
		styles.Register(draculaPro)
		
		// Register custom Nord theme
		nordTheme := chroma.MustNewStyle("nord-custom", chroma.StyleEntries{
			chroma.Background:    "#2e3440 bg:#2e3440",
			chroma.Text:          "#d8dee9",
			chroma.Comment:       "#4c566a italic",
			chroma.Keyword:       "#81a1c1 bold",
			chroma.KeywordType:   "#8fbcbb",
			chroma.LiteralString: "#a3be8c",
			chroma.LiteralNumber: "#b48ead",
			chroma.NameFunction:  "#88c0d0 bold",
			chroma.NameClass:     "#88c0d0 bold",
			chroma.NameTag:       "#81a1c1",
			chroma.Operator:      "#81a1c1",
		})
		styles.Register(nordTheme)
	})
}

// GetAvailableThemes returns list of all available themes including custom ones
func GetAvailableThemes() []string {
	RegisterThemes()
	
	themes := styles.Names()
	
	// Add custom themes that might not be in the default registry
	customThemes := []string{
		"fruity",
		"fruity-light",
		"fruity-darker",
		"monokai-pro",
		"dracula-pro",
		"nord-custom",
	}
	
	themeMap := make(map[string]bool)
	for _, t := range themes {
		themeMap[t] = true
	}
	
	for _, t := range customThemes {
		if !themeMap[t] {
			themes = append(themes, t)
		}
	}
	
	return themes
}

// IsThemeAvailable checks if a specific theme is available
func IsThemeAvailable(themeName string) bool {
	RegisterThemes()
	
	for _, theme := range GetAvailableThemes() {
		if theme == themeName {
			return true
		}
	}
	
	return false
}
