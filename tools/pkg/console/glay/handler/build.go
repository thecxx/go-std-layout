package handler

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/thecxx/go-std-layout/tools/pkg/internal"
	"github.com/thecxx/go-std-layout/tools/pkg/service"
)

type BuildFlags struct {
	*GlobalFlags
	Install bool
	Output  string
}

func NewBuildFlags(gflags *GlobalFlags) (flags *BuildFlags) {
	flags = &BuildFlags{GlobalFlags: gflags}
	// Install
	flags.Install = false
	// Output
	flags.Output = "bin"
	return
}

func OnBuildHandler(ctx context.Context, flags *BuildFlags, args []string) (err error) {
	if len(args) <= 0 {
		return fmt.Errorf("command not found")
	}

	ws, err := os.Getwd()
	if err != nil {
		return err
	}
	gp, err := service.NewProject(ws)
	if err != nil {
		return err
	}

	var (
		cmd = strings.ToLower(args[0])
		out = flags.Output
	)

	if out == "" {
		out = fmt.Sprintf("%s/bin", ws)
	}

	if flags.Install {
		_, err = internal.Shell("go", "install", fmt.Sprintf("%s/cmd/%s", gp.Module, cmd))
	} else {
		_, err = internal.Shell("go", "build", "-o", fmt.Sprintf("%s/%s", out, cmd), fmt.Sprintf("%s/cmd/%s", gp.Module, cmd))
	}
	return

}
