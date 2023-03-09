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
)

type GlobalFlags struct {
}

func NewGlobalFlags() (flags *GlobalFlags) {
	flags = &GlobalFlags{}
	return
}

type GlayFlags struct {
	*GlobalFlags
}

func NewGlayFlags(gflags *GlobalFlags) (flags *GlayFlags) {
	flags = &GlayFlags{GlobalFlags: gflags}
	return
}

func OnGlayHandler(ctx context.Context, flags *GlayFlags, args []string) (err error) {
	panic("not implemented")
}
