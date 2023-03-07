// Copyright 2023 Kami
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"context"
	"fmt"

	"github.com/thecxx/go-std-layout/tools/pkg/internal"
	"github.com/thecxx/go-std-layout/tools/pkg/service"
)

type CmdFlags struct {
	*GlobalFlags
	Install bool
	Remove  bool
	Parent  string
}

// OnCmdHandler
func OnCmdHandler(ctx context.Context, flags *CmdFlags, args []string) (err error) {
	workspace, err := flags.GetWorkspace()
	if err != nil {
		return err
	}
	gp, err := service.GetGoProject(workspace)
	if err != nil {
		return err
	}
	if len(args) <= 0 {
		return fmt.Errorf("command not found")
	}

	gp.License.Header = internal.TryReadFile(fmt.Sprintf("%s/HEADER", workspace))
	gp.License.Description = internal.TryReadFile(fmt.Sprintf("%s/LICENSE", workspace))

	if flags.Install {
		return service.Command.Install(ctx, gp, args[0], flags.Parent)
	} else if flags.Remove {
		return service.Command.Remove(ctx, gp, args[0], flags.Parent)
	}
	return
}
