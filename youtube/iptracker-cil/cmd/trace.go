package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(traceCmd)
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace the IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Please provide IP to trace")
		}
	},
}

type IP struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func showData(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := IP{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)

	c.Println("DATA FOUND :")

	fmt.Printf("IP :%s\nCITY: %s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nTIMEZONE :%s\nPOSTAL: %s\n", data.IP, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response")
	}

	return responseByte
}
