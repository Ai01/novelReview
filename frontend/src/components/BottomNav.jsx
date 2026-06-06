import { useNavigate } from 'react-router-dom';
import { Compass, Library, Plus, Bookmark, User } from 'lucide-react';

const tabs = [
  { key: 'explore', label: '探索', icon: Compass, path: '/explore' },
  { key: 'library', label: '书库', icon: Library, path: '/library' },
  { key: 'add', label: '', icon: Plus, path: null },
  { key: 'booklists', label: '书单', icon: Bookmark, path: '/booklists' },
  { key: 'profile', label: '个人', icon: User, path: '/profile' },
];

const BottomNav = ({ activeTab = 'explore' }) => {
  const navigate = useNavigate();

  const handleClick = (tab) => {
    if (tab.key === 'add') {
      // TODO: 打开底部弹窗
      return;
    }
    navigate(tab.path);
  };

  return (
    <nav className="fixed bottom-0 left-0 right-0 bg-white border-t border-gray-100 z-50">
      <div className="max-w-lg mx-auto flex justify-around items-center h-14 px-2">
        {tabs.map((tab) => {
          const Icon = tab.icon;
          const isActive = activeTab === tab.key;
          const isAdd = tab.key === 'add';

          return (
            <button
              key={tab.key}
              onClick={() => handleClick(tab)}
              className={`flex flex-col items-center justify-center gap-0.5 flex-1 h-full ${
                isAdd
                  ? ''
                  : isActive
                    ? 'text-amber-600'
                    : 'text-gray-400 hover:text-gray-600'
              }`}
            >
              {isAdd ? (
                <div className="w-10 h-10 rounded-full bg-amber-500 flex items-center justify-center text-white -mt-3 shadow-md">
                  <Icon className="w-5 h-5" />
                </div>
              ) : (
                <Icon className="w-5 h-5" />
              )}
              {tab.label && (
                <span className="text-xs">{tab.label}</span>
              )}
            </button>
          );
        })}
      </div>
    </nav>
  );
};

export default BottomNav;
