package info

import (
	"fmt"
	"runtime"
)

var gitCommit string

type Info struct {
	GitCommit string
	GoOS      string
	GoArch    string
	GoVersion string
}

var info Info

func init() {
	info.GitCommit = gitCommit
	info.GoOS = runtime.GOOS
	info.GoArch = runtime.GOARCH
	info.GoVersion = runtime.Version()
	fmt.Printf("\033[32mCurrent executable binary information: %+v\033[0m\n", info)
}

func Infos() Info {
	return info
}
