/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// waitCmd represents the wait command
var waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait for hosts healthcheks and execute an entrypoint",
	Long:  `Wait for hosts healthcheks and execute an entrypoint`,
	Run: func(cmd *cobra.Command, args []string) {

		// Extract flags to variable
		timeout, _ := cmd.Flags().GetInt("timeout")
		hosts, _ := cmd.Flags().GetString("host")
		entrypoint, _ := cmd.Flags().GetString("entrypoint")

		hostsArray := strings.Split(hosts, ",")

		// Validate hosts urls
		hostsArray = validateHosts(hostsArray)

		// Wait for hosts
		executeStatus := waitTimeout(hostsArray, timeout)

		// Execute entrypoint
		if executeStatus {
			executeEntrypoint(entrypoint)
		}
	},
}

func init() {
	rootCmd.AddCommand(waitCmd)

	// Here you will define your flags and configuration settings.
	waitCmd.PersistentFlags().String("host", "", "wait for host, expected URLs entrypoints, comma separated")
	waitCmd.MarkPersistentFlagRequired("host")

	waitCmd.PersistentFlags().Int("timeout", 15, "timeout in seconds, zero for no timeout")

	waitCmd.PersistentFlags().String("entrypoint", "", "entrypoint to execute")
	waitCmd.MarkPersistentFlagRequired("entrypoint")
}

func validateHosts(hosts []string) []string {
	validatedHosts := hosts[:0]
	for _, host := range hosts {
		_, err := url.ParseRequestURI(host)
		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			validatedHosts = append(validatedHosts, host)
		}
	}

	return validatedHosts
}

func checkHostHealthCheck(wg *sync.WaitGroup, host string) {
	defer wg.Done()
	for {
		resp, err := http.Get(host)

		if err != nil {
			fmt.Printf("%s\n", err)
			time.Sleep(2 * time.Second)
			continue

		} else if resp.StatusCode == 200 {
			fmt.Printf("%s return status code 200\n", host)
			break
		}
	}
}

func waitForHosts(ch chan bool, hosts []string) {
	wg := new(sync.WaitGroup)
	for _, host := range hosts {
		wg.Add(1)
		go checkHostHealthCheck(wg, host)
	}
	wg.Wait()
	ch <- true
}

func waitTimeout(hosts []string, timeout int) bool {

	var durationTimeout time.Duration

	if timeout == 0 {
		durationTimeout = time.Duration(86400) * time.Second // 24 hours
	} else {
		durationTimeout = time.Duration(timeout) * time.Second
	}

	fmt.Printf("Wait for hosts %v with %v timeout duration\n", hosts, durationTimeout)

	ch := make(chan bool)

	go waitForHosts(ch, hosts)

	select {
	case <-ch:
		return true // All host available
	case <-time.After(durationTimeout):
		fmt.Printf("Aborted. %s timeout over\n", durationTimeout)
		return false // Timeout gone
	}
}

func executeEntrypoint(entrypoint string) {
	// Arrange entrypoint for executing
	entrypointParts := strings.Split(entrypoint, " ")

	// Execute entrypoint
	cmd := exec.Command(entrypointParts[0], entrypointParts[1:]...)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
