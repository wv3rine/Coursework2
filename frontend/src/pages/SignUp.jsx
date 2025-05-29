import { useState } from 'react'
import api from '../api'
import { useNavigate } from 'react-router-dom'

export default function SignUp() {
  const [form, setForm] = useState({ login: '', password: '', repeatPassword: '' })
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value })
    setError('')
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    if (form.password !== form.repeatPassword) {
      setError('Пароли не совпадают')
      return
    }
    try {
      // const { login: saveUser } = useAuth()

      await api.post('/user/sign_up', {
        login: form.login,
        password: form.password,
        role: 'user'
      })
      // saveUser()
      navigate('/posts')
    } catch (err) {
      setError('Ошибка регистрации')
    }
  }

  return (
    <div className="container">
      <h2>Регистрация</h2>
      <form onSubmit={handleSubmit}>
        <input
          name="login"
          placeholder="Login"
          value={form.login}
          onChange={handleChange}
          type="text"
        />
        <input
          name="password"
          type="password"
          placeholder="Password"
          value={form.password}
          onChange={handleChange}
        />
        <input
          name="repeatPassword"
          type="password"
          placeholder="Повторите пароль"
          value={form.repeatPassword}
          onChange={handleChange}
        />
        {error && <p style={{ color: 'red', margin: 0 }}>{error}</p>}
        <button type="submit">Зарегистрироваться</button>
        <a className="button-secondary" href="/login">Войти</a>
      </form>
    </div>
  )
}
