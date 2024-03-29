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
	port = &cobra.Command{}
)

func init() {
	var (
		flags = &handler.PortFlags{GlobalFlags: gflags}
	)
	port.Use = "port"
	port.Short = "Startup an online server"
	port.Long = "Startup an online server"
	// Events
	port.RunE = func(cmd *cobra.Command, args []string) error {
		return handler.OnPortHandler(cmd.Context(), flags, args)
	}
	// Flags
	if f := port.Flags(); f != nil {
		f.IntVarP(&flags.Port, "port", "p", 8090, "a port for server to bind")
	}
}
