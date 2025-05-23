import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { useAuth } from './context/AuthContext'
import Login from './pages/Login'
import CreatePost from './pages/CreatePost'
import PostList from './pages/PostList'
import SignUp from './pages/SignUp'
import PendingPosts from './pages/PendingPosts'


function App() {
  const { user, logout, loading } = useAuth()

  if (loading) return <div className="container"><p>Загрузка...</p></div>

  return (
    <BrowserRouter> 
      
      <nav>
        <div className="nav-left">
          <a href="/posts" className="primary">На главную</a>
          {user?.role === 'editor' && (
            <a href="/pending" className="secondary">Посмотреть предложенные тексты</a>
          )}
        </div>
        <div className="nav-right">
          {user ? (
            <>
              <a className="secondary" href="/create">Создать пост</a>
              <button className="logout-button" onClick={logout}>Выйти</button>
            </>
          ) : (
            <>
              <a className="primary" href="/login">Войти</a>
              <a className="secondary" href="/signup">Регистрация</a>
            </>
          )}
        </div>
      </nav>
      <div className="container">
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<SignUp />} />
          <Route path="/create" element={user ? <CreatePost /> : <Navigate to="/login" />} />
          <Route path="/posts" element={<PostList />} />
          <Route path="/pending" element={user?.role === 'editor' ? <PendingPosts /> : <Navigate to="/posts" />} />
          <Route path="*" element={<Navigate to="/posts" />} />
        </Routes>
      </div>
    </BrowserRouter>
  )
}

export default App
