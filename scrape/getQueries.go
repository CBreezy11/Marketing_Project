package scrape

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

func RemoveShows(showList, idList []string) ([]string, []string) {
	var removes string
	var keep bool = true
	var answer string

	fmt.Println("Here is a list of all the shows from your search: \n")

	for i, s := range(showList) {
		fmt.Printf("%d: %s\n", i+1, s)
	}

	fmt.Println("\nDo you want to eliminate any of these shows from the data? Enter 'y' or 'n': ")
	fmt.Scanln(&answer)

	if strings.ToLower(answer) == "y" {
		for keep {
			fmt.Println("\nPlease enter the corresponding number of the show you wish to remove, these must be done one at a time: ")
			fmt.Scanln(&removes)	
			removesInt, err := strconv.Atoi(removes)

			if err != nil {
				fmt.Println("Please enter a valid number")
				continue
			}
			if removesInt < 1 || removesInt > len(showList) {
				fmt.Println("The number you entered doesn't match a show")
				continue
			}
			removed := showList[removesInt-1]
			showList = append(showList[:removesInt-1], showList[removesInt:]...)
			idList = append(idList[:removesInt-1], idList[removesInt:]...)
			fmt.Println(removed, "HAS BEEN REMOVED")
			fmt.Println("Remove another show? Enter 'y' or 'n': ")
			fmt.Scanln(&removes)
			if strings.ToLower(removes) == "n" {
				keep = false
			}
			for i, s := range(showList) {
				fmt.Printf("%d: %s\n", i+1, s)
			}
		}
	}
	return showList, idList
}