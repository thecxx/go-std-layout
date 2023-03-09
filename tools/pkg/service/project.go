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

import (
	"fmt"
	"path"
	"strings"

	"github.com/thecxx/go-std-layout/tools/pkg/internal"
)

type Project struct {
	GoMod     string
	Module    string
	Curplace  string
	Workspace string
	License   struct {
		Header      string
		Description string
	}
}

func NewProject(ws string) (gp *Project, err error) {
	gp = &Project{
		Workspace: ws,
	}
	if err = gp.initModule(); err != nil {
		return nil, err
	}
	if err = gp.initLicense(); err != nil {
		return nil, err
	}
	return
}

func (gp *Project) initModule() (err error) {
	if gp.GoMod, gp.Module, err = internal.FindModulePath(gp.Workspace); err != nil {
		return err
	}
	var (
		realws = path.Dir(gp.GoMod)
	)
	if realws != gp.Workspace {
		if !strings.HasPrefix(gp.Workspace, realws) {
			return fmt.Errorf("please re-execute it at the root of the project")
		}
		gp.Curplace, gp.Workspace = gp.Workspace, realws
	}
	return
}

func (gp *Project) initLicense() (err error) {
	gp.License.Header = fmt.Sprintf("%s/HEADER", gp.Workspace)
	gp.License.Description = fmt.Sprintf("%s/LICENSE", gp.Workspace)
	return
}
