//Copyright © 2022 Ugo Landini <ugo.landini@gmail.com>
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/ugol/jr/pkg/constants"
)

var templateShowCmd = &cobra.Command{
	Use:   "show [template]",
	Short: "Show a template",
	Long:  `Show a template. Templates must be in templates directory, which is '$JR_SYSTEM_DIR/templates'`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Error().Msg("Template missing. Try the list command to see available templates")
			os.Exit(1)
		}

		noColor, _ := cmd.Flags().GetBool("nocolor")
		templateDir := os.ExpandEnv(fmt.Sprintf("%s/%s", constants.JR_SYSTEM_DIR, "templates"))
		templatePath := fmt.Sprintf("%s/%s.tpl", templateDir, args[0])
		templateScript, err := os.ReadFile(templatePath)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to ReadFile")
		}
		valid, err := isValidTemplate([]byte(templateScript))
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to read a template")
		}
		templateString := string(templateScript)

		var Reset = "\033[0m"
		if runtime.GOOS != "windows" && !noColor {
			var Cyan = "\033[36m"
			coloredOpeningBracket := fmt.Sprintf("%s%s", Cyan, "{{")
			coloredClosingBracket := fmt.Sprintf("%s%s", "}}", Reset)
			templateString = strings.ReplaceAll(templateString, "{{", coloredOpeningBracket)
			templateString = strings.ReplaceAll(templateString, "}}", coloredClosingBracket)
		}
		fmt.Println(templateString)
		fmt.Print(Reset)
		if !valid {
			log.Fatal().Msg("Invalid template")
		}
	},
}

func init() {
	templateCmd.AddCommand(templateShowCmd)
	templateShowCmd.Flags().BoolP("nocolor", "n", false, "Do not color output")
}
