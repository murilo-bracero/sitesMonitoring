package main

import (
	"fmt"
	"os"
	"time"

	"sitesMonitoring/lib"
)

const monitorings = 3
const delay = 5

func showIntro() {
	name := "Carlos"
	version := 1.1

	fmt.Println("Olá, sr", name)
	fmt.Println("version ", version)
}

func showMenu() {

	fmt.Println("==:- Menu -:==")
	fmt.Println("1 - Iniciar Monitoramento dos sites")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func receiveCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func monitoring() {

	sites := lib.ReadSitesFromFile()

	for i := 0; i < monitorings; i++ {
		fmt.Println("Teste rodada", i+1)
		fmt.Println("")

		for i, site := range sites {
			fmt.Println("Testando site", i, site)
			lib.TestSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func main() {
	showIntro()

	for {
		showMenu()
		command := receiveCommand()

		switch command {
		case 1:
			monitoring()
		case 2:
			lib.ShowLogs()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço este comando")
			os.Exit(-1)
		}
	}
}
