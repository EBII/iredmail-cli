// Copyright © 2018 Christian Nolte
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

package cmd

import (
	"errors"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/drlogout/iredmail-cli/iredmail"
	"github.com/spf13/cobra"
)

// aliasAddCmd represents the 'alias add' command
var aliasAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an alias",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires alias email")
		}

		var err error

		if !govalidator.IsEmail(args[0]) {
			return fmt.Errorf("Invalid alias email format: \"%v\"", args[0])
		}
		args[0], err = govalidator.NormalizeEmail(args[0])

		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		server, err := iredmail.New()
		if err != nil {
			fatal("%v\n", err)
		}
		defer server.Close()

		err = server.AliasAdd(args[0])
		if err != nil {
			fatal("%v\n", err)
		}

		success("Successfully added alias %v\n", args[0])
	},
}

func init() {
	aliasCmd.AddCommand(aliasAddCmd)

	aliasAddCmd.SetUsageTemplate(usageTemplate("alias add [ALIAS_EMAIL]"))
}
