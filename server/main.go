package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type HubID int64

type Hub struct {
	Id      HubID     `json:"id"`
	Title   string    `title:"title"`
	Clients []*Client `clients:"clients"`
}

type HubResponse struct {
	Id    HubID  `json:"id"`
	Title string `title:"title"`
}

type Client struct {
	conn *websocket.Conn
	hub  *Hub
}

var sessionStore map[string]string

var hubs []Hub = []Hub{
	{Id: 1, Title: "School", Clients: []*Client{}},
	{Id: 0, Title: "Kolya"},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello\n"))
	})

	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Can't read body", http.StatusInternalServerError)
			return
		}
		log.Println(string(body))
		defer r.Body.Close()

		w.Write([]byte("hello\n"))
	})

	mux.HandleFunc("/hubs", func(w http.ResponseWriter, r *http.Request) {
		res := make([]HubResponse, len(hubs))
		for i, u := range hubs {
			res[i] = HubResponse{
				Id:    u.Id,
				Title: u.Title,
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/hubs/{id}/connect", func(w http.ResponseWriter, r *http.Request) {
		hubId := r.PathValue("id")
		log.Println(hubId)

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}

		for {
			conn.WriteMessage(websocket.TextMessage, []byte("hello :)"))
			time.Sleep(time.Second)
		}
	})

	if err := http.ListenAndServe(":8080", cors(mux)); err != nil {
		log.Println("Failed to start server", err)
	}
}
