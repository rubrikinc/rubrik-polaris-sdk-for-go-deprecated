package rubrikpolaris

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	operationNamePrefix = "SdkGoLang"
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

// GetStringFromSlice is useful to emulate default function params:
// it returns s[index] if it exists; if it doesn't, it returns defaultVal;
// if replaceEmpty is true and s[index] is empty, it also returns defaultVal.
func GetStringFromSlice(s []string, index int, replaceEmpty bool,
	defaultVal string) string {
	if len(s) > index && !(replaceEmpty && s[index] == "") {
		return s[index]
	}
	return defaultVal
}

// OperationNamePrefix returns a prefix string to be used when sending up
// operation names to Rubrik. The default is "SdkGoLang".
// So for instance if sending up the query "RadarEventsPerTimePeriod",
// the operation name will be "SdkGoLangRadarEventsPerTimePeriod".
//
// A less common use case is to specify a *second* prefix: say
// operationNameSecondPrefix="XYZ", then for the query above, the operation
// name sent up to Rubrik would be "SdkGoLangXYZRadarEventsPerTimePeriod"
func OperationNamePrefix(operationNameSecondPrefix ...string) string {
	if len(operationNameSecondPrefix) == 0 {
		return operationNamePrefix
	}
	op := operationNameSecondPrefix[0]
	if !strings.HasPrefix(op, operationNamePrefix) {
		return operationNamePrefix + op
	}
	return op
}
