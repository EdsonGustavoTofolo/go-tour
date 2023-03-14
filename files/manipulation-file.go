package files

import (
	"bufio"
	"fmt"
	"os"
)

func ManipulateFile() {
	writeFile()
	readFile()
	readPartialFile()
	removeFile()
	checkIfFileExists()
}

func checkIfFileExists() {
	if _, err := os.Stat("arquivo.txt"); os.IsNotExist(err) {
		fmt.Println("arquivo.txt not exists")
	} else {
		fmt.Println("arquivo.txt exists")
	}
}

func removeFile() {
	err := os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}

func readPartialFile() {
	arquivo, err := os.Open("arquivo_grande.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo)

	buffer := make([]byte, 10)

	for {
		// irá ler somente o tamanho especificado no buffer, ou seja, vai ler de 10 em 10 bytes
		// isso serve principalmente para arquivos muito grandes, assim, nao sobrecarrega memeory
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	arquivo.Close()
}

func readFile() {
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))
}

func writeFile() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	//tamanho, err := f.WriteString("Hello World")

	tamanho, err := f.Write([]byte("Escrevendo bytes no arquivo"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso. Tamanho %d bytes\n", tamanho)

	f.Close()
}
