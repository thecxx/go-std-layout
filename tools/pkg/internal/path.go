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
	"io/ioutil"
	"os"
)

func IsFile(name string) (bool, error) {
	i, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return !i.IsDir(), nil
}

func IsDir(name string) (bool, error) {
	i, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return i.IsDir(), nil
}

func TryReadFile(filename string) string {
	b, err := IsFile(filename)
	if err != nil && !b {
		return ""
	}
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(buf)
}
