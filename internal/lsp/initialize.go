package lsp

type InitializeRequst struct {
    Request
    Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
    ClientInfo ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
    Name string `json:"name"`
    Version string `json:"version"`
}

type InitializeResponse struct {
    Response
    Result InitializeResult `json:"result"`
}

type InitializeResult struct {
    Capabilities ServerCapabilities `json:"capabilities"`
    ServerInfo ServerInfo `json:"serverInfo"`
}

type ServerCapabilities struct {
    TextDocumentSync int `json:"textDocumentSync"`
    TextDocumentHover bool `json:"textDocumentHover"`
    CompletionProvider map[string]any `json:"completionProvider"`
}

type ServerInfo struct {
    Name string `json:"name"`
    Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
    return InitializeResponse {
        Response: Response {
            Rpc: "2.0",
            Id: id,
        },
        Result: InitializeResult {
            Capabilities: ServerCapabilities{
                TextDocumentSync: 1,
                CompletionProvider: map[string]any{},
            },
            ServerInfo: ServerInfo{
                Name: "my-lsp-server",
                Version: "0.1",
            },
        },
    }
}
