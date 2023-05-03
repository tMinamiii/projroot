package projroot

import (
	"os"
	"path/filepath"
)

func Find() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		_, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				// at the root
				break
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			return ""
		}
		break
	}
	return currentDir
}
