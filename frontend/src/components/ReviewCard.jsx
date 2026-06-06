import { useNavigate } from 'react-router-dom';
import { Heart, MessageCircle, User } from 'lucide-react';

const formatTime = (dateStr) => {
  const date = new Date(dateStr);
  const now = new Date();
  const diff = now - date;
  const minutes = Math.floor(diff / 60000);
  const hours = Math.floor(diff / 3600000);
  const days = Math.floor(diff / 86400000);

  if (minutes < 1) return '刚刚';
  if (minutes < 60) return `${minutes}分钟前`;
  if (hours < 24) return `${hours}小时前`;
  if (days < 30) return `${days}天前`;
  return date.toLocaleDateString('zh-CN');
};

const ReviewCard = ({ item }) => {
  const navigate = useNavigate();
  const { book, user, content, likes, created_at } = item;

  return (
    <div
      className="bg-white rounded-xl p-4 shadow-sm cursor-pointer hover:shadow-md transition-shadow"
      onClick={() => navigate(`/book/${book.id}`)}
    >
      {/* 书籍信息和封面 */}
      <div className="flex gap-3 mb-3">
        <img
          src={book.cover}
          alt={book.title}
          className="w-16 h-20 rounded-lg object-cover bg-gray-100 flex-shrink-0"
          onError={(e) => { e.target.src = ''; e.target.className += ' hidden'; }}
        />
        <div className="flex-1 min-w-0">
          <h3 className="font-semibold text-gray-900 text-base truncate">{book.title}</h3>
          <p className="text-sm text-gray-500">{book.author}</p>
          <span className="inline-block mt-1 px-2 py-0.5 text-xs rounded-full bg-amber-50 text-amber-700">
            {book.category || '未分类'}
          </span>
        </div>
      </div>

      {/* 评论内容 */}
      <p className="text-sm text-gray-700 leading-relaxed mb-3 line-clamp-3">
        {content}
      </p>

      {/* 底部：用户信息和互动数据 */}
      <div className="flex items-center justify-between text-xs text-gray-400">
        <div className="flex items-center gap-2">
          {user.avatar ? (
            <img src={user.avatar} alt="" className="w-5 h-5 rounded-full" />
          ) : (
            <User className="w-5 h-5 text-gray-300" />
          )}
          <span>{user.username}</span>
          <span>{formatTime(created_at)}</span>
        </div>
        <div className="flex items-center gap-3">
          <span className="flex items-center gap-1">
            <Heart className="w-4 h-4" />
            {likes}
          </span>
          <span className="flex items-center gap-1">
            <MessageCircle className="w-4 h-4" />
            0
          </span>
        </div>
      </div>
    </div>
  );
};

export default ReviewCard;
