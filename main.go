package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"io"
	"strconv"
)

const port int = 9090
const logFile string = "./application.log"

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	log.Println("path", r.URL.Path)
	log.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello world!")
}

func server() {
	http.HandleFunc("/", sayhelloName) //アクセスのルーティングを設定します。
	log.Println("start server at port:", strconv.Itoa(port))
	err := http.ListenAndServe(":" + strconv.Itoa(port), nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	logfile, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open " + logFile + ":" + err.Error())
	}
	defer logfile.Close()
	// io.MultiWriteで、
	// 標準出力とファイルの両方を束ねて、
	// logの出力先に設定する
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))

	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("Log!!")

	server()
}