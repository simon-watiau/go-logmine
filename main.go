package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/simon-watiau/go-logmine/clustering"
	"github.com/simon-watiau/go-logmine/tokenizer"
)

func main() {
	clusterSet := clustering.NewClusterAggregate(
		context.Background(),
		[]float64{
			0.01,
			0.1,
			0.3,
		})

	readFile, err := os.Open("good.log")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		clusterSet.AddLog(tokenizer.NewTokenizedLogFromRawString(fileScanner.Text()))
	}

	readFile.Close()

	logs := clusterSet.Aggregate()

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].Weight > logs[j].Weight
	})

	total := 0
	for _, c := range logs {
		fmt.Println(fmt.Sprint(c.Weight) + ";" + fmt.Sprint(c.Pattern))
		total += c.Weight
	}

	fmt.Println(total)
}
