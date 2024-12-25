package exception

import (
	"fmt"
	"samplecode/cmd/util"
)

func CatchUp() {
	if err := recover(); err != nil {
		util.LogError(fmt.Sprintf("Internal server error or '%v'", err))
	}
}
