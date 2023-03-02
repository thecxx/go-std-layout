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

package service

import (
	"context"
	"fmt"

	"github.com/thecxx/go-std-layout/examples/demo/pkg/dao/fs"
)

func WelcomeMessage(ctx context.Context) (ret string) {
	return fmt.Sprintf(`
<pre>
Welcome to demo
File: %s
Size: %d
</pre>`, fs.GetExecutorName(), fs.GetExecutorSize())
}
