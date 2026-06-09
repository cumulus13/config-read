package highlighter

import (
	"github.com/alecthomas/chroma/v2"
)

// GetFruityTheme returns a custom theme inspired by Python Pygments' Fruity style
// Based on the original Fruity theme from Pygments library
func GetFruityTheme() *chroma.Style {
	return chroma.MustNewStyle("fruity", chroma.StyleEntries{
		// Core
		chroma.Background:    "#111111 bg:#111111",
		chroma.Text:          "#ffffff",
		chroma.Other:         "#ffffff",
		
		// Error
		chroma.Error:         "#ffffff bg:#aa0000",
		
		// Comments
		chroma.Comment:            "#008800 italic",
		chroma.CommentMultiline:   "#008800 italic",
		chroma.CommentPreproc:     "#ff0007 bold",
		chroma.CommentSingle:      "#008800 italic",
		chroma.CommentSpecial:     "#008800 italic bold",
		
		// Generic
		chroma.GenericDeleted:    "#ff0007 bg:#440000",
		chroma.GenericEmph:       "italic",
		chroma.GenericError:      "#ffffff bg:#aa0000",
		chroma.GenericHeading:    "#ffffff bold",
		chroma.GenericInserted:   "#0086d2 bg:#003350",
		chroma.GenericOutput:     "#888888",
		chroma.GenericPrompt:     "#ffffff bold",
		chroma.GenericStrong:     "bold",
		chroma.GenericSubheading: "#ffffff bold",
		chroma.GenericTraceback:  "#ffffff underline",
		
		// Keywords - Orange family
		chroma.Keyword:            "#fb660a bold",
		chroma.KeywordConstant:    "#0086d2",
		chroma.KeywordDeclaration: "#fb660a bold",
		chroma.KeywordNamespace:   "#ff0086 bold",
		chroma.KeywordPseudo:      "#0086d2",
		chroma.KeywordReserved:    "#fb660a bold",
		chroma.KeywordType:        "#cdcaa9 bold",
		
		// Literals
		chroma.Literal:                "#0086d2",
		chroma.LiteralDate:            "#0086d2",
		chroma.LiteralNumber:          "#0086f7 bold",
		chroma.LiteralNumberBin:       "#0086f7 bold",
		chroma.LiteralNumberFloat:     "#0086f7 bold",
		chroma.LiteralNumberHex:       "#0086f7 bold",
		chroma.LiteralNumberInteger:   "#0086f7 bold",
		chroma.LiteralNumberOct:       "#0086f7 bold",
		chroma.LiteralString:          "#0086d2",
		chroma.LiteralStringAffix:     "#0086d2",
		chroma.LiteralStringBacktick:  "#0086d2",
		chroma.LiteralStringChar:      "#0086d2",
		chroma.LiteralStringDelimiter: "#0086d2",
		chroma.LiteralStringDoc:       "#0086d2 italic",
		chroma.LiteralStringDouble:    "#0086d2",
		chroma.LiteralStringEscape:    "#0086d2 bold",
		chroma.LiteralStringHeredoc:   "#0086d2",
		chroma.LiteralStringInterpol:  "#0086d2 bold",
		chroma.LiteralStringOther:     "#0086d2",
		chroma.LiteralStringRegex:     "#0086d2",
		chroma.LiteralStringSingle:    "#0086d2",
		chroma.LiteralStringSymbol:    "#0086d2",
		
		// Names
		chroma.Name:              "#ffffff",
		chroma.NameAttribute:     "#ff0086 bold",
		chroma.NameBuiltin:       "#ff0086 bold",
		chroma.NameBuiltinPseudo: "#fb660a",
		chroma.NameClass:         "#ff0086 bold",
		chroma.NameConstant:      "#0086d2",
		chroma.NameDecorator:     "#ffffff bold",
		chroma.NameEntity:        "#ff0086 bold",
		chroma.NameException:     "#ffffff bold",
		chroma.NameFunction:      "#ff0086 bold",
		chroma.NameFunctionMagic: "#ff0086 bold",
		chroma.NameLabel:         "#ff0086 bold",
		chroma.NameNamespace:     "#ff0086 bold",
		chroma.NameOther:         "#ffffff",
		chroma.NameProperty:      "#ffffff",
		chroma.NameTag:           "#fb660a bold",
		chroma.NameVariable:      "#fb660a",
		chroma.NameVariableClass: "#fb660a",
		chroma.NameVariableGlobal:"#fb660a",
		chroma.NameVariableInstance: "#fb660a",
		chroma.NameVariableMagic: "#fb660a",
		
		// Operators
		chroma.Operator:     "#ffffff",
		chroma.OperatorWord: "#fb660a bold",
		
		// Punctuation
		chroma.Punctuation: "#ffffff",
		
		// Whitespace
		chroma.TextWhitespace: "#888888",
	})
}

// GetAdditionalFruityVariants returns variant themes based on Fruity
func GetAdditionalFruityVariants() map[string]*chroma.Style {
	variants := make(map[string]*chroma.Style)
	
	// Fruity Light variant
	lightFruity := chroma.MustNewStyle("fruity-light", chroma.StyleEntries{
		chroma.Background:    "#ffffff bg:#ffffff",
		chroma.Text:          "#333333",
		chroma.Error:         "#cc0000 bg:#ffeeee",
		chroma.Comment:       "#008800 italic",
		chroma.CommentPreproc: "#cc0000 bold",
		chroma.Keyword:       "#ff6600 bold",
		chroma.KeywordType:   "#996633 bold",
		chroma.LiteralString: "#0066cc",
		chroma.LiteralNumber: "#0066ff bold",
		chroma.NameFunction:  "#cc0066 bold",
		chroma.NameClass:     "#cc0066 bold",
		chroma.NameTag:       "#ff6600 bold",
	})
	variants["fruity-light"] = lightFruity
	
	// Fruity Darker variant
	darkerFruity := chroma.MustNewStyle("fruity-darker", chroma.StyleEntries{
		chroma.Background:    "#000000 bg:#000000",
		chroma.Text:          "#e0e0e0",
		chroma.Error:         "#ff0007 bg:#330000",
		chroma.Comment:       "#006600 italic",
		chroma.CommentPreproc: "#cc0000 bold",
		chroma.Keyword:       "#e65c00 bold",
		chroma.KeywordType:   "#b8a87c bold",
		chroma.LiteralString: "#0066cc",
		chroma.LiteralNumber: "#0055cc bold",
		chroma.NameFunction:  "#e60073 bold",
		chroma.NameClass:     "#e60073 bold",
		chroma.NameTag:       "#e65c00 bold",
	})
	variants["fruity-darker"] = darkerFruity
	
	return variants
}
