CREATE DATABASE IF NOT EXISTS sgr_pass
    DEFAULT CHARACTER SET utf8mb4
    DEFAULT COLLATE utf8mb4_general_ci;

USE sgr_pass;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for devops_projects
-- ----------------------------
DROP TABLE IF EXISTS `devops_projects`;
CREATE TABLE `devops_projects` (
                                   `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                   `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '项目名称',
                                   `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
                                   `creation_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `soft_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '删除标记',
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='devops 项目管理';



-- ----------------------------
-- Table structure for devops_projects_apps
-- ----------------------------
DROP TABLE IF EXISTS `devops_projects_apps`;
CREATE TABLE `devops_projects_apps` (
                                        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                        `project_id` bigint unsigned NOT NULL COMMENT '项目ID',
                                        `application_id` bigint unsigned NOT NULL COMMENT '应用 ID',
                                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='devops 项目应用对应表';


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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用于保存其他服务或者第三方组件所依赖的资源，例如连接字符串，ssh秘钥，git连接等等';

-- ----------------------------
-- Records of service_connection
-- ----------------------------
BEGIN;
INSERT INTO `service_connection` (`id`, `tenant_id`, `name`, `service_url`, `service_type`, `update_time`, `creation_time`) VALUES (6, 1, 'Gitlab私有库', '', 1, '2022-04-20 19:55:45', '2022-04-20 19:55:45');
INSERT INTO `service_connection` (`id`, `tenant_id`, `name`, `service_url`, `service_type`, `update_time`, `creation_time`) VALUES (7, 1, 'Gogs私有库', '', 1, '2022-04-20 19:55:41', '2022-04-20 19:55:41');
INSERT INTO `service_connection` (`id`, `tenant_id`, `name`, `service_url`, `service_type`, `update_time`, `creation_time`) VALUES (10, 1, 'Harbor私有仓库', '', 2, '2022-08-10 14:07:46', '2022-04-26 16:53:14');
INSERT INTO `service_connection` (`id`, `tenant_id`, `name`, `service_url`, `service_type`, `update_time`, `creation_time`) VALUES (11, 1, '系统部署Webhook', '', 4, '2022-08-10 14:02:48', '2022-08-10 14:02:23');
INSERT INTO `service_connection` (`id`, `tenant_id`, `name`, `service_url`, `service_type`, `update_time`, `creation_time`) VALUES (12, 1, 'Jenkins流水线引擎', '', 3, '2022-08-10 14:03:52', '2022-08-10 14:03:52');
COMMIT;

-- ----------------------------
-- Table structure for service_connection_credentials
-- ----------------------------
DROP TABLE IF EXISTS `service_connection_credentials`;
CREATE TABLE `service_connection_credentials` (
                                                  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                                  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '凭据名称',
                                                  `type` int NOT NULL COMMENT '凭证类型 1. 用户密码 2.token',
                                                  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '凭证用户名',
                                                  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '凭证密码',
                                                  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '凭证TOKEN',
                                                  `creation_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='常用的连接凭证，例如token';

-- ----------------------------
-- Records of service_connection_credentials
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for service_connection_details
-- ----------------------------
DROP TABLE IF EXISTS `service_connection_details`;
CREATE TABLE `service_connection_details` (
                                              `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                              `main_id` bigint unsigned NOT NULL COMMENT '主数据id',
                                              `type` int NOT NULL COMMENT '连接类型 vcs type:\n1: ''github'',\n2: ''gitlab'',\n3: ''gogs'',\n4: ''gitee''\n',
                                              `detail` varchar(500) COLLATE utf8mb4_general_ci DEFAULT NULL,
                                              `enable` tinyint unsigned DEFAULT '1',
                                              `creation_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                              `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='连接的详细信息，例如mysql连接字符串';

-- ----------------------------
-- Records of service_connection_details
-- ----------------------------
BEGIN;
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `enable`, `creation_time`, `update_time`) VALUES (4, 7, 3, '{\"name\":\"Gogs私有库\",\"repo\":\"gogs.xxx.com\",\"userName\":\"sgr\",\"password\":\"admin123\",\"token\":\"\",\"type\":3}', 1, '2022-04-20 15:30:03', '2022-04-20 19:51:42');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `enable`, `creation_time`, `update_time`) VALUES (6, 6, 2, '{\"name\":\"Gitlab私有库\",\"repo\":\"gitlab.xxx.com\",\"userName\":\"yoyofx\",\"password\":\"1234abcd\",\"token\":\"\",\"type\":2}', 1, '2022-04-12 15:03:11', '2022-04-20 19:51:30');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `enable`, `creation_time`, `update_time`) VALUES (9, 10, 5, '{\"name\":\"Harbor私有仓库\",\"repo\":\"harbor.xxx.com\",\"userName\":\"admin\",\"token\":\"=\",\"type\":5}', 1, '2022-04-26 16:53:14', '2022-08-10 14:07:46');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `enable`, `creation_time`, `update_time`) VALUES (10, 11, 6, '{\"name\":\"系统部署Webhook\",\"repo\":\"http://prod-kubelilin-apiserver-cls-hbktlqm5-svc-cluster-sgr.yoyogo:8080\",\"type\":6}', 1, '2022-08-10 14:02:23', '2022-08-10 14:02:48');
INSERT INTO `service_connection_details` (`id`, `main_id`, `type`, `detail`, `enable`, `creation_time`, `update_time`) VALUES (11, 12, 5, '{\"name\":\"Jenkins流水线引擎\",\"repo\":\"http://jenkins:32001\",\"userName\":\"jenkins\",\"password\":\"sgr-ci\",\"token\":\"\",\"type\":5}', 1, '2022-08-10 14:03:52', '2022-08-10 14:03:52');
COMMIT;

-- ----------------------------
-- Table structure for service_connection_type_code
-- ----------------------------
DROP TABLE IF EXISTS `service_connection_type_code`;
CREATE TABLE `service_connection_type_code` (
                                                `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                                `service_type` int unsigned NOT NULL,
                                                `type_code` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                                                `type_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
                                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='服务连接类型';

-- ----------------------------
-- Records of service_connection_type_code
-- ----------------------------
BEGIN;
INSERT INTO `service_connection_type_code` (`id`, `service_type`, `type_code`, `type_name`) VALUES (1, 1, 'vcs', 'GIT仓库');
INSERT INTO `service_connection_type_code` (`id`, `service_type`, `type_code`, `type_name`) VALUES (2, 2, 'hub', '镜像仓库');
INSERT INTO `service_connection_type_code` (`id`, `service_type`, `type_code`, `type_name`) VALUES (3, 3, 'pipeline', '流水线引擎');
INSERT INTO `service_connection_type_code` (`id`, `service_type`, `type_code`, `type_name`) VALUES (4, 4, 'system', '系统回调');
COMMIT;

-- ----------------------------
-- Table structure for service_connection_type_list
-- ----------------------------
DROP TABLE IF EXISTS `service_connection_type_list`;
CREATE TABLE `service_connection_type_list` (
                                                `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                                `service_type` int unsigned NOT NULL,
                                                `value` int unsigned NOT NULL,
                                                `name` varchar(100) NOT NULL,
                                                `default` tinyint unsigned NOT NULL,
                                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of service_connection_type_list
-- ----------------------------
BEGIN;
INSERT INTO `service_connection_type_list` (`id`, `service_type`, `value`, `name`, `default`) VALUES (1, 1, 1, 'Github', 0);
INSERT INTO `service_connection_type_list` (`id`, `service_type`, `value`, `name`, `default`) VALUES (2, 1, 2, 'Gitlab', 1);
INSERT INTO `service_connection_type_list` (`id`, `service_type`, `value`, `name`, `default`) VALUES (3, 1, 3, 'Gogs', 0);
INSERT INTO `service_connection_type_list` (`id`, `service_type`, `value`, `name`, `default`) VALUES (4, 1, 4, 'Gitee', 0);
INSERT INTO `service_connection_type_list` (`id`, `service_type`, `value`, `name`, `default`) VALUES (5, 2, 0, 'Docker Registy', 0);
INSERT INTO `service_connection_type_list` (`id`, `service_type`, `value`, `name`, `default`) VALUES (6, 3, 5, 'Jenkins', 1);
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
                                `menu_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '编码',
                                `menu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '目录名称',
                                `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图标',
                                `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由路径',
                                `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'react组件路径',
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
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (31, 0, '', '服务治理', '', '/applications/service', './applications/service', 0, 27, 0, 1, '2021-12-24 07:03:05', '2022-06-30 14:21:26');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (32, 0, '', 'DevOps', 'ProjectOutlined', '/devops', './devops', 1, 0, 0, 1, '2021-12-24 07:30:43', '2021-12-24 07:30:43');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (34, 0, '', '网络中心', 'RocketOutlined', '/components', './components', 1, 0, 0, 1, '2021-12-24 07:33:40', '2022-01-11 04:04:25');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (35, 0, '', '网关(API Gateway)', '', '/components/apigateway', './components/apigateway', 0, 34, 0, 1, '2021-12-24 07:44:50', '2022-03-30 15:29:51');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (36, 0, '', '监控中心', 'RadarChartOutlined', '/monitor', './monitor', 1, 0, 0, 1, '2021-12-24 07:51:59', '2021-12-24 07:51:59');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (37, 0, '', '部署环境', '', '/applications/info/deployments', './applications/info/deployments', 0, 27, 0, 1, '2021-12-27 13:07:42', '2021-12-27 13:07:42');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (38, 0, '', '服务(Service)', '', '/applications/serviceconfig', './applications/serviceconfig', 0, 34, 2, 1, '2022-01-11 04:05:42', '2022-06-24 19:44:21');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (40, 0, '', '服务集成', '', '/resources/serviceConnection', './resources/serviceConnection', 0, 24, 0, 1, '2022-03-24 16:49:43', '2022-08-10 20:19:02');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES (41, 0, '', '项目管理', '', '/devops/projects', './devops/projects', 0, 32, 0, 1, '2022-03-25 11:48:43', '2022-03-31 18:19:51');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant`;
CREATE TABLE `sgr_tenant` (
                              `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                              `t_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '租户名称',
                              `t_code` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '租户编码',
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
                                          `nickname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '应用中文名称',
                                          `remarks` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '集群应用备注',
                                          `git` varchar(500) CHARACTER SET utf8 NOT NULL COMMENT '集群应用绑定的git地址',
                                          `imagehub` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '集群应用绑定镜像仓库地址',
                                          `level` smallint unsigned NOT NULL COMMENT '应用级别',
                                          `language` smallint unsigned NOT NULL COMMENT '开发语言',
                                          `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
                                          `create_time` datetime DEFAULT NULL COMMENT '创建时间',
                                          `update_time` datetime DEFAULT NULL COMMENT '更新时间',
                                          `labels` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '应用标签',
                                          `git_type` varchar(20) CHARACTER SET utf8 NOT NULL COMMENT 'git类型 github/ gitee/ gogs/gitlab',
                                          `sc_id` bigint unsigned DEFAULT '0' COMMENT '服务连接git类型的凭据ID',
                                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='集群应用';

-- ----------------------------
-- Records of sgr_tenant_application
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_application_pipelines
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_application_pipelines`;
CREATE TABLE `sgr_tenant_application_pipelines` (
                                                    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'Pipeline ID',
                                                    `appid` bigint unsigned NOT NULL COMMENT '应用ID',
                                                    `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '流水线名称, appid 下唯一',
                                                    `dsl` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '流水线DSL',
                                                    `taskStatus` int unsigned DEFAULT NULL COMMENT '流水线任务状态( ready=0 , running=1, success=2, fail=3,  )',
                                                    `lastTaskId` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '最后一次任务执行ID',
                                                    `status` tinyint unsigned NOT NULL,
                                                    `creation_time` datetime DEFAULT NULL,
                                                    `update_time` datetime DEFAULT NULL,
                                                    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='应用流水线';

-- ----------------------------
-- Records of sgr_tenant_application_pipelines
-- ----------------------------
BEGIN;
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
) ENGINE=InnoDB AUTO_INCREMENT=155 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='部署发布记录';

-- ----------------------------
-- Records of sgr_tenant_deployment_record
-- ----------------------------
BEGIN;
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
                                          `image_hub` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
                                          `app_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
                                          `workload_type` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
                                          `replicas` int unsigned NOT NULL DEFAULT '1',
                                          `service_enable` tinyint unsigned DEFAULT NULL,
                                          `service_name` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
                                          `service_away` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
                                          `service_port` int unsigned NOT NULL DEFAULT '0',
                                          `service_port_type` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
                                          `last_image` varchar(350) CHARACTER SET utf8 DEFAULT NULL,
                                          `level` varchar(8) CHARACTER SET utf8 DEFAULT NULL,
                                          PRIMARY KEY (`id`),
                                          KEY `levev_idx` (`level`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='集群部署';

-- ----------------------------
-- Records of sgr_tenant_deployments
-- ----------------------------
BEGIN;
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
                                                     `environments` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
                                                     `workdir` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
                                                     `run_cmd` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
                                                     `run_params` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
                                                     `podstop` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
                                                     `liveness` varchar(300) CHARACTER SET utf8 DEFAULT NULL,
                                                     `readness` varchar(300) CHARACTER SET utf8 DEFAULT NULL,
                                                     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='应用部署容器配置';

-- ----------------------------
-- Records of sgr_tenant_deployments_containers
-- ----------------------------
BEGIN;
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
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_role
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_role`;
CREATE TABLE `sgr_tenant_role` (
                                   `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                   `role_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色编码',
                                   `role_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
                                   `description` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色描述',
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
                                   `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
                                   `account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '账号',
                                   `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
                                   `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机',
                                   `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
                                   `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
                                   `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                   `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息';

-- ----------------------------
-- Records of sgr_tenant_user
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES (1, 1, 'admin', '平台管理员', 'ef73781effc5774100f87fe2f437a435', '13111111111', 'xx11@hotmail.com', 1, NULL, '2022-08-10 10:26:35');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES (11, 1, 'user1', '111', 'ef73781effc5774100f87fe2f437a435', '18631111111', '111', 1, NULL, '2022-08-10 10:27:02');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES (15, 1, 'ghost', 'ghost', 'ef73781effc5774100f87fe2f437a435', '18630111111', '9022@qq.com', 1, NULL, '2022-08-10 10:27:26');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES (17, 39, 'com-dev-admin', '公司研发团队admin', 'ef73781effc5774100f87fe2f437a435', '18630111111', '9022@qq.com', 1, '2022-01-11 19:26:55', '2022-08-10 10:28:45');
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
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户角色';

-- ----------------------------
-- Records of sgr_tenant_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (17, 0, 2, '2022-01-10 21:21:20', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (22, 16, 1, '2022-01-11 19:03:33', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (23, 17, 2, '2022-01-11 19:26:55', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (27, 1, 1, '2022-08-10 10:26:35', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (28, 11, 2, '2022-08-10 10:27:02', NULL);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES (29, 15, 2, '2022-08-10 10:27:26', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
