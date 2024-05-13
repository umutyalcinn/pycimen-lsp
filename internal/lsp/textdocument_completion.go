package lsp

type TextDocumentCompletionRequest struct {
    Request
    Params TextDocumentCompletionParams `json:"params"`
}

type TextDocumentCompletionParams struct {
    Context CompletionContext `json:"context"`
}

type CompletionContext struct {
    TriggerKind int `json:"triggerKind"`
    TriggerCharacter string `json:"triggerCharacter"`
}

type TextDocumentCompletionResponse struct  {
    Response
    Result CompletionList `json:"result"`
}

type CompletionList struct {
    IsIncomplete bool `json:"inInComplete"`
    Items []CompletionItem `json:"items"`
}

type CompletionItem struct {
    Label string `json:"label"`
    InsertText string `json:"insertText"`
    Kind int `json:"kind"`
}

func NewKeywordCompletion(keyword string) CompletionItem{
    return CompletionItem {
        Label: keyword,
        InsertText: keyword,
        Kind: KEYWORD,
    }
}

func NewCompletionResponse(id int) TextDocumentCompletionResponse {
    return TextDocumentCompletionResponse {
        Response: Response {
            Rpc: "2.0",
            Id: id,
        },
        Result: CompletionList {
            IsIncomplete: false,
            Items: []CompletionItem {
                NewKeywordCompletion("print"),
                NewKeywordCompletion("True"),
                NewKeywordCompletion("False"),
                NewKeywordCompletion("None"),
                NewKeywordCompletion("and"),
                NewKeywordCompletion("or"),
                NewKeywordCompletion("not"),
                NewKeywordCompletion("while"),
                NewKeywordCompletion("if"),
                NewKeywordCompletion("elif"),
                NewKeywordCompletion("def"),
                NewKeywordCompletion("return"),
                NewKeywordCompletion("break"),
                NewKeywordCompletion("continue"),
                NewKeywordCompletion("pass"),
                NewKeywordCompletion("class"),
                NewKeywordCompletion("for"),
                NewKeywordCompletion("in"),
                NewKeywordCompletion("import"),
                NewKeywordCompletion("numpy"),
                NewKeywordCompletion("numpy.mean"),
                NewKeywordCompletion("numpy.median"),
                NewKeywordCompletion("numpy.std"),
                NewKeywordCompletion("numpy.array"),
            },
        },
    }
}

const (
	TEXT = 1;
	METHOD = 2;
	FUNCTION = 3;
	CONSTRUCTOR = 4;
	FIELD = 5;
	VARIABLE = 6;
	CLASS = 7;
	INTERFACE = 8;
	MODULE = 9;
	PROPERTY = 10;
	UNIT = 11;
	VALUE = 12;
	ENUM = 13;
	KEYWORD = 14;
	SNIPPET = 15;
	COLOR = 16;
	FILE = 17;
	REFERENCE = 18;
	FOLDER = 19;
	ENUMMEMBER = 20;
	CONSTANT = 21;
	STRUCT = 22;
	EVENT = 23;
	OPERATOR = 24;
	TYPEPARAMETER = 25;
)
