package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <program> <host> <command> [option]")
		os.Exit(1)
	}

	host := os.Args[1]
	command := os.Args[2]
	option := ""
	if len(os.Args) >= 4 {
		option = os.Args[3]
	}

	switch command {
	case "show-shards":
		executeSSHCommand(host, "show-shards", option)
	case "show":
		executeSSHCommand(host, "show", "")
	case "copy-shard":
		executeSSHCommand(host, "copy-shard", option)
	default:
		fmt.Println("Invalid command:", command)
		os.Exit(1)
	}
}

func executeSSHCommand(host, command, option string) {
	influxMetaCommand := "/influxd-ctl -auth-type jwt -secret $(cat /influxdb/conf/influxdb-meta.conf | grep intern | awk '{print $3}' | sed 's/\"//g') " + command
	if command == "show-shards" && option != "" {
		influxMetaCommand += " " + option
	}

	sshCommand := "ssh"
	sshArgs := []string{
		"-i", "~/.ssh/influx-cloud.pem",
		"-oStrictHostKeyChecking=no",
		"core@" + host,
		"docker exec influxd-meta /bin/sh -c \"" + influxMetaCommand + "\"",
	}

	cmd := exec.Command(sshCommand, sshArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}
