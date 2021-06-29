package main

import (
	"bufio"
	"fmt"
	"github.com/nyudlts/go-aspace"
	"os"
	"strconv"
	"strings"
)

var (
	err    error
	client *aspace.ASClient
)

func init() {
	client, err = aspace.NewClient("/etc/sysconfig/go-aspace.yml", "dev", 2)
	if err != nil {
		panic(err)
	}
}

type Corp struct {
	ID          int
	AuthorityID string
}

func main() {
	updates := make(map[int]string)
	tsv, err := os.Open("AuthorityIDOnly.tsv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(tsv)
	scanner.Scan()
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "\t")
		i, _ := strconv.Atoi(split[0])
		updates[i] = split[10]
	}

	for k, _ := range updates {
		agent, err := client.GetAgent("corporate", k)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(agent)
	}
}
