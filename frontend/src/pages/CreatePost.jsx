import { useState } from 'react'
import api from '../api'
import { useNavigate } from 'react-router-dom'

export default function CreatePost() {
  const [form, setForm] = useState({ name: '', author: '', genre: '', content: '', tag_id: 1 })
  const navigate = useNavigate()

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    await api.post('/post/create_post', form)
    navigate('/posts')
  }

  return (
    <form onSubmit={handleSubmit}>
      <h2>Создать пост</h2>
      {['name', 'author', 'genre', 'content'].map(field => (
        <input
          key={field}
          name={field}
          placeholder={field}
          value={form[field]}
          onChange={handleChange}
        />
      ))}
      <input
        name="tag_id"
        placeholder="tag_id"
        type="number"
        value={form.tag_id}
        onChange={handleChange}
      />
      <button type="submit">Создать</button>
    </form>
  )
}
