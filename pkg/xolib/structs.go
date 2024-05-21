package xolib

type (
	// Params are the websocket paramteres
	Params map[string]interface{}

	// MessageRequest is the message request format
	MessageRequest struct {
		Method string  `json:"method"`
		Params *Params `json:"params,omitempty"`
	}

	messageRequest struct {
		ID      string `json:"id"`
		Jsonrpc string `json:"jsonrpc"`
		*MessageRequest
	}

	// MessageError represents an error message from xoa
	MessageError struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	// MessageResult is the result of a websocket response if successful
	MessageResult interface{}

	// MessageResponse is the websocket message response
	MessageResponse struct {
		ID      string         `json:"id"`
		Jsonrpc string         `json:"jsonrpc"`
		Error   *MessageError  `json:"error"`
		Result  *MessageResult `json:"result"`
	}
)
