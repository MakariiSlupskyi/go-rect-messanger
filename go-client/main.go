package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"
)

type wsMsg string

type model struct {
	conn     *websocket.Conn
	messages []string

	width    int
	height   int
	quitting bool
}

func listenWebsocket(conn *websocket.Conn) tea.Cmd {
	return func() tea.Msg {
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("%v", err)
		}
		return wsMsg(data)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(listenWebsocket(m.conn))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case wsMsg:
		m.messages = append(m.messages, string(msg))
		return m, listenWebsocket(m.conn)
	case tea.KeyMsg:
		m.quitting = true
		return m, tea.Quit
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	default:
		return m, nil
	}
}

func (m model) View() string {
	return strings.Join(m.messages, "\n")
}

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	p := tea.NewProgram(model{
		conn:     c,
		width:    10,
		height:   10,
		quitting: false,
	}, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Could not start program:", err)
		os.Exit(1)
	}
}
