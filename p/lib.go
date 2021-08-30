package p

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func Version() string {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		version := bi.Main.Version
		version = strings.Trim(version, "()")
		return fmt.Sprintf("%s@%s\n", bi.Main.Path, version)
	}
	return "unknown"
}
