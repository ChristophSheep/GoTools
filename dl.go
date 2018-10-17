/*

https://readwrite.com/2013/10/02/github-for-beginners-part-2/

> git remote add origin https://github.com/username/myproject.git
> git remote -v

git remote add origin https://github.com/MySheep/GoTools.git
git push -u origin master

*/
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func getTs(i int) string {
	return "https://cs6.rbmbtnx.net/v1/STV/s/1/X3/QT/64/8H/5N/11/" + strconv.Itoa(i) + ".ts"
}

func getCurlCmd(i int) string {
	return "curl " + getTs(i) + " --output " + strconv.Itoa(i) + ".ts"
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

func dlScript(i, j int) {

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
}

func joinScript(i, j int) {

	fmt.Printf("Create jn.bat for i=%d, j=%d\n", i, j)

	// Open dl.bat file
	file, err := os.Create("jn.bat")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Println(getJoinCmd(i, j))
	fmt.Fprintln(file, getJoinCmd(i, j))
}

func removeScript(i, j int) {

	fmt.Printf("Create rm.bat for i=%d, j=%d\n", i, j)

	// Open dl.bat file
	file, err := os.Create("rs.bat")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	tss := createTss(i, j)

	rmCmd := "rm"
	if runtime.GOOS == "windows" {
		rmCmd = "del"
	}

	for _, ts := range tss {
		fmt.Fprintln(file, rmCmd+" "+ts)
	}
}

func main() {

	i, err := strconv.Atoi(os.Args[1])
	j, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal("Cannot parse args", err)
	}

	dlScript(i, j)
	joinScript(i, j)
	removeScript(i, j)
}
