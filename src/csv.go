package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func writeToCSV (fileName string, data []string)(err error){
	f, err := os.OpenFile(fmt.Sprintf("%s.csv",fileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error in opening file . Err : " , err)
		return
	}
	w := csv.NewWriter(f)
	w.Write(data)
	w.Flush()
	return
}

func findDataCSV (fileName , userName string ) (password string , err error){
	recordFile, err := os.Open(fmt.Sprintf("%s.csv",fileName) )
	if err != nil {
		log.Println("An error encountered in opeing the file ::", err)
		return
	}
	// Setup the reader
	reader := csv.NewReader(recordFile)

	// Read the records
	allRecords, err := reader.ReadAll()
	if err != nil {
		log.Println("Error in reading all records. Err :", err)
		return
	}

	isPresent := false
	password = ""

	for num , row := range allRecords{
		if num == 0 {
			continue
		}
		if row[0] == userName{
			isPresent = true
			password = row[1]
		}
	}

	if !isPresent{
		err = fmt.Errorf(userNotFoundError)
	}

	return
}


func createCSVFile(fileName string )(err error){
		_, err = os.Stat(fmt.Sprintf("%s.csv",fileName))

		if os.IsNotExist(err) {
			_, err = os.Create(fmt.Sprintf("%s.csv",fileName))
			if err != nil{
				log.Println("error in os.create ")
				return err
			}
			err = writeToCSV("USER-DATA" , []string{"USERNAME" , "PASSWORD"});
			if err != nil{
				log.Println("Error in writing headers." )
				return err
			}
		}
	return 
}
