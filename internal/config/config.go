package config

import (
	"os"
	"path/filepath"
	"runtime"
)

const (
	aliases        = "aliases.json"
	appName        = "clias"
	appData        = "AppData"
	ghNotConfigDir = "GHNOT_CONFIG_DIR"
	localAppData   = "LocalAppData"
	xdgConfigHome  = "XDG_CONFIG_HOME"
	xdgStateHome   = "XDG_STATE_HOME"
)

func Dir() string {
	var path string
	if a := os.Getenv(ghNotConfigDir); a != "" {
		path = a
	} else if b := os.Getenv(xdgConfigHome); b != "" {
		path = filepath.Join(b, appName)
	} else if c := os.Getenv(appData); runtime.GOOS == "windows" && c != "" {
		path = filepath.Join(c, appName)
	} else {
		d, _ := os.UserHomeDir()
		path = filepath.Join(d, ".config", appName)
	}

	return filepath.Join(path, aliases)
}
