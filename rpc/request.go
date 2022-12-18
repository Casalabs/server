package rpc

type Tx struct {
	Transaction string `json:"Transaction"`
}
type Cursor struct {
	TxSeq    int `json:"txSeq"`
	EventSeq int `json:"eventSeq"`
}
