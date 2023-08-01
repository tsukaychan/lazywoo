/**
 * @author tsukiyo
 * @date 2023-08-01 23:35
 */

package cmd

import (
	"github.com/spf13/cobra"
	"lazywoo/internal/word"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "Please enter the word content")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "Please enter the format for word conversion")
}

var desc = strings.Join([]string{
	"This sub instruction supports various word format conversions, with the following modes: ",
	"1: Convert words to uppercase words",
	"2: Convert words to lowercase words",
	"3: Convert underlined words to uppercase words",
	"4: Convert underlined words to lowercase words",
	"5: Convert camel case words to underlined words",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word format conversion",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("The conversion format is currently not supported. Please execute 'help word' to view the help document")
		}
		log.Printf("output result: %s", content)
	},
}
