package main

import (
	"crypto/sha256"
	"os"
	"fmt"
	"strings"
	"os/exec"
)


func getSha256Hash(str string)(hash string){
	sum := sha256.Sum256([]byte(str))
	return string(sum[:])
}

func getIp()(ipAddress string, err error){

	choice := os.Getenv("TO_RUN")
	if choice == "localhost"{
		ipAddress = "localhost"
	}else{
	curl := exec.Command("curl", "ifconfig.me")
    out, err := curl.Output()
        if err != nil {
			return "" , err
		}
		ipAddress = string(out)
	}
	fmt.Println("ip address  = " , ipAddress)
	return ipAddress, nil
}

func extrapolateByteSlice(data []byte, kv map[string]string)([]byte){
	stringToExtrapolate := string(data)
	for key , val := range kv{
		stringToExtrapolate = strings.Replace(stringToExtrapolate,key, val , -1)
	}
    return []byte(stringToExtrapolate)
}

