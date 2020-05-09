package scrape

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetQueries() string {
	var showQuery string
	fmt.Print("Enter headliner: ")
	in := bufio.NewScanner(os.Stdin)

	if in.Scan() {
		showQuery = strings.ToLower(in.Text())
	}

	if err := in.Err(); err != nil {
		os.Exit(2)
	}

	return showQuery
}