package queue_handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (m *mqController) QueueBindHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal(body, QueueBindEntity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = RabbitMQ.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer RabbitMQ.Close()

		if r.Method == "POST" {
			if err = RabbitMQ.BindQueue(QueueBindEntity.Queue, QueueBindEntity.Exchange, QueueBindEntity.Keys, QueueBindEntity.NoWait); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("bind queue ok"))
		} else if r.Method == "DELETE" {
			if err = RabbitMQ.UnBindQueue(QueueBindEntity.Queue, QueueBindEntity.Exchange, QueueBindEntity.Keys); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("unbind queue ok"))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
