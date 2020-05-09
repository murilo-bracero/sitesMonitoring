package lib

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//ReadSitesFromFile reads all sites of sites.txt and return them in a string array
func ReadSitesFromFile() []string {
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error while reading sites file")
		fmt.Println(err)
		os.Exit(-1)
	}

	cursor := bufio.NewReader(file)

	var sites []string

	for {
		site, err := cursor.ReadString('\n')
		site = strings.TrimSpace(site)

		sites = append(sites, site)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

//TestSite test and print sites in sites.txt
func TestSite(url string) {

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error while calling url")
		fmt.Println(err)
		os.Exit(-1)
	}

	if res.StatusCode == 200 {
		fmt.Println("site", url, "carregado com sucesso!")
		WriteLog(url, true)
	} else {
		fmt.Println("site", url, "com problemas. Status Code:", res.StatusCode)
		WriteLog(url, false)
	}
}
