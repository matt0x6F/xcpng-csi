package xolib

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func rawCall(ws *websocket.Conn, rawReq *MessageRequest) (*MessageResult, error) {
	id := uuid.New()
	req := &messageRequest{
		ID:             id.String(),
		Jsonrpc:        "2.0",
		MessageRequest: rawReq,
	}

	raw, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var data *MessageResponse
	timeouterr := &MessageResponse{
		ID: id.String(),
		Error: &MessageError{
			Message: "Timeout Error",
			Code:    404,
		},
	}

	if err := ws.WriteMessage(websocket.TextMessage, raw); err != nil {
		return nil, err
	}

	//ws.SetReadDeadline(time.Now().Add(5 * time.Second))
	timeout := time.Now().Add(30 * time.Second)

	for timeout.After(time.Now()) {
		if _, message, err := ws.ReadMessage(); err == nil {
			raw := MessageResponse{}
			if err = json.Unmarshal(message, &raw); err == nil {
				if raw.ID != "" && raw.ID == req.ID {
					data = &raw
					break
				}
			}
		}
	}

	if data == nil {
		data = timeouterr
	}

	if data.Error != nil {
		return nil, errors.New(data.Error.Message)
	}

	return data.Result, nil
}

func (xo *xolib) getLogin() *MessageRequest {
	params := Params{}

	var method string

	method = "session.signIn"
	params["email"] = xo.Config.Username
	params["password"] = xo.Config.Password

	return &MessageRequest{
		Method: method,
		Params: &params,
	}
}

// Call is used to execute calls to the xolib api
func (xo *xolib) Call(req *MessageRequest) (*MessageResult, error) {
	ws, err := xo.getWS()
	if err != nil {
		return nil, err
	}

	login := xo.getLogin()

	if _, err := rawCall(ws, login); err != nil {
		return nil, err
	}

	result, err := rawCall(ws, req)

	return result, err
}

// Init check if login works
func (xo *xolib) Init() error {
	ws, err := xo.getWS()
	if err != nil {
		return err
	}

	login := xo.getLogin()

	if _, err := rawCall(ws, login); err != nil {
		return err
	}

	return nil
}
