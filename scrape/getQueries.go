package scrape

import (
	"bufio"
	"fmt"
	"os"
)

func GetQueries() string {
	var showQuery string
	fmt.Print("Enter headliner: ")
	in := bufio.NewScanner(os.Stdin)

	if in.Scan() {
		showQuery = in.Text()
	}

	if err := in.Err(); err != nil {
		os.Exit(2)
	}

	return showQuery
}