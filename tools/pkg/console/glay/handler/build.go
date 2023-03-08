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
	// Test string
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
	)
	_, err = internal.Shell("go", "build", "-o", fmt.Sprintf("%s/bin/%s", ws, cmd), fmt.Sprintf("%s/cmd/%s", gp.Module, cmd))
	return

}
