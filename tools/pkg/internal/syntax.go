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

package internal

import (
	"fmt"
	"strings"
)

func Comment(source string) (tpl string) {
	var (
		i     = 0
		lines = strings.Split(source, "\n")
	)
	for ; i < len(lines)-1; i++ {
		tpl += fmt.Sprintf("// %s\n", lines[i])
	}
	tpl += fmt.Sprintf("// %s", lines[i])
	return
}

func InsertLicenseHeader(source, license string) (tpl string) {
	if license = strings.Trim(license, "\n\r\t "); license != "" {
		tpl += Comment(license)
		tpl += "\n\n"
	}
	tpl += source
	return
}
