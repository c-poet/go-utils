package os

import (
	"fmt"
	"log/slog"
	stdos "os"
	"path/filepath"
)

// GetProgramPath returns the path of the current executable. If the executable
// path cannot be determined, it falls back to the first program argument.
func GetProgramPath() string {
	programPath, err := stdos.Executable()
	if err != nil {
		slog.Error("get executable path", "error", err)
		return stdos.Args[0]
	}
	return programPath
}

// GetProgramExtPath returns the path of a named executable located beside the
// current program. On Windows, the .exe extension is appended automatically.
// It returns an error when the executable does not exist.
func GetProgramExtPath(name string) (string, error) {
	programDir := filepath.Dir(GetProgramPath())
	extension := ""
	if IsWindows() {
		extension = ".exe"
	}

	programPath := filepath.Join(programDir, name+extension)
	if _, err := stdos.Stat(programPath); stdos.IsNotExist(err) {
		return "", fmt.Errorf("executable %q does not exist: %s", name, programPath)
	}
	return programPath, nil
}
