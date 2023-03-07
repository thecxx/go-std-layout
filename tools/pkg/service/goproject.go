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

package service

import "github.com/thecxx/go-std-layout/tools/pkg/internal"

type License struct {
	Header      string
	Description string
}

type GoProject struct {
	Module    string
	Workspace string
	License   License
}

func GetGoProject(ws string) (gp *GoProject, err error) {
	gp = &GoProject{
		Workspace: ws,
	}
	if err = gp.init(); err != nil {
		return nil, err
	}
	return
}

func (gp *GoProject) init() (err error) {
	if gp.Module, err = internal.FindModulePath(gp.Workspace); err != nil {
		return err
	}
	return
}
