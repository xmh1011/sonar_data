package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xmh1011/sonar_data/pkg"
	"io"
	"os"
	"time"
)

type Options struct {
	Stderr io.Writer
	Stdout io.Writer

	startTime int64
	endTime   int64
	step      int64
}

func NewOptions() *Options {
	return &Options{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
}

var opt = NewOptions()

func Sonar_data() *cobra.Command {
	c := &cobra.Command{
		Use:   "generate",
		Short: "Generate sonar_data for benchmark test",
		Run: func(cmd *cobra.Command, args []string) {
			var startTime, endTime, step int64
			startTime = opt.startTime
			endTime = opt.endTime
			step = opt.step
			for i := startTime; i <= endTime; i = i + step {
				timestamp_1 := 1e9 * i
				timestamp_2 := 1e9*i + 1e8*2
				timestamp_3 := 1e9*i + 1e8*4
				timestamp_4 := 1e9*i + 1e8*6
				timestamp_5 := 1e9*i + 1e8*8
				fmt.Fprintf(os.Stdout, "sonar,sonar_id=%d velocity=%f,longitude=\"%fE\",latitude=\"%fN\",angle=%f,distance=%f,signal=%f %d\n",
					//pkg.RandRange(1,4) //submarine_id
					pkg.RandRange(1, 10),   //sonar_id
					pkg.RandValue(0, 18),   //velocity
					pkg.RandValue(0, 30),   //longitude
					pkg.RandValue(90, 130), //latitude
					pkg.RandValue(0, 360),  //angle
					pkg.RandValue(0, 100),  //distance
					pkg.RandValue(0, 180),  //signal
					timestamp_1,
				)
				fmt.Fprintf(os.Stdout, "sonar,sonar_id=%d velocity=%f,longitude=\"%fE\",latitude=\"%fN\",angle=%f,distance=%f,signal=%f %d\n",
					//pkg.RandRange(1,4) //submarine_id
					pkg.RandRange(1, 10),   //sonar_id
					pkg.RandValue(0, 18),   //velocity
					pkg.RandValue(0, 30),   //longitude
					pkg.RandValue(90, 130), //latitude
					pkg.RandValue(0, 360),  //angle
					pkg.RandValue(0, 100),  //distance
					pkg.RandValue(0, 180),  //signal
					timestamp_1,
				)
				fmt.Fprintf(os.Stdout, "sonar,sonar_id=%d velocity=%f,longitude=\"%fE\",latitude=\"%fN\",angle=%f,distance=%f,signal=%f %d\n",
					//pkg.RandRange(1,4) //submarine_id
					pkg.RandRange(1, 10),   //sonar_id
					pkg.RandValue(0, 18),   //velocity
					pkg.RandValue(0, 30),   //longitude
					pkg.RandValue(90, 130), //latitude
					pkg.RandValue(0, 360),  //angle
					pkg.RandValue(0, 100),  //distance
					pkg.RandValue(0, 180),  //signal
					timestamp_2,
				)
				fmt.Fprintf(os.Stdout, "sonar,sonar_id=%d velocity=%f,longitude=\"%fE\",latitude=\"%fN\",angle=%f,distance=%f,signal=%f %d\n",
					//pkg.RandRange(1,4) //submarine_id
					pkg.RandRange(1, 10),   //sonar_id
					pkg.RandValue(0, 18),   //velocity
					pkg.RandValue(0, 30),   //longitude
					pkg.RandValue(90, 130), //latitude
					pkg.RandValue(0, 360),  //angle
					pkg.RandValue(0, 100),  //distance
					pkg.RandValue(0, 180),  //signal
					timestamp_3,
				)
				fmt.Fprintf(os.Stdout, "sonar,sonar_id=%d velocity=%f,longitude=\"%fE\",latitude=\"%fN\",angle=%f,distance=%f,signal=%f %d\n",
					//pkg.RandRange(1,4) //submarine_id
					pkg.RandRange(1, 10),   //sonar_id
					pkg.RandValue(0, 18),   //velocity
					pkg.RandValue(0, 30),   //longitude
					pkg.RandValue(90, 130), //latitude
					pkg.RandValue(0, 360),  //angle
					pkg.RandValue(0, 100),  //distance
					pkg.RandValue(0, 180),  //signal
					timestamp_4,
				)
				fmt.Fprintf(os.Stdout, "sonar,sonar_id=%d velocity=%f,longitude=\"%fE\",latitude=\"%fN\",angle=%f,distance=%f,signal=%f %d\n",
					//pkg.RandRange(1,4) //submarine_id
					pkg.RandRange(1, 10),   //sonar_id
					pkg.RandValue(0, 18),   //velocity
					pkg.RandValue(0, 30),   //longitude
					pkg.RandValue(90, 130), //latitude
					pkg.RandValue(0, 360),  //angle
					pkg.RandValue(0, 100),  //distance
					pkg.RandValue(0, 180),  //signal
					timestamp_5,
				)
			}
		},
	}
	c.PersistentFlags().Int64Var(&opt.startTime, "startTime", time.Now().Unix(), "Set start time")
	c.PersistentFlags().Int64Var(&opt.endTime, "endTime", time.Now().Unix(), "Set end time")
	c.PersistentFlags().Int64Var(&opt.step, "step", 1, "Set the step")
	return c
}
