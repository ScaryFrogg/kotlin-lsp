package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func Encode(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		//TODO handle with parsing error code
		//{"jsonrpc": "2.0", "error": {"code": -32700, "message": "Parse error"}, "id": null}
		panic("asd")
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func Decode(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Did not find separator")
	}
	length, err := strconv.Atoi(string(header[16:]))
	if err != nil {
		return "", nil, err
	}
	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:length], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content, nil
}

type BaseMessage struct {
	Method string `json:"method"`
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}
	length, err := strconv.Atoi(string(header[16:]))
	if err != nil {
		return 0, nil, err
	}
	if len(content) < length {
		return 0, nil, nil
	}
	totalLength := 20 + length // 16+ header + 4 for separator + content
	return totalLength, data[:totalLength], nil
}
