package entity

const (
	symbolsRegex = "[^0-9]"
)

type CPF struct {
	value string
}

func SameNumber(cpf string) bool {
	for i := 0; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}

	return true
}
