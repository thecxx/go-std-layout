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
	print = &cobra.Command{}
)

func init() {
	var (
		flags = &handler.PrintFlags{GlobalFlags: gflags}
	)
	print.Use = "print"
	print.Short = "Print any message"
	print.Long = "Print any message"
	// Events
	print.RunE = func(cmd *cobra.Command, args []string) error {
		return handler.OnPrintHandler(cmd.Context(), flags, args)
	}
	// Flags
	if f := print.Flags(); f != nil {
		f.StringVarP(&flags.Message, "message", "m", "", "message")
	}
}
