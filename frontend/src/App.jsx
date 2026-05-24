import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import LoginPage from './pages/LoginPage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/explore" element={<div className="p-8"><h1>探索页 (开发中)</h1><button onClick={() => { localStorage.clear(); window.location.href='/login' }} className="mt-4 p-2 bg-red-500 text-white rounded">退出登录</button></div>} />
        <Route path="/" element={<Navigate to="/login" replace />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
