package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"osapp/cmd"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "oss",
		Short: "this is a simple oss cli",
		Long: `A Fast and Flexible oss cli in Go.
                Complete documentation is available at liudong11`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			log.Infof("root cmd")
		},
	}

	cmd.RegisterOSSCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
