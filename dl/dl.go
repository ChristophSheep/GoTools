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

func getTs(i int, tsBaseUrl string) string {
	//  e.g.:
	// "https://cs6.rbmbtnx.net/v1/STV/s/1/X3/QT/64/8H/5N/11/" + strconv.Itoa(i) + ".ts"
	//
	return tsBaseUrl + strconv.Itoa(i) + ".ts"
}

func getCurlCmd(i int, tsBaseUrl string) string {
	return "curl " + getTs(i, tsBaseUrl) + " --output " + strconv.Itoa(i) + ".ts"
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

func dlScript(i, j int, tsBaseUrl string) {

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
		curlCmd := getCurlCmd(ii, tsBaseUrl)
		fmt.Println(curlCmd)
		fmt.Fprintln(file, curlCmd)
	}
}

func joinScriptWin(i, j int) {

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

func joinScriptMac(i, j int) {
	// e.g. cat datei1 datei2 > ausgabedatei
	file, err := os.Create("jn.sh")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	tss := createTss(i, j)

	// Join strings into one
	// e.g. copy /b 1.ts+2.ts all.ts
	//
	ts := strings.Join(tss, " ")

	cmd := "cat " + ts + " > movie.ts"

	fmt.Fprintln(file, cmd)
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

	// To find tsBaseUrl:
	//  - Use FireFox
	//  - Start Video
	//  - Look at WebDeveloperTools "NetzwerkAnalyse" and search for
	//
	// State 	Method 	File 	Host														Type
	// -------- ------- ------- ----------------------------------------------------------- -----
	// 200 		GET 	1.ts 	https://cs6.rbmbtnx.net/v1/STV/s/1/X3/QT/64/8H/5N/11/.1ts 	mp2t

	tsBaseUrl := "https://cs6.rbmbtnx.net/v1/STV/s/1/X3/QT/64/8H/5N/11/" // Tulpen aus Ammerland #ts 1 937

	i, err := strconv.Atoi(os.Args[1])
	j, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal("Cannot parse args", err)
	}

	dlScript(i, j, tsBaseUrl)
	joinScriptWin(i, j)
	joinScriptMac(i, j)
	removeScript(i, j)
}
