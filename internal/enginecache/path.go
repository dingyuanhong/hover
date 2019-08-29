package enginecache

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
)

func cachePath() string {
	homePath := os.Getenv("HOVER_HOME")
	if homePath == "" {
		path, err := homedir.Dir()
		if err != nil {
			fmt.Printf("hover: Failed to resolve home path: %v\n", err)
			os.Exit(1)
		}
		homePath = path
	}

	var p string
	switch runtime.GOOS {
	case "linux":
		p = filepath.Join(homePath, ".cache")
	case "darwin":
		p = filepath.Join(homePath, "Library", "Caches")
	case "windows":
		p = filepath.Join(homePath, "AppData", "Local")
	default:
		fmt.Printf("hover: cannot run on %s, enginecache not implemented.\n", runtime.GOOS)
		os.Exit(1)
	}
	return p
}
