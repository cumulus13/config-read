package main

import (
	"fmt"
	"os"

	"github.com/cumulus13/config-read/internal/config"
	"github.com/cumulus13/config-read/internal/highlighter"
	"github.com/cumulus13/config-read/internal/masker"
	"github.com/cumulus13/config-read/internal/pager"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var (
	cfgFile     string
	noPager     bool
	themeName   string
	sensitivePatterns []string
	version     = "1.0.0 by Hadi Cahyadi <cumulus13@gmail.com>"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "config-read [file]",
		Short: "Read config files with syntax highlighting and sensitive data masking",
		Long: `config-read is a cross-platform CLI tool that reads configuration files
with beautiful syntax highlighting (default: Fruity theme), pagination support,
and automatic masking of sensitive data like passwords, tokens, and API keys.

Examples:
  config-read config.yaml
  config-read .env --no-pager
  config-read appsettings.json --theme monokai
  config-read config.toml --sensitive-patterns "password,secret,key"`,
		Args: cobra.MinimumNArgs(1),
		RunE: runConfigRead,
	}

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config-read.yaml)")
	rootCmd.PersistentFlags().BoolVar(&noPager, "no-pager", false, "disable paging")
	rootCmd.PersistentFlags().StringVar(&themeName, "theme", "fruity", "color theme (fruity, monokai, dracula, nord, solarized-dark, solarized-light)")
	rootCmd.PersistentFlags().StringArrayVar(&sensitivePatterns, "sensitive-patterns", []string{}, "additional sensitive key patterns to mask")

	// Version command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("config-read version %s\n", version)
		},
	})

	// Config command
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
	}
	configCmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.LoadConfig(cfgFile)
			fmt.Printf("Theme: %s\n", cfg.Theme)
			fmt.Printf("Pager Enabled: %v\n", cfg.PagerEnabled)
			fmt.Printf("Sensitive Patterns: %v\n", cfg.SensitivePatterns)
		},
	})
	configCmd.AddCommand(&cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set configuration value",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return config.SetConfigValue(cfgFile, args[0], args[1])
		},
	})
	rootCmd.AddCommand(configCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runConfigRead(cmd *cobra.Command, args []string) error {
	filePath := args[0]

	// Load configuration
	cfg := config.LoadConfig(cfgFile)

	// Override with CLI flags
	if themeName != "fruity" {
		cfg.Theme = themeName
	}
	if noPager {
		cfg.PagerEnabled = false
	}
	if len(sensitivePatterns) > 0 {
		cfg.SensitivePatterns = append(cfg.SensitivePatterns, sensitivePatterns...)
	}

	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Detect file type
	fileType := highlighter.DetectFileType(filePath)

	// Create highlighter with theme
	h, err := highlighter.NewHighlighter(cfg.Theme, fileType)
	if err != nil {
		return fmt.Errorf("failed to create highlighter: %w", err)
	}

	// Create masker
	m := masker.NewMasker(cfg.SensitivePatterns)

	// Process content: mask sensitive data then highlight
	processedContent := string(content)
	processedContent = m.MaskContent(processedContent, fileType)
	highlightedContent := h.Highlight(processedContent)

	// Display with or without pager
	if cfg.PagerEnabled && !noPager {
		pgr := pager.NewPager()
		return pgr.Display(highlightedContent)
	}

	fmt.Print(highlightedContent)
	return nil
}