package main

import (
	"io/ioutil"
	"fmt"
	"log"
)


func getFileData(fileName string)(fileData []byte, err error){

	fileData, err  = ioutil.ReadFile(fileName)
	if err !=  nil{
		log.Println("Error in ioutil.ReadFile")
		return nil, err
	}
	return

}



func getExtrapolatedFileData(fileName string, extratapolationMap map[string]string)(fileData []byte, err error){

		data , err := getFileData(fileName)
		if err != nil{
			log.Println("Error in getting file data")
			return nil , err
		}

		extrapolatedFileData := extrapolateByteSlice(data , extratapolationMap)

		return extrapolatedFileData, nil
}


func getFilesExtrapolatedFileMap(folderName string, extrapolationMap map[string]string)  (fileByteMap map[string]([]byte), err error) {

	files, err := ioutil.ReadDir(fmt.Sprintf("./%s",folderName))
	if err != nil{
		log.Println("Error in reading directory")
		return nil, err
	}

	fileByteMap = make(map[string]([]byte))
	for _, itr := range files {

		if (itr.IsDir()) || (len(itr.Name()) > 5 && itr.Name()[len(itr.Name())-5:len(itr.Name())] != ".html"){
			continue
		}

		fileDataByte, err := getExtrapolatedFileData(fmt.Sprintf("./%s/%s",folderName, itr.Name()), extrapolationMap)
		if err != nil{
			log.Println("Error in extrapolating file data")
			return nil, err
		}
	fileByteMap[itr.Name()] = fileDataByte
	}

	return
}
