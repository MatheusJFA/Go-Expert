package main

import (
	"GoExpert/utils"
	"os"
)

func main() {
	nomeArquivo := "file.txt"
	file := utils.CreateFile(nomeArquivo)

	utils.WriteFile(file, "Hello, World!")
	println(utils.ReadFile(nomeArquivo))

	file.Close()

	file = utils.OpenFile(nomeArquivo)
	utils.WriteFile(file, "Matheus ama a Pâmella")
	println(utils.ReadFile(nomeArquivo))

	file, err := os.Open(nomeArquivo)

	if err != nil {
		println(err)
	}

	println("Lendo arquivo com buffer:")
	utils.ReadFileBuffer(file)

	exists := utils.FileExists(nomeArquivo)

	if exists {
		println("Deletando arquivo...")
		utils.DeleteFile(nomeArquivo)
	}
}
