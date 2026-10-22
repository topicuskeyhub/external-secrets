/*
Copyright © The ESO Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// The binary entry point for the controller manager.
package main

import (
	"log/slog"
	"os"

	"github.com/KimMachineGun/automemlimit/memlimit"

	"github.com/external-secrets/external-secrets/cmd/controller"

	_ "github.com/external-secrets/external-secrets/pkg/register" // Register all providers and generators
)

// init sets GOMEMLIMIT from the cgroup memory limit, but only when the
// AUTOMEMLIMIT environment variable is set, so the default behavior is unchanged.
// AUTOMEMLIMIT takes a ratio in (0.0, 1.0] or "off". An explicit GOMEMLIMIT takes precedence.
func init() {
	if _, ok := os.LookupEnv("AUTOMEMLIMIT"); !ok {
		return
	}
	_, _ = memlimit.Set(memlimit.WithLogger(slog.Default()))
}

func main() {
	controller.Execute()
}
