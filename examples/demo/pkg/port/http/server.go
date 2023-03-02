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

package http

import (
	"context"
	httpPkg "net/http"

	"github.com/thecxx/go-std-layout/examples/demo/pkg/port/http/handler"
)

func Listen(addr string) (start, stop func() error, shutdown func(context.Context)) {
	mux := httpPkg.ServeMux{}
	srv := httpPkg.Server{Addr: addr, Handler: &mux}

	mux.Handle("/welcome", httpPkg.HandlerFunc(handler.OnWelcomeHandler))

	start = func() error { return srv.ListenAndServe() }
	stop = func() error { return srv.Close() }
	shutdown = func(ctx context.Context) { srv.Shutdown(ctx) }

	return
}
