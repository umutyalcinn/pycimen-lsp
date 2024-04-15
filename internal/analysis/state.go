package analysis

type State struct {
    Documents map[string]string
}

func NewState() State {
    return State { 
        Documents: map[string]string{},
    }
}

func (s State) SetDocument(document string, content string) {
    s.Documents[document] = content
}
