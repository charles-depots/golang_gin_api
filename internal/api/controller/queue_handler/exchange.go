package queue_handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (m *mqController) ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = json.Unmarshal(body, ExchangeEntity); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err = RabbitMQ.Connect(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer RabbitMQ.Close()

		if r.Method == "POST" {
			if err = RabbitMQ.DeclareExchange(ExchangeEntity.Name, ExchangeEntity.Type, ExchangeEntity.Durable, ExchangeEntity.AutoDelete, ExchangeEntity.NoWait); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("declare exchange ok"))
		} else if r.Method == "DELETE" {
			if err = RabbitMQ.DeleteExchange(ExchangeEntity.Name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("delete exchange ok"))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
