/*
SQLyog 企业版 - MySQL GUI v8.14 
MySQL - 5.6.17 : Database - dawdler
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`dawdler` /*!40100 DEFAULT CHARACTER SET latin1 */;

/*Table structure for table `pub_template` */

DROP TABLE IF EXISTS `pub_template`;

CREATE TABLE `pub_template` (
  `ID` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `code` varchar(64) NOT NULL COMMENT '编码',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `datasource` varchar(64) NOT NULL COMMENT '数据库',
  `tablename` varchar(64) NOT NULL COMMENT '表名',
  PRIMARY KEY (`ID`),
  UNIQUE KEY `pub_template_code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

/*Data for the table `pub_template` */

LOCK TABLES `pub_template` WRITE;

UNLOCK TABLES;

/*Table structure for table `pub_template_item` */

DROP TABLE IF EXISTS `pub_template_item`;

CREATE TABLE `pub_template_item` (
  `id` int(11) NOT NULL COMMENT '编号',
  `pub_template_id` int(11) NOT NULL COMMENT '模版ID',
  `en` varchar(64) NOT NULL COMMENT '英文名称',
  `cn` varchar(64) NOT NULL COMMENT '中文名称',
  `type` varchar(256) NOT NULL DEFAULT 'string' COMMENT '类型（添加一个字典类型)',
  `sort` int(11) NOT NULL DEFAULT '1' COMMENT '位置',
  `isInsert` tinyint(2) NOT NULL DEFAULT '1' COMMENT '是否插入',
  `isEdit` tinyint(2) NOT NULL DEFAULT '1' COMMENT '是否编辑',
  `isQuery` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否查询',
  `isShowTable` tinyint(2) NOT NULL DEFAULT '1' COMMENT '是否显示列表',
  `isShowCard` tinyint(2) NOT NULL DEFAULT '1' COMMENT '是否显示面版',
  `isPrint` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否打印',
  `isSort` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否支持排序',
  `comment` varchar(256) NOT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `pub_template_item` */

LOCK TABLES `pub_template_item` WRITE;

UNLOCK TABLES;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
