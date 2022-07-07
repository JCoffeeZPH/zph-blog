-- MySQL dump 10.13  Distrib 8.0.25, for macos11 (x86_64)
--
-- Host: localhost    Database: m_blob
-- ------------------------------------------------------
-- Server version	8.0.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `about_me_tab`
--

DROP TABLE IF EXISTS `about_me_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `about_me_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `music_id` bigint NOT NULL COMMENT '网易云音乐id',
  `comment_enabled` tinyint(1) NOT NULL DEFAULT '1' COMMENT '评论开关',
  `content` longtext NOT NULL COMMENT '内容',
  `user_id` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `about_me_tab`
--

LOCK TABLES `about_me_tab` WRITE;
/*!40000 ALTER TABLE `about_me_tab` DISABLE KEYS */;
INSERT INTO `about_me_tab` VALUES (1,'my about meaaadss',423015580,1,'hhhhsadasad',2);
/*!40000 ALTER TABLE `about_me_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `basic_setting_tab`
--

DROP TABLE IF EXISTS `basic_setting_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `basic_setting_tab` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name_en` varchar(256) NOT NULL COMMENT '博客名',
  `name_zh` varchar(256) NOT NULL,
  `value` longtext,
  `user_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `basic_setting_tab`
--

LOCK TABLES `basic_setting_tab` WRITE;
/*!40000 ALTER TABLE `basic_setting_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `basic_setting_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_tab`
--

DROP TABLE IF EXISTS `blog_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blog_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
  `first_picture` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章首图，用于随机文章展示',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章正文',
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `is_published` tinyint(1) NOT NULL DEFAULT '0' COMMENT '公开或私密',
  `is_recommend` tinyint(1) NOT NULL DEFAULT '0' COMMENT '推荐开关',
  `is_appreciation` tinyint(1) NOT NULL DEFAULT '0' COMMENT '赞赏开关',
  `is_comment_enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '评论开关',
  `create_time` int unsigned NOT NULL COMMENT '创建时间',
  `update_time` int unsigned NOT NULL COMMENT '更新时间',
  `views` int NOT NULL DEFAULT '0' COMMENT '浏览次数',
  `words` int NOT NULL COMMENT '文章字数',
  `read_time` int NOT NULL COMMENT '阅读时长(分钟)',
  `category_id` bigint NOT NULL COMMENT '文章分类',
  `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '密码保护',
  `user_id` bigint DEFAULT NULL COMMENT '文章作者',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `blog_tab_title_uindex` (`title`),
  KEY `type_id` (`category_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_tab`
--

LOCK TABLES `blog_tab` WRITE;
/*!40000 ALTER TABLE `blog_tab` DISABLE KEYS */;
INSERT INTO `blog_tab` VALUES (15,'test','asdas','asdas','asdassad',0,0,0,0,1639129833,1639390949,0,5,1,4,0,'',2),(16,'哈哈哈','aaa','asdas','阿斯达',1,1,1,0,1639130009,1639130009,0,5,1,6,0,'zhangphh=520',2),(17,'这是一个测试','abc.com','> asdasd\n\n|column1|column2|column3|\n|-|-|-|\n|content1|content2|content3|\n','```\ngo test\n\n```\n\n\n',1,1,1,1,1639391252,1639391252,0,73,1,7,1,'',2),(18,'测试','啊啊','==asdas==\n# 一级标题','阿斯达',0,0,0,0,1639462444,1639463345,0,24,1,18,0,'搜索',2),(19,'test操作日志','好好好','萨达==++标记++==','萨达萨阿三```\nlanguage\n\n```\n',1,1,0,1,1639721772,1639721772,0,20,1,22,0,'',2),(20,'go sync包下的RWmutex','https://s4.ax1x.com/2021/12/29/Tc6k5V.png','### RWMutex\n\n|      |  读  |  写  |\n| :--: | :--: | :--: |\n|  读  |  Y   |  N   |\n|  写  |  N   |  N   |\n\n读写锁在go中的syncRWMutex实现，其结构体如下：\n\n```go\ntype RWMutex struct {\n	w           Mutex  // held if there are pending writers\n	writerSem   uint32 // semaphore for writers to wait for completing readers\n	readerSem   uint32 // semaphore for readers to wait for completing writers\n	readerCount int32  // number of pending readers\n	readerWait  int32  // number of departing readers\n}\n```\n\n发现除了有一个互斥锁之外还持有两个信号量，一个是写等待读，另一个是读等待写\n\nreaderCount字段存储了当前正在执行的读操作的数量，最后的readerWait表示当写操作被阻塞时等待的读操作的个数\n\n#### 1、读锁\n\n```go\nfunc (rw *RWMutex) RLock() {\n	if race.Enabled {\n		_ = rw.w.state\n		race.Disable()\n	}\n	if atomic.AddInt32(&rw.readerCount, 1) < 0 {\n		// A writer is pending, wait for it.\n		runtime_SemacquireMutex(&rw.readerSem, false, 0)\n	}\n	if race.Enabled {\n		race.Enable()\n		race.Acquire(unsafe.Pointer(&rw.readerSem))\n	}\n}\n```\n\n读锁的加锁方式比较简单，通过atomic.AddInt32方法为readerCount加一，如果该方法返回了负数说明当前有goroutine获得了写锁，当前goroutine就会调用runtime_SemacquireMutex方法陷入休眠等待唤醒。\n\n如果没有写操作获取当前互斥锁，当前方法就会在readerCount加一后返回，当goroutine想要释放读锁时会调用RUnlock方法：\n\n```go\nfunc (rw *RWMutex) RUnlock() {\n	if race.Enabled {\n		_ = rw.w.state\n		race.ReleaseMerge(unsafe.Pointer(&rw.writerSem))\n		race.Disable()\n	}\n	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {\n		// Outlined slow-path to allow the fast-path to be inlined\n		rw.rUnlockSlow(r)\n	}\n	if race.Enabled {\n		race.Enable()\n	}\n}\n\nfunc (rw *RWMutex) rUnlockSlow(r int32) {\n	if r+1 == 0 || r+1 == -rwmutexMaxReaders {\n		race.Enable()\n		throw(\"sync: RUnlock of unlocked RWMutex\")\n	}\n	// A writer is pending.\n	if atomic.AddInt32(&rw.readerWait, -1) == 0 {\n		// The last reader unblocks the writer.\n		runtime_Semrelease(&rw.writerSem, false, 1)\n	}\n}\n```\n\n该方法会在减少正在读资源的readerCount，当前方法如果遇到了返回值小于0的情况，说明有一个正在进行的写操作，在这时就应该通过rUnlockSlow方法减少当前写操作等待的读操作数readerWait，并在所有的读操作都被释放后触发写操作的信号量writerSem，writerSem在被触发后，尝试获取读写锁的进程就会被唤醒并获得锁。\n\n#### 2、读写锁\n\n当资源想要获取读写锁时，就需要通过Lock方法了，在Lock方法中首先调用了读写互斥持有的Mutex的Lock方法，以此保证其他获取读写锁的goroutine进入等待状态，随后调用atomic.AddInt32方法来阻塞后续的读操作。\n\n```go\nfunc (rw *RWMutex) Lock() {\n	if race.Enabled {\n		_ = rw.w.state\n		race.Disable()\n	}\n	// First, resolve competition with other writers.\n	rw.w.Lock()\n	// Announce to readers there is a pending writer.\n	r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders\n	// Wait for active readers.\n	if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {\n		runtime_SemacquireMutex(&rw.writerSem, false, 0)\n	}\n	if race.Enabled {\n		race.Enable()\n		race.Acquire(unsafe.Pointer(&rw.readerSem))\n		race.Acquire(unsafe.Pointer(&rw.writerSem))\n	}\n}\n```\n\n如果当前仍有其他goroutine持有互斥的读锁，该goroutine就会调用runtime_SemacquireMutex方法就如休眠状态，等待读锁释放时触发writerSem信号量将当前协程唤醒。\n\n对资源的读写操作完成之后就会通过atomic.AddInt32方法变回正数并通过for循环触发所有由于获取读锁而陷入等待的goroutine：\n\n```go\nfunc (rw *RWMutex) Unlock() {\n	if race.Enabled {\n		_ = rw.w.state\n		race.Release(unsafe.Pointer(&rw.readerSem))\n		race.Disable()\n	}\n\n	// Announce to readers there is no active writer.\n	r := atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)\n	if r >= rwmutexMaxReaders {\n		race.Enable()\n		throw(\"sync: Unlock of unlocked RWMutex\")\n	}\n	// Unblock blocked readers, if any.\n	for i := 0; i < int(r); i++ {\n		runtime_Semrelease(&rw.readerSem, false, 0)\n	}\n	// Allow other writers to proceed.\n	rw.w.Unlock()\n	if race.Enabled {\n		race.Enable()\n	}\n}\n```\n\n在释放锁的方法最后，释放持有的互斥锁让其他的协程能够重新获取读写锁\n\nRWMutex站在了Mutex的肩膀上，因此实现简单很多\n\n**总结：**\n\n* readerSem，读写锁释放时通知由于获取读锁等待的goroutine\n* writerSem，读锁释放时通知由于获取读写锁等待的goroutine\n* w互斥锁，保证写操作之间的互斥\n* readerCount，统计当前进行读操作的协程数，触发写锁时会将其减少rwmutexMaxReaders阻塞后续的读操作\n* readerWait，当前读写锁等待的进行读操作的协程数，再出发Lock之后的每次RUnlock都会将其减一，当它归零时该goroutine就会获得读写锁\n* 当读写锁被释放Unlock时首先会通知所有的读操作，然后才会释放持有的互斥锁，这样能够保证读操作不会被连续的写操作饿死。\n\n\n\n参考文章：[Golang 并发编程之同步原语](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&amp;mid=2247484379&amp;idx=1&amp;sn=1a2abc6f639a34e62f3a5a0fcd774a71&amp;chksm=fa80d24ccdf75b5a70d45168ad9e3a55dd258c1dd57147166a86062ee70d909ff1e5b0ba2bcb&amp;token=183756123&amp;lang=zh_CN#rd)\n\n\n\n','<p>{&quot;markdown&quot;:&quot;\\u003cp\\u003ego的读写锁\\u003c/p\\u003e\\n\\n\\u003cp\\u003e\\u003cimg src=&quot;https://s4.ax1x.com/2021/12/29/Tc6k5V.png&quot; alt=&quot;go&quot; /\\u003e\\u003c/p\\u003e\\n\\n\\u003cp\\u003e加一张图\\u003c/p\\u003e\\n&quot;}</p>\n',1,1,0,1,1640762513,1641547612,22,5488,19,4,1,'',2);
/*!40000 ALTER TABLE `blog_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_tag`
--

DROP TABLE IF EXISTS `blog_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blog_tag` (
  `blog_id` bigint NOT NULL,
  `tag_id` bigint NOT NULL,
  KEY `blog_tag_index` (`blog_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_tag`
--

LOCK TABLES `blog_tag` WRITE;
/*!40000 ALTER TABLE `blog_tag` DISABLE KEYS */;
INSERT INTO `blog_tag` VALUES (2,4),(10,12),(10,14),(10,8),(11,12),(11,14),(11,8),(13,12),(13,14),(13,8),(14,6),(15,7),(15,8),(16,0),(16,28),(17,28),(17,14),(17,29),(18,6),(18,8),(18,31),(18,12),(19,32),(19,33),(20,34);
/*!40000 ALTER TABLE `blog_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category_tab`
--

DROP TABLE IF EXISTS `category_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `category_tab_category_name_uindex` (`category_name`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category_tab`
--

LOCK TABLES `category_tab` WRITE;
/*!40000 ALTER TABLE `category_tab` DISABLE KEYS */;
INSERT INTO `category_tab` VALUES (3,'java',1),(4,'go',2),(5,'学习笔记',2),(6,'ssss',2),(7,'test_go',1),(17,'admin',2),(18,'自定',2),(19,'asa ',2),(22,'新建',2);
/*!40000 ALTER TABLE `category_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `city_visitor_tab`
--

DROP TABLE IF EXISTS `city_visitor_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `city_visitor_tab` (
  `city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '城市名称',
  `pv` int NOT NULL COMMENT '页面访问量',
  `user_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `city_visitor_tab`
--

LOCK TABLES `city_visitor_tab` WRITE;
/*!40000 ALTER TABLE `city_visitor_tab` DISABLE KEYS */;
INSERT INTO `city_visitor_tab` VALUES ('深圳市',186,2),('北京市',5,1);
/*!40000 ALTER TABLE `city_visitor_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment_tab`
--

DROP TABLE IF EXISTS `comment_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像(图片路径)',
  `create_time` int unsigned NOT NULL COMMENT '评论时间',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '评论者ip地址',
  `is_published` tinyint(1) NOT NULL DEFAULT '1' COMMENT '公开或回收站',
  `is_admin_comment` tinyint(1) NOT NULL DEFAULT '0' COMMENT '博主回复',
  `from_page` int NOT NULL COMMENT '0普通文章，1关于我页面，2友链页面',
  `is_notice` tinyint(1) NOT NULL DEFAULT '0' COMMENT '接收邮件提醒',
  `blog_id` bigint DEFAULT NULL COMMENT '所属的文章',
  `parent_comment_id` bigint NOT NULL COMMENT '父评论id，-1为根评论',
  `website` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '个人网站',
  `qq` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '如果评论昵称为QQ号，则将昵称和头像置为QQ昵称和QQ头像，并将此字段置为QQ号备份',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment_tab`
--

LOCK TABLES `comment_tab` WRITE;
/*!40000 ALTER TABLE `comment_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `comment_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exception_log_tab`
--

DROP TABLE IF EXISTS `exception_log_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exception_log_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方式',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作描述',
  `errors` text COLLATE utf8mb4_general_ci COMMENT '异常信息',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip',
  `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip来源',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作系统',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '浏览器',
  `create_time` int unsigned NOT NULL COMMENT '操作时间',
  `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'user-agent用户代理',
  `user_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exception_log_tab`
--

LOCK TABLES `exception_log_tab` WRITE;
/*!40000 ALTER TABLE `exception_log_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `exception_log_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `friend_tab`
--

DROP TABLE IF EXISTS `friend_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `friend_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '描述',
  `website` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '站点',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `is_published` tinyint(1) NOT NULL DEFAULT '1' COMMENT '公开或隐藏',
  `views` int NOT NULL DEFAULT '0' COMMENT '点击次数',
  `create_time` int unsigned NOT NULL COMMENT '创建时间',
  `user_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `friend_tab`
--

LOCK TABLES `friend_tab` WRITE;
/*!40000 ALTER TABLE `friend_tab` DISABLE KEYS */;
INSERT INTO `friend_tab` VALUES (2,'zhangphh1','test1','http://jcoffeezph.top','https://static.lty.fun/weblogo/my.jp',0,0,1639102906,2),(3,'zhangphh2','test2','http://www.baidu.com','https://uestcxxs.com/img/img.png',1,0,1639987173,2);
/*!40000 ALTER TABLE `friend_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `login_log_tab`
--

DROP TABLE IF EXISTS `login_log_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `login_log_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名称',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip',
  `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip来源',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作系统',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '浏览器',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '登录状态',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作描述',
  `create_time` int unsigned NOT NULL COMMENT '登录时间',
  `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'user-agent用户代理',
  `user_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `login_log_tab`
--

LOCK TABLES `login_log_tab` WRITE;
/*!40000 ALTER TABLE `login_log_tab` DISABLE KEYS */;
INSERT INTO `login_log_tab` VALUES (3,'Admin',' 58.250.178.130','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1,'登录成功',1639735359,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(4,'Admin',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1,'登录成功',1640160183,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(5,'Admin',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1,'登录成功',1640602109,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2);
/*!40000 ALTER TABLE `login_log_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `moment_tab`
--

DROP TABLE IF EXISTS `moment_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `moment_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '动态内容',
  `create_time` int unsigned NOT NULL COMMENT '创建时间',
  `likes` int NOT NULL DEFAULT '0' COMMENT '点赞数量',
  `is_published` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否公开',
  `update_time` int unsigned NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `moment_tab`
--

LOCK TABLES `moment_tab` WRITE;
/*!40000 ALTER TABLE `moment_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `moment_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `operation_log_tab`
--

DROP TABLE IF EXISTS `operation_log_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `operation_log_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作者用户名',
  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方式',
  `param` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求参数',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作描述',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip',
  `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip来源',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作系统',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '浏览器',
  `times` int NOT NULL COMMENT '请求耗时（毫秒）',
  `create_time` int unsigned NOT NULL COMMENT '操作时间',
  `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'user-agent用户代理',
  `user_id` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=277 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `operation_log_tab`
--

LOCK TABLES `operation_log_tab` WRITE;
/*!40000 ALTER TABLE `operation_log_tab` DISABLE KEYS */;
INSERT INTO `operation_log_tab` VALUES (2,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 58.250.178.130','中国-广东-深圳','\"macOS\"','\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"',12665,1639707659,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(3,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 58.250.178.130','中国-广东-深圳','\"macOS\"','\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"',185697,1639707832,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(4,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 58.250.178.130','中国-广东-深圳','','\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Microsoft Edge\";v=\"92\"',391219,1639708037,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84',2),(5,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 58.250.178.130','中国-广东-深圳','','\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Microsoft Edge\";v=\"92\"',411167,1639708057,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84',2),(6,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 121.15.15.162','中国-广东-深圳','\"macOS\"','\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"',460,1639708237,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(7,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 121.15.15.162','中国-广东-深圳','\"macOS\"','\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\"',263,1639708252,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(8,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 58.250.178.130','中国-广东-深圳','','\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Microsoft Edge\";v=\"92\"',264,1639708681,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36 Edg/92.0.902.84',2),(9,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 121.15.15.162','中国-广东-深圳','macOS 10.15.7(Catalina)','Chrome 96.0.4664.110',11291,1639713466,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',1),(10,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 121.15.15.162','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',574,1639713550,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',1),(11,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 121.15.15.162','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',591,1639713623,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',1),(12,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','获取全部分类',' 121.15.15.162','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',884,1639713638,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(13,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',1644,1639721727,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(14,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','创建blog',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',560,1639721773,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(15,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',559,1639721773,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(16,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',8203,1639726296,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(17,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',513,1639726307,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(18,'Admin','/blog/api/v1/admin/operation_logs?page=2&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',588,1639726488,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(19,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',537,1639726491,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(20,'Admin','/blog/api/v1/admin/operation_logs?page=2&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',534,1639726495,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(21,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',822,1639726584,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(22,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',553,1639726678,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(23,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',646,1639726763,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(24,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',565,1639726802,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(25,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',548,1639726816,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(26,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',802,1639726850,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(27,'Admin','/blog/api/v1/admin/operation_log/1','DELETE','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',525,1639726855,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(28,'Admin','/blog/api/v1/admin/operation_logs?page=1&per_page=10','GET','','',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',541,1639726855,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(30,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',547,1639727012,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(33,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.130','中国-广东-深圳','macOS 10.15.7','Chrome 96.0.4664.110',4775,1639727066,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(34,'Admin','/blog/api/v1/admin/categories?page=1&per_page=10','GET','','分页获取分裂',' 58.250.178.130','|中国|广东|深圳','macOS 10.15.7(Catalina)','Chrome 96.0.4664.110',798,1639735024,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(35,'Admin','/blog/api/v1/admin/categories?page=1&per_page=10','GET','','分页获取分裂',' 58.250.178.130','|中国|广东|深圳','macOS 10.15.7(Catalina)','Chrome 96.0.4664.110',496,1639735051,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(36,'Admin','/blog/api/v1/admin/tags?page=1&per_page=10','GET','','分页获取标签',' 58.250.178.130','|中国|广东|深圳','macOS 10.15.7(Catalina)','Chrome 96.0.4664.110',7509,1639735118,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(37,'Admin','/blog/api/v1/admin/categories?page=1&per_page=10','GET','','分页获取分分类',' 58.250.178.130','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',4641,1639735337,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(38,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',6418,1639969273,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(39,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',527,1639969280,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(40,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',541,1639969375,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(41,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',498,1639969380,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(42,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',488,1639969400,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(43,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',498,1639969436,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(44,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',514,1639969444,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(45,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',985,1639969987,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(46,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',520,1639970243,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(47,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',500,1639970284,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(48,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',516,1639970398,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(49,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',647,1639970581,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(50,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',532,1639970629,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(51,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1046,1639971976,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(52,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',471,1639972024,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(53,'Admin','/blog/api/v1/admin/site_settings','POST','','更新站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',6502,1639972185,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(54,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',477,1639972275,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(55,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',460,1639972291,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(56,'Admin','/blog/api/v1/admin/site_settings','GET','','更新站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',503,1639972291,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(57,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',739,1639972366,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(58,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1108,1639972541,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(59,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',492,1639972560,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(60,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',26132,1639972618,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(61,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',2937,1639972772,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(62,'Admin','/blog/api/v1/admin/site_settings','POST','','更新站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',66986,1639972887,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(63,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',473,1639972913,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(64,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',574,1639972943,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(65,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1125,1639973061,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(66,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',440,1639973085,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(67,'Admin','/blog/api/v1/admin/site_settings','POST','','更新站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',523,1639973085,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(68,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',489,1639973090,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(69,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',475,1639973111,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(70,'','/blog/api/v1/admin/site_settings','OPTIONS','','更新站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',525,1639973111,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',0),(71,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',440,1639973115,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(72,'','/blog/api/v1/admin/site_settings','OPTIONS','','更新站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',516,1639973124,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',0),(73,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',463,1639973124,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(74,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',510,1639973127,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(75,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',465,1639973429,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(76,'Admin','/blog/api/v1/admin/site_settings','POST','','更新站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',512,1639973429,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(77,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',484,1639973441,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(78,'Admin','/blog/api/v1/admin/site_settings','POST','','更新站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',533,1639973441,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(79,'Admin','/blog/api/v1/admin/moments?page=1&per_page=10','GET','','分页获取动态',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',784,1639973729,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(80,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1219,1639981210,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(81,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',761,1639981251,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(82,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',468,1639981253,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(83,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',494,1639981283,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(84,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',475,1639981294,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(85,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1008,1639986767,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(86,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',492,1639986777,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(87,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',481,1639986779,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(88,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',477,1639986825,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(89,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',498,1639986825,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(90,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',467,1639987016,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(91,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',485,1639987016,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(92,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',466,1639987028,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(93,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',491,1639987028,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(94,'Admin','/blog/api/v1/admin/friend_chain/1','PUT','','更新友链信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',462,1639987081,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(95,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',465,1639987081,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(96,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',481,1639987114,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(97,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',626,1639987114,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(98,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',460,1639987125,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(99,'Admin','/blog/api/v1/admin/friend_chain/1','PUT','','更新友链信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',510,1639987125,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(100,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','新加友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',475,1639987174,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(101,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',479,1639987174,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(102,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',523,1639987196,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(103,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','删除友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',558,1639987196,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(104,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',927,1639988353,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(105,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',931,1639988353,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(106,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',499,1639988509,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(107,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',513,1639988509,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(108,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',489,1639988513,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(109,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',516,1639988513,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(110,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',486,1639988657,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(111,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',510,1639988657,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(112,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',463,1639988664,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(113,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',459,1639988671,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(114,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',482,1639988671,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(115,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',914,1639989840,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(116,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',928,1639989840,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(117,'Admin','/blog/api/v1/admin/friend_chain/content','PUT','','修改友链页面信息',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',475,1639989852,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(118,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',480,1639989852,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(119,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',542,1639989938,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(120,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',554,1639989938,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(121,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',727,1639990155,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(122,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',758,1639990156,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(123,'Admin','/blog/api/v1/admin/friend/is_published/2?is_published=false','PUT','','修改友链是否公开开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',6357,1639990242,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(124,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',500,1639990247,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(125,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',508,1639990247,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(126,'Admin','/blog/api/v1/admin/friend/is_published/2?is_published=true','PUT','','修改友链是否公开开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',3660,1639990254,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(127,'Admin','/blog/api/v1/admin/friend/is_published/2?is_published=false','PUT','','修改友链是否公开开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',491,1639990258,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(128,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1052,1639990410,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(129,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1102,1639990410,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(130,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=false','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',698,1639990483,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(131,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=true','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',472,1639990488,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(132,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=false','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',456,1639990489,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(133,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=true','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',446,1639990517,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(134,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=false','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',43842,1639990574,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(135,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',456,1639990607,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(136,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',495,1639990609,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(137,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',707,1639990755,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(138,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',4352,1639990761,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(139,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',4392,1639990761,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(140,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=false','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',4348,1639990771,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(141,'Admin','/blog/api/v1/admin/site_settings','GET','','获取站点配置',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',461,1639990777,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(142,'Admin','/blog/api/v1/admin/friend_info','GET','','获取友链页面信息',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',449,1639990778,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(143,'Admin','/blog/api/v1/admin/friend_chains?page=1&per_page=10','GET','','获取友链',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',478,1639990778,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(144,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=true','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',460,1639990782,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(145,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=true','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',483,1639990782,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(146,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=true','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',459,1639990783,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(147,'Admin','/blog/api/v1/admin/friend_chain/comment_enabled?comment_enabled=false','PUT','','修改友链页面评论开关',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',454,1639990791,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(148,'Admin','/blog/api/v1/admin/about_me','GET','','获取AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1042,1639996364,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(149,'Admin','/blog/api/v1/admin/about_me','PUT','','修稿AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',480,1639996375,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(150,'Admin','/blog/api/v1/admin/about_me','GET','','获取AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',542,1639996380,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(151,'Admin','/blog/api/v1/admin/about_me','GET','','获取AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',480,1639996398,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(152,'Admin','/blog/api/v1/admin/about_me','PUT','','修稿AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',742,1639996443,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(153,'Admin','/blog/api/v1/admin/about_me','GET','','获取AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',529,1639996633,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(154,'Admin','/blog/api/v1/admin/about_me','PUT','','修稿AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',465,1639996647,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(155,'Admin','/blog/api/v1/admin/about_me','GET','','获取AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',563,1639996651,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(156,'Admin','/blog/api/v1/admin/about_me','PUT','','修稿AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',501,1639996681,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(157,'Admin','/blog/api/v1/admin/about_me','PUT','','修稿AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',2508,1639996773,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(158,'Admin','/blog/api/v1/admin/about_me','GET','','获取AboutMe信息',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',474,1639996777,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(159,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.131','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',490,1640160298,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(160,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1113,1640602109,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(161,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',493,1640602127,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(162,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',505,1640602146,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(163,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',480,1640602166,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(164,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',488,1640602189,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(165,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',498,1640602324,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(166,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',480,1640602440,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(167,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1224,1640602895,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(168,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',500,1640603159,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(169,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',528,1640603172,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(170,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',551,1640603182,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(171,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',486,1640603250,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(172,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',524,1640603285,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(173,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',484,1640603299,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(174,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',4402,1640603464,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(175,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',481,1640603475,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(176,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',490,1640603485,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(177,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',796,1640603595,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(178,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1226,1640762243,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(179,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',478,1640762250,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(180,'Admin','/blog/api/v1/admin/blog','POST','','创建blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',574,1640762514,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(181,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',489,1640762514,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(182,'Admin','/blog/api/v1/admin/tags?page=1&per_page=10','GET','','分页获取标签',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',478,1640762535,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(183,'Admin','/blog/api/v1/admin/tags?page=2&per_page=10','GET','','分页获取标签',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',438,1640762539,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(184,'Admin','/blog/api/v1/admin/tag','PUT','','更新标签信息',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',467,1640762544,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(185,'Admin','/blog/api/v1/admin/tags?page=2&per_page=10','GET','','分页获取标签',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',445,1640762544,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(186,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',596,1640762738,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(187,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',462,1640762740,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(188,'Admin','/blog/api/v1/admin/blog/20','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',487,1640762741,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(189,'','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','OPTIONS','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',469,1640762771,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',0),(190,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',494,1640762771,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(191,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',468,1640762819,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(192,'Admin','/blog/api/v1/admin/blog/20','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',479,1640762819,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(193,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',486,1640762858,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(194,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',477,1640762858,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(195,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',482,1640762868,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(196,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',494,1640762868,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(197,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',472,1640762931,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(198,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',466,1640762931,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(199,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',442,1640762945,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(200,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',469,1640762945,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(201,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',576,1640763026,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(202,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',461,1640763026,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(203,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',469,1640763039,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(204,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',482,1640763039,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(205,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',458,1640763061,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(206,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',483,1640763062,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(207,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',503,1640763140,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(208,'Admin','/blog/api/v1/admin/moments?page=1&per_page=10','GET','','分页获取动态',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',474,1640763145,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(209,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',475,1640763146,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(210,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',982,1640766564,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(211,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1007,1640766564,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(212,'','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','OPTIONS','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',486,1640766569,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',0),(213,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',483,1640766570,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(214,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',471,1640766587,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(215,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',484,1640766587,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(216,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',449,1640766595,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(217,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',474,1640766596,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(218,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','更新blog状态',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',478,1640766600,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(219,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',474,1640766600,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(220,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',823,1640845693,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(221,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',730,1640845775,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(222,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',567,1640845849,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(223,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',498,1640845851,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(224,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',515,1640845851,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(225,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',477,1640845866,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(226,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',448,1640845866,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(227,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','更新blog状态',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',467,1640845880,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(228,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',460,1640845880,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(229,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',457,1640845938,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(230,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',487,1640845938,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(231,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',640,1640846552,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(232,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',490,1640846552,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(233,'','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','OPTIONS','','更新blog状态',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',477,1640846595,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',0),(234,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',483,1640846595,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(235,'','/blog/api/v1/admin/blog/20','OPTIONS','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',460,1640846600,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',0),(236,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',470,1640846600,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(237,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',489,1640846611,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(238,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',501,1640846611,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(239,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',482,1640846653,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(240,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',494,1640846653,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(241,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',59056,1640846721,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(242,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',2175,1640846727,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(243,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',456,1640846727,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(244,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',461,1640846815,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(245,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',477,1640846815,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(246,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',2641,1640846824,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(247,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',444,1640846824,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(248,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',460,1640846851,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(249,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',465,1640846851,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(250,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1005,1640846879,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(251,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',23546,1640846879,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(252,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',467,1640846879,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(253,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',733,1640846887,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(254,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',738,1640846887,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(255,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1789,1640846895,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(256,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.133','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',462,1640846896,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(257,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',14056,1641283208,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(258,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',24640,1641284069,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(259,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',69773,1641284143,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(260,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',4189,1641284946,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(261,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',12604,1641285798,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(262,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',20031,1641285825,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(263,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',163967,1641286009,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(264,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',214245,1641286236,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(265,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',1032,1641288014,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(266,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',470,1641288018,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(267,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',472,1641288018,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(268,'Admin','/blog/api/v1/admin/dashboard','GET','','查看首页仪表盘',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',3825,1641547071,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(269,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',542,1641547073,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(270,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',481,1641547075,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(271,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 121.15.15.162','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',522,1641547075,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(272,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',63463,1641547602,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(273,'Admin','/blog/api/v1/admin/blog/20','PUT','','更新blog',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',482,1641547612,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(274,'Admin','/blog/api/v1/admin/blogs?title=&page=1&per_page=10','GET','','按分页获取blogs',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',479,1641547613,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(275,'Admin','/blog/api/v1/admin/tags_categories','GET','','获取全部标签和分类',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',465,1641547615,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2),(276,'Admin','/blog/api/v1/admin/blog/20','GET','','根据Id获取blog',' 58.250.178.136','|中国|广东|深圳','macOS 10.15.7','Chrome 96.0.4664.110',452,1641547615,'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36',2);
/*!40000 ALTER TABLE `operation_log_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedule_job`
--

DROP TABLE IF EXISTS `schedule_job`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schedule_job` (
  `job_id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务id',
  `bean_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'spring bean名称',
  `method_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '方法名',
  `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数',
  `cron` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'cron表达式',
  `status` tinyint DEFAULT NULL COMMENT '任务状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`job_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedule_job`
--

LOCK TABLES `schedule_job` WRITE;
/*!40000 ALTER TABLE `schedule_job` DISABLE KEYS */;
/*!40000 ALTER TABLE `schedule_job` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedule_job_log`
--

DROP TABLE IF EXISTS `schedule_job_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schedule_job_log` (
  `log_id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务日志id',
  `job_id` bigint NOT NULL COMMENT '任务id',
  `bean_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'spring bean名称',
  `method_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '方法名',
  `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数',
  `status` tinyint NOT NULL COMMENT '任务执行结果',
  `error` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '异常信息',
  `times` int NOT NULL COMMENT '耗时（单位：毫秒）',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`log_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedule_job_log`
--

LOCK TABLES `schedule_job_log` WRITE;
/*!40000 ALTER TABLE `schedule_job_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `schedule_job_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `site_setting_tab`
--

DROP TABLE IF EXISTS `site_setting_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `site_setting_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name_zh` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `type` int DEFAULT NULL COMMENT '1基础设置，2页脚徽标，3资料卡，4友链信息',
  `user_id` int NOT NULL DEFAULT '2',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `site_setting_tab`
--

LOCK TABLES `site_setting_tab` WRITE;
/*!40000 ALTER TABLE `site_setting_tab` DISABLE KEYS */;
INSERT INTO `site_setting_tab` VALUES (1,'webTitleSuffix','网页标题后缀',' - Naccl\'s Blog',1,2),(2,'blogName','博客名称','Naccl\'s Blog',1,2),(3,'footerImgTitle','页脚图片标题','手机看本站',1,2),(4,'footerImgUrl','页脚图片路径','/img/qr.png',1,2),(5,'copyright','Copyright','{\"title\":\"Copyright © 2019 - 2020\",\"siteName\":\"NACCL\'S BLOG\"}',1,2),(6,'beian','ICP备案号','',1,2),(7,'badge','徽标','{\"title\":\"由 Spring Boot 强力驱动\",\"url\":\"https://spring.io/projects/spring-boot/\",\"subject\":\"Powered\",\"value\":\"Spring Boot\",\"color\":\"blue\"}',2,2),(8,'badge','徽标','{\"title\":\"Vue.js 客户端渲染\",\"url\":\"https://cn.vuejs.org/\",\"subject\":\"SPA\",\"value\":\"Vue.js\",\"color\":\"brightgreen\"}',2,2),(9,'badge','徽标','{\"title\":\"UI 框架 Semantic-UI\",\"url\":\"https://semantic-ui.com/\",\"subject\":\"UI\",\"value\":\"Semantic-UI\",\"color\":\"semantic-ui\"}',2,2),(10,'badge','徽标','{\"title\":\"阿里云提供服务器及域名相关服务\",\"url\":\"https://www.aliyun.com/\",\"subject\":\"VPS & DNS\",\"value\":\"Aliyun\",\"color\":\"blueviolet\"}',2,2),(11,'badge','徽标','{\"title\":\"jsDelivr 提供 CDN 加速服务\",\"url\":\"https://www.jsdelivr.com/\",\"subject\":\"CDN\",\"value\":\"jsDelivr\",\"color\":\"orange\"}',2,2),(12,'badge','徽标','{\"title\":\"GitHub 提供图床\",\"url\":\"https://github.com/\",\"subject\":\"OSS\",\"value\":\"GitHub\",\"color\":\"github\"}',2,2),(13,'badge','徽标','{\"title\":\"本站点采用 CC BY 4.0 国际许可协议进行许可\",\"url\":\"https://creativecommons.org/licenses/by/4.0/\",\"subject\":\"CC\",\"value\":\"BY 4.0\",\"color\":\"lightgray\"}',2,2),(14,'avatar','图片路径','/img/avatar.jpg',3,2),(15,'name','昵称','Naccl',3,2),(16,'rollText','滚动个签','云鹤当归天，天不迎我妙木仙;游龙当归海，海不迎我自来也',3,2),(17,'github','GitHub地址','https://github.com/Naccl',3,2),(18,'qq','QQ链接','http://sighttp.qq.com/authd?IDKEY=',3,2),(19,'bilibili','bilibili链接','https://space.bilibili.com/',3,2),(20,'netease','网易云音乐','https://music.163.com/#/user/home?id=',3,2),(21,'email','email','mailto:i@naccl.top',3,2),(22,'favorite','自定义','{\"title\":\"最喜欢的动漫 ?\",\"content\":\"异度侵入、春物语、NO GAME NO LIFE、实力至上主义的教室、辉夜大小姐、青春猪头少年不会梦到兔女郎学姐、路人女主、Re0、魔禁、超炮、俺妹、在下坂本、散华礼弥、OVERLORD、慎勇、人渣的本愿、白色相簿2、死亡笔记、DARLING in the FRANXX、鬼灭之刃\"}',3,2),(23,'favorite','自定义','{\"title\":\"最喜欢我的女孩子们 ?\",\"content\":\"芙兰达、土间埋、食蜂操祈、佐天泪爷、樱岛麻衣、桐崎千棘、02、亚丝娜、高坂桐乃、五更琉璃、安乐冈花火、一色彩羽、英梨梨、珈百璃、时崎狂三、可儿那由多、和泉纱雾、早坂爱\"}',3,2),(25,'reward','赞赏码路径','/img/reward.jpg',1,2),(26,'commentAdminFlag','博主评论标识','咕咕',1,2),(27,'friendContent','友链页面信息','随机排序，不分先后。欢迎交换友链~(￣▽￣)~*\n\n* 昵称：Naccl\n* 一句话：游龙当归海，海不迎我自来也。\n* 网址：[https://naccl.top](https://naccl.top)\n* 头像URL：[https://naccl.top/img/avatar.jpg](https://naccl.top/img/avatar.jpg)\n\n仅凭个人喜好添加友链，请在收到我的回复邮件后再于贵站添加本站链接。原则上已添加的友链不会删除，如果你发现自己被移除了，恕不另行通知，只需和我一样做就好。\nhhh\n',4,2),(28,'friendCommentEnabled','友链页面评论开关','0',4,2),(33,'favorite','自定义','{\"title\":\"zz\",\"content\":\"zz\"}',3,2),(34,'badge','徽标','{\"color\":\"a\",\"subject\":\"a\",\"title\":\"a\",\"url\":\"a\",\"value\":\"a\"}',2,2);
/*!40000 ALTER TABLE `site_setting_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag_tab`
--

DROP TABLE IF EXISTS `tag_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'red' COMMENT '标签颜色(可选)',
  `user_id` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `tag_tab_tag_name_uindex` (`tag_name`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag_tab`
--

LOCK TABLES `tag_tab` WRITE;
/*!40000 ALTER TABLE `tag_tab` DISABLE KEYS */;
INSERT INTO `tag_tab` VALUES (1,'go哈哈哈','orange',1),(2,'redis','teal',2),(3,'kafka','blue',2),(4,'mysql','black',2),(6,'算法','blue',2),(7,'go并发','orange',1),(8,'spring','pink',2),(12,'粉红色','pink',2),(14,'灰色色','grey',1),(28,'sda奥术大师大所','grey',1),(29,'test_ggg','pink',1),(30,'自定义','orange',2),(31,'啊啊啊','blue',2),(32,'aaasss ','blue',2),(33,'日志','blue',2),(34,'sync','red',2);
/*!40000 ALTER TABLE `tag_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_tab`
--

DROP TABLE IF EXISTS `user_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名，唯一',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像地址',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `create_time` int unsigned NOT NULL COMMENT '创建时间',
  `update_time` int unsigned NOT NULL COMMENT '更新时间',
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色访问权限',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_username_uindex` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_tab`
--

LOCK TABLES `user_tab` WRITE;
/*!40000 ALTER TABLE `user_tab` DISABLE KEYS */;
INSERT INTO `user_tab` VALUES (2,'Admin','admin','Admin','/img/avatar.jpg','zhangphh@test.com',1639102906,1639102906,'ROLE_admin');
/*!40000 ALTER TABLE `user_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `visit_log`
--

DROP TABLE IF EXISTS `visit_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `visit_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '访客标识码',
  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方式',
  `param` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求参数',
  `behavior` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '访问行为',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '访问内容',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip',
  `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip来源',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作系统',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '浏览器',
  `times` int NOT NULL COMMENT '请求耗时（毫秒）',
  `create_time` datetime NOT NULL COMMENT '访问时间',
  `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'user-agent用户代理',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `visit_log`
--

LOCK TABLES `visit_log` WRITE;
/*!40000 ALTER TABLE `visit_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `visit_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `visit_record_tab`
--

DROP TABLE IF EXISTS `visit_record_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `visit_record_tab` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `pv` int NOT NULL COMMENT '访问量',
  `uv` int NOT NULL COMMENT '独立用户',
  `date` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '日期"20210101"',
  `user_id` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `visit_record_tab`
--

LOCK TABLES `visit_record_tab` WRITE;
/*!40000 ALTER TABLE `visit_record_tab` DISABLE KEYS */;
INSERT INTO `visit_record_tab` VALUES (3,16,2,'20211227',2),(5,100,2,'20211227',3),(7,34,10,'20211226',2),(8,65,12,'20211225',2);
/*!40000 ALTER TABLE `visit_record_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `visitor`
--

DROP TABLE IF EXISTS `visitor`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `visitor` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '访客标识码',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip',
  `ip_source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip来源',
  `os` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作系统',
  `browser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '浏览器',
  `create_time` datetime NOT NULL COMMENT '首次访问时间',
  `last_time` datetime NOT NULL COMMENT '最后访问时间',
  `pv` int DEFAULT NULL COMMENT '访问页数统计',
  `user_agent` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'user-agent用户代理',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `visitor`
--

LOCK TABLES `visitor` WRITE;
/*!40000 ALTER TABLE `visitor` DISABLE KEYS */;
/*!40000 ALTER TABLE `visitor` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-10 18:10:00
