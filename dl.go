package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getTs(i int) string {
	return "https://cs6.rbmbtnx.net/v1/STV/s/1/X3/QT/64/8H/5N/11/" + strconv.Itoa(i) + ".ts"
}

func getCurlCmd(i int) string {
	return "curl " + getTs(i) + " --output " + strconv.Itoa(i) + ".ts"
}

/*
see https://stackoverflow.com/questions/39868029/how-to-generate-a-sequence-of-numbers-in-golang

a := makeRange(10, 20)
fmt.Println(a)

[10 11 12 13 14 15 16 17 18 19 20]
*/
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func createTss(i, j int) []string {

	// TODO MAKE SMALLER
	//
	tss := make([]string, j-i+1)
	jj := 0

	for ii := i; ii <= j; ii++ {
		ts := strconv.Itoa(ii) + ".ts"
		tss[jj] = ts
		jj++
	}

	return tss
}

/*
see https://praxistipps.chip.de/ts-dateien-zusammenfuegen-so-gehts_29381
*/
func getJoinCmd(i, j int) string {

	// Create list
	// 1.ts
	// 10.ts
	tss := createTss(i, j)

	// Join strings into one
	// e.g. copy /b 1.ts+2.ts all.ts
	//
	ts := strings.Join(tss, "+")

	return "copy /b " + ts + " all.ts"
}

func main() {

	i, err := strconv.Atoi(os.Args[1])
	j, err := strconv.Atoi(os.Args[2])

	fmt.Printf("Create dl.bat for i=%d, j=%d\n", i, j)

	// Open dl.bat file
	file, err := os.Create("dl.bat")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	// Create curls .. "curl https://cs6.rbmbtnx.net/v1/STV/s/1/X3/QT/64/8H/5N/11/1.ts --output 1.ts"
	//
	for ii := i; ii <= j; ii++ {
		fmt.Println(getCurlCmd(ii))
		fmt.Fprintln(file, getCurlCmd(ii))
	}

	// Create join tss .. "copy /b 1.ts+2.ts all.ts""
	//
	fmt.Println(getJoinCmd(i, j))
	fmt.Fprintln(file, getJoinCmd(i, j))
}
