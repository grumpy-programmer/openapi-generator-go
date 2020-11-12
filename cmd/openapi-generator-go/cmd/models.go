/*
Copyright © 2020 Tino Rusch

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/contiamo/openapi-generator-go/pkg/generators/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// modelsCmd represents the models command
var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "generate a models",
	Long:  `generate a models.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("spec")
		outputDirectory, _ := cmd.Flags().GetString("output")
		packageName, _ := cmd.Flags().GetString("package-name")
		bs, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal().Str("spec-file", file).Err(err).Msg("failed to read the spec file")
		}
		err = os.MkdirAll(outputDirectory, 0755)
		if err != nil {
			log.Fatal().Str("output", outputDirectory).Err(err).Msg("failed to create output folder")
		}
		reader := bytes.NewReader(bs)
		err = models.GenerateModels(reader, outputDirectory, models.Options{
			PackageName: packageName,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("failed to generate models")
		}
		reader = bytes.NewReader(bs)
		err = models.GenerateEnums(reader, outputDirectory, models.Options{
			PackageName: packageName,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("failed to generate enums")
		}
	},
}

func init() {
	modelsCmd.Flags().Bool("fail-no-group", false, "fail when there is no x-handler-group defined for any of the endpoints")
	modelsCmd.Flags().Bool("fail-no-operation-id", false, "fail when there is no operationId defined for any of the methods")
	generateCmd.AddCommand(modelsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
