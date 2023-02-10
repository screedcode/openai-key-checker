package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "USAGE:openai-key-checker FILE")
		os.Exit(1)
	}

	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open file: %s, %v", os.Args[1], err)
		os.Exit(1)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		key := strings.TrimSpace(s.Text())
		client := gpt3.NewClient(key)
		_, err := client.Engines(context.Background())
		if err == nil {
			fmt.Println(key)
		}
	}
}
