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
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

func FindModulePath(wd string) (string, error) {
	var (
		paths []string
		gomod = GoEnv("GOMOD")
	)
	if path.Base(gomod) == "go.mod" {
		return GetModulePathFromGoMod(gomod)
	}
	if path := GoEnv("GOPATH"); path != "" {
		paths = append(paths, path)
	}
	if path := GoEnv("GOROOT"); path != "" {
		paths = append(paths, path)
	}
	for _, path := range paths {
		if src := fmt.Sprintf("%s/src/", path); strings.HasPrefix(wd, src) {
			return strings.TrimPrefix(wd, src), nil
		}
	}
	return "", fmt.Errorf("module not found")
}

func GetModulePathFromGoMod(file string) (string, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	matches := regexp.MustCompile(`^module\s+([^\r\n]+)`).FindStringSubmatch(string(buf))
	if len(matches) <= 0 {
		return "", fmt.Errorf("failed to match")
	}
	return matches[1], nil
}
