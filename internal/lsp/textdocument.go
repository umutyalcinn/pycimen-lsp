package lsp

type TextDocumentItem struct {
    Uri string `json:"uri"`
    LanguageId string `json:"languageId"`
    Version int `json:"version"`
    Text string `json:"text"`
}

type VersionedTextDocumentIdentifier struct {
    Uri string `json:"uri"`
    Version int `json:"version"`
}

type TextDocumentContentChangeEvent struct {
    Text string `json:"text"`
} 

type Position struct {
    Line uint `json:"line"`
    Character uint `json:"character"`
}
