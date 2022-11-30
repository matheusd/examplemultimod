package main

import (
	"fmt"
	"runtime/debug"

	"github.com/matheusd/examplemultimod/libone/v4"
	"github.com/matheusd/examplemultimod/libtwo"
)

func main() {
	fmt.Println(libone.Name(), libone.Version())
	fmt.Println(libtwo.Name(), libtwo.Version())
	fmt.Println("  ", libtwo.Deps())

	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Build info not available")
		return
	}
	fmt.Println("")
	fmt.Println("Go Version:", info.GoVersion)
	fmt.Println("Path:", info.Path)
	fmt.Println("Main:", info.Main.Path, "@", info.Main.Version)
	fmt.Println("Deps:")
	for _, dep := range info.Deps {
		replacedBy := "<nil>"
		if dep.Replace != nil {
			replacedBy = fmt.Sprintf("%s@%v", dep.Replace.Path,
				dep.Replace.Version)
		}
		fmt.Println("  Path:", dep.Path, " Version:", dep.Version,
			" Replace:", replacedBy)
	}
	fmt.Println("Build Settings:")
	for _, set := range info.Settings {
		fmt.Println(set.Key, "-", set.Value)
	}
}
