package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	ip, err := getIp()
	if err != nil {
		log.Println("Error in getting ip. Err : ", err)
	}

	dataMap := &fileNameMap{}
	dataMap.fileByteMap = make(map[string]([]byte))
	extrapolationMap := make(map[string]string)
	extrapolationMap[ipExtrapolationVariable] = ip

	dataMap.fileByteMap, err = getFilesExtrapolatedFileMap("static", extrapolationMap)
	if err != nil {
		log.Println("Error while getting extrapolated file map. Err : ", err)
	}

	err = createCSVFile(userPasswordDataStore)
	if err != nil {
		log.Println("Error in creating CSV file . Err : ", err)
		return
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/Login", dataMap.Login)
	http.HandleFunc("/SignUp", dataMap.SignUp)
	http.HandleFunc("/Review", dataMap.Review)
	http.HandleFunc("/Movie", dataMap.Movie)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!", r.URL.Path[1:])
}

func (dataMap *fileNameMap) Login(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Write(dataMap.fileByteMap["index.html"])
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		pWord, err := findDataCSV(userPasswordDataStore, username)
		if err != nil {
			if err.Error() == userNotFoundError {
				w.Write([]byte("User not found\n"))
				return
			} else {
				log.Println("Error in finding CSV data. Err : ", err)
				return
			}
		}

		if getSha256Hash(password) != pWord {
			w.Write([]byte("wrong password\n"))
			return
		}

		w.Write(dataMap.fileByteMap["home.html"])
	default:
		w.Write([]byte("Wrong API call"))
	}
}

func (dataMap *fileNameMap) SignUp(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Write(dataMap.fileByteMap["signup.html"])
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := findDataCSV(userPasswordDataStore, username)
		if err == nil {
			w.Write([]byte("User already exists\n"))
			return
		}

		err = writeToCSV(userPasswordDataStore, []string{username, getSha256Hash(password)})
		if err != nil {
			log.Println("Error in writing headers. Err : ", err)
			return
		}

		w.Write(dataMap.fileByteMap["index.html"])

	default:
		w.Write([]byte("Wrong API call"))

	}

}

func (dataMap *fileNameMap) Review(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Write(dataMap.fileByteMap["review.html"])
	default:
		w.Write([]byte("Wrong API call"))

	}
}

func (dataMap *fileNameMap) Movie(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		w.Write(dataMap.fileByteMap["home.html"])
	default:
		w.Write([]byte("Wrong API call"))

	}
}
