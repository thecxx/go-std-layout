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

package handler

import (
	"context"
	"fmt"

	"github.com/thecxx/go-std-layout/examples/demo/pkg/port/http"
)

type PortFlags struct {
	*GlobalFlags
	Port int
}

func OnPortHandler(ctx context.Context, flags *PortFlags, args []string) (err error) {
	start, stop, shutdown := http.Listen(fmt.Sprintf(":%d", flags.Port))

	// Startup http server
	go func() {
		start()
	}()
	// Stop http server
	defer func() {
		stop()
	}()

	fmt.Printf("Who: %s\n", flags.Who)

	fmt.Printf("Access: http://127.0.0.1:%d/welcome\n", flags.Port)

	fmt.Printf("Directive:\n")
	fmt.Printf("    shutdown: Shutdown the server\n")
	fmt.Printf("\n")

	for {
		var (
			directive string
		)
		fmt.Printf("> ")
		_, err = fmt.Scan(&directive)
		if err != nil {
			fmt.Printf("Input error: %s\n", err.Error())
			continue
		}
		switch directive {
		case "shutdown":
			// Shutdown http server
			shutdown(ctx)
			fmt.Printf("Bye!!!\n")
			return
		default:
			fmt.Printf("Unsupported directive: %s\n", directive)
		}
	}

}
