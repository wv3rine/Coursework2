import { useEffect, useState } from 'react'
import api from '../api'

export default function PendingPosts() {
  const [posts, setPosts] = useState([])

  useEffect(() => {
    api.post('/post/get_posts', { statuses: ['on_check'] })
      .then(res => setPosts(res.data.data.posts))
      .catch(console.error)
  }, [])

  const approvePost = async (postId) => {
    try {
      await api.post('/post/approve_post', { post_id: postId })
      setPosts(posts.filter(p => p.post_id !== postId))
    } catch (e) {
      alert('Ошибка при одобрении поста')
    }
  }

  const rejectPost = async (postId) => {
    try {
      // Предположим, у тебя есть эндпоинт для отклонения
      await api.post('/post/reject_post', { post_id: postId })
      setPosts(posts.filter(p => p.post_id !== postId))
    } catch (e) {
      alert('Ошибка при отклонении поста')
    }
  }

  return (
    <div className="container">
      <h2>Предложенные тексты на проверке</h2>
      {posts.length === 0 && <p>Нет постов для проверки.</p>}
      {posts.map(post => (
        <div key={post.post_id} className="post">
          <h3>{post.name}</h3>
          <p><strong>Автор:</strong> {post.author}</p>
          <p>{post.content}</p>
          <div className="post-meta">
            Жанр: {post.genre} | Тэг: {post.tag_name}
          </div>
          <div style={{ marginTop: '1rem' }}>
            <button style={{ marginRight: 10, backgroundColor: 'green', color: 'white' }}
              onClick={() => approvePost(post.post_id)}>Одобрить</button>
            <button style={{ backgroundColor: 'red', color: 'white' }}
              onClick={() => rejectPost(post.post_id)}>Отклонить</button>
          </div>
        </div>
      ))}
    </div>
  )
}
