package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ScaryFrogg/kotlin-lsp/lsp"
	"github.com/ScaryFrogg/kotlin-lsp/rpc"
)

func main() {
	logger := getLogger("/home/vjn/Workspace/log.txt")
	logger.Println("Kotlin LSP started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	writer := os.Stdout
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.Decode(msg)
		if err != nil {
			logger.Printf("error: %v", err)
		}
		handleMessage(logger, writer, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, method string, contents []byte) {
	logger.Printf("Handling [%s] -> %s:", method, contents)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, we couldn't parse this: %s", err)
		}

		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.Id)
		writeResponse(writer, msg)
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/hover: %s", err)
			return
		}

		//TODO get proper stuff
		response := lsp.HoverResponse{
			Response: lsp.Response{
				RPC: "2.0",
				Id:  request.Id,
			},
			Result: lsp.HoverResult{
				Contents: fmt.Sprintf("File: %s ", request.Params.TextDocument.URI),
			},
		}

		// Write it back
		writeResponse(writer, response)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
	}
	return log.New(logfile, "[kotlin-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.Encode(msg)
	writer.Write([]byte(reply))

}
