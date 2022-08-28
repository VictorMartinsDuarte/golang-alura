package main

import "fmt"
// import "reflect"
import "os"
import "net/http"

func main() {

	exibeIntroducao()

	// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))

	exibeMenu()

	comando := leComando()

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa.")
	// } else {
	// 	fmt.Println("Não conheço este comando.")
	// }

	switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa.")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Victor"
	// var versao float32 = 1.1
	versao := 1.1
	fmt.Println("Olá, Sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento.")
	fmt.Println("2- Exibir logs.")
	fmt.Println("0- Sair do programa.")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	// fmt.Println("O endereço da minha variável comando é", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	// resp, err := http.Get(site)

}