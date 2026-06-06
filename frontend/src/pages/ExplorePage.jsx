import { useState, useEffect, useRef, useCallback } from 'react';
import { Search, Loader2, RefreshCw } from 'lucide-react';
import axios from 'axios';
import ReviewCard from '../components/ReviewCard';
import BottomNav from '../components/BottomNav';

const ExplorePage = () => {
  const [items, setItems] = useState([]);
  const [cursor, setCursor] = useState(0);
  const [hasMore, setHasMore] = useState(true);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');
  const [searchQuery, setSearchQuery] = useState('');
  const [inputValue, setInputValue] = useState('');
  const sentinelRef = useRef(null);
  const debounceRef = useRef(null);
  const loadingRef = useRef(false);
  const apiBaseUrl = import.meta.env.VITE_API_BASE_URL || '';

  // 获取数据
  const fetchData = useCallback(async (cursorVal, query, reset) => {
    if (loadingRef.current) return;
    loadingRef.current = true;
    setIsLoading(true);
    setError('');

    try {
      const endpoint = query
        ? `${apiBaseUrl}/api/v1/explore/search`
        : `${apiBaseUrl}/api/v1/explore`;

      const params = { limit: 20, cursor: cursorVal };
      if (query) params.q = query;

      const response = await axios.get(endpoint, { params });
      const data = response.data;

      if (reset) {
        setItems(data.items);
      } else {
        setItems((prev) => [...prev, ...data.items]);
      }
      setCursor(data.next_cursor);
      setHasMore(data.has_more);
    } catch (err) {
      setError('加载失败，请重试');
    } finally {
      loadingRef.current = false;
      setIsLoading(false);
    }
  }, [apiBaseUrl]);

  // 首次加载
  useEffect(() => {
    fetchData(0, '', true);
  }, []);

  // IntersectionObserver 无限滚动 (Task 3.5)
  useEffect(() => {
    const sentinel = sentinelRef.current;
    if (!sentinel) return;

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0].isIntersecting && hasMore && !isLoading) {
          fetchData(cursor, searchQuery, false);
        }
      },
      { threshold: 0.1 }
    );

    observer.observe(sentinel);
    return () => observer.disconnect();
  }, [cursor, hasMore, isLoading, searchQuery, fetchData]);

  // 搜索防抖处理 (Task 3.6)
  const handleSearchInput = (e) => {
    const val = e.target.value;
    setInputValue(val);

    if (debounceRef.current) {
      clearTimeout(debounceRef.current);
    }

    debounceRef.current = setTimeout(() => {
      setSearchQuery(val);
      if (val.trim()) {
        setCursor(0);
        setHasMore(true);
        fetchData(0, val, true);
      } else {
        setCursor(0);
        setHasMore(true);
        fetchData(0, '', true);
      }
    }, 300);
  };

  const clearSearch = () => {
    setInputValue('');
    setSearchQuery('');
    setCursor(0);
    setHasMore(true);
    fetchData(0, '', true);
  };

  const handleRetry = () => {
    fetchData(cursor, searchQuery, items.length === 0);
  };

  return (
    <div className="min-h-screen bg-amber-50 pb-16">
      {/* 顶部：标题 + 搜索框 */}
      <header className="sticky top-0 bg-amber-50/95 backdrop-blur-sm z-40 px-4 pt-4 pb-3">
        <h1 className="text-2xl font-bold text-gray-900 mb-3">探索</h1>
        <div className="relative">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
          <input
            type="text"
            value={inputValue}
            onChange={handleSearchInput}
            placeholder="搜索书籍或评论..."
            className="w-full pl-10 pr-8 py-2.5 bg-white border border-gray-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-amber-400 focus:border-transparent"
          />
          {inputValue && (
            <button
              onClick={clearSearch}
              className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
            >
              ✕
            </button>
          )}
        </div>
      </header>

      {/* 内容：评论信息流列表 */}
      <main className="px-4">
        {/* 加载中状态 */}
        {isLoading && items.length === 0 && (
          <div className="flex justify-center py-20">
            <Loader2 className="animate-spin w-8 h-8 text-amber-500" />
          </div>
        )}

        {/* 错误状态 */}
        {error && items.length === 0 && (
          <div className="text-center py-20">
            <p className="text-gray-500 mb-4">{error}</p>
            <button
              onClick={handleRetry}
              className="inline-flex items-center gap-2 px-4 py-2 bg-amber-500 text-white rounded-lg hover:bg-amber-600"
            >
              <RefreshCw className="w-4 h-4" />
              重试
            </button>
          </div>
        )}

        {/* 空状态 */}
        {!isLoading && !error && items.length === 0 && (
          <div className="text-center py-20">
            <p className="text-gray-400 text-lg">
              {searchQuery ? '暂无搜索结果' : '暂无内容'}
            </p>
          </div>
        )}

        {/* 评论卡片列表 */}
        {items.length > 0 && (
          <div className="space-y-3 pb-4">
            {items.map((item) => (
              <ReviewCard key={item.comment_id} item={item} />
            ))}
          </div>
        )}

        {/* 无限滚动 sentinel */}
        {hasMore && !error && (
          <div
            ref={sentinelRef}
            className="flex justify-center py-6"
          >
            {isLoading && <Loader2 className="animate-spin w-6 h-6 text-amber-400" />}
          </div>
        )}

        {/* 没有更多了 */}
        {!hasMore && items.length > 0 && (
          <div className="text-center py-6 text-gray-400 text-sm">
            没有更多了
          </div>
        )}
      </main>

      {/* 底部导航栏 */}
      <BottomNav activeTab="explore" />
    </div>
  );
};

export default ExplorePage;
