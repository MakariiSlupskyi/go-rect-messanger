import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"

const Rooms = () => {
  const [messages, setMessages] = useState([])
  const [hubs, setHubs] = useState([])
  const [socket, setSocket] = useState(null);
  const navigate = useNavigate()

  useEffect(() => {
    const fetchData = async () => {
      const resp = await fetch("http://localhost:8080/hubs")
      const result = await resp.json()
      console.log(result)
    }
    fetchData()

    const newSocket = new WebSocket("ws://localhost:8080/hubs/1/connect")
    setSocket(newSocket)

    newSocket.onopen = () => {
    }

    newSocket.onmessage = (event) => {
      setMessages(prev => [...prev, event.data]);
    }

    newSocket.onclose = () => {
    }

    newSocket.onerror = () => {
    }

    return () => {
      newSocket.close();
    };
  }, [])

  return (
    <div className="h-screen">
      {hubs.map((h, i) => <p key={i}>{h}</p>)}
      {messages.map((m, i) => <p key={i}>{m}</p>)}
    </div>
  )
}

export default Rooms