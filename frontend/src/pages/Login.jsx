import { useState } from 'react'
import api from '../api'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'

export default function Login() {
  const [login, setLogin] = useState('')
  const [password, setPassword] = useState('')
  const { login: saveToken } = useAuth()
  const navigate = useNavigate()

  const { login: saveUser } = useAuth()

  const handleSubmit = async (e) => {
    e.preventDefault()
    await api.post('/user/sign_in', { login, password })
    saveUser() // проверяем сессию
    navigate('/posts')
  }

 return (
  <div className="container">
    <h2>Вход</h2>
    <form onSubmit={handleSubmit}>
      <input placeholder="Login" value={login} onChange={(e) => setLogin(e.target.value)} type="text" />
      <input placeholder="Password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button type="submit">Войти</button>
      <a className="button-secondary" style={{ textAlign: 'center' }} href="/signup">
        Зарегистрироваться
      </a>
    </form>
  </div>
  )

}
