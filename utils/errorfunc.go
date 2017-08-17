package utils

import "log"

func CheckErr(erro error, msg ...string) {
	if erro != nil {
		log.Fatal(erro, msg)
	}
}
