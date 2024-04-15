package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/umutyalcinn/lsp/internal/analysis"
	"github.com/umutyalcinn/lsp/internal/lsp"
	"github.com/umutyalcinn/lsp/internal/rpc"
)


func main(){
    logger := getLogger("/Users/umutyalcinn/source/umutyalcinn/lsp/logs/log.txt")
    
    /*
    listener, err := net.Listen("tcp", ":2000")

    if err != nil {
        log.Fatal(err)
    }
    */

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)

    state := analysis.NewState()

    logger.Println("starting lsp server")

    for scanner.Scan() {
        msg := scanner.Bytes()
        handleMessage(logger, msg, state)
    }

    logger.Println("Hey, i'm done")

    /*

    defer listener.Close()

    for {
        conn, err := listener.Accept()


        if err != nil {
            log.Fatal(err)
        }

        go func (c net.Conn) {
            defer c.Close()

            buffer := make([]byte, 999999)

            for {
                n, err := conn.Read(buffer)

                if err != nil {
                    log.Print(err)
                    return
                } 

                method, contents, err := rpc.DecodeMessage(buffer[:n])

                if err != nil {
                    log.Print(err)
                    continue
                }

                handleMessage(logger, method, contents, state, conn)

                logger.Printf("%v", state)

            }
        }(conn)
    }
    */
}

func getLogger(filepath string) *log.Logger {
    file, err := os.OpenFile(filepath, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0666)

    if err != nil {
        panic("unable to open log file")
    }

    return log.New(file, "[my-lsp]", log.Ldate | log.Ltime | log.Lshortfile)
}

func handleMessage(logger *log.Logger, message []byte, state analysis.State) {

    method, contents, err := rpc.DecodeMessage(message)

    logger.Printf("Received method: %s\n", method)
    logger.Printf("Contents: %s\n", contents)

    writer := os.Stdout

    if err != nil {
        logger.Printf("error decoding message, %s", err)
    }

    switch method {
    case "initialize" :
        var request lsp.InitializeRequst
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("Error when unmarshal initialize: %s\n", err)
            return
        }

        logger.Printf("Connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

        response := lsp.NewInitializeResponse(request.Id)

        logger.Printf("Response: %s", rpc.EncodeMessage(response))

        writer.Write([]byte(rpc.EncodeMessage(response)))

        break;

    case "textDocument/didOpen":
        var request lsp.DidOpenTextDocumentNotification
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("Error when unmarshal didopen: %s\n", err)
            return
        }

        logger.Printf("Document opened: %s\nDocument Contents:\n%s\n", request.Params.TextDocument.Uri, request.Params.TextDocument.Text)

        state.SetDocument(request.Params.TextDocument.Uri, request.Params.TextDocument.Text)
        break;

    case "textDocument/didChange":
        var request lsp.DidChangeTextDocumentNotification
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("Error when unmarshal didchange: %s\n", err)
            return
        }

        logger.Printf("Document changed: %s\nDocument Contents:\n%s\n", request.Params.TextDocument.Uri, request.Params.ContentChanges[0].Text)

        state.SetDocument(request.Params.TextDocument.Uri, request.Params.ContentChanges[0].Text)

        break;

    case "textDocument/completion":
        var request lsp.TextDocumentCompletionRequest
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("Error when unmarshal completion: %s\n", err)
            return
        }

        response := lsp.NewCompletionResponse(request.Id)

        logger.Printf("Completion result: %v", response)

        writer.Write([]byte(rpc.EncodeMessage(response)))

        break;
    }
}
