package rubrikpolaris

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ExpandTildePath will replace a "~/" prefix with the user home directory.
// Caveat: only "~/" is supported (~user/path is not).
func ExpandTildePath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path, fmt.Errorf("could not resolve home dir: %v", err)
		}
		return filepath.Join(home, path[2:]), nil
	}
	return path, nil
}

// GetStringFromSlice is a workaround for Go's lack of default func params:
// it returns s[index] if it exists; if it doesn't, it returns defaultVal;
// if replaceEmpty is true and s[index] is empty, it also returns defaultVal.
func GetStringFromSlice(s []string, index int, replaceEmpty bool,
	defaultVal string) string {
	if len(s) > index && !(replaceEmpty && s[index] == "") {
		return s[index]
	}
	return defaultVal
}
