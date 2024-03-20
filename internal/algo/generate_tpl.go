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
	generateTplCmd.Flags().StringVarP(&fileName, "name", "n", "main", "template's file name")
	generateTplCmd.Flags().StringVarP(&funcName, "func", "f", "solve", "template's func name")
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

var (
	fileName string
	funcName string
)

func GenerateTpl() {
	tpl := template.New("algoTpl")
	tpl, _ = tpl.Parse(algoTpl)
	args := map[string]string{
		"func": funcName,
	}

	outFile, err := os.Create(fmt.Sprintf("%v.go", fileName))
	if err != nil {
		log.Fatalf("generate file failed: %v", err)
	}
	defer outFile.Close()

	if err := tpl.Execute(outFile, args); err != nil {
		log.Fatalf("generate template into file failed: %v", err)
	}
}
