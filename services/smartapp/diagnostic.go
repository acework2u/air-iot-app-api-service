package smartapp

// DiagnosticService
type DiagnosticService interface {
	CheckDiagnosticBoard(filter *DiagnosticFilter) (*DiagnosticResponse, error)
	DiagnosticBoards() ([]*DiagnosticResponse, error)
	CheckCompBoard(btu int64, compModel string) (*DiagnosticResponse, error)
}
type DiagnosticResponse struct {
	Btu       int64  `json:"btu"`
	CompId    string `json:"compId"`
	CompItem  string `json:"compItem"`
	CompModel string `json:"compModel"`
}
type DiagnosticFilter struct {
	Btu    int64  `json:"btu"`
	CompId string `json:"compId"`
}
