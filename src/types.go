package main

const userNotFoundError="DATA NOT FOUND"
const userPasswordDataStore="USER-DATA"
const ipExtrapolationVariable="${IP_ADDRESS}"

type fileNameMap struct{
	fileByteMap map[string]([]byte)
}
