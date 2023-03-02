// Copyright 2022 Kami
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
	"github.com/thecxx/go-std-layout/examples/demo/pkg/console/demo/handler"
)

var (
	demo   = &cobra.Command{}
	gflags = &handler.GlobalFlags{}
)

func init() {
	demo.Use = "demo"
	demo.Short = "go-std-layout-demo"
	demo.Long = "go-std-layout-demo"
	demo.Version = "1.0.0"
	demo.SilenceUsage = true
	demo.CompletionOptions.HiddenDefaultCmd = true
	// Events
	demo.RunE = func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	}
	// Flags
	if pf := demo.PersistentFlags(); pf != nil {
		pf.StringVarP(&gflags.Who, "who", "w", "nobody", "execute with specified user privileges")
	}
}

func main() {
	var (
		cmds []*cobra.Command
	)

	// Register sub commands
	cmds = append(cmds, port)
	cmds = append(cmds, print)

	demo.AddCommand(cmds...)
	defer func() {
		demo.RemoveCommand(cmds...)
	}()

	demo.Execute()
}
