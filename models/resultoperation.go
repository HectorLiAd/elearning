package models

/*ResultOperation salida en json para verificar la salida de un proceso */
type ResultOperation struct {
	NameMessage  string `json:"Message"`
	Code         int64  `json:"Code"`
	RowAffected  int64  `json:"RowAffected"`
	ErrorType    string `json:"ErrorType,omitempty"`
	ErrorMessage string `json:"ErrorMessage,omitempty"`
}
