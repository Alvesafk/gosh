package utils

import (
	"os"
	"os/user"
	"log"
	"strings"
)

func GetUserHomeDir() string {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return home_dir
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}

	return hostname
}

func GetWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	return wd
}

func GetUser() *user.User {
	u, err := user.Current()
	if err != nil {
		log.Fatal("User not found!")
		return nil
	}

	return u
}

func ReverseStringArray(input []string) []string {
	result := input
	for i, s := 0, len(input)-1; i < s; i, s = i+1, s-1 {
		result[i], result[s] = result[s], result[i]
	}

	return result
}

func GetFormattedWorkingDirectory(depth int) string {
	wd_array := ReverseStringArray([]string{GetWorkingDirectory()})
	wd_formatted := strings.Split(wd_array[0], "/")

	if depth < 1 || len(wd_formatted) <= 2 {
		return GetWorkingDirectory()
	}

	var result string

	for i := depth; i > 0; i-- {
		result += wd_formatted[len(wd_formatted)-i] + "/"
	}

	return result
}
