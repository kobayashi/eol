package commands

import (
	"fmt"
)

const version = "v0.1.0"

// RunVersion shows the current version
func RunVersion() {
	fmt.Print(version)
}
