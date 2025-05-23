import { useEffect, useState } from 'react'
import api from '../api'

export default function PostList() {
  const [posts, setPosts] = useState([])

  useEffect(() => {
    api.post('/post/get_posts', { statuses: ['approved'] })
      .then(res => setPosts(res.data.data.posts))
      .catch(err => console.error(err))
  }, [])

  return (
    <div>
      <h2>Посты</h2>
      {posts.map(post => (
        <div className="post" key={post.post_id}>
          <h3>{post.name}</h3>
          <p><strong>Автор:</strong> {post.author}</p>
          <p>{post.content}</p>
          <div className="post-meta">
            Жанр: {post.genre} | Тэг: {post.tag_name}
          </div>
        </div>
      ))}
    </div>
  )
}
