package pager

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Pager struct {
	command string
}

func NewPager() *Pager {
	return &Pager{
		command: getDefaultPager(),
	}
}

func getDefaultPager() string {
	// Check environment variables
	if pager := os.Getenv("PAGER"); pager != "" {
		return pager
	}
	
	// Platform-specific defaults
	switch runtime.GOOS {
	case "windows":
		return "more" // Windows doesn't have less by default
	default:
		// Check if less is available
		if _, err := exec.LookPath("less"); err == nil {
			return "less -R -F -X" // -R for raw control chars (colors), -F quit if one screen, -X no termcap init/deinit
		}
		return "more"
	}
}

func (p *Pager) Display(content string) error {
	// If no pager command, just print
	if p.command == "" {
		fmt.Print(content)
		return nil
	}

	// Parse pager command
	parts := strings.Fields(p.command)
	if len(parts) == 0 {
		fmt.Print(content)
		return nil
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Create pipe for input
	stdin, err := cmd.StdinPipe()
	if err != nil {
		// Fallback to direct print
		fmt.Print(content)
		return nil
	}

	// Start command
	if err := cmd.Start(); err != nil {
		// Fallback to direct print
		fmt.Print(content)
		return nil
	}

	// Write content
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, content)
	}()

	// Wait for command to finish
	return cmd.Wait()
}

// IsPagerAvailable checks if a pager is available on the system
func IsPagerAvailable() bool {
	return getDefaultPager() != ""
}