package main

// Importando o pacote utils dentro de GoExpert (que é o nome do módulo -> go mod init GoExpert)
import (
	"GoExpert/utils"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// Chamando a função VerificaModulo11 do pacote utils
	fmt.Println("O número a seguir possuí o módulo 11, como: ", utils.VerificaModulo11("5478235482"))

	// fmt.Println(utils.padLeft("12345", 11, "0")) // Não é possível chamar essa função porque ela é privada (inicia com letra minúscula)
	// fmt.Println(utils.onlyZero("00000000000")) // Não é possível chamar essa função porque ela é privada (inicia com letra minúscula)

	// Funções, variáveis, structs, constantes e tipos que iniciam com letra maiúscula são públicos
	// e podem ser acessados por outros pacotes (módulos),
	// já os que iniciam com letra minúscula são privados e só podem ser acessados dentro do pacote

	// Testando o pacote color
	// Para importar um pacote de terceiros, basta usar o comando go get -u github.com/fatih/color ou simplesmente colocar a url do pacote no import
	// O parâmetro -u serve para atualizar o pacote

	color.Blue("CRUZEIROOOOO")
}
