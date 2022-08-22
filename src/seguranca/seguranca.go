package seguranca

import "golang.org/x/crypto/bcrypt"

func Hash(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
}

func VerificarSenha(senhaString, senhaHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
