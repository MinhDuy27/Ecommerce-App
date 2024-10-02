package helper

import (
	"crypto/rand"
	"log"
	"strconv"
)

func GenNumber(length int) (uint, error) {
	const number = "0123456789"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return 0,err
	}

	numlen := len(number)
	for i:=0 ;i < length; i++ {
		buffer[i] = number[int(buffer[i]) % numlen]
	}
	value,err := strconv.ParseUint(string(buffer),10,64)
	if err != nil {
		log.Fatalln(err)
	}
	return	uint(value),nil
}

