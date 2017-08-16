package utils

import "log"

func CheckErr(erro error) {
	if erro != nil {
		log.Fatal(erro)
	}
}
