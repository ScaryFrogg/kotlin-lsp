package main

import (
	"bufio"
	"github.com/ScaryFrogg/kotlin-lsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/vjn/Workspace/log.txt")
	logger.Println("Kotlin LSP started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
	}
	return log.New(logfile, "[kotlin-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
