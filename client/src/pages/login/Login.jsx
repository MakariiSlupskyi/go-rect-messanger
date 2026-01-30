import { useNavigate } from "react-router-dom"
import { Eye, EyeClosed, EyeOff } from 'lucide-react';
import { useState } from "react";

const Login = () => {
  const [name, setName] = useState(false)
  const [password, setPassword] = useState(false)
  const [showPassword, setShowPassword] = useState(false)
  const navigate = useNavigate()

  const HandleLogin = async (e) => {
    e.preventDefault()
    const resp = await fetch("http://localhost:8080/login", {
      method: "POST",
      body: JSON.stringify({username: name, password: password})
    })
    console.log(resp)
  }

  return (
    <div className="flex h-screen items-center justify-center bg-gray-50">
      <div className="px-10 py-6 border bg-white border-gray-100 rounded-2xl shadow-lg">
        <form onSubmit={e => HandleLogin(e)}>
          <p className="text-2xl font-bold mb-6">Login</p>
          
          <p className="text-gray-500 mb-2">Your username:</p>
          <input
            className="block mb-3 px-1 border-2 border-gray-300 rounded-lg w-50"
            autoComplete="on"
            value={name}
            onChange={e => setName(e.target.value)}
          />
          
          <p className="text-gray-500 mb-2">Your password:</p>
          <div>
            <input
              type={showPassword ? "text" : "password"}
              className="mb-6 mr-2 px-1 border-2 border-gray-300 rounded-lg w-50"
              autoComplete="on"
              value={password}
              onChange={e => setPassword(e.target.value)}
            />
            { showPassword
              ? <Eye onClick={() => setShowPassword(false)} className="inline text-gray-500" />
              : <EyeOff onClick={() => setShowPassword(true)} className="inline text-gray-500"  />
            }
          </div>
          
          <button className="px-3 py-2 bg-amber-400 text-white font-medium rounded-xl shadow" type="submit">Login</button>
        </form>
      </div>
    </div>
  )
}

export default Login