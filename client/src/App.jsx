import { useState } from 'react'
import { Route, Routes } from 'react-router-dom'
import Login from './pages/login/Login'
import Rooms from './pages/rooms/Rooms'

function App() {
  const [count, setCount] = useState(0)

  return (
    <Routes>
      <Route path='/login' element={<Login />} />
      <Route path='/rooms' element={<Rooms />} />
    </Routes>
  )
}

export default App
