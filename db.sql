-- -------------------------------------------------------------
-- TablePlus 5.9.6(546)
--
-- https://tableplus.com/
--
-- Database: api_manager
-- Generation Time: 2024-04-18 11:47:09.2320
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `api` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '接口编号',
  `aid` int(11) NOT NULL DEFAULT 0 COMMENT '接口分类id',
  `num` varchar(100) NOT NULL DEFAULT '' COMMENT '接口编号',
  `url` varchar(240) NOT NULL DEFAULT '' COMMENT '请求地址',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '接口名',
  `des` varchar(300) NOT NULL DEFAULT '' COMMENT '接口描述',
  `parameter` text DEFAULT NULL COMMENT '请求参数{所有的主求参数,以json格式在此存放}',
  `parameter_text` longtext DEFAULT NULL COMMENT '请求参数 存body请求体等',
  `memo` longtext DEFAULT NULL COMMENT '备注',
  `re` longtext DEFAULT NULL COMMENT '返回值',
  `lasttime` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '提后操作时间',
  `lastuid` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '最后修改uid',
  `isdel` tinyint(4) NOT NULL DEFAULT 0 COMMENT '{0:正常,1:删除}',
  `type` char(11) NOT NULL DEFAULT '' COMMENT '请求方式',
  `ord` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '排序(值越大,越靠前)',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=479 DEFAULT CHARSET=utf8 COMMENT='接口明细表';

CREATE TABLE `auth` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '用户',
  `aid` int(11) NOT NULL DEFAULT 0 COMMENT '接口分类权限',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='权限表 - 若用户为普通管理员时，读此表获取权限';

CREATE TABLE `cate` (
  `aid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `cname` varchar(200) NOT NULL DEFAULT '' COMMENT '分类名称',
  `cdesc` varchar(200) NOT NULL DEFAULT '' COMMENT '分类描述',
  `isdel` int(11) NOT NULL DEFAULT 0 COMMENT '是否删除{0:正常,1删除}',
  `addtime` int(11) NOT NULL DEFAULT 0 COMMENT '添加时间',
  `ord` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  PRIMARY KEY (`aid`)
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='接口分类表';

CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `nice_name` char(20) NOT NULL DEFAULT '' COMMENT '昵称',
  `login_name` char(30) NOT NULL DEFAULT '' COMMENT '登录名',
  `last_time` int(11) NOT NULL DEFAULT 0 COMMENT '最近登录时间',
  `login_pwd` varchar(32) NOT NULL DEFAULT '' COMMENT '登录密码',
  `isdel` tinyint(11) NOT NULL DEFAULT 0 COMMENT '状态 {0正常,1:删除}',
  `role` tinyint(11) NOT NULL DEFAULT 3 COMMENT '角色 {1:超级管理员,2:管理员,3:游客}',
  PRIMARY KEY (`id`),
  UNIQUE KEY `login_name` (`login_name`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='用户表';



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
