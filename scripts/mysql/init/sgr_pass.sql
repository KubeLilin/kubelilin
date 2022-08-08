CREATE DATABASE IF NOT EXISTS sgr_pass
    DEFAULT CHARACTER SET utf8mb4
    DEFAULT COLLATE utf8mb4_general_ci;

USE sgr_pass;
/*
 Navicat Premium Data Transfer

 Source Server         : PaaS
 Source Server Type    : MySQL
 Source Server Version : 80025
 Date: 01/06/2022 11:44:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for code_service_connection
-- ----------------------------
DROP TABLE IF EXISTS `code_service_connection`;
CREATE TABLE `code_service_connection` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(20) CHARACTER SET utf8 NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='服务连接类型';

-- ----------------------------
-- Records of code_service_connection
-- ----------------------------
BEGIN;
INSERT INTO `code_service_connection` (`id`, `code`, `name`) VALUES (1, 'vcs', 'GIT仓库');
INSERT INTO `code_service_connection` (`id`, `code`, `name`) VALUES (2, 'hub', '镜像仓库');
INSERT INTO `code_service_connection` (`id`, `code`, `name`) VALUES (3, 'pipeline', '流水线引擎');
INSERT INTO `code_service_connection` (`id`, `code`, `name`) VALUES (4, 'system', '系统回调');
COMMIT;

-- ----------------------------
-- Table structure for service_connection
-- ----------------------------
DROP TABLE IF EXISTS `service_connection`;
CREATE TABLE `service_connection` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户id',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '连接名称',
  `service_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '服务连接地址',
  `service_type` int NOT NULL COMMENT '连接类型: 1: vcs 2. hub 3. pipline 4.url',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用于保存其他服务或者第三方组件所依赖的资源，例如连接字符串，ssh秘钥，git连接等等';

-- ----------------------------
-- Records of service_connection
-- ----------------------------
BEGIN;
INSERT INTO `service_connection` (`id`, `tenant_id`, `name`, `service_url`, `service_type`, `update_time`, `creation_time`) VALUES (6, 1, 'Gitlab私有库', '', 1, '2022-04-20 19:55:45', '2022-04-20 19:55:45');
INSERT INTO `service_connection` (`id`, `tenant_id`, `name`, `service_url`, `service_type`, `update_time`, `creation_time`) VALUES (7, 1, 'Gogs私有库', '', 1, '2022-04-20 19:55:41', '2022-04-20 19:55:41');
INSERT INTO `service_connection` (`id`, `tenant_id`, `name`, `service_url`, `service_type`, `update_time`, `creation_time`) VALUES (9, 1, 'DockerHub私有库', '', 2, '2022-04-26 16:04:28', '2022-04-26 16:04:28');
INSERT INTO `service_connection` (`id`, `tenant_id`, `name`, `service_url`, `service_type`, `update_time`, `creation_time`) VALUES (10, 1, 'Harbor私有仓库', '', 2, '2022-04-26 16:53:14', '2022-04-26 16:53:14');
COMMIT;

-- ----------------------------
-- Table structure for service_connection_credentials
-- ----------------------------
DROP TABLE IF EXISTS `service_connection_credentials`;
CREATE TABLE `service_connection_credentials` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '凭据名称',
  `type` int NOT NULL COMMENT '凭证类型 1. 用户密码 2.token',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '凭证用户名',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '凭证密码',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '凭证TOKEN',
  `creation_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='常用的连接凭证，例如token';

-- ----------------------------
-- Records of service_connection_credentials
-- ----------------------------
BEGIN;
INSERT INTO `service_connection_credentials` (`id`, `name`, `type`, `username`, `password`, `token`, `creation_time`, `update_time`) VALUES (1, '自建GOGS凭证', 2, 'sgr', 'admin123', 'd2911632a4ac8db4ce13a3135700b58a5c8d7772', '2022-04-20 14:51:31', '2022-04-20 14:51:31');
INSERT INTO `service_connection_credentials` (`id`, `name`, `type`, `username`, `password`, `token`, `creation_time`, `update_time`) VALUES (2, '自建GitLab凭证', 2, 'yoyofx', '1234abcd', 'VZkAcPMATFad7B6ebgHT', '2022-04-20 15:20:39', '2022-04-20 15:20:39');
COMMIT;

-- ----------------------------
-- Table structure for service_connection_details
-- ----------------------------
DROP TABLE IF EXISTS `service_connection_details`;
CREATE TABLE `service_connection_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `main_id` bigint unsigned NOT NULL COMMENT '主数据id',
  `type` int NOT NULL COMMENT '连接类型',
  `detail` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='连接的详细信息，例如mysql连接字符串';

-- ----------------------------
-- Records of service_connection_details
-- ----------------------------
BEGIN;
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `creation_time`, `update_time`) VALUES (2, 2, 3, '{\"name\":\"234123\",\"repo\":\"1341\",\"userName\":\"1341\",\"password\":\"13213\",\"type\":3}', '2022-04-12 14:38:53', '2022-04-12 14:38:53');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `creation_time`, `update_time`) VALUES (3, 0, 3, '{\"name\":\"去玩儿群无\",\"repo\":\"气儿群\",\"userName\":\"企鹅王若\",\"password\":\"曲儿\",\"token\":\"曲儿\",\"type\":3}', '2022-04-11 15:53:13', '2022-04-11 15:53:13');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `creation_time`, `update_time`) VALUES (4, 7, 3, '{\"name\":\"Gogs私有库\",\"repo\":\"gogs.xiaocui.site\",\"userName\":\"sgr\",\"password\":\"admin123\",\"token\":\"d2911632a4ac8db4ce13a3135700b58a5c8d7772\",\"type\":3}', '2022-04-20 15:30:03', '2022-04-20 19:51:42');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `creation_time`, `update_time`) VALUES (5, 5, 2, '{\"name\":\"GITLAB连接\",\"repo\":\"http://1231213\",\"userName\":\"xiaocui\",\"password\":\"6666\",\"token\":\"token\",\"type\":2}', '2022-04-12 14:41:55', '2022-04-12 14:41:55');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `creation_time`, `update_time`) VALUES (6, 6, 2, '{\"name\":\"Gitlab私有库\",\"repo\":\"gitlab.yoyogo.run\",\"userName\":\"yoyofx\",\"password\":\"1234abcd\",\"token\":\"VZkAcPMATFad7B6ebgHT\",\"type\":2}', '2022-04-12 15:03:11', '2022-04-20 19:51:30');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `creation_time`, `update_time`) VALUES (7, 8, 1, '{\"name\":\"Github公开网络\",\"repo\":\"api.github.com\",\"type\":1}', '2022-04-20 19:51:07', '2022-04-20 19:51:07');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `creation_time`, `update_time`) VALUES (8, 9, 5, '{\"name\":\"DockerHub私有库\",\"repo\":\"https://index.docker.io/v1/\",\"userName\":\"yoyofx\",\"password\":\"\",\"token\":\"eW95b2Z4OnpsMTI1MzMwMw==\",\"type\":4}', '2022-04-26 16:51:33', '2022-04-26 16:51:33');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `creation_time`, `update_time`) VALUES (9, 10, 5, '{\"name\":\"Harbor私有仓库\",\"repo\":\"harbor.xiaocui.site\",\"userName\":\"admin\",\"token\":\"YWRtaW46SGFyYm9yMTIzNDU=\",\"type\":5}', '2022-04-26 16:53:14', '2022-04-26 16:53:14');
COMMIT;

-- ----------------------------
-- Table structure for sgr_code_application_language
-- ----------------------------
DROP TABLE IF EXISTS `sgr_code_application_language`;
CREATE TABLE `sgr_code_application_language` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(8) CHARACTER SET utf8 DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8 NOT NULL,
  `sort` smallint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典-应用开发语言';

-- ----------------------------
-- Records of sgr_code_application_language
-- ----------------------------
BEGIN;
INSERT INTO `sgr_code_application_language` (`id`, `code`, `name`, `sort`) VALUES (5, '0001', 'go', 0);
INSERT INTO `sgr_code_application_language` (`id`, `code`, `name`, `sort`) VALUES (6, '0002', 'java', 0);
INSERT INTO `sgr_code_application_language` (`id`, `code`, `name`, `sort`) VALUES (7, '0003', 'nodejs', 0);
INSERT INTO `sgr_code_application_language` (`id`, `code`, `name`, `sort`) VALUES (8, '0004', 'python', 0);
INSERT INTO `sgr_code_application_language` (`id`, `code`, `name`, `sort`) VALUES (9, '0005', '.net', 0);
COMMIT;

-- ----------------------------
-- Table structure for sgr_code_application_level
-- ----------------------------
DROP TABLE IF EXISTS `sgr_code_application_level`;
CREATE TABLE `sgr_code_application_level` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(8) CHARACTER SET utf8 DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8 NOT NULL,
  `sort` smallint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='字典-应用级别';

-- ----------------------------
-- Records of sgr_code_application_level
-- ----------------------------
BEGIN;
INSERT INTO `sgr_code_application_level` (`id`, `code`, `name`, `sort`) VALUES (5, '0001', 'P0', 0);
INSERT INTO `sgr_code_application_level` (`id`, `code`, `name`, `sort`) VALUES (6, '0002', 'P1', 0);
INSERT INTO `sgr_code_application_level` (`id`, `code`, `name`, `sort`) VALUES (7, '0003', 'P2', 0);
INSERT INTO `sgr_code_application_level` (`id`, `code`, `name`, `sort`) VALUES (8, '0004', 'p3', 0);
COMMIT;

-- ----------------------------
-- Table structure for sgr_code_deployment_level
-- ----------------------------
DROP TABLE IF EXISTS `sgr_code_deployment_level`;
CREATE TABLE `sgr_code_deployment_level` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 NOT NULL,
  `sort` smallint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='部署环境';

-- ----------------------------
-- Records of sgr_code_deployment_level
-- ----------------------------
BEGIN;
INSERT INTO `sgr_code_deployment_level` (`id`, `code`, `name`, `sort`) VALUES (1, 'dev', '开发环境', 0);
INSERT INTO `sgr_code_deployment_level` (`id`, `code`, `name`, `sort`) VALUES (2, 'test', '测试环境', 1);
INSERT INTO `sgr_code_deployment_level` (`id`, `code`, `name`, `sort`) VALUES (3, 'release', '预发布环境', 2);
INSERT INTO `sgr_code_deployment_level` (`id`, `code`, `name`, `sort`) VALUES (4, 'prod', '生产环境', 3);
COMMIT;

-- ----------------------------
-- Table structure for sgr_role_menu_map
-- ----------------------------
DROP TABLE IF EXISTS `sgr_role_menu_map`;
CREATE TABLE `sgr_role_menu_map` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `menu_id` bigint unsigned NOT NULL COMMENT '菜单ID',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=692 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色菜单权限影射';

-- ----------------------------
-- Records of sgr_role_menu_map
-- ----------------------------
BEGIN;
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (49, 10, 1, '2021-09-24 16:45:46', '2021-09-24 16:45:46');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (50, 11, 1, '2021-09-24 16:51:16', '2021-09-24 16:51:16');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (248, 12, 1, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (249, 12, 27, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (250, 12, 29, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (251, 12, 31, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (252, 12, 37, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (253, 12, 28, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (254, 12, 32, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (255, 12, 33, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (256, 12, 34, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (257, 12, 35, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (258, 12, 38, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (259, 12, 36, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (557, 2, 1, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (558, 2, 27, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (559, 2, 29, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (560, 2, 31, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (561, 2, 37, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (562, 2, 28, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (563, 2, 32, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (564, 2, 33, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (565, 2, 34, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (566, 2, 35, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (567, 2, 38, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (568, 2, 7, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (569, 2, 5, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (570, 2, 3, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (671, 1, 1, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (672, 1, 3, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (673, 1, 4, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (674, 1, 7, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (675, 1, 5, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (676, 1, 6, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (677, 1, 24, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (678, 1, 25, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (679, 1, 26, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (680, 1, 29, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (681, 1, 31, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (682, 1, 37, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (683, 1, 28, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (684, 1, 32, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (685, 1, 34, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (686, 1, 35, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (687, 1, 38, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (688, 1, 36, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (689, 1, 40, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (690, 1, 41, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES (691, 1, 27, '2022-03-29 14:25:22', '2022-03-29 14:25:22');
COMMIT;

-- ----------------------------
-- Table structure for sgr_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sgr_sys_menu`;
CREATE TABLE `sgr_sys_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `menu_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '编码',
  `menu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '目录名称',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '图标',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '路由路径',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'react组件路径',
  `is_root` tinyint NOT NULL DEFAULT '0' COMMENT '是否是根目录',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父层级id',
  `sort` int NOT NULL DEFAULT '0' COMMENT '权重，正序排序',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单';

-- ----------------------------
-- Records of sgr_sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (1, 0, '001', '仪表盘', 'dashboard', '/dashboard/analysis', './dashboard/analysis', 1, 0, 0, 1, NULL, '2022-03-17 00:45:32');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (3, 0, '002', '管理中心', 'SettingOutlined', '/account', './account', 1, 0, 99, 1, '2021-09-17 14:15:51', '2022-01-11 20:02:36');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (4, 0, '003', '菜单管理', 'smile', '/account/route', './account/route', 0, 3, 2, 1, '2021-09-17 14:16:30', '2022-01-11 20:02:34');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (5, 0, '004', '用户管理', 'smile', '/account/manage', './account/manage', 0, 3, 4, 1, '2021-09-17 14:16:43', '2022-01-11 20:02:33');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (6, 0, '005', '租户管理', 'smile', '/account/tenant', './account/tenant', 0, 3, 5, 1, '2021-09-17 14:16:50', '2022-01-11 20:02:31');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (7, 0, '006', '角色管理', 'user', '/account/role', './account/role', 0, 3, 3, 1, '2021-09-17 14:17:01', '2022-01-11 20:02:29');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (24, 0, '', '资源中心', 'ClusterOutlined', '/resources', './resources', 1, 0, 0, 1, '2021-12-24 06:32:51', '2021-12-24 07:21:09');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (25, 0, '', '集群管理', '', '/resources/clusters', './resources/clusters', 0, 24, 0, 1, '2021-12-24 06:43:14', '2021-12-24 06:43:14');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (26, 0, '', '团队空间', '', '/resources/namespaces', './resources/namespaces', 0, 24, 0, 1, '2021-12-24 06:43:44', '2022-03-24 16:19:04');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (27, 0, '', '应用中心', 'SendOutlined', '/applications', './applications', 1, 0, 0, 1, '2021-12-24 06:47:01', '2021-12-24 07:24:01');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (28, 0, '', '应用管理', '', '/applications/apps', './applications/apps', 0, 27, 10, 1, '2021-12-24 06:48:52', '2021-12-28 03:48:22');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (29, 0, '', '配置管理', '', '/applications/configs', './applications/configs', 0, 27, 0, 1, '2021-12-24 06:50:20', '2021-12-24 06:50:20');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (31, 0, '', '服务治理', '', '/applications/serviceconfig', './applications/serviceconfig', 0, 27, 0, 1, '2021-12-24 07:03:05', '2021-12-24 07:03:05');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (32, 0, '', 'DevOps', 'ProjectOutlined', '/devops', './devops', 1, 0, 0, 1, '2021-12-24 07:30:43', '2021-12-24 07:30:43');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (34, 0, '', '网络中心', 'RocketOutlined', '/components', './components', 1, 0, 0, 1, '2021-12-24 07:33:40', '2022-01-11 04:04:25');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (35, 0, '', '网关(API Gateway)', '', '/components/apigateway', './components/apigateway', 0, 34, 0, 1, '2021-12-24 07:44:50', '2022-03-30 15:29:51');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (36, 0, '', '监控中心', 'RadarChartOutlined', '/monitor', './monitor', 1, 0, 0, 1, '2021-12-24 07:51:59', '2021-12-24 07:51:59');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (37, 0, '', '部署环境', '', '/applications/info/deployments', './applications/info/deployments', 0, 27, 0, 1, '2021-12-27 13:07:42', '2021-12-27 13:07:42');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (38, 0, '', '服务(Service)', '', '/resources/services', './resources/services', 0, 34, 2, 1, '2022-01-11 04:05:42', '2022-03-30 15:29:17');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (40, 0, '', '服务连接器', '', '/resources/serviceConnection', './resources/serviceConnection', 0, 24, 0, 1, '2022-03-24 16:49:43', '2022-03-29 14:23:10');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (41, 0, '', '项目管理', '', '/devops/projects', './devops/projects', 0, 32, 0, 1, '2022-03-25 11:48:43', '2022-03-31 18:19:51');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant`;
CREATE TABLE `sgr_tenant` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `t_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '租户名称',
  `t_code` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '租户编码',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `creation_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_code` (`t_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='租户';

-- ----------------------------
-- Records of sgr_tenant
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant` (`id`, `t_name`, `t_code`, `status`, `creation_time`, `update_time`) VALUES (1, '平台研发团队', 'administration', 1, '2022-03-20 00:10:45', '2022-03-20 00:10:45');
INSERT INTO `sgr_tenant` (`id`, `t_name`, `t_code`, `status`, `creation_time`, `update_time`) VALUES (39, '公司研发团队', 'com-dev', 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_application
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_application`;
CREATE TABLE `sgr_tenant_application` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_Id` bigint unsigned NOT NULL COMMENT '租户ID',
  `name` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '集群应用名称(英文唯一)',
  `nickname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '应用中文名称',
  `remarks` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '集群应用备注',
  `git` varchar(500) CHARACTER SET utf8 NOT NULL COMMENT '集群应用绑定的git地址',
  `imagehub` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '集群应用绑定镜像仓库地址',
  `level` smallint unsigned NOT NULL COMMENT '应用级别',
  `language` smallint unsigned NOT NULL COMMENT '开发语言',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `labels` varchar(100) CHARACTER SET utf8 DEFAULT NULL COMMENT '应用标签',
  `git_type` varchar(20) CHARACTER SET utf8 NOT NULL COMMENT 'git类型 github/ gitee/ gogs/gitlab',
  `sc_id` bigint unsigned DEFAULT '0' COMMENT '服务连接git类型的凭据ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='集群应用';

-- ----------------------------
-- Records of sgr_tenant_application
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_application` (`id`, `tenant_Id`, `name`, `nickname`, `remarks`, `git`, `imagehub`, `level`, `language`, `status`, `create_time`, `update_time`, `labels`, `git_type`, `sc_id`) VALUES (1, 1, 'nginx', NULL, '', 'https://gogs.xiaocui.site/administration/nginx.git', 'https://harbor.xiaocui.site/apps/', 5, 7, 1, NULL, NULL, 'web', 'gogs', 7);
INSERT INTO `sgr_tenant_application` (`id`, `tenant_Id`, `name`, `nickname`, `remarks`, `git`, `imagehub`, `level`, `language`, `status`, `create_time`, `update_time`, `labels`, `git_type`, `sc_id`) VALUES (2, 1, 'yoyogodemo', NULL, '', 'http://gitlab.yoyogo.run/yoyofx/yoyogo.git', 'https://harbor.xiaocui.site/apps/', 5, 5, 1, NULL, NULL, 'web api', 'gitlab', 6);
INSERT INTO `sgr_tenant_application` (`id`, `tenant_Id`, `name`, `nickname`, `remarks`, `git`, `imagehub`, `level`, `language`, `status`, `create_time`, `update_time`, `labels`, `git_type`, `sc_id`) VALUES (15, 1, 'kubelilin-apiserver', '', '', 'http://gitlab.yoyogo.run/kubelilin/kubelilin-api.git', '', 5, 5, 1, NULL, NULL, 'kubelilin', 'gitlab', 6);
INSERT INTO `sgr_tenant_application` (`id`, `tenant_Id`, `name`, `nickname`, `remarks`, `git`, `imagehub`, `level`, `language`, `status`, `create_time`, `update_time`, `labels`, `git_type`, `sc_id`) VALUES (16, 1, 'kubelilin-dashbroad', '', '', 'http://gitlab.yoyogo.run/kubelilin/dashboard.git', '', 5, 7, 1, NULL, NULL, 'kubelilin', 'gitlab', 6);
INSERT INTO `sgr_tenant_application` (`id`, `tenant_Id`, `name`, `nickname`, `remarks`, `git`, `imagehub`, `level`, `language`, `status`, `create_time`, `update_time`, `labels`, `git_type`, `sc_id`) VALUES (17, 1, 'vue-demo', '', '', 'https://gitee.com/yoyofx/vue-docker-container.git', '', 7, 7, 1, NULL, NULL, 'vue', 'gitee', 0);
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_application_pipelines
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_application_pipelines`;
CREATE TABLE `sgr_tenant_application_pipelines` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'Pipeline ID',
  `appid` bigint unsigned NOT NULL COMMENT '应用ID',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '流水线名称, appid 下唯一',
  `dsl` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '流水线DSL',
  `taskStatus` int unsigned DEFAULT NULL COMMENT '流水线任务状态( ready=0 , running=1, success=2, fail=3,  )',
  `lastTaskId` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '最后一次任务执行ID',
  `status` tinyint unsigned NOT NULL,
  `creation_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='应用流水线';

-- ----------------------------
-- Records of sgr_tenant_application_pipelines
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_application_pipelines` (`id`, `appid`, `name`, `dsl`, `taskStatus`, `lastTaskId`, `status`, `creation_time`, `update_time`) VALUES (2, 1, 'ewfwewef', '', 1, '0', 0, '2022-02-21 11:24:26', '2022-03-08 17:11:56');
INSERT INTO `sgr_tenant_application_pipelines` (`id`, `appid`, `name`, `dsl`, `taskStatus`, `lastTaskId`, `status`, `creation_time`, `update_time`) VALUES (3, 1, 'wefwefszzzzv2', '[{\"name\":\"代码\",\"steps\":[{\"name\":\"拉取代码\",\"key\":\"git_pull\",\"save\":true,\"content\":{\"git\":\"https://gogs.xiaocui.site/administration/nginx.git\",\"branch\":\"dev\"}}]},{\"name\":\"编译构建\",\"steps\":[{\"name\":\"编译命令\",\"key\":\"code_build\",\"save\":true,\"content\":{\"buildEnv\":\"golang\",\"buildScript\":\"# 编译命令，注：当前已在代码根路径下\\ngo env -w GOPROXY=https://goproxy.cn,direct\\ngo build -ldflags=\\\"-s -w\\\" -o app .\\n\",\"buildFile\":\"./Dockerfile\"}},{\"name\":\"命令执行\",\"key\":\"cmd_shell\",\"save\":true,\"content\":{\"shell\":\"# bash\"}}]},{\"name\":\"部署\",\"steps\":[{\"name\":\"应用部署\",\"key\":\"app_deploy\",\"save\":true,\"content\":{\"depolyment\":1}}]},{\"name\":\"通知\",\"steps\":[{\"name\":\"命令执行\",\"key\":\"cmd_shell\",\"save\":true,\"content\":{\"shell\":\"# bash\"}}]}]', 1, '0', 0, '2022-02-21 11:46:36', '2022-03-23 20:18:40');
INSERT INTO `sgr_tenant_application_pipelines` (`id`, `appid`, `name`, `dsl`, `taskStatus`, `lastTaskId`, `status`, `creation_time`, `update_time`) VALUES (4, 2, 'TEST环境部署', '[{\"name\":\"代码\",\"steps\":[{\"name\":\"拉取代码\",\"key\":\"git_pull\",\"save\":true,\"content\":{\"git\":\"https://gitee.com/yoyofx/yoyogo.git\",\"branch\":\"master\"}}]},{\"name\":\"编译构建\",\"steps\":[{\"name\":\"编译命令\",\"key\":\"code_build\",\"save\":true,\"content\":{\"buildEnv\":\"golang\",\"buildScript\":\"# 编译命令，注：当前已在代码根路径下\\ngo env -w GOPROXY=https://goproxy.cn,direct\\ncd examples/simpleweb\\ngo mod download\\ngo mod tidy \\nCGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .\\n\",\"buildFile\":\"./examples/simpleweb/Dockerfile_NoCompile\"}}]},{\"name\":\"部署\",\"steps\":[{\"name\":\"应用部署\",\"key\":\"app_deploy\",\"save\":true,\"content\":{\"depolyment\":10}}]},{\"name\":\"通知\",\"steps\":[{\"name\":\"发布通知\",\"key\":\"publish_notify\",\"save\":true,\"content\":{\"notifyType\":\"dingtalk\",\"notifyKey\":\"a5e6519d74b1b3f9486d05e3ab765ec7bda31ea6fa4051b39e29a5bfde1a6b59\"}}]}]', 1, '57', 1, '2022-02-28 05:04:56', '2022-04-26 19:30:22');
INSERT INTO `sgr_tenant_application_pipelines` (`id`, `appid`, `name`, `dsl`, `taskStatus`, `lastTaskId`, `status`, `creation_time`, `update_time`) VALUES (5, 15, '腾讯云k8s部署', '[{\"name\":\"代码\",\"steps\":[{\"name\":\"拉取代码\",\"key\":\"git_pull\",\"save\":true,\"content\":{\"git\":\"http://gitlab.yoyogo.run/kubelilin/kubelilin-api.git\",\"branch\":\"dev\"}}]},{\"name\":\"编译构建\",\"steps\":[{\"name\":\"编译命令\",\"key\":\"code_build\",\"save\":true,\"content\":{\"buildEnv\":\"golang\",\"buildScript\":\"# 编译命令，注：当前已在代码根路径下\\ngo env -w GOPROXY=https://goproxy.cn,direct\\ncd src\\ngo mod download\\ngo mod tidy \\nCGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags=\\\"-s -w\\\" -o app .\\n\",\"buildFile\":\"./src/Dockerfile_Prod\"}}]},{\"name\":\"部署\",\"steps\":[{\"name\":\"应用部署\",\"key\":\"app_deploy\",\"save\":true,\"content\":{\"depolyment\":11}}]},{\"name\":\"通知\",\"steps\":[{\"name\":\"发布通知\",\"key\":\"publish_notify\",\"save\":true,\"content\":{\"notifyType\":\"dingtalk\",\"notifyKey\":\"a5e6519d74b1b3f9486d05e3ab765ec7bda31ea6fa4051b39e29a5bfde1a6b59\"}}]}]', 1, '67', 1, '2022-03-03 17:03:11', '2022-05-31 17:39:42');
INSERT INTO `sgr_tenant_application_pipelines` (`id`, `appid`, `name`, `dsl`, `taskStatus`, `lastTaskId`, `status`, `creation_time`, `update_time`) VALUES (6, 16, '腾讯云正式部署流水线', '[{\"name\":\"代码\",\"steps\":[{\"name\":\"拉取代码\",\"key\":\"git_pull\",\"save\":true,\"content\":{\"git\":\"http://gitlab.yoyogo.run/kubelilin/dashboard.git\",\"branch\":\"dev\"}}]},{\"name\":\"编译构建\",\"steps\":[{\"name\":\"编译命令\",\"key\":\"code_build\",\"save\":true,\"content\":{\"buildEnv\":\"nodejs\",\"buildScript\":\"# 编译命令，注：当前已在代码根路径下\\nnpm config set registry https://registry.npm.taobao.org --global\\ncd src\\nnpm install --force\\nnpm run build\\n\",\"buildFile\":\"./src/Dockerfile\"}}]},{\"name\":\"部署\",\"steps\":[{\"name\":\"应用部署\",\"key\":\"app_deploy\",\"save\":true,\"content\":{\"depolyment\":12}}]},{\"name\":\"通知\",\"steps\":[{\"name\":\"发布通知\",\"key\":\"publish_notify\",\"save\":true,\"content\":{\"notifyType\":\"dingtalk\",\"notifyKey\":\"a5e6519d74b1b3f9486d05e3ab765ec7bda31ea6fa4051b39e29a5bfde1a6b59\"}}]}]', 1, '81', 1, '2022-03-07 13:38:21', '2022-04-27 11:23:40');
INSERT INTO `sgr_tenant_application_pipelines` (`id`, `appid`, `name`, `dsl`, `taskStatus`, `lastTaskId`, `status`, `creation_time`, `update_time`) VALUES (7, 2, 'Yoyogo-demo-Dev腾讯云部署', '[{\"name\":\"代码\",\"steps\":[{\"name\":\"拉取代码\",\"key\":\"git_pull\",\"save\":true,\"content\":{\"git\":\"http://gitlab.yoyogo.run/yoyofx/yoyogo.git\",\"branch\":\"dev\"}}]},{\"name\":\"编译构建\",\"steps\":[{\"name\":\"编译命令\",\"key\":\"code_build\",\"save\":true,\"content\":{\"buildEnv\":\"golang\",\"buildScript\":\"# 编译命令，注：当前已在代码根路径下\\ngo env -w GOPROXY=https://goproxy.cn,direct\\ncd examples/simpleweb\\ngo mod download\\ngo mod tidy \\nCGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .\\n\",\"buildFile\":\"./examples/simpleweb/Dockerfile_NoCompile\"}}]},{\"name\":\"部署\",\"steps\":[{\"name\":\"应用部署\",\"key\":\"app_deploy\",\"save\":true,\"content\":{\"depolyment\":4}}]},{\"name\":\"通知\",\"steps\":[{\"name\":\"命令执行\",\"key\":\"publish_notify\",\"save\":true,\"content\":{\"notifyType\":\"wechat\",\"notifyKey\":\"428c53f6-261a-404d-8315-7ca368598a06\"}}]}]', 1, '4', 1, '2022-03-23 19:21:39', '2022-05-16 14:31:50');
INSERT INTO `sgr_tenant_application_pipelines` (`id`, `appid`, `name`, `dsl`, `taskStatus`, `lastTaskId`, `status`, `creation_time`, `update_time`) VALUES (8, 17, 'vuebuilder', '[{\"name\":\"代码\",\"steps\":[{\"name\":\"拉取代码\",\"key\":\"git_pull\",\"save\":true,\"content\":{\"git\":\"https://gitee.com/yoyofx/vue-docker-container.git\",\"branch\":\"master\"}}]},{\"name\":\"编译构建\",\"steps\":[{\"name\":\"编译命令\",\"key\":\"code_build\",\"save\":true,\"content\":{\"buildEnv\":\"nodejs\",\"buildScript\":\"# 编译命令，注：当前已在代码根路径下\\nnpm config set registry https://registry.npm.taobao.org --global\\nnpm install\\nnpm run build\\n\",\"buildFile\":\"./Dockerfile\"}}]},{\"name\":\"部署\",\"steps\":[{\"name\":\"应用部署\",\"key\":\"app_deploy\",\"save\":true,\"content\":{\"depolyment\":15}}]}]', 1, '10', 1, '2022-04-21 15:34:08', '2022-04-27 11:39:33');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_cluster
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_cluster`;
CREATE TABLE `sgr_tenant_cluster` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `nickname` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '别名',
  `name` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '集群名称',
  `version` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT 'k8s 版本号',
  `distribution` varchar(30) CHARACTER SET utf8 DEFAULT NULL COMMENT '来源',
  `config` text CHARACTER SET utf8 NOT NULL COMMENT 'k8s config text',
  `sort` int DEFAULT NULL COMMENT '排序',
  `status` tinyint NOT NULL COMMENT '状态',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='集群信息';

-- ----------------------------
-- Records of sgr_tenant_cluster
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_cluster` (`id`, `tenant_id`, `nickname`, `name`, `version`, `distribution`, `config`, `sort`, `status`, `create_time`, `update_time`) VALUES (3, 1, '腾讯云集群', 'cls-hbktlqm5', 'v1.18.4-tke.11', '', 'apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJd01UQXlOakExTkRFek5sb1hEVE13TVRBeU5EQTFOREV6Tmxvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTlVCCmZCNDJHVkxGaXVxSnZmcGdWSVYzT3h5VjNJU0l2SU5aaFdvMW9IRFF6bnFVOXU1b0hnNEZmbkxHU3BRSVdLN1gKOXNNRTBQc21DaCtwY2ZiMEszNW9OVnlFUWU3dndYVFVaU3Znb0lWb0xNWENYRlB5L2xkTVFTQURuZWw4Vnl0dgpYREVua0pvSTdIUEtmZ3E3czZ4cjh0VXFDQUZPMGErLzZHYWtFZXI0SlNVREQxZDFyY3dNVWd4VS9IaGQxVkhSClNQOEh3d1EvWHdIUzdpWUF2bmJOQ3pkNWxLWmNEeWV2ZVVsb2JrYmRxUllRMFplY3dGWjU4bDgyZFdyUWJkUFIKem9DcGRCdTJHQmtmZUtpUjhUb0Q2Q0IxYlFvb0JNSXQ4dG96dzNCL2JDMTRJeEZwUFgxdUNCWWJyK2NzUmNrTQoxOUM3NldqaUFCZ0hpMXNvTkgwQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0tVTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFMWll6VjRkZkF2Y1VPYXVFL0dGT2tvOVRQdWMKZ3h5NHVJczhhOURtTno4S3J0YjRsZG41bTljWnQ0RFRFUTJwbWwwOEY5cVBYRy9EY0FjWWE3ckFlVkZWUWNyMgpXMmI5dEVyMWZjTEhUWlF1aFpzQVVuT3FkcTYreDRkQTBTcTJLbm1hMWlnQmJoUWJCc3cyVGdIOEFtS0NxMXVXClhQVU1XWFE5NlB0eUYvWjEvTUdmcC9CMi9LZHdpVWd0WCtSQnlhWGsxaXJsZzZLb1owcVZDK1lxMTdOZG92QTYKa2habC9XZXlYaEZnaUhaclAwNHR1MzlhVHEyTkhsWk4vWjBvOXVYVUZ2RnVZYURwWGRNakxHNDFhU3p6KzJLVwpwZ0NkVGZXQkpGMzJMV2Npc21lcG5YVENxS2xEeWxDZW13bEg5OWljUGtUd2tOd2xXYlIzOFpKOFlsTT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\n    server: https://cls-hbktlqm5.ccs.tencent-cloud.com\n  name: cls-hbktlqm5\ncontexts:\n- context:\n    cluster: cls-hbktlqm5\n    user: \"86509022\"\n  name: cls-hbktlqm5-86509022-context-default\ncurrent-context: cls-hbktlqm5-86509022-context-default\nkind: Config\npreferences: {}\nusers:\n- name: \"86509022\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURDRENDQWZDZ0F3SUJBZ0lJUnJvbVFheWpCRFF3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TURFeU1qRXhOakEyTVRGYUZ3MDBNREV5TWpFeE5qQTJNVEZhTURJeApFakFRQmdOVkJBb1RDWFJyWlRwMWMyVnljekVjTUJvR0ExVUVBeE1UT0RZMU1Ea3dNakl0TVRZd09EVTJOamMzCk1UQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQUw2NXRiZDdsNjlDTTdsL1dNTHAKSTJSZ3ZXTmtic3BjcGllMlpVUXNIZ0JGcE1jc1JvZFVNUW5LbmZrZHh3NlIrSVAzTWVtVnVEZHYwR0dEalFFVAp6K0V6aU13dmorUkdqZkxrc0E1WWVxRjZnVUJjOFArSkMvVW11SDVNM0taK29IYXVZT2M1VTREeXBscVFpWTdOCmdNdk1SZ29LcmNOaVE1V3BxZ2k2UzNtTnF2Z0I3azFsR2hFeFhoRDZERExmWFRvaGhiMGQ0Q1lZUnlLSkZiWUUKMGFHTSt0TjVTOUdOZXlQZGZORTNFcVVzbHp4emJDajZseWxMQ3NUTGw1L0pLUGRvblBWd2JoV0RlZ1ZtNEp5LwpzbUpoSW81MHdLRDZDdVUxa1JYZXgzcGMrbUZDMUwvV3VMdUVOVWcxSVRPdWh0YjlsdXVqYVlnVUk3V2JvMkVOCi9Wa0NBd0VBQWFNL01EMHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHQVFVRkJ3TUMKQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUFqRFgybQp2c0N4T05VSmhLUDZWRHBFZy9MWTB6Ni9GMHBkblpoVFBFb3pwMDg1R01EbWZCdUxuSmxIbE54czZjc0t5Qk53CmtDMWtUM3Q1NndkWXUxVjBtN0tvZHFudjloWGxBYWZ6S0pmT1F3NlVzZmtqSjJZL0wwT2FMaWhLZ0IzYnpMWVcKZE0veVgwYlpVNzRlSHRraS9ydTRzOGFLamg4UzVMQnpzQ25Yc1dsOXdwdCt6ZXBkVnJVbDZPWmhOOVdTLzU2awpYcWZxMjAzeDMvWk1kdG53R0huOWZxQisyb2Yrc3BwY0dYdGJnanV0LzBXQVVhZWdTcUFqRGxLekcrMkJrTGJmCjB6dFp1UUFBVXBBSVBZeU1PZU1YOUdObmVSWWM5QUQ2ZnpVMWdTaEdIaEY3RDZobDl0NXpjRmZueUdGRmVSR2IKVVBEWVc0MlNJL0I5Vk1NNwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBdnJtMXQzdVhyMEl6dVg5WXd1a2paR0M5WTJSdXlseW1KN1psUkN3ZUFFV2t4eXhHCmgxUXhDY3FkK1IzSERwSDRnL2N4NlpXNE4yL1FZWU9OQVJQUDRUT0l6QytQNUVhTjh1U3dEbGg2b1hxQlFGencKLzRrTDlTYTRma3pjcG42Z2RxNWc1emxUZ1BLbVdwQ0pqczJBeTh4R0NncXR3MkpEbGFtcUNMcExlWTJxK0FIdQpUV1VhRVRGZUVQb01NdDlkT2lHRnZSM2dKaGhISW9rVnRnVFJvWXo2MDNsTDBZMTdJOTE4MFRjU3BTeVhQSE5zCktQcVhLVXNLeE11WG44a285MmljOVhCdUZZTjZCV2JnbkwreVltRWlqblRBb1BvSzVUV1JGZDdIZWx6NllVTFUKdjlhNHU0UTFTRFVoTTY2RzF2Mlc2Nk5waUJRanRadWpZUTM5V1FJREFRQUJBb0lCQUVXVWNMUG9sZlR0UFB1TwpkdTVjcVhuRVJUT09mMUM2TGkvTXZmTDUrVlAyRkdCSlNjMnpMRlM3STVpdmdXQlNab3lXVVJJN2VjSlh1M2puCnlqZzdaeHBzZDVxdU8xdDNWZS9uK0VhemhzR0VkTVRyWTB5R1RlTjQ1ZFBGN0xXYytxTnhpSTZ5ZmtGTHhON0QKWGp2SHd4WVdodkxBNUpXa01xM1dBTjlBUDZxdkh0N0FPVW4zb3lLZklWUTRHVlNXaWVPL3hBR0dMMUVMamZKVwpzYmR4NG92eHdHb1JKMGpod0x3THViMVp1ZnFGQkY0TldKa3NRRG94SW9RNzc0R2VCdWlPQnlBbEZJVW1ydERnCkhyS21Pci92cjVZelNuZjc3UFNrUUlwWnNETG5JLzRMRmxhc0xpQ0U0dVlMYnh5aDNoQVh4dllNL3E1RWUyWHIKakdjUXZvVUNnWUVBK2VTOSt2MitvalRUcFpGMmJLWmtIbzVac3FSSlFGcVB5ZVh1aHdSOWxkeTJvV1M2Z25ibQpnV0RQaVVJUGsvOG4vaGEvU1BVRWpLbjdEUEl3NWp2MEduYU1ucEhZVDdpTWlyajJWMGZEZWF6eUJoZVEvenRqClFBSXZLeXAvcmN6cGttTUpBTWJpamVVQ3dvNE9CQnh1bFpMdUlyc3BMeTBsbCtOUlo5NGs4KzhDZ1lFQXcyTFUKZ2tCR2dqT0ZPR2RCOHg5UHdlcm5VVG00b3huajZ5dGJzY2VacG1wQzZBNG8zcytqc1U3YXRXOVJDOHk4U0ZhOApQR2IrV0haL0cwajlXam9XOFl0d1ZlaUxTMzNsNE9GeFRSM2xxdkZZNDFUbjk0cGYvazdIMnVQTWkwazNhaDh4ClZVUG81OHJRTEpobVlaV2d1b3FjdkZONElkZnRPNVp0SFAveXV6Y0NnWUJLeTh5UmM2RzdxMVF1R252M3lWWHUKVDIxSnF5TEJ3Rm1KZE9rUVFLZldVMW5XdE8rZVhUaGhRVGpkUElpdEk1STAyMW9sM0RDZ2FjQmEvNkxqUnM2cApuUkk1NUMxNnJ5Smg0enJZcFFJOVNTYW43Q1hhUDB4VnZGR2grZlo5YnZmNmVPb1k4VzZlU2cycGJodUQzMzY2CkJtQ0F4TVJ1K25SbUlnUWJzc0ljd1FLQmdRQ2hBakpJMjVxKzlLSFZweEdydmFQR0UwTm5wZjlIT0xDZlBPMmwKQk96VFBFSENaTmk5TTdLYkRIWWlpWWJxQ0Y4bjVZSGM3S3F3VDRYVEFFVDNNMk53elExWFhmaGJ6M1c5NlVtcQoyUFpIOWZiZjd6bnd2WEQ1YWdZN2xQa2Ixc3Y1Z1pidndyU05QbzVxRVhSYytpYW5VazV5eXYzMk5hL2pLTVRsCjN1MHg1UUtCZ1FDVWkyeWVMOTYvcUREZ3p3amdyOWNWZFhGdU8xVFZtNXl2QmMzK015ZFh1dzlYaGFmRGowMXYKUmcvT2loNnptZmRkOEwxb05GNTY2ZmVja3ZKYzRCTWJVeDRRd2o3c1A3SStibGYyUzlLc3B3U00yWTN4SFBkSApsZU02eldnajF1bFBUbWt4SkpWYk4rL2RzZGZHTnpFQU9Qc2dKYkVKbDJFVnBVZWhlOWZ4bVE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\n', 0, 1, '2022-03-28 22:30:15', '2022-03-28 22:30:15');
INSERT INTO `sgr_tenant_cluster` (`id`, `tenant_id`, `nickname`, `name`, `version`, `distribution`, `config`, `sort`, `status`, `create_time`, `update_time`) VALUES (4, 1, '线上测试集群', 'microk8s-cluster', 'v1.22.4-3+adc4115d990346', '', 'apiVersion: v1\r\nclusters:\r\n- cluster:\r\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREekNDQWZlZ0F3SUJBZ0lVQzdPVGlodEx1ZWFVYUNOZ3VFQ1F1aWRmV1ljd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0Z6RVZNQk1HQTFVRUF3d01NVEF1TVRVeUxqRTRNeTR4TUI0WERUSXlNREl3TVRFek16RTBOVm9YRFRNeQpNREV6TURFek16RTBOVm93RnpFVk1CTUdBMVVFQXd3TU1UQXVNVFV5TGpFNE15NHhNSUlCSWpBTkJna3Foa2lHCjl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFuWHVTbmVwc3NIRUp0UjlWMUpjQTJaQmxnQk0vTVFXeDAwcU8KQzJuVWk2UCtUZDloOENtam1wS2p5TURwYzBOV2tVQVp3L0JOd1VYN0dHWG5EcjBsaW5oc1lTWkxydkdWOGw3TgpFR3dhd1VsY3FzZXova3dOZDI5Q2pBaTU4TUFaOWtKNUpianJLN2lXU0tPeHdkK0pvcWtSSVh3clJCUzU4VEFlCjVTNXFNK3dlUXhtcy9oSUpnZFNSL3ZSNC8yMERZd3JKT3ZkR0kzb1I1cUxjdm8weldYZ0NvNFZkNUQwYUQzMGMKMW50eVViODNjc2JPcGtvcU9HNnRsejBtNk5tZUNNLy9FVW5ncEtRbkIyUXNwaFhMYjFsT0VEZWF0K3ZyU1RMUApKQzUreG5xK1J0d3VXL0Fod0tJb1pBN2NiT2Y1MHljYlFqclArUTBqWHVZMi9uY0REd0lEQVFBQm8xTXdVVEFkCkJnTlZIUTRFRmdRVWQxbG9jOEpRaVo2R2pLMTZNQ0Z6OSs2bkllOHdId1lEVlIwakJCZ3dGb0FVZDFsb2M4SlEKaVo2R2pLMTZNQ0Z6OSs2bkllOHdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQwpBUUVBVmpBVDFkT1BEN3VMNE9DcXNCTVBTQ2JVeTFBS0wyT3AreFNJQ0NDRUV6V1RhNVk3bkRuYlh0bnVZS1hkCnhOdkdrSStZMFh4TE1oZm1PSklaNUEwTjRwdmRuYzdxSWZ0d01uWHNVU3dMTENkTE5sa1VpcWVRZ1F6UWtUeWcKaHB3aUU4REtlL2xvblMwczh2VUkvdVpEb0xiOFF0d2U2U3ZTcmpoajJsNUJLZUNaYjB6RHl1ZGNRZ1VDYWNJYwpMcWo5SXZ0amZaUG9selAxSTV6aDE0VG5pUFpERGFhYXplNGkxU1VET1JPeFp0TjhFSXlDWWxPN0t0b0txWFNsCmhqb0dzeEVnY1N2NmZBMlhVVGhKd1IvYUZPemxXZW4vOWV2VHlrZjRhSzN0YnFEMWVzL3M4cHVXY1g2ZExwNlYKK1NYc3E4RTl2WjhZRGRlQkRJZS9lSXlnN3c9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\r\n    server: https://47.100.213.49:8325\r\n  name: microk8s-cluster\r\ncontexts:\r\n- context:\r\n    cluster: microk8s-cluster\r\n    user: admin\r\n  name: microk8s\r\ncurrent-context: microk8s\r\nkind: Config\r\npreferences: {}\r\nusers:\r\n- name: admin\r\n  user:\r\n    token: VXJKWEp1ZmJJS09JVklxbDBoNFpoSitkY0g5b3lqcVUvTUxqcFlWbHAxQT0K', 0, 1, '2022-03-28 22:30:16', '2022-03-28 22:30:16');
INSERT INTO `sgr_tenant_cluster` (`id`, `tenant_id`, `nickname`, `name`, `version`, `distribution`, `config`, `sort`, `status`, `create_time`, `update_time`) VALUES (5, 1, '本地开发集群', 'kind-kind', 'v1.23.4', '', 'apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvakNDQWVhZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJeU1ETXlPREV6TlRnMU4xb1hEVE15TURNeU5URXpOVGcxTjFvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTFQwCjJiWUFJYk9pY1ljL2cyUTllTDNQdWdvYldVMS9nNnlCTW82TkRTNDdITWh3bXF4MDc1cWNUdzVNcVU2dFNHQVUKZ1BOYTZrdGwweENmUGdDK3RnZ1VMUnZ2a0ppQkI1VEM0Ti81UENocExwREdhYm11MzFoUHRaTzMyM2x5cDdSZwozdHhTRVJXRkpIRlM4R3AzZkVRL01aWDJRWHdxRHhvS2U2ekVjOUlqMGxGT0VlNllUSG15SnJ6Q1pvZVJMZkR0ClJUbkloUnVRWmVUazhOb3psZ0FXM0tPQ2luYnNLL1FMUjB1UGlKbmE3ZHdYRzhHeGJ6NllOelMrTVV2R0NJQloKM2w4UVExelV2YnFWQTd5dm5ZVGxZQ3kzMkRIM1UrZGlKbGVFc3FzNzVOdHdFZmJ3YUdtS0dCNzkvazAyeXdmbQpxQ3VrSnNEUXlLYWVDTnE4N1pFQ0F3RUFBYU5aTUZjd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZOOExoTFlGMVYrbHhIWnJwT1RPckxyMWpJU2VNQlVHQTFVZEVRUU8KTUF5Q0NtdDFZbVZ5Ym1WMFpYTXdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBSVdtUHYyRW5wUURqMzArTUhaZQpRZEsySmtpVHk1NmlGV3hVVUNHM015Umtnb1VRcmt5cjgyMUwxd01ER05Xd3ZYeld6WEtTR1RtdVh3aHVwYlJ6CnMrK1N2cEpHYnR3Vm5WUnQ0amN0aDNZYW13dnZRZVc4dCttVmRQWFo5SERIckRiRCtVNHI3U1RJVnRjZUljK0IKeURqMFp4Tm5ZUVUrTWdydG9CZXUvTlpoNkZWS3dMNTNQQTM4bzRnNW55dDJFd2VCSkZsdDk5QUZwVlBCUGE1egpPRXpMNVJpVFhweHhYbGozNXhMOWNhbTViVTlpNVNiMVNMbUZnbTdQY1RMb0JBQ0o3bTRrWEJjNEovRUhJZjdpCkRIcngyM2hzYU5mSjExbStPQXI5enFueHk4bFlHeUhFTFoxaEVScUFZV1FPUTNObllUTFphSkpLc3MxY2V4Z0gKc1R3PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    server: https://kubeapi.yoyogo.run:6443\n  name: kind-kind\ncontexts:\n- context:\n    cluster: kind-kind\n    user: kind-kind\n  name: kind-kind\ncurrent-context: kind-kind\nkind: Config\npreferences: {}\nusers:\n- name: kind-kind\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURJVENDQWdtZ0F3SUJBZ0lJY3hNUkc1UGdFRzh3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TWpBek1qZ3hNelU0TlRkYUZ3MHlNekF6TWpneE16VTROVGhhTURReApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sa3dGd1lEVlFRREV4QnJkV0psY201bGRHVnpMV0ZrCmJXbHVNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXFUcEFiWGphNGprOEZoc0IKUFB4NjVWdDM5OEh3TXFudG5TNXlpWXdyQTgyemxtQmJRVGZoL0JhU1phZjF6eVRTSUpBS2I5M2xyNU42UnVZTAo2b2FiVzVhWEFNeWt4dWdsMm50SGFGcmR1d2J0ZzZVZTVFL1dLc2h6d2VFZm5GME1UVmpiT1M2NVlqOXQ5OEJMCjdMWHgxMEw0Z3JrdUdrR05CSlJ0T1RGRVo0ZUxqOWkrYUkvamU1cjRra1hubHpaNjNPaDZ5SDhlWTZseU5zaVEKbEJLOS92M2N0dWNqRFV6UWFGandKU2dFUXBFbkJZWkpFWXZMMWw4dWJtckRTQ2RVZ0hOZzNrZEluK1hESXhnbAptQmJUWUN2TDBpRlNJbis0Qm5HYy8xRnlSbGxNTStxWmg2UTdhb1lTeUxpYWxUUTdRZ3FoN2h0LytJOXhlQk15Ck51OGszUUlEQVFBQm8xWXdWREFPQmdOVkhROEJBZjhFQkFNQ0JhQXdFd1lEVlIwbEJBd3dDZ1lJS3dZQkJRVUgKQXdJd0RBWURWUjBUQVFIL0JBSXdBREFmQmdOVkhTTUVHREFXZ0JUZkM0UzJCZFZmcGNSMmE2VGt6cXk2OVl5RQpuakFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBakdLNFNyRXZnNGwyekY3cmZaNkVZZUtBWjNTdFd3UThXZ1B5CnRoSDExZHl4RkZZV2lWSlRNK0hob1hFME11NHJsMGVmVGlsZS91RVozTWtSTyszT3orT1R1SWxvUFVqVDl3bG8KenJHcW8wRnR6RzFyT203VWp2cmJWcis5WjROdHdKenFDem93MXducHZMMGpEenFMK2VCOUVsek1JUDYvYy9kTAoyZXVVS0hNR1FEWFY5SzJUbnp5amVRaFZONmJmTHFQQnZOMFJpZEJINERXVmpIekdGUFZ6d1MxRTg5R3pZQlFTCkEvOVlpU29XMXV2R0ZzbHZMRm0zVS82MU5SU1BwS21aelR2MklCMTlrcmZyWHdXeHJ5VEk0UWs1ZHVTRGNETFMKNWxlLzdSNTZXR1YyNERONHd0eGxVbmdPMkMvWi9sSU1Mc2Vkc2praGRhdGNSeElXU1E9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBcVRwQWJYamE0ams4RmhzQlBQeDY1VnQzOThId01xbnRuUzV5aVl3ckE4MnpsbUJiClFUZmgvQmFTWmFmMXp5VFNJSkFLYjkzbHI1TjZSdVlMNm9hYlc1YVhBTXlreHVnbDJudEhhRnJkdXdidGc2VWUKNUUvV0tzaHp3ZUVmbkYwTVRWamJPUzY1WWo5dDk4Qkw3TFh4MTBMNGdya3VHa0dOQkpSdE9URkVaNGVMajlpKwphSS9qZTVyNGtrWG5selo2M09oNnlIOGVZNmx5TnNpUWxCSzkvdjNjdHVjakRVelFhRmp3SlNnRVFwRW5CWVpKCkVZdkwxbDh1Ym1yRFNDZFVnSE5nM2tkSW4rWERJeGdsbUJiVFlDdkwwaUZTSW4rNEJuR2MvMUZ5UmxsTU0rcVoKaDZRN2FvWVN5TGlhbFRRN1FncWg3aHQvK0k5eGVCTXlOdThrM1FJREFRQUJBb0lCQVFDY1J5azhiNEZqclF0cQpOcUljR2VXOGJZNEtoVlUyMUdxVTMxSHk2RGpKR0d1aWtYVjBjeEVFNWl0b0tZWDlON205VElmMjZ1cTBDL08rCkNiQXpjditCd3ozTUJUQ2VaOENkVlgrS1JXL0N5aEVKbzdFMWt1enJNWmFGRTB0SDhUM3RLQU83ZVFUR1VHRlQKaklnUkg2MG81N1hyc3V0Nlg5TEV5ZlZiL2IxNEJOcFBBTUhYY3JGc0tySDVHLzErU1ZsTWtBcXIxR1ZveTlsegpVRDR4Q0FEYUpYNjBtMlBWcy94VlpUWTBnelpNN1oyUVpSaVZlMmwvQkFMYXBtdjRyb2trN3FHeVlqNTc3L1VzClJEeDJTWUkwbnNjUXVFYnhQRzE4L3pTakVudVB6YjRWSDFxNGtrYmZJVFhYeU92bVZSM1dFZ084NTQ3MnNnTnYKRFBlSDcvakJBb0dCQU5KNGttZjFmQVlLejFvTmh6RmVidHNheG9iTmJtZDkyUno2ems1V3BKTU0ySWZMd3cxRgpRUDZaVU9lWGxsS0lnK2gzOE4wT1prc0hPaXZqd25iWHlwZWZ0d2lXUFpSRU5kNGJVNkRBWm81WWQwejVwam5WCll1SURaTVBrV1gxbGlZY0FGb040dnUybjZ2SnpkTTJhMURQYS9rOW9xL0lyZGhjbXIrTVFZYmJGQW9HQkFNM1YKdFJKdzI5L3pFeVZLbDhrNG43ck44ZmtCb0NLeDd2S2JGVXhUdUk0QUVhOUNkQlpBSkN6THQ1TVhQS0RRcm8rWQo1bWNld0NqdGVZeHc0amhCL2dnbGVOVmhNRjdMcndLWW5pSlBUeThFYjNsT2Z6VXhlalAzMkxpQ1FkZFlFa2YrCnB5ZzlERHpnaHhicjhCd2haVVBJZUpkQXI3ME1Ob1Y3d3ROWTF0YzVBb0dCQUk4M1VsQ0JQN0tRMm9XcG5PdmcKR1Vqb2tGeWNIMXMyekcwbTBRbUhkWmRHYjNVQUZ5ckpqQzl5dmhYSkZaWDFwT1NqOFBkSzNCcUR3WGFxRmVKdQpka0gwZ1VUTGp2MTc2TGZKR24vUlREZmlSTlBSTU9ZN3FIeW8ySUZmdWZGVkRWOGFsVGREQzVDRjZaNnZKN1RjCmZUY0g5bktzaEF4V3ZKV0dVcnNNbzZFTkFvR0FLMHVOVjhDbmJ4YTQ2em9pYlMvYzRUVWFkWTd5K1BsR1VuOXUKWG15cHdDeWhpNlRGdUR4cm40U0dLNGxTdGx6T3F1TzFhdkZiNS8vemRpZjBYWEw4RlFpVXZ3VS9ZR0VsZ2IxSwpYZHNpdzdnQ0hwajFSdUlPVzBvQnF1V3pKYkdmdnNjQXNBVVBRdXJUWDIzblJuTjcwbU1qZ21VWDBnTDBvT1VrCldNSXlCbmtDZ1lBSGdOY2o4cE9IY1M0a2ptVHRWdXlVajJzeCtPMVNZTEFOc0hVdjJsbEJUanVQTEQyQmZBUUMKeER3WnY3SDVkYmIrSEs5ZzZFNnlpdFlMQS9xa0MvRUFtN2RhQ0d2c1dZOXdSbElHNTZIdEhTR2FxS3pHdzdPUApqRHZOQXhZVUJzQzlsMlhwNkdwQnN6U3RlOVIxaHZmSnpMMDlBUU01Q0JPZlRkbVppRkFpZEE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=', NULL, 1, '2022-03-28 22:30:17', '2022-03-28 22:30:17');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_deployment_record
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_deployment_record`;
CREATE TABLE `sgr_tenant_deployment_record` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `app_id` bigint unsigned NOT NULL,
  `deployment_id` bigint unsigned NOT NULL,
  `apply_image` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `ops_type` char(20) COLLATE utf8mb4_general_ci NOT NULL,
  `operator` bigint unsigned DEFAULT NULL,
  `creation_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `state` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `remark` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=146 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='部署发布记录';

-- ----------------------------
-- Records of sgr_tenant_deployment_record
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (2, 1, 1, 'docker.io/library/nginx:alpine', 'githook', NULL, '2022-03-02 19:10:01', NULL, NULL, '2022-03-02 19:10:01');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (3, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', NULL, '2022-03-02 16:42:59', '成功', '', '2022-03-02 16:43:07');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (4, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', NULL, '2022-03-02 16:43:03', '成功', '', '2022-03-02 16:43:07');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (5, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', 0, '2022-03-02 08:44:47', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (6, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', 0, '2022-03-02 08:45:29', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (7, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', 0, '2022-03-02 08:57:37', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (8, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v36', 'githook', NULL, '2022-03-02 19:09:58', NULL, NULL, '2022-03-02 19:09:58');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (9, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v38', 'githook', NULL, '2022-03-02 19:09:59', NULL, NULL, '2022-03-02 19:09:59');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (10, 1, 1, 'docker.io/library/nginx:alpine', 'githook', 0, '2022-03-02 11:13:30', NULL, NULL, '2022-03-02 11:13:30');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (11, 1, 1, 'docker.io/library/nginx:alpine', 'githook', 0, '2022-03-02 11:13:42', NULL, NULL, '2022-03-02 11:13:42');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (12, 1, 1, 'docker.io/library/nginx:alpine', 'githook', 0, '2022-03-02 11:14:41', NULL, NULL, '2022-03-02 11:14:41');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (13, 1, 1, 'docker.io/library/nginx:alpine', 'githook', 0, '2022-03-02 11:16:22', NULL, NULL, '2022-03-02 11:16:22');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (14, 2, 10, 'yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-03 14:08:30', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (15, 2, 10, 'yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-03 15:24:13', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (16, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v25', 'manual', 0, '2022-03-04 17:39:00', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (17, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v26', 'manual', 0, '2022-03-04 18:12:32', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (18, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v27', 'manual', 0, '2022-03-04 18:44:43', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (19, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v34', 'manual', 0, '2022-03-04 19:36:09', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (20, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v36', '', 0, '2022-03-07 12:00:14', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (21, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v36', '', 0, '2022-03-07 21:16:13', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (22, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v38', '', 0, '2022-03-08 01:22:04', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (23, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v39', '', 0, '2022-03-08 01:53:05', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (24, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v37', '', 0, '2022-03-08 01:59:52', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (25, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v41', '', 0, '2022-03-08 02:21:15', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (26, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v42', '', 0, '2022-03-08 17:19:28', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (27, 1, 1, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-10 18:24:42', '失败', 'Deployment.apps \"dev-nginx-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"1\", \"clusterId\":\"3\", \"k8s-app\":\"dev-nginx-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (28, 1, 13, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-10 18:27:48', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (29, 1, 2, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-10 18:28:12', '失败', 'Deployment.apps \"test-nginx-microk8s-cluster\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"1\", \"clusterId\":\"4\", \"k8s-app\":\"test-nginx-microk8s-cluster\", \"kubelilin-default\":\"true\", \"namespace\":\"sukt-core\", \"namespaceId\":\"2\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (30, 2, 10, 'docker.io/yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-10 18:31:08', '失败', 'Deployment.apps \"test-yoyogodemo-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"2\", \"clusterId\":\"3\", \"k8s-app\":\"test-yoyogodemo-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (31, 2, 10, 'docker.io/yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-10 18:32:00', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (32, 1, 1, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-10 19:19:43', '失败', 'Deployment.apps \"dev-nginx-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"1\", \"clusterId\":\"3\", \"k8s-app\":\"dev-nginx-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (33, 1, 1, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-10 19:20:30', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (34, 1, 1, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-11 10:19:41', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (35, 1, 2, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-11 10:19:45', '失败', 'Deployment.apps \"test-nginx-microk8s-cluster\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"1\", \"clusterId\":\"4\", \"k8s-app\":\"test-nginx-microk8s-cluster\", \"kubelilin-default\":\"true\", \"namespace\":\"sukt-core\", \"namespaceId\":\"2\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (36, 1, 2, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-11 13:42:55', '失败', 'Deployment.apps \"test-nginx-microk8s-cluster\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"1\", \"clusterId\":\"4\", \"k8s-app\":\"test-nginx-microk8s-cluster\", \"kubelilin-default\":\"true\", \"namespace\":\"sukt-core\", \"namespaceId\":\"2\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (37, 1, 2, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-11 15:50:59', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (38, 1, 3, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-11 15:51:25', '失败', 'Deployment.apps \"prod-nginx-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"1\", \"clusterId\":\"3\", \"k8s-app\":\"prod-nginx-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (39, 1, 3, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-11 15:51:35', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (40, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', 0, '2022-03-11 15:51:45', '失败', 'Deployment.apps \"test-nginx-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"1\", \"clusterId\":\"3\", \"k8s-app\":\"test-nginx-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (41, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', 0, '2022-03-11 15:51:58', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (42, 1, 13, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-11 15:52:14', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (43, 2, 4, 'yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-11 16:06:10', '失败', 'Deployment.apps \"dev-yoyogodemo-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"2\", \"clusterId\":\"3\", \"k8s-app\":\"dev-yoyogodemo-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"qcloud-app\":\"dev-yoyogodemo-cls-hbktlqm5\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (44, 2, 4, 'yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-11 16:06:23', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (45, 2, 10, 'docker.io/yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-11 16:06:44', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (46, 2, 6, 'docker.io/yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-11 16:06:47', '失败', 'Deployment.apps \"prod-yoyogodemo-microk8s-cluster\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"2\", \"clusterId\":\"4\", \"k8s-app\":\"prod-yoyogodemo-microk8s-cluster\", \"kubelilin-default\":\"true\", \"namespace\":\"sukt-core\", \"namespaceId\":\"2\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (47, 2, 6, 'docker.io/yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-11 16:07:00', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (48, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v39', 'manual', 0, '2022-03-11 16:07:57', '失败', 'Deployment.apps \"prod-kubelilin-apiserver-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"15\", \"clusterId\":\"3\", \"k8s-app\":\"prod-kubelilin-apiserver-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (49, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v39', 'manual', 0, '2022-03-11 16:08:20', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (50, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v37', 'manual', 0, '2022-03-11 16:10:39', '失败', 'Deployment.apps \"prod-kubelilin-dashbroad-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"16\", \"clusterId\":\"3\", \"k8s-app\":\"prod-kubelilin-dashbroad-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (51, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v37', 'manual', 0, '2022-03-11 16:11:22', '失败', 'Deployment.apps \"prod-kubelilin-dashbroad-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"16\", \"clusterId\":\"3\", \"k8s-app\":\"prod-kubelilin-dashbroad-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (52, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v37', 'manual', 0, '2022-03-11 16:11:35', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (53, 1, 1, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-14 11:20:45', '失败', 'Deployment.apps \"dev-nginx-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"1\", \"clusterId\":\"3\", \"k8s-app\":\"dev-nginx-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"profileLevel\":\"dev\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (54, 1, 1, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-14 11:20:58', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (55, 1, 2, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-14 11:21:09', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (56, 1, 3, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-14 11:21:17', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (57, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', 0, '2022-03-14 11:21:29', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (58, 1, 13, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-14 11:21:38', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (59, 2, 4, 'yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-14 11:23:17', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (60, 2, 6, 'docker.io/yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-14 11:23:28', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (61, 2, 10, 'docker.io/yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-14 11:23:37', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (62, 2, 6, 'docker.io/yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-14 11:33:33', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (63, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v39', 'manual', 0, '2022-03-14 11:33:56', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (64, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v37', 'manual', 0, '2022-03-14 11:34:11', '失败', 'Deployment.apps \"prod-kubelilin-dashbroad-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"appId\":\"16\", \"clusterId\":\"3\", \"k8s-app\":\"prod-kubelilin-dashbroad-cls-hbktlqm5\", \"kubelilin-default\":\"true\", \"namespace\":\"yoyogo\", \"namespaceId\":\"1\", \"profileLevel\":\"prod\", \"tenantId\":\"1\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (65, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v37', 'manual', 0, '2022-03-14 11:34:25', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (66, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v40', '', 0, '2022-03-16 00:32:27', '失败', 'Deployment.apps \"prod-kubelilin-apiserver-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"k8s-app\":\"prod-kubelilin-apiserver-cls-hbktlqm5\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (67, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v38', '', 0, '2022-03-16 00:38:06', '失败', 'Deployment.apps \"prod-kubelilin-dashbroad-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"k8s-app\":\"prod-kubelilin-dashbroad-cls-hbktlqm5\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (68, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v39', '', 0, '2022-03-16 00:50:13', '失败', 'Deployment.apps \"prod-kubelilin-dashbroad-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"k8s-app\":\"prod-kubelilin-dashbroad-cls-hbktlqm5\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (69, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v40', 'manual', 0, '2022-03-16 01:24:12', '失败', 'Deployment.apps \"prod-kubelilin-apiserver-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"k8s-app\":\"prod-kubelilin-apiserver-cls-hbktlqm5\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (70, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v40', 'manual', 0, '2022-03-16 01:24:19', '失败', 'Deployment.apps \"prod-kubelilin-apiserver-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"k8s-app\":\"prod-kubelilin-apiserver-cls-hbktlqm5\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (71, 1, 1, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-16 01:26:57', '失败', 'Deployment.apps \"dev-nginx-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"k8s-app\":\"dev-nginx-cls-hbktlqm5\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (72, 1, 1, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-16 01:27:03', '失败', 'Deployment.apps \"dev-nginx-cls-hbktlqm5\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"k8s-app\":\"dev-nginx-cls-hbktlqm5\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (73, 1, 1, 'docker.io/library/nginx:alpine', 'manual', 0, '2022-03-16 01:27:42', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (74, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v40', 'manual', 0, '2022-03-16 01:28:18', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (75, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v40', '', 0, '2022-03-16 01:29:07', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (76, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v41', '', 0, '2022-03-16 21:41:45', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (77, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v42', '', 0, '2022-03-17 00:22:51', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (78, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v41', '', 0, '2022-03-17 17:58:42', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (79, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v44', '', 0, '2022-03-17 18:35:06', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (80, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v42', '', 0, '2022-03-17 21:05:33', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (81, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v45', '', 0, '2022-03-17 21:07:01', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (82, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v46', '', 0, '2022-03-18 00:15:57', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (83, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v47', '', 0, '2022-03-18 01:00:27', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (84, 2, 10, 'harbor.yoyogo.run/apps/pipeline-4-app-2:v1', '', 0, '2022-03-21 11:56:05', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (85, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v44', '', 0, '2022-03-22 18:06:20', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (86, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v48', '', 0, '2022-03-22 18:18:31', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (87, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v47', '', 0, '2022-03-22 18:46:04', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (88, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v48', '', 0, '2022-03-22 19:20:49', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (89, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v48', '', 0, '2022-03-23 10:47:40', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (90, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v52', '', 0, '2022-03-23 13:54:24', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (91, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v53', '', 0, '2022-03-23 14:11:37', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (92, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v49', '', 0, '2022-03-23 14:24:20', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (93, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v50', '', 0, '2022-03-23 16:00:51', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (94, 1, 1, 'docker.io/library/nginx:alpine', 'rollback', 1, '2022-03-23 17:23:45', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (95, 1, 1, 'docker.io/library/nginx:alpine', 'rollback', 1, '2022-03-23 17:26:28', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (96, 1, 1, 'docker.io/library/nginx:alpine', 'rollback', 1, '2022-03-23 17:29:18', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (97, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v51', '', 0, '2022-03-23 20:49:44', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (98, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v49', '', 0, '2022-03-23 20:55:30', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (99, 1, 1, 'docker.io/library/nginx:alpine', 'rollback', 1, '2022-03-24 20:19:31', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (100, 1, 1, 'docker.io/library/nginx:alpine', 'rollback', 1, '2022-03-24 20:21:45', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (101, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v52', 'rollback', 1, '2022-03-24 20:23:48', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (102, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v53', 'rollback', 1, '2022-03-24 20:26:43', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (103, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v52', '', 0, '2022-03-24 20:39:49', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (104, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v60', '', 0, '2022-03-25 01:36:13', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (105, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v53', '', 0, '2022-03-26 11:04:52', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (106, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v61', '', 0, '2022-03-26 11:19:27', '成功', '', '2022-03-26 11:19:27');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (107, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v62', '', 0, '2022-03-26 12:38:44', '成功', '', '2022-03-26 12:38:44');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (108, 2, 14, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v53', 'manual', 0, '2022-03-26 16:04:58', '成功', '', '2022-03-26 16:04:58');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (109, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v54', '', 0, '2022-03-27 11:54:46', '成功', '', '2022-03-27 11:54:45');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (110, 2, 14, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v53', 'manual', 0, '2022-03-27 23:51:00', '成功', '', '2022-03-27 23:50:59');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (111, 2, 14, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v53', 'manual', 0, '2022-03-28 00:03:53', '失败', 'Get \"https://kubeapi.yoyogo.run:6443/version?timeout=32s\": dial tcp 49.232.111.253:6443: connect: connection refused', '2022-03-28 00:03:53');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (112, 2, 14, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v53', 'manual', 0, '2022-03-28 00:04:07', '失败', 'Get \"https://kubeapi.yoyogo.run:6443/version?timeout=32s\": dial tcp 49.232.111.253:6443: connect: connection refused', '2022-03-28 00:04:07');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (113, 2, 14, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v53', 'manual', 0, '2022-03-28 00:05:09', '成功', '', '2022-03-28 00:05:09');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (114, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v55', '', 0, '2022-03-28 20:15:50', '成功', '', '2022-03-28 20:15:50');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (115, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v63', '', 0, '2022-03-28 20:16:34', '成功', '', '2022-03-28 20:16:34');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (116, 2, 14, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v53', 'manual', 0, '2022-03-29 10:19:41', '成功', '', '2022-03-29 10:19:41');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (117, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v64', '', 0, '2022-03-29 10:34:38', '成功', '', '2022-03-29 10:34:37');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (118, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v56', '', 0, '2022-03-29 20:21:35', '成功', '', '2022-03-29 20:21:34');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (119, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v65', '', 0, '2022-03-29 20:32:09', '成功', '', '2022-03-29 20:32:08');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (120, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v57', '', 0, '2022-03-29 23:04:04', '成功', '', '2022-03-29 23:04:04');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (121, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v67', '', 0, '2022-03-30 00:11:24', '成功', '', '2022-03-30 00:11:23');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (122, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v68', '', 0, '2022-03-30 10:38:06', '成功', '', '2022-03-30 10:38:05');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (123, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v69', '', 0, '2022-03-30 13:08:06', '成功', '', '2022-03-30 13:08:06');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (124, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v70', '', 0, '2022-03-30 13:43:04', '成功', '', '2022-03-30 13:43:03');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (125, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v58', '', 0, '2022-03-30 13:55:26', '成功', '', '2022-03-30 13:55:25');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (126, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v71', '', 0, '2022-03-30 13:55:49', '成功', '', '2022-03-30 13:55:49');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (127, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v59', '', 0, '2022-03-30 18:16:43', '成功', '', '2022-03-30 18:16:43');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (128, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v72', '', 0, '2022-03-30 18:17:08', '成功', '', '2022-03-30 18:17:08');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (129, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v73', '', 0, '2022-03-30 23:50:14', '成功', '', '2022-03-30 23:50:14');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (130, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v74', '', 0, '2022-03-31 00:08:49', '成功', '', '2022-03-31 00:08:48');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (131, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v60', '', 0, '2022-04-25 18:34:56', '成功', '', '2022-04-25 18:34:55');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (132, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v61', '', 0, '2022-04-25 18:41:26', '成功', '', '2022-04-25 18:41:26');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (133, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v80', '', 0, '2022-04-25 19:12:48', '成功', '', '2022-04-25 19:12:48');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (134, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v57', '', 0, '2022-04-26 19:36:05', '成功', '', '2022-04-26 19:36:05');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (135, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v62', '', 0, '2022-04-27 10:53:17', '成功', '', '2022-04-27 10:53:17');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (136, 16, 12, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v81', '', 0, '2022-04-27 11:30:59', '成功', '', '2022-04-27 11:30:58');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (137, 17, 15, 'harbor.xiaocui.site/apps/pipeline-8-app-17:v10', '', 0, '2022-04-27 11:49:58', '成功', '', '2022-04-27 11:49:57');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (138, 2, 4, 'harbor.xiaocui.site/apps/pipeline-7-app-2:v4', '', 0, '2022-05-16 14:37:54', '成功', '', '2022-05-16 14:37:53');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (139, 2, 4, 'harbor.xiaocui.site/apps/pipeline-7-app-2:v5', '', 0, '2022-05-16 14:45:28', '成功', '', '2022-05-16 14:45:28');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (140, 2, 4, 'harbor.xiaocui.site/apps/pipeline-7-app-2:v7', '', 0, '2022-05-16 15:09:17', '成功', '', '2022-05-16 15:09:16');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (141, 2, 4, 'harbor.xiaocui.site/apps/pipeline-7-app-2:v8', '', 0, '2022-05-16 15:16:27', '成功', '', '2022-05-16 15:16:27');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (142, 2, 4, 'harbor.xiaocui.site/apps/pipeline-7-app-2:v12', '', 0, '2022-05-16 15:32:52', '成功', '', '2022-05-16 15:32:52');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (143, 2, 4, 'harbor.xiaocui.site/apps/pipeline-7-app-2:v13', '', 0, '2022-05-16 15:53:35', '成功', '', '2022-05-16 15:53:34');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (144, 2, 4, 'harbor.xiaocui.site/apps/pipeline-7-app-2:v18', '', 0, '2022-05-16 17:01:31', '成功', '', '2022-05-16 17:01:31');
INSERT INTO `sgr_tenant_deployment_record` (`id`, `app_id`, `deployment_id`, `apply_image`, `ops_type`, `operator`, `creation_time`, `state`, `remark`, `update_time`) VALUES (145, 15, 11, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v67', '', 0, '2022-05-31 17:44:55', '成功', '', '2022-05-31 17:44:55');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_deployments
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_deployments`;
CREATE TABLE `sgr_tenant_deployments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '部署名称(英文唯一)',
  `nickname` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '部署中文名称',
  `cluster_id` bigint unsigned NOT NULL COMMENT '集群ID',
  `namespace_id` bigint unsigned NOT NULL COMMENT '命名空间ID',
  `app_id` bigint unsigned NOT NULL COMMENT '应用ID',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `image_hub` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `app_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `workload_type` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `replicas` int unsigned NOT NULL DEFAULT '1',
  `service_enable` tinyint unsigned DEFAULT NULL,
  `service_name` varchar(150) CHARACTER SET utf8 DEFAULT NULL,
  `service_away` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `service_port` int unsigned NOT NULL DEFAULT '0',
  `service_port_type` varchar(8) CHARACTER SET utf8 DEFAULT NULL,
  `last_image` varchar(350) CHARACTER SET utf8 DEFAULT NULL,
  `level` varchar(8) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `levev_idx` (`level`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='集群部署';

-- ----------------------------
-- Records of sgr_tenant_deployments
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (1, 1, 'dev-nginx-cls-hbktlqm5', 'nginx', 3, 1, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'dev-nginx-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 80, '', '', 'dev');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (2, 1, 'test-nginx-microk8s-cluster', 'nginx', 4, 2, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'svc-test-nginx-microk8s-cluster', 'ClusterPort', 80, '', '', 'test');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (3, 1, 'prod-nginx-cls-hbktlqm5', 'prod-nginx', 3, 1, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'prod-nginx-cls-hbktlqm-svc-cluster-sgr', 'ClusterPort', 80, '', '', 'prod');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (4, 1, 'dev-yoyogodemo-cls-hbktlqm5', 'yoyogo-demo', 3, 1, 2, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'dev-yoyogodemo-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'dev');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (5, 1, 'test-nginx-cls-hbktlqm5', 'prism-desgin', 3, 1, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'test-nginx-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8092, '', '', 'test');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (6, 1, 'prod-yoyogodemo-microk8s-cluster', 'yoyogo-demo正式环境', 4, 2, 2, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'prod-yoyogodemo-microk8s-cluster-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'prod');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (10, 1, 'test-yoyogodemo-cls-hbktlqm5', 'YoyoGo框架demo测试环境', 3, 1, 2, 1, NULL, NULL, '', '', 'Deployment', 1, 0, 'test-yoyogodemo-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'test');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (11, 1, 'prod-kubelilin-apiserver-cls-hbktlqm5', 'Kubelilin-腾讯云正式部署', 3, 1, 15, 1, NULL, NULL, '', '', 'Deployment', 1, 0, 'prod-kubelilin-apiserver-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'prod');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (12, 1, 'prod-kubelilin-dashbroad-cls-hbktlqm5', 'Dashborad-腾讯云正式部署', 3, 1, 16, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'prod-kubelilin-dashbroad-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8092, '', '', 'prod');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (13, 1, 'dev-nginx-microk8s-cluster', '私有云开发环境Nginx', 4, 13, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'dev-nginx-microk8s-cluster-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'dev');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (14, 1, 'dev-yoyogodemo-kind-kind', 'localhost', 5, 15, 2, 1, NULL, NULL, '', '', 'Deployment', 3, 1, 'dev-yoyogodemo-kind-kind-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'dev');
INSERT INTO `sgr_tenant_deployments` (`id`, `tenant_id`, `name`, `nickname`, `cluster_id`, `namespace_id`, `app_id`, `status`, `create_time`, `update_time`, `image_hub`, `app_name`, `workload_type`, `replicas`, `service_enable`, `service_name`, `service_away`, `service_port`, `service_port_type`, `last_image`, `level`) VALUES (15, 1, 'dev-vue-demo-kind-kind', 'vue环境部署', 5, 15, 17, 1, NULL, NULL, '', '', 'Deployment', 1, 0, 'dev-vue-demo-kind-kind-svc-cluster-sgr', 'ClusterPort', 8092, '', '', 'dev');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_deployments_containers
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_deployments_containers`;
CREATE TABLE `sgr_tenant_deployments_containers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `deploy_id` bigint unsigned NOT NULL,
  `is_main` tinyint unsigned NOT NULL,
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `image_version` varchar(40) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `image_pull_strategy` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `request_cpu` decimal(4,2) NOT NULL,
  `request_memory` decimal(5,0) NOT NULL,
  `limit_cpu` decimal(4,2) NOT NULL,
  `limit_memory` decimal(5,0) NOT NULL,
  `environments` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `workdir` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `run_cmd` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `run_params` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `podstop` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `liveness` varchar(300) CHARACTER SET utf8 DEFAULT NULL,
  `readness` varchar(300) CHARACTER SET utf8 DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='应用部署容器配置';

-- ----------------------------
-- Records of sgr_tenant_deployments_containers
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (1, '', 1, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, '[{\"key\":\"test\",\"value\":\"1\"}]', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (2, '', 2, 1, 'docker.io/library/nginx:alpine', '', '', 0.05, 64, 0.07, 128, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (3, '', 3, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (4, '', 4, 1, 'harbor.xiaocui.site/apps/pipeline-7-app-2:v18', '', '', 0.10, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (5, '', 5, 1, 'yoyofx/prism-desgin:v0.1', '', '', 0.25, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (6, '', 6, 1, 'docker.io/yoyofx/yoyogo-demo:v0.1', '', '', 0.25, 128, 0.25, 256, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (7, '', 7, 1, '', '', '', 0.25, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (8, '', 8, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, '[{\"key\":\"es\",\"value\":\"123\"},{\"key\":\"se\",\"value\":\"133\"}]', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (9, '', 9, 1, '', '', '', 0.25, 128, 0.25, 256, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (10, '', 10, 1, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v57', '', '', 0.25, 128, 0.25, 256, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (11, '', 11, 1, 'harbor.xiaocui.site/apps/pipeline-5-app-15:v67', '', '', 0.25, 128, 0.25, 256, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (12, '', 12, 1, 'harbor.xiaocui.site/apps/pipeline-6-app-16:v81', '', '', 0.05, 64, 0.15, 128, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (13, '', 13, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (14, '', 14, 1, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v53', '', '', 0.15, 128, 0.25, 256, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` (`id`, `name`, `deploy_id`, `is_main`, `image`, `image_version`, `image_pull_strategy`, `request_cpu`, `request_memory`, `limit_cpu`, `limit_memory`, `environments`, `workdir`, `run_cmd`, `run_params`, `podstop`, `liveness`, `readness`) VALUES (15, '', 15, 1, 'harbor.xiaocui.site/apps/pipeline-8-app-17:v10', '', '', 0.05, 50, 0.10, 100, 'null', '', '', '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_namespace
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_namespace`;
CREATE TABLE `sgr_tenant_namespace` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `cluster_id` bigint unsigned NOT NULL COMMENT '集群ID',
  `namespace` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '命名空间名称',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `status` tinyint NOT NULL COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='集群_命名空间';

-- ----------------------------
-- Records of sgr_tenant_namespace
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_namespace` (`id`, `tenant_id`, `cluster_id`, `namespace`, `create_time`, `update_time`, `status`) VALUES (1, 1, 3, 'yoyogo', '2021-12-24 16:24:21', '2021-12-24 16:24:23', 1);
INSERT INTO `sgr_tenant_namespace` (`id`, `tenant_id`, `cluster_id`, `namespace`, `create_time`, `update_time`, `status`) VALUES (2, 1, 4, 'sukt-core', '2021-12-24 16:54:47', '2021-12-24 16:54:49', 1);
INSERT INTO `sgr_tenant_namespace` (`id`, `tenant_id`, `cluster_id`, `namespace`, `create_time`, `update_time`, `status`) VALUES (11, 1, 3, 'klns-administration', '2022-03-10 13:37:43', '2022-03-10 13:37:43', 1);
INSERT INTO `sgr_tenant_namespace` (`id`, `tenant_id`, `cluster_id`, `namespace`, `create_time`, `update_time`, `status`) VALUES (12, 39, 3, 'klns-com-dev', '2022-03-10 13:38:09', '2022-03-10 13:38:09', 1);
INSERT INTO `sgr_tenant_namespace` (`id`, `tenant_id`, `cluster_id`, `namespace`, `create_time`, `update_time`, `status`) VALUES (13, 39, 4, 'klns-com-dev', '2022-03-10 13:58:14', '2022-03-10 13:58:14', 1);
INSERT INTO `sgr_tenant_namespace` (`id`, `tenant_id`, `cluster_id`, `namespace`, `create_time`, `update_time`, `status`) VALUES (14, 1, 4, 'klns-administration', '2022-03-24 20:07:33', '2022-03-24 20:07:33', 1);
INSERT INTO `sgr_tenant_namespace` (`id`, `tenant_id`, `cluster_id`, `namespace`, `create_time`, `update_time`, `status`) VALUES (15, 1, 5, 'klns-administration', '2022-03-26 16:03:37', '2022-03-26 16:03:37', 1);
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_role
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_role`;
CREATE TABLE `sgr_tenant_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色编码',
  `role_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名称',
  `description` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '角色描述',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_role_code_name` (`role_code`,`role_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='租户角色';

-- ----------------------------
-- Records of sgr_tenant_role
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_role` (`id`, `role_code`, `role_name`, `description`, `status`, `tenant_id`, `creation_time`, `update_time`) VALUES (1, 'PlatformAdmin', '平台管理员', NULL, 1, 1, '2021-09-24 08:48:09', '2021-09-24 08:48:09');
INSERT INTO `sgr_tenant_role` (`id`, `role_code`, `role_name`, `description`, `status`, `tenant_id`, `creation_time`, `update_time`) VALUES (2, 'TenantAdmin', '租户管理员', NULL, 1, 1, '2021-09-24 08:47:16', '2021-09-24 08:47:16');
INSERT INTO `sgr_tenant_role` (`id`, `role_code`, `role_name`, `description`, `status`, `tenant_id`, `creation_time`, `update_time`) VALUES (12, 'tuser', '用户', '', 1, 39, '2022-01-11 11:46:09', '2022-01-11 11:46:09');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_user
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_user`;
CREATE TABLE `sgr_tenant_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名',
  `account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '手机',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息';

-- ----------------------------
-- Records of sgr_tenant_user
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES (1, 1, 'admin', '平台管理员', 'ef73781effc5774100f87fe2f437a435', '13877668829', 'zl.hxd@hotmail.com', 1, NULL, '2022-03-30 18:02:31');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES (11, 1, 'user1', '111', 'e10adc3949ba59abbe56e057f20f883e', '18630535890', '111', 1, NULL, '2022-03-30 18:03:06');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES (15, 1, 'ghost', 'ghost', 'e10adc3949ba59abbe56e057f20f883e', '18630535890', '86509022@qq.com', 1, NULL, '2022-03-30 18:03:17');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES (17, 39, 'com-dev-admin', '公司研发团队admin', 'e10adc3949ba59abbe56e057f20f883e', '', '', 1, '2022-01-11 19:26:55', '2022-03-30 18:03:37');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_user_role`;
CREATE TABLE `sgr_tenant_user_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL COMMENT '用户id',
  `role_id` bigint NOT NULL COMMENT '角色id',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户角色';

-- ----------------------------
-- Records of sgr_tenant_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (17, 0, 2, '2022-01-10 21:21:20', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (22, 16, 1, '2022-01-11 19:03:33', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (23, 17, 2, '2022-01-11 19:26:55', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (24, 1, 1, '2022-03-30 18:02:31', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (25, 11, 2, '2022-03-30 18:03:07', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (26, 15, 2, '2022-03-30 18:03:18', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
