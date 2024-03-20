package algo

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

func init() {
	generateTplCmd.Flags().StringVarP(&title, "title", "t", "solve", "template's title")
}

var generateTplCmd = &cobra.Command{
	Use:   "algo",
	Short: "generate algo template for golang",
	Long:  "generate algo template for golang",
	Run: func(cmd *cobra.Command, args []string) {
		GenerateTpl()
	},
}

func Register(root *cobra.Command) {
	root.AddCommand(generateTplCmd)
}

//go:embed algo_tpl.txt
var algoTpl string

var title string

func GenerateTpl() {
	tpl := template.New("algoTpl")
	tpl, _ = tpl.Parse(algoTpl)
	args := map[string]string{
		"title": title,
	}

	outFile, err := os.Create(fmt.Sprintf("%v.go", title))
	if err != nil {
		log.Fatalf("generate file failed: %v", err)
	}
	defer outFile.Close()

	if err := tpl.Execute(outFile, args); err != nil {
		log.Fatalf("generate template into file failed: %v", err)
	}
}
