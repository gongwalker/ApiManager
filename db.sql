# ************************************************************
# Sequel Pro SQL dump
# Version 5438
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.30)
# Database: apidoc
# Generation Time: 2021-03-25 06:58:43 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table api
# ------------------------------------------------------------

DROP TABLE IF EXISTS `api`;

CREATE TABLE `api` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '接口编号',
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '接口分类id',
  `num` varchar(100) NOT NULL DEFAULT '' COMMENT '接口编号',
  `url` varchar(240) NOT NULL DEFAULT '' COMMENT '请求地址',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '接口名',
  `des` varchar(300) NOT NULL DEFAULT '' COMMENT '接口描述',
  `parameter` text COMMENT '请求参数{所有的主求参数,以json格式在此存放}',
  `parameter_text` longtext COMMENT '请求参数 存body请求体等',
  `memo` longtext COMMENT '备注',
  `re` longtext COMMENT '返回值',
  `lasttime` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '提后操作时间',
  `lastuid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最后修改uid',
  `isdel` tinyint(4) NOT NULL DEFAULT '0' COMMENT '{0:正常,1:删除}',
  `type` char(11) NOT NULL DEFAULT '' COMMENT '请求方式',
  `ord` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序(值越大,越靠前)',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='接口明细表';

LOCK TABLES `api` WRITE;
/*!40000 ALTER TABLE `api` DISABLE KEYS */;

INSERT INTO `api` (`id`, `aid`, `num`, `url`, `name`, `des`, `parameter`, `parameter_text`, `memo`, `re`, `lasttime`, `lastuid`, `isdel`, `type`, `ord`)
VALUES
	(1,1,'001','http://www.test.com/user/register','用户注册','用户注册接用此接口','{\"param_name\":[\"login_name\",\"password\",\"sms_code\"],\"param_type\":[\"string\",\"string\",\"string\"],\"param_cate\":[\"Y\",\"Y\",\"Y\"],\"param_default\":[\"\",\"\",\"\"],\"param_des\":[\"用户登录名\",\"用户密码\",\"短信验证码\"]}','','','',1615122392,1,0,'POST',2),
	(2,1,'002','http://www.test.com/user/info','用户详情','用于查看某个用户的详细资料','{\"param_name\":[\"user_id\"],\"param_type\":[\"string\"],\"param_cate\":[\"Y\"],\"param_default\":[\"\"],\"param_des\":[\"用户id\"]}','','','',1614954891,1,0,'GET',1);

/*!40000 ALTER TABLE `api` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table auth
# ------------------------------------------------------------

DROP TABLE IF EXISTS `auth`;

CREATE TABLE `auth` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户',
  `aid` int(11) NOT NULL DEFAULT '0' COMMENT '接口分类权限',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='权限表 - 若用户为普通管理员时，读此表获取权限';



# Dump of table cate
# ------------------------------------------------------------

DROP TABLE IF EXISTS `cate`;

CREATE TABLE `cate` (
  `aid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `cname` varchar(200) NOT NULL DEFAULT '' COMMENT '分类名称',
  `cdesc` varchar(200) NOT NULL DEFAULT '' COMMENT '分类描述',
  `isdel` int(11) NOT NULL DEFAULT '0' COMMENT '是否删除{0:正常,1删除}',
  `addtime` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `ord` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`aid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='接口分类表';

LOCK TABLES `cate` WRITE;
/*!40000 ALTER TABLE `cate` DISABLE KEYS */;

INSERT INTO `cate` (`aid`, `cname`, `cdesc`, `isdel`, `addtime`, `ord`)
VALUES
	(1,'演示分类一','这是一个演示分类',0,0,1),
	(2,'演示分类二','这个是分类描述',0,0,0);

/*!40000 ALTER TABLE `cate` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `nice_name` char(20) NOT NULL DEFAULT '' COMMENT '昵称',
  `login_name` char(30) NOT NULL DEFAULT '' COMMENT '登录名',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '最近登录时间',
  `login_pwd` varchar(32) NOT NULL DEFAULT '' COMMENT '登录密码',
  `isdel` tinyint(11) NOT NULL DEFAULT '0' COMMENT '状态 {0正常,1:删除}',
  `role` tinyint(11) NOT NULL DEFAULT '3' COMMENT '角色 {1:超级管理员,2:管理员,3:游客}',
  PRIMARY KEY (`id`),
  UNIQUE KEY `login_name` (`login_name`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='用户表';

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`id`, `nice_name`, `login_name`, `last_time`, `login_pwd`, `isdel`, `role`)
VALUES
	(1,'root','root',1614954208,'e10adc3949ba59abbe56e057f20f883e',0,1),
	(2,'admin','admin',1614954208,'e10adc3949ba59abbe56e057f20f883e',0,2),
	(3,'guest','guest',1614954208,'e10adc3949ba59abbe56e057f20f883e',0,3);

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
