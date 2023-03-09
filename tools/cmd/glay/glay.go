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
	"github.com/spf13/cobra"
	"github.com/thecxx/go-std-layout/tools/pkg/console/glay/handler"
)

var (
	glayc  = &cobra.Command{}
	gflags = handler.NewGlobalFlags()
)

func init() {
	// var (
	// 	flags = handler.NewGlayFlags(gflags)
	// )
	glayc.Use = "glay"
	glayc.Short = "A tool use for managing Go project which use the go-std-layout directory structure"
	glayc.Long = `A tool use for managing Go project which use the go-std-layout directory structure.`
	glayc.Version = "1.0.0"
	glayc.SilenceUsage = true
	glayc.CompletionOptions.HiddenDefaultCmd = true
	// Events
	glayc.RunE = func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
		// return handler.OnGlayHandler(cmd.Context(), flags, args)
	}
	// Flags
	// if f := glayc.Flags(); f != nil {
	//     f.StringVarP(&flags.Test, "test", "t", flags.Test, "a test flag")
	// }
	// if pf := glayc.PersistentFlags(); pf != nil {
	// }
}

func main() {
	var (
		cmds []*cobra.Command
	)

	// Register sub commands
	cmds = append(cmds, cmdc)
	cmds = append(cmds, buildc)
	cmds = append(cmds, licensec)
	// sub command placeholder

	glayc.AddCommand(cmds...)
	defer func() {
		glayc.RemoveCommand(cmds...)
	}()

	glayc.Execute()
}
