package lsp

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Id     int    `json:"id"`
	Method string `json:"method"`
}
type Request struct {
	RPC    string `json:"jsonrpc"`
	Id     int    `json:"id"`
	Method string `json:"method"`
}
type Response struct {
	RPC string `json:"jsonrpc"`
	Id  int    `json:"id"`
}
