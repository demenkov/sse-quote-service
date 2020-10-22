package ssequote

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	socketio "github.com/googollee/go-socket.io"
	"strings"
)

func sseHandler(s socketio.Conn, msg string) {
	res := Response{
		ServerTime: Timestamp(),
	}
	var r Request
	err := json.Unmarshal([]byte(msg), &r)
	if err != nil {
		res.Error = &ErrorType{
			Code:    1101400,
			Message: fmt.Sprintf("Undefined request. %s", msg),
		}
	}
	if "" == r.Symbols {
		res.Error = &ErrorType{
			Code:    1103400,
			Message: fmt.Sprintf("Missing symbols attribute %s", msg),
		}
	}

	if nil != res.Error {
		s.Emit("sse-response", res)
		return
	}
	var symbols []string
	sp := strings.Split(r.Symbols, ",")

	for _, symbol := range sp {
		if ok, _ := in_array(symbol, server.Rooms("/")); ok {
			continue
		}

		symbols = append(symbols, symbol)
	}

	for _, symbol := range sp {
		s.Join(symbol)

		if ok, _ := in_array(symbol, symbols); !ok {
			c, e := Redis.Get(context.Background(), fmt.Sprintf("merostm:cloud:sse:%s", symbol)).Result()

			if e != redis.Nil {
				var q []Quote
				_ = json.Unmarshal([]byte(c), &q)
				s.Emit("sse-response", q)
			}
		}
	}

	if len(symbols) == 0 {
		return
	}

	r.Symbols = strings.Join(symbols, ",")
	go sseEvent(r)
}
