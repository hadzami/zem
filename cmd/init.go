package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func AddAlias(alias string, command string) error {
	zshrcPath := os.Getenv("HOME") + "/.zshrc"

	aliasLine := fmt.Sprintf("alias %s='%s'\n", alias, command)

	file, err := os.OpenFile(zshrcPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open .zshrc: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(aliasLine); err != nil {
		return fmt.Errorf("failed to write alias to .zshrc: %w", err)
	}

	cmd := exec.Command("zsh", "-c", "source "+zshrcPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to source .zshrc: %w", err)
	}

	return nil
}
