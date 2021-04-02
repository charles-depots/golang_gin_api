package queue_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (m *mqController) QueueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "DELETE" {
		if r.Body == nil {
			fmt.Println("missing form body")
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal(body, QueueEntity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = RabbitMQ.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer RabbitMQ.Close()

		if r.Method == "POST" {
			if err = RabbitMQ.DeclareQueue(QueueEntity.Name, QueueEntity.Durable, QueueEntity.AutoDelete, QueueEntity.Exclusive, QueueEntity.NoWait); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("declare queue ok"))
		} else if r.Method == "DELETE" {
			if err = RabbitMQ.DeleteQueue(QueueEntity.Name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("delete queue ok"))
		}
	} else if r.Method == "GET" {
		r.ParseForm()
		if err := RabbitMQ.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer RabbitMQ.Close()

		message := make(chan []byte)

		for _, name := range r.Form["name"] {
			if err := RabbitMQ.ConsumeQueue(name, message); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.Write([]byte(""))
		w.(http.Flusher).Flush()

		for {
			fmt.Printf(" Received message %s\n", <-message)
			//fmt.Fprintf(w, "%s\n", <-message)
			w.(http.Flusher).Flush()
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
