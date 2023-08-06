/**
 * @author tsukiyo
 * @date 2023-08-03 1:22
 */

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tsukaychan/lazywoo/internal/timer"
	"log"
	"strconv"
	"strings"
	"time"
)

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "The time to be calculated, valid units are timestamp or formatted time")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "Duration, effective Unit of time is \"ns\", \"us\" (or\"μS\"), \"ms\", \"s\", \"m\", \"h\"")
}

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Time format processing",
	Long:  "Time format processing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "Get current time",
	Long:  "Get current time",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := time.Now()
		log.Printf("current time is: %s, %d", nowTime.Format(time.DateTime), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculate the required time",
	Long:  "Calculate the required time",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = time.DateTime
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = time.DateOnly
			} else if space == 1 {
				layout = time.DateTime
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("output result: %s， %d", t.Format(layout), t.Unix())
	},
}
