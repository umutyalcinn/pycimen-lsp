package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func Split(data []byte, _ bool) (advance int, token []byte, err error){
    header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
    if !found {
        return 0, nil, nil
    }

    contentLenghtBytes := header[len("Content-Lenght: "):]
    contentLength, err := strconv.Atoi(string(contentLenghtBytes))

    if err != nil {
        return 0, nil, err
    }

    if len(content) < contentLength {
        return 0, nil, nil
    }

    totalLength := len(header) + 4 + contentLength

    return totalLength, data[:totalLength], nil
}

func EncodeMessage(msg any) string {
    content, err := json.Marshal(msg)

    if err != nil {
        panic(err)
    }

    return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (string, []byte, error) {
    _, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

    if !found {
        return "", nil, errors.New("Did not found seperator")
    }

    var baseMessage BaseMessage
    err := json.Unmarshal(content, &baseMessage)
    if err != nil {
        return "", nil, err
    }
    return baseMessage.Method, content, nil
}

type BaseMessage struct {
    Method string `json:"method"`
}
