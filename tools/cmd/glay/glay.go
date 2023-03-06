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

package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/thecxx/go-std-layout/tools/pkg/console/glay/handler"
)

var (
	glay   = &cobra.Command{}
	gflags = &handler.GlobalFlags{}
)

func init() {
	// var (
	// 	flags = &handler.GlayFlags{GlobalFlags: gflags}
	// )
	glay.Use = "layout"
	glay.Short = "A short description"
	glay.Long = "A long description"
	glay.Version = "1.0.0"
	glay.SilenceUsage = true
	glay.CompletionOptions.HiddenDefaultCmd = true
	// Events
	glay.RunE = func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
		// return handler.OnGlayHandler(cmd.Context(), flags, args)
	}
	// Flags
	// if f := glay.Flags(); f != nil {
	//     f.StringVarP(&flags.Test, "test", "t", "", "a test flag")
	// }
	if pf := glay.PersistentFlags(); pf != nil {
		wd, _ := os.Getwd()
		pf.StringVarP(&gflags.Workspace, "workspace", "w", wd, "a test flag")
	}
}

func main() {
	var (
		cmds []*cobra.Command
	)

	// Register sub commands
	cmds = append(cmds, cmdc)

	glay.AddCommand(cmds...)
	defer func() {
		glay.RemoveCommand(cmds...)
	}()

	glay.Execute()
}
