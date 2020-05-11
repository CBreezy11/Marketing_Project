package data

import (
	"fmt"
	"strconv"
	"strings"
)

func splitDollarString(s string) int {
	n := strings.Split(s,"$")
	v := strings.Split(n[1], ".")[0]
	x, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(err)
	}	
	return x
}

func dollarToString(s int) string {
	v := "$"+strconv.Itoa(s)+".00"
	return v
}

func convertToNumber(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

func getAverages(total, length int) int {
	return total / length
}

func convertIntToString(i int) string {
	v := strconv.Itoa(i)
	return v
}

func percentageToString(s int) string {
	v := strconv.Itoa(s)+"%"
	return v
}