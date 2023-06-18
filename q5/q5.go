package q5

import "strings"

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

func IsPalindrome(s string) bool {
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	for i := 0; i < len(s)/2; i++ {
		for p := len(s) / 2; p > 0; p-- {
			if s[p] != s[i] {
				return false
			}
		}
	}
	return true
}
