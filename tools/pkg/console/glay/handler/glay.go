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

package handler

import (
	"context"
	"fmt"

	"github.com/thecxx/go-std-layout/tools/pkg/internal"
)

type GlobalFlags struct {
	Workspace string
}

func (f GlobalFlags) GetWorkspace() (string, error) {
	if f.Workspace == "" {
		return "", fmt.Errorf("workspace not found")
	}
	b, err := internal.IsDir(f.Workspace)
	if err != nil {
		return "", err
	}
	if !b {
		return "", fmt.Errorf("invalid workspace")
	}
	return f.Workspace, nil
}

type GlayFlags struct {
	*GlobalFlags
}

func OnGlayHandler(ctx context.Context, flags *GlayFlags, args []string) (err error) {
	panic("not implemented")
}
