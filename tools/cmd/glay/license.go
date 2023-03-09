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
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thecxx/go-std-layout/tools/pkg/common"
	"github.com/thecxx/go-std-layout/tools/pkg/console/glay/handler"
)

var (
	licensec = &cobra.Command{}
)

func init() {
	var (
		flags = handler.NewLicenseFlags(gflags)
	)
	licensec.Use = fmt.Sprintf("license [flags] {%s}", strings.Join(common.LicenseKeys, "|"))
	licensec.Short = "Generating license"
	licensec.Long = "Generating license for our project."
	licensec.Example = `  glay license apache2
  glay license --header=false apache2`
	licensec.Aliases = []string{"lic"}
	licensec.ValidArgs = common.LicenseKeys
	licensec.Args = cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)
	// Events
	licensec.RunE = func(cmd *cobra.Command, args []string) error {
		return handler.OnLicenseHandler(cmd.Context(), flags, args)
	}
	// Flags
	if f := licensec.Flags(); f != nil {
		f.StringVarP(&flags.Owner, "owner", "", flags.Owner, "the license owner")
		f.StringVarP(&flags.Year, "year", "", flags.Year, "the license year")
		f.BoolVarP(&flags.Header, "header", "", flags.Header, "whether to generate a file header")
	}
}
