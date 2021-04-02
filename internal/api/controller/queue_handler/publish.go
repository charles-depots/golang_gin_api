package queue_handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (m *mqController) PublishHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal(body, MessageEntity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = RabbitMQ.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer RabbitMQ.Close()

		if err = RabbitMQ.Publish(MessageEntity.Exchange, MessageEntity.Key, MessageEntity.DeliveryMode, MessageEntity.Priority, MessageEntity.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("publish message ok => " + MessageEntity.Body))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
