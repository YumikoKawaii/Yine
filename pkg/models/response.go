package models

type Response struct {
	Table    string `json:"table"`
	Operator string `json:"operator"`
	Action   string `json:"action"`
	Target   string `json:"target"`
	Content  string `json:"content"`
}

func (r Response) NewResponse(table string, operator string, action string, target string, content string) Response {

	return Response{
		Table:    table,
		Operator: operator,
		Action:   action,
		Target:   target,
		Content:  content,
	}

}
