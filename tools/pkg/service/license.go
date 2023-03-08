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
	"context"
	"fmt"
	"os"

	"github.com/thecxx/go-std-layout/tools/pkg/common"
	"github.com/thecxx/go-std-layout/tools/pkg/internal"
)

type license interface {
	ValidateLicense(ctx context.Context, lic string) (err error)
	GenerateHeader(ctx context.Context, gp *Project, lic string, year int, owner string) (err error)
	GenerateLicense(ctx context.Context, gp *Project, lic string, year int, owner string) (err error)
}

type licenseImpl struct {
}

// ValidateLicense implements license
func (*licenseImpl) ValidateLicense(ctx context.Context, lic string) (err error) {
	if _, ok := common.Licenses[lic]; !ok {
		err = fmt.Errorf("license %s not supported", lic)
	}
	return
}

// GenerateHeader implements license
func (l *licenseImpl) GenerateHeader(ctx context.Context, gp *Project, lic string, year int, owner string) (err error) {
	return os.WriteFile(gp.License.Header, []byte(internal.GenerateApacheLicense2Header(year, owner)), 0644)
}

// GenerateLicense implements license
func (l *licenseImpl) GenerateLicense(ctx context.Context, gp *Project, lic string, year int, owner string) (err error) {
	return os.WriteFile(gp.License.Description, []byte(internal.GenerateApacheLicense2(year, owner)), 0644)
}

var (
	License license = &licenseImpl{}
)
