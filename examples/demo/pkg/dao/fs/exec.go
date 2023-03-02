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

package fs

import (
	"os"
)

var (
	er error
	ex *os.File
)

func GetExecutorName() (name string) {
	if f, err := lazyLoad(); err == nil {
		name = f.Name()
	}
	return
}

func GetExecutorSize() (size int64) {
	if f, err := lazyLoad(); err == nil {
		if info, err := f.Stat(); err == nil {
			size = info.Size()
		}
	}
	return
}

func lazyLoad() (*os.File, error) {
	if ex == nil && er == nil {
		ex, er = os.Open(os.Args[0])
	}
	return ex, er
}
