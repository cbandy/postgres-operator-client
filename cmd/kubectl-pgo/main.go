// Copyright 2021 - 2023 Crunchy Data Solutions, Inc.
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

package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/crunchydata/postgres-operator-client/internal/cmd"
)

func main() {
	name := filepath.Base(os.Args[0])

	flags := pflag.NewFlagSet(name, pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewPGOCommand(name, os.Stdin, os.Stdout, os.Stderr)
	cobra.CheckErr(root.Execute())
}
