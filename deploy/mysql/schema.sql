CREATE TABLE IF NOT EXISTS users (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  username VARCHAR(191) NOT NULL,
  email VARCHAR(191) NOT NULL,
  password VARCHAR(255) NOT NULL,
  avatar VARCHAR(255) NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME NULL,
  PRIMARY KEY (id),
  UNIQUE KEY uq_users_username (username),
  UNIQUE KEY uq_users_email (email),
  KEY idx_users_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS books (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  cover VARCHAR(512) NULL,
  description TEXT NULL,
  category VARCHAR(100) NULL,
  tags VARCHAR(255) NULL,
  status VARCHAR(50) NOT NULL DEFAULT '连载中',
  last_updated DATETIME NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FULLTEXT KEY ft_books_title_author (title, author)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS comments (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  book_id BIGINT UNSIGNED NOT NULL,
  user_id BIGINT UNSIGNED NOT NULL,
  content TEXT NOT NULL,
  likes INT NOT NULL DEFAULT 0,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY idx_comments_book_created (book_id, created_at),
  KEY idx_comments_created (created_at),
  CONSTRAINT fk_comments_book FOREIGN KEY (book_id) REFERENCES books(id),
  CONSTRAINT fk_comments_user FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- =============================================
-- Seed Data: 种子用户
-- =============================================
INSERT IGNORE INTO users (id, username, email, password, avatar) VALUES
(1, 'reader01', 'reader01@novel.com', '$2a$10$dummyhashforseed', 'https://api.dicebear.com/7.x/avataaars/svg?seed=reader01'),
(2, 'bookworm', 'bookworm@novel.com', '$2a$10$dummyhashforseed', 'https://api.dicebear.com/7.x/avataaars/svg?seed=bookworm'),
(3, '小说迷', 'novelfan@novel.com', '$2a$10$dummyhashforseed', 'https://api.dicebear.com/7.x/avataaars/svg?seed=novelfan'),
(4, 'reading123', 'reading123@novel.com', '$2a$10$dummyhashforseed', 'https://api.dicebear.com/7.x/avataaars/svg?seed=reading123'),
(5, '文学青年', 'literary@novel.com', '$2a$10$dummyhashforseed', 'https://api.dicebear.com/7.x/avataaars/svg?seed=literary');

-- =============================================
-- Seed Data: 种子书籍 (20本)
-- =============================================
INSERT IGNORE INTO books (id, title, author, cover, description, category, tags, status, last_updated) VALUES
(1, '三体', '刘慈欣', 'https://img2.doubanio.com/view/subject/l/public/s27683781.jpg', '文化大革命如火如荼进行的同时，军方探寻外星文明的绝秘计划"红岸工程"取得了突破性进展。', '科幻', '科幻,硬科幻,外星文明', '已完结', '2026-05-01 10:00:00'),
(2, '活着', '余华', 'https://img2.doubanio.com/view/subject/l/public/s29053580.jpg', '地主少爷富贵嗜赌成性，终于赌光了家业一贫如洗，穷困之中的富贵因为母亲生病前去求医。', '文学', '文学,现实主义,人生', '已完结', '2026-04-28 09:00:00'),
(3, '百年孤独', '加西亚·马尔克斯', 'https://img2.doubanio.com/view/subject/l/public/s6384944.jpg', '《百年孤独》是魔幻现实主义文学的代表作，描写了布恩迪亚家族七代人的传奇故事。', '文学', '魔幻现实主义,拉美文学,经典', '已完结', '2026-04-25 08:00:00'),
(4, '平凡的世界', '路遥', 'https://img2.doubanio.com/view/subject/l/public/s29651121.jpg', '以中国70年代中期到80年代中期十年间为背景，以孙少安和孙少平两兄弟为中心。', '文学', '现实主义,农村,奋斗', '已完结', '2026-05-15 14:00:00'),
(5, '盗墓笔记', '南派三叔', 'https://img2.doubanio.com/view/subject/l/public/s29619565.jpg', '五十年前，一群长沙土夫子挖到一部战国帛书，残篇中记载了一座奇特的战国古墓的位置。', '悬疑', '悬疑,盗墓,探险', '已完结', '2026-05-20 16:00:00'),
(6, '哈利·波特与魔法石', 'J.K.罗琳', 'https://img2.doubanio.com/view/subject/l/public/s1070959.jpg', '一岁的哈利·波特失去父母后，在姨父家饱受欺凌，但在他十一岁生日那天一切都发生了变化。', '奇幻', '奇幻,魔法,青少年', '已完结', '2026-03-15 07:00:00'),
(7, '围城', '钱钟书', 'https://img2.doubanio.com/view/subject/l/public/s1070222.jpg', '《围城》是钱钟书所著的长篇小说，是中国现代文学史上一部风格独特的讽刺小说。', '文学', '讽刺,婚姻,知识分子', '已完结', '2026-05-05 11:00:00'),
(8, '白夜行', '东野圭吾', 'https://img2.doubanio.com/view/subject/l/public/s4610502.jpg', '1973年，大阪的一栋废弃建筑内发现了一具男尸，此后19年，嫌疑人之女与被害者之子走上截然不同的人生道路。', '悬疑', '悬疑,推理,日系', '已完结', '2026-05-10 13:00:00'),
(9, '红楼梦', '曹雪芹', 'https://img2.doubanio.com/view/subject/l/public/s1070959.jpg', '以贾、史、王、薛四大家族的兴衰为背景，以富贵公子贾宝玉为视角，描绘了一批闺阁佳人的人生百态。', '古典', '古典文学,四大名著,爱情', '已完结', '2026-01-01 00:00:00'),
(10, '挪威的森林', '村上春树', 'https://img2.doubanio.com/view/subject/l/public/s1028292.jpg', '这是一部动人心弦的、平缓舒雅的、略带感伤的恋爱小说。', '文学', '爱情,青春,日系文学', '已完结', '2026-04-20 10:00:00'),
(11, '鬼吹灯之精绝古城', '天下霸唱', 'https://img2.doubanio.com/view/subject/l/public/s27683781.jpg', '胡八一上山下乡来到云南，在一片原始森林中遭遇了一系列诡异事件。', '悬疑', '悬疑,盗墓,探险,恐怖', '已完结', '2026-05-18 15:00:00'),
(12, '诛仙', '萧鼎', 'https://img2.doubanio.com/view/subject/l/public/s29619565.jpg', '天地不仁，以万物为刍狗！这世间本是没有什么神仙，但自太古以来，人类眼见周遭世界。', '仙侠', '仙侠,修真,爱情', '已完结', '2026-05-22 17:00:00'),
(13, '庆余年', '猫腻', 'https://img2.doubanio.com/view/subject/l/public/s29651121.jpg', '一个年轻的病人，因为一次毫不意外的经历，重生到一个完全不同的世界。', '玄幻', '玄幻,权谋,穿越', '已完结', '2026-05-25 18:00:00'),
(14, '斗破苍穹', '天蚕土豆', 'https://img2.doubanio.com/view/subject/l/public/s29053580.jpg', '这里是属于斗气的世界，没有花俏艳丽的魔法，有的仅仅是繁衍到巅峰的斗气！', '玄幻', '玄幻,升级流,热血', '已完结', '2026-06-01 19:00:00'),
(15, '撒哈拉的故事', '三毛', 'https://img2.doubanio.com/view/subject/l/public/s6384944.jpg', '三毛作品中脍炙人口的作品，由12篇精彩动人的散文结合而成。', '散文', '散文,旅行,生活', '已完结', '2026-03-20 06:00:00'),
(16, '小王子', '安托万·德·圣-埃克苏佩里', 'https://img2.doubanio.com/view/subject/l/public/s1103152.jpg', '以一位飞行员作为故事叙述者，讲述了小王子从自己星球出发前往地球的过程中，所经历的各种历险。', '童话', '童话,哲学,经典', '已完结', '2026-02-14 05:00:00'),
(17, '天龙八部', '金庸', 'https://img2.doubanio.com/view/subject/l/public/s1070959.jpg', '以北宋哲宗时代为背景，通过宋、辽、大理、西夏、吐蕃等王国之间的武林恩怨和民族矛盾。', '武侠', '武侠,金庸,江湖', '已完结', '2026-04-10 12:00:00'),
(18, '雪中悍刀行', '烽火戏诸侯', 'https://img2.doubanio.com/view/subject/l/public/s27683781.jpg', '有个白狐儿脸，佩双刀绣冬春雷，要做那天下第一。', '玄幻', '玄幻,武侠,江湖', '连载中', '2026-06-05 20:00:00'),
(19, '三体II·黑暗森林', '刘慈欣', 'https://img2.doubanio.com/view/subject/l/public/s29619565.jpg', '三体人在利用魔法般的科技锁死了地球人的科学之后，庞大的宇宙舰队开始向地球进发。', '科幻', '科幻,硬科幻,宇宙', '已完结', '2026-05-02 10:00:00'),
(20, '长安十二时辰', '马伯庸', 'https://img2.doubanio.com/view/subject/l/public/s29651121.jpg', '唐天宝三年，元月十四日，长安。上元节辉煌灯火亮起之时，等待他们的，将是场吞噬一切的劫难。', '历史', '历史,悬疑,唐朝', '已完结', '2026-05-28 21:00:00');

-- =============================================
-- Seed Data: 种子评论 (40条)
-- =============================================
INSERT IGNORE INTO comments (id, book_id, user_id, content, likes, created_at) VALUES
(1, 1, 1, '这本书是我读过的科幻小说中最震撼的一本！黑暗森林法则太精彩了。', 128, '2026-06-05 08:30:00'),
(2, 1, 2, '刘慈欣的想象力令人叹服，三体世界的构建非常宏大。', 95, '2026-06-04 15:20:00'),
(3, 1, 3, '读完三体后，我对宇宙有了全新的认识。强烈推荐！', 76, '2026-06-03 12:45:00'),
(4, 2, 2, '活着让我泪流满面，余华把人生的苦难写得如此真实。', 203, '2026-06-05 10:10:00'),
(5, 2, 4, '福贵的一生就是中国近现代史的缩影，太感人了。', 156, '2026-06-04 18:00:00'),
(6, 3, 5, '百年孤独是一部需要静下心来慢慢品味的巨著。', 89, '2026-06-02 09:15:00'),
(7, 4, 1, '平凡的世界教会了我什么是坚韧不拔的精神。', 167, '2026-06-05 14:30:00'),
(8, 4, 3, '孙少平和孙少安两兄弟的故事激励了无数人。', 142, '2026-06-04 11:00:00'),
(9, 5, 4, '盗墓笔记比鬼吹灯更注重悬疑氛围的营造，很过瘾！', 234, '2026-06-05 16:45:00'),
(10, 5, 2, '小哥太帅了！整个系列我都刷了三遍了。', 198, '2026-06-04 20:20:00'),
(11, 6, 3, '哈利波特是我童年的回忆，永远的神作。', 321, '2026-06-05 07:00:00'),
(12, 6, 5, 'J.K.罗琳创造了一个令人向往的魔法世界。', 267, '2026-06-03 19:30:00'),
(13, 7, 1, '围城里的比喻简直绝了，钱钟书的文笔让人佩服。', 176, '2026-06-02 14:00:00'),
(14, 7, 4, '婚姻就像一座围城，外面的人想进去，里面的人想出来。', 145, '2026-06-01 10:30:00'),
(15, 8, 2, '东野圭吾的巅峰之作，结局让人细思极恐。', 289, '2026-06-05 21:00:00'),
(16, 8, 3, '亮司和雪穗之间的感情太复杂了，看完久久不能平静。', 198, '2026-06-04 16:15:00'),
(17, 9, 5, '红楼梦每次读都有新的感悟，真是中国文学的巅峰。', 156, '2026-06-03 08:00:00'),
(18, 9, 1, '大观园里的故事百读不厌，曹雪芹的笔力无人能及。', 134, '2026-06-02 12:00:00'),
(19, 10, 2, '村上春树的文字有一种独特的忧郁美感。', 178, '2026-06-05 13:20:00'),
(20, 10, 4, '挪威的森林陪我度过了大学时光，每次重读都有不同感受。', 145, '2026-06-04 09:45:00'),
(21, 11, 3, '天下霸唱的文笔比南派三叔更粗犷，各有特色。', 167, '2026-06-05 11:30:00'),
(22, 11, 5, '精绝古城这一段写得最吓人，晚上都不敢关灯了！', 234, '2026-06-04 22:00:00'),
(23, 12, 1, '诛仙是我看的第一本网络小说，从此入坑网文。', 456, '2026-06-05 17:00:00'),
(24, 12, 4, '碧瑶和小凡的感情线写得太好哭了。', 389, '2026-06-04 14:30:00'),
(25, 13, 2, '猫腻的文笔在网络作家里是顶尖的，庆余年格局很大。', 298, '2026-06-05 19:15:00'),
(26, 13, 3, '范闲这个角色塑造得太好了，聪明又有人情味。', 245, '2026-06-04 08:00:00'),
(27, 14, 4, '斗破苍穹虽然有些套路，但就是看得停不下来。', 567, '2026-06-05 22:30:00'),
(28, 14, 5, '三十年河东三十年河西，莫欺少年穷！这句太经典了。', 432, '2026-06-04 19:00:00'),
(29, 15, 1, '三毛的文字有一种让人安静下来的力量。', 234, '2026-06-03 15:40:00'),
(30, 15, 2, '撒哈拉的故事让我对沙漠生活充满了向往。', 189, '2026-06-02 16:20:00'),
(31, 16, 3, '小王子是一本给成年人看的童话，每次读都有新感悟。', 345, '2026-06-05 06:30:00'),
(32, 16, 5, '"只有用心才能看清，本质的东西眼睛是看不见的。" 最喜欢这句话。', 298, '2026-06-04 12:00:00'),
(33, 17, 1, '金庸先生的小说里，天龙八部的格局最为宏大。', 234, '2026-06-05 15:00:00'),
(34, 17, 4, '乔峰是真正的悲剧英雄，每次看到结局都忍不住落泪。', 198, '2026-06-04 10:30:00'),
(35, 18, 2, '雪中悍刀行的江湖气太足了，烽火戏诸侯是懂武侠的。', 345, '2026-06-06 08:00:00'),
(36, 18, 3, '徐凤年这个角色越看越有魅力，期待后续更新！', 267, '2026-06-06 09:30:00'),
(37, 19, 1, '黑暗森林法则的设定太精妙了，比第一部更上一层楼。', 189, '2026-06-05 20:00:00'),
(38, 19, 5, '面壁计划和破壁人的对决写得真精彩！', 156, '2026-06-04 17:45:00'),
(39, 20, 2, '马伯庸的历史小说细节考究，十二时辰节奏紧凑。', 234, '2026-06-06 10:00:00'),
(40, 20, 4, '长安城的一天写得惊心动魄，拍成剧也很还原！', 189, '2026-06-06 11:15:00');
