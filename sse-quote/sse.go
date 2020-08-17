package ssequote

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/r3labs/sse"
	"log"
	"strings"
	"time"
)

func sseEvent(r Request) {
	url := fmt.Sprintf("%s?symbols=%s&token=%s",
		Conf.Cloud.Url, r.Symbols, Conf.Cloud.Key)

	client := sse.NewClient(url)

	_ = client.Subscribe("messages", func(msg *sse.Event) {
		var q []Quote
		_ = json.Unmarshal(msg.Data, &q)

		log.Println(string(msg.Data))
		if len(q) > 0 {
			s := strings.ToLower(q[0].Symbol)
			server.BroadcastToRoom("/", s, "sse-response", q)

			j, _ := json.Marshal(q)
			Redis.Set(context.Background(),
				fmt.Sprintf("merostm:cloud:sse:%s", s),
				j,
				time.Duration(Conf.Rest.CacheInterval)*time.Second,
			)
		}
	})
}
