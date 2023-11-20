package helpers

import "log"

func PanicIfErrSystem(err error) {
	if err != nil {
		log.Default().Panic(err)
	}
}
