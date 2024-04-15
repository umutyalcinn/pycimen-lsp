package lsp

type Request struct {
    Rpc string `json:"jsonrpc"`
    Id int `json:"id"`
    Method string `json:"method"`

    //Params
}

type Response struct {
    Rpc string `json:"jsonrpc"`
    Id int `json:"id"`

    // Result or Error
}

type Notification struct {
    Method string `json:"method"`

    // Params
}

