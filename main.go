/**
 * @author tsukiyo
 * @date 2023-08-01 23:34
 */

package main

import (
	"lazywoo/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
