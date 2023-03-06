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

func GetApacheLicense2(year int, owners ...string) (tpl string) {
	if len(owners) <= 0 {
		panic("owner no found")
	}
	tpl += fmt.Sprintf("// Copyright %d %s\n", year, strings.Join(owners, ", "))
	tpl += "//\n"
	tpl += "// Licensed under the Apache License, Version 2.0 (the \"License\");\n"
	tpl += "// you may not use this file except in compliance with the License.\n"
	tpl += "// You may obtain a copy of the License at\n"
	tpl += "//\n"
	tpl += "//     http://www.apache.org/licenses/LICENSE-2.0\n"
	tpl += "//\n"
	tpl += "// Unless required by applicable law or agreed to in writing, software\n"
	tpl += "// distributed under the License is distributed on an \"AS IS\" BASIS,\n"
	tpl += "// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n"
	tpl += "// See the License for the specific language governing permissions and\n"
	tpl += "// limitations under the License.\n"
	return
}
