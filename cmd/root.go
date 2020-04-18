/*
Copyright © 2020 MAGNUS FURUGÅRD <magnus.furugard@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/magnusfurugard/texp/parse"
	"github.com/magnusfurugard/texp/printer"

	"github.com/spf13/cobra"
)

// Flags
var (
	tokens       *[]string
	outputFormat *string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   `texp "<token-string>" [flags]`,
	Short: "Token-based string expansion",
	Long:  `Create multiple strings out of a single, token-filled one.`,
	Example: `  texp "@a@b" -t @a=x -t @a=y -t @b=1
  texp "@a@a@b" -t @a=1 -t @a=2 @a=3 -t @b=5
  texp "Good @time, how are you @person?" \
    -t @time=day \
    -t @time=evening \
    -t @person=friend \
    -t @person=pal \
    -t @person=''`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if fi, _ := os.Stdin.Stat(); fi.Size() > 0 {
			scanner := bufio.NewScanner(os.Stdin)
			text := ""
			for scanner.Scan() {
				text += scanner.Text() + "\n"
			}
			text = strings.Trim(text, "\n")
			if len(text) > 0 {
				args = append(args, text)
			}
		}

		if len(args) > 1 {
			return fmt.Errorf("expected 1 token string, got %v: %v", len(args), args)
		}

		if len(args) != 1 {
			cmd.Help()
			return nil
		}

		tokens, result := parse.Tokens(args[0], tokens)
		err := printer.Print(*outputFormat, tokens, result)
		if err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	tokens = rootCmd.PersistentFlags().StringArrayP("token", "t", []string{}, "tokens to replace with format key=value")
	outputFormat = rootCmd.PersistentFlags().StringP("output", "o", "json", "output format: json, yaml and raw")
}
