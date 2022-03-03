/*
 Navicat Premium Data Transfer

 Source Server         : 47.100.213.49
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : 47.100.213.49:3306
 Source Schema         : sgr_pass

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 03/03/2022 17:28:26
*/
CREATE DATABASE IF NOT EXISTS sgr_pass
     DEFAULT CHARACTER SET utf8
     DEFAULT COLLATE utf8_general_ci;


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sgr_code_application_language
-- ----------------------------
DROP TABLE IF EXISTS `sgr_code_application_language`;
CREATE TABLE `sgr_code_application_language` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(8) DEFAULT NULL,
  `name` varchar(50) NOT NULL,
  `sort` smallint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3 COMMENT='字典-应用开发语言';

-- ----------------------------
-- Records of sgr_code_application_language
-- ----------------------------
BEGIN;
INSERT INTO `sgr_code_application_language` VALUES (5, '0001', 'go', 0);
INSERT INTO `sgr_code_application_language` VALUES (6, '0002', 'java', 0);
INSERT INTO `sgr_code_application_language` VALUES (7, '0003', 'nodejs', 0);
INSERT INTO `sgr_code_application_language` VALUES (8, '0004', 'python', 0);
INSERT INTO `sgr_code_application_language` VALUES (9, '0005', '.net', 0);
COMMIT;

-- ----------------------------
-- Table structure for sgr_code_application_level
-- ----------------------------
DROP TABLE IF EXISTS `sgr_code_application_level`;
CREATE TABLE `sgr_code_application_level` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(8) DEFAULT NULL,
  `name` varchar(50) NOT NULL,
  `sort` smallint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3 COMMENT='字典-应用级别';

-- ----------------------------
-- Records of sgr_code_application_level
-- ----------------------------
BEGIN;
INSERT INTO `sgr_code_application_level` VALUES (5, '0001', 'P0', 0);
INSERT INTO `sgr_code_application_level` VALUES (6, '0002', 'P1', 0);
INSERT INTO `sgr_code_application_level` VALUES (7, '0003', 'P2', 0);
INSERT INTO `sgr_code_application_level` VALUES (8, '0004', 'p3', 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=571 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色菜单权限影射';

-- ----------------------------
-- Records of sgr_role_menu_map
-- ----------------------------
BEGIN;
INSERT INTO `sgr_role_menu_map` VALUES (49, 10, 1, '2021-09-24 16:45:46', '2021-09-24 16:45:46');
INSERT INTO `sgr_role_menu_map` VALUES (50, 11, 1, '2021-09-24 16:51:16', '2021-09-24 16:51:16');
INSERT INTO `sgr_role_menu_map` VALUES (136, 1, 3, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (137, 1, 4, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (138, 1, 5, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (139, 1, 6, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (140, 1, 7, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (141, 1, 1, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (142, 1, 24, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (143, 1, 25, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (144, 1, 26, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (145, 1, 30, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (146, 1, 32, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (147, 1, 33, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (148, 1, 34, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (149, 1, 35, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (150, 1, 36, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (151, 1, 27, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (152, 1, 28, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (153, 1, 29, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (154, 1, 31, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (155, 1, 37, '2021-12-27 21:07:53', '2021-12-27 21:07:53');
INSERT INTO `sgr_role_menu_map` VALUES (248, 12, 1, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (249, 12, 27, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (250, 12, 29, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (251, 12, 31, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (252, 12, 37, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (253, 12, 28, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (254, 12, 32, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (255, 12, 33, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (256, 12, 34, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (257, 12, 35, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (258, 12, 38, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (259, 12, 36, '2022-01-11 19:58:25', '2022-01-11 19:58:25');
INSERT INTO `sgr_role_menu_map` VALUES (557, 2, 1, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (558, 2, 27, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (559, 2, 29, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (560, 2, 31, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (561, 2, 37, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (562, 2, 28, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (563, 2, 32, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (564, 2, 33, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (565, 2, 34, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (566, 2, 35, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (567, 2, 38, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (568, 2, 7, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (569, 2, 5, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
INSERT INTO `sgr_role_menu_map` VALUES (570, 2, 3, '2022-02-21 17:00:39', '2022-02-21 17:00:39');
COMMIT;

-- ----------------------------
-- Table structure for sgr_sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sgr_sys_menu`;
CREATE TABLE `sgr_sys_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `menu_code` varchar(100) NOT NULL COMMENT '编码',
  `menu_name` varchar(50) NOT NULL COMMENT '目录名称',
  `icon` varchar(50) DEFAULT NULL COMMENT '图标',
  `path` varchar(100) NOT NULL COMMENT '路由路径',
  `component` varchar(100) DEFAULT NULL COMMENT 'react组件路径',
  `is_root` tinyint NOT NULL DEFAULT '0' COMMENT '是否是根目录',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父层级id',
  `sort` int NOT NULL DEFAULT '0' COMMENT '权重，正序排序',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='菜单';

-- ----------------------------
-- Records of sgr_sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sgr_sys_menu` VALUES (1, 0, '001', '概览', 'dashboard', '/dashboard/analysis', './dashboard/analysis', 1, 0, 0, 1, NULL, '2022-01-11 20:02:38');
INSERT INTO `sgr_sys_menu` VALUES (3, 0, '002', '管理中心', 'SettingOutlined', '/account', './account', 1, 0, 99, 1, '2021-09-17 14:15:51', '2022-01-11 20:02:36');
INSERT INTO `sgr_sys_menu` VALUES (4, 0, '003', '菜单管理', 'smile', '/account/route', './account/route', 0, 3, 2, 1, '2021-09-17 14:16:30', '2022-01-11 20:02:34');
INSERT INTO `sgr_sys_menu` VALUES (5, 0, '004', '用户管理', 'smile', '/account/manage', './account/manage', 0, 3, 4, 1, '2021-09-17 14:16:43', '2022-01-11 20:02:33');
INSERT INTO `sgr_sys_menu` VALUES (6, 0, '005', '租户管理', 'smile', '/account/tenant', './account/tenant', 0, 3, 5, 1, '2021-09-17 14:16:50', '2022-01-11 20:02:31');
INSERT INTO `sgr_sys_menu` VALUES (7, 0, '006', '角色管理', 'user', '/account/role', './account/role', 0, 3, 3, 1, '2021-09-17 14:17:01', '2022-01-11 20:02:29');
INSERT INTO `sgr_sys_menu` VALUES (24, 0, '', '资源中心', 'ClusterOutlined', '/resources', './resources', 1, 0, 0, 1, '2021-12-24 06:32:51', '2021-12-24 07:21:09');
INSERT INTO `sgr_sys_menu` VALUES (25, 0, '', '集群管理', '', '/resources/clusters', './resources/clusters', 0, 24, 0, 1, '2021-12-24 06:43:14', '2021-12-24 06:43:14');
INSERT INTO `sgr_sys_menu` VALUES (26, 0, '', '命名空间', '', '/resources/namespaces', './resources/namespaces', 0, 24, 0, 1, '2021-12-24 06:43:44', '2021-12-24 06:43:44');
INSERT INTO `sgr_sys_menu` VALUES (27, 0, '', '应用中心', 'SendOutlined', '/applications', './applications', 1, 0, 0, 1, '2021-12-24 06:47:01', '2021-12-24 07:24:01');
INSERT INTO `sgr_sys_menu` VALUES (28, 0, '', '应用管理', '', '/applications/apps', './applications/apps', 0, 27, 10, 1, '2021-12-24 06:48:52', '2021-12-28 03:48:22');
INSERT INTO `sgr_sys_menu` VALUES (29, 0, '', '配置管理', '', '/applications/configs', './applications/configs', 0, 27, 0, 1, '2021-12-24 06:50:20', '2021-12-24 06:50:20');
INSERT INTO `sgr_sys_menu` VALUES (31, 0, '', '服务治理', '', '/applications/serviceconfig', './applications/serviceconfig', 0, 27, 0, 1, '2021-12-24 07:03:05', '2021-12-24 07:03:05');
INSERT INTO `sgr_sys_menu` VALUES (32, 0, '', 'DevOps', 'ProjectOutlined', '/devops', './devops', 1, 0, 0, 1, '2021-12-24 07:30:43', '2021-12-24 07:30:43');
INSERT INTO `sgr_sys_menu` VALUES (33, 0, '', '流水线管理', '', '/devops/pipeline', './devops/pipeline', 0, 32, 0, 1, '2021-12-24 07:31:08', '2022-02-10 06:46:21');
INSERT INTO `sgr_sys_menu` VALUES (34, 0, '', '网络中心', 'RocketOutlined', '/components', './components', 1, 0, 0, 1, '2021-12-24 07:33:40', '2022-01-11 04:04:25');
INSERT INTO `sgr_sys_menu` VALUES (35, 0, '', '微服务网关', '', '/components/apigateway', './components/apigateway', 0, 34, 0, 1, '2021-12-24 07:44:50', '2021-12-24 07:44:50');
INSERT INTO `sgr_sys_menu` VALUES (36, 0, '', '监控中心', 'RadarChartOutlined', '/monitor', './monitor', 1, 0, 0, 1, '2021-12-24 07:51:59', '2021-12-24 07:51:59');
INSERT INTO `sgr_sys_menu` VALUES (37, 0, '', '部署环境', '', '/applications/info/deployments', './applications/info/deployments', 0, 27, 0, 1, '2021-12-27 13:07:42', '2021-12-27 13:07:42');
INSERT INTO `sgr_sys_menu` VALUES (38, 0, '', '服务与路由', '', '/resources/services', './resources/services', 0, 34, 2, 1, '2022-01-11 04:05:42', '2022-01-11 04:05:42');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant`;
CREATE TABLE `sgr_tenant` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `t_name` varchar(50) NOT NULL COMMENT '租户名称',
  `t_code` varchar(16) NOT NULL COMMENT '租户编码',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `creation_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_code` (`t_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='租户';

-- ----------------------------
-- Records of sgr_tenant
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant` VALUES (1, '平台研发团队', 'administration', 1, '2022-01-25 11:13:39', '2022-01-25 11:13:39');
INSERT INTO `sgr_tenant` VALUES (39, '公司研发团队', 'com-dev', 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_application
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_application`;
CREATE TABLE `sgr_tenant_application` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_Id` bigint unsigned NOT NULL COMMENT '租户ID',
  `name` varchar(50) NOT NULL COMMENT '集群应用名称(英文唯一)',
  `nickname` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '应用中文名称',
  `remarks` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '集群应用备注',
  `git` varchar(500) NOT NULL COMMENT '集群应用绑定的git地址',
  `imagehub` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '集群应用绑定镜像仓库地址',
  `level` smallint unsigned NOT NULL COMMENT '应用级别',
  `language` smallint unsigned NOT NULL COMMENT '开发语言',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `labels` varchar(100) DEFAULT NULL COMMENT '应用标签',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb3 COMMENT='集群应用';

-- ----------------------------
-- Records of sgr_tenant_application
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_application` VALUES (1, 1, 'nginx', NULL, '', 'https://gogs.xiaocui.site/administration/nginx.git', 'https://harbor.xiaocui.site/apps/', 5, 7, 1, NULL, NULL, 'web');
INSERT INTO `sgr_tenant_application` VALUES (2, 1, 'yoyogodemo', NULL, '', 'https://gitee.com/yoyofx/yoyogo.git', 'https://harbor.xiaocui.site/apps/', 5, 5, 1, NULL, NULL, 'web api');
INSERT INTO `sgr_tenant_application` VALUES (15, 1, 'kubelilin-apiserver', '', '', 'https://github.com/yoyofxteam/sgr-platform-api.git', '', 5, 5, 1, NULL, NULL, 'apiserver');
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3 COMMENT='应用流水线';

-- ----------------------------
-- Records of sgr_tenant_application_pipelines
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_application_pipelines` VALUES (2, 1, 'ewfwewef', '', 0, '', 1, '2022-02-21 11:24:26', '2022-02-28 10:54:42');
INSERT INTO `sgr_tenant_application_pipelines` VALUES (3, 1, 'wefwefszzzzv2', '[{\"name\":\"代码\",\"steps\":[{\"name\":\"拉取代码\",\"key\":\"git_pull\",\"save\":true,\"content\":{\"git\":\"https://gogs.xiaocui.site/administration/nginx.git\",\"branch\":\"dev\"}}]},{\"name\":\"编译构建\",\"steps\":[{\"name\":\"编译命令\",\"key\":\"code_build\",\"save\":true,\"content\":{\"buildEnv\":\"golang\",\"buildScript\":\"# 编译命令，注：当前已在代码根路径下\\ngo env -w GOPROXY=https://goproxy.cn,direct\\ngo build -ldflags=\\\"-s -w\\\" -o app .\\n\",\"buildFile\":\"./Dockerfile\"}},{\"name\":\"命令执行\",\"key\":\"cmd_shell\",\"save\":true,\"content\":{\"shell\":\"# bash\"}}]},{\"name\":\"部署\",\"steps\":[{\"name\":\"应用部署\",\"key\":\"app_deploy\",\"save\":true,\"content\":{\"depolyment\":1}}]},{\"name\":\"通知\",\"steps\":[{\"name\":\"命令执行\",\"key\":\"cmd_shell\",\"save\":true,\"content\":{\"shell\":\"# bash\"}}]}]', NULL, '', 1, '2022-02-21 11:46:36', '2022-02-25 07:02:11');
INSERT INTO `sgr_tenant_application_pipelines` VALUES (4, 2, 'TEST环境部署', '[{\"name\":\"代码\",\"steps\":[{\"name\":\"拉取代码\",\"key\":\"git_pull\",\"save\":true,\"content\":{\"git\":\"https://gitee.com/yoyofx/yoyogo.git\",\"branch\":\"master\"}}]},{\"name\":\"编译构建\",\"steps\":[{\"name\":\"编译命令\",\"key\":\"code_build\",\"save\":true,\"content\":{\"buildEnv\":\"golang\",\"buildScript\":\"# 编译命令，注：当前已在代码根路径下\\ngo env -w GOPROXY=https://goproxy.cn,direct\\n\\n\",\"buildFile\":\"./examples/simpleweb/Dockerfile\"}}]},{\"name\":\"部署\",\"steps\":[{\"name\":\"应用部署\",\"key\":\"app_deploy\",\"save\":true,\"content\":{\"depolyment\":10}}]},{\"name\":\"通知\",\"steps\":[{\"name\":\"命令执行\",\"key\":\"cmd_shell\",\"save\":true,\"content\":{\"shell\":\"# bash\\necho \'helloworld\'\"}}]}]', 1, '38', 1, '2022-02-28 05:04:56', '2022-03-02 10:12:15');
INSERT INTO `sgr_tenant_application_pipelines` VALUES (5, 15, '腾讯云k8s部署', '[{\"name\":\"代码\",\"steps\":[{\"name\":\"拉取代码\",\"key\":\"git_pull\",\"save\":true,\"content\":{\"git\":\"https://github.com/yoyofxteam/sgr-platform-api.git\",\"branch\":\"master\"}}]},{\"name\":\"编译构建\",\"steps\":[{\"name\":\"编译命令\",\"key\":\"code_build\",\"save\":true,\"content\":{\"buildEnv\":\"golang\",\"buildScript\":\"# 编译命令，注：当前已在代码根路径下\\ngo env -w GOPROXY=https://goproxy.cn,direct\\n# go build -ldflags=\\\"-s -w\\\" -o app .\\n\",\"buildFile\":\"./src/Dockerfile_Prod\"}}]},{\"name\":\"部署\",\"steps\":[{\"name\":\"应用部署\",\"key\":\"app_deploy\",\"save\":true,\"content\":{\"depolyment\":11}}]},{\"name\":\"通知\",\"steps\":[{\"name\":\"命令执行\",\"key\":\"cmd_shell\",\"save\":true,\"content\":{\"shell\":\"# bash\"}}]}]', 0, '', 1, '2022-03-03 17:03:11', '2022-03-03 17:21:55');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_cluster
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_cluster`;
CREATE TABLE `sgr_tenant_cluster` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `nickname` varchar(50) NOT NULL COMMENT '别名',
  `name` varchar(50) NOT NULL COMMENT '集群名称',
  `version` varchar(50) DEFAULT NULL COMMENT 'k8s 版本号',
  `distribution` varchar(30) DEFAULT NULL COMMENT '来源',
  `config` text NOT NULL COMMENT 'k8s config text',
  `sort` int DEFAULT NULL COMMENT '排序',
  `status` tinyint NOT NULL COMMENT '状态',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3 COMMENT='集群信息';

-- ----------------------------
-- Records of sgr_tenant_cluster
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_cluster` VALUES (3, 1, '', 'cls-hbktlqm5', 'v1.18.4-tke.11', '', 'apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJd01UQXlOakExTkRFek5sb1hEVE13TVRBeU5EQTFOREV6Tmxvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTlVCCmZCNDJHVkxGaXVxSnZmcGdWSVYzT3h5VjNJU0l2SU5aaFdvMW9IRFF6bnFVOXU1b0hnNEZmbkxHU3BRSVdLN1gKOXNNRTBQc21DaCtwY2ZiMEszNW9OVnlFUWU3dndYVFVaU3Znb0lWb0xNWENYRlB5L2xkTVFTQURuZWw4Vnl0dgpYREVua0pvSTdIUEtmZ3E3czZ4cjh0VXFDQUZPMGErLzZHYWtFZXI0SlNVREQxZDFyY3dNVWd4VS9IaGQxVkhSClNQOEh3d1EvWHdIUzdpWUF2bmJOQ3pkNWxLWmNEeWV2ZVVsb2JrYmRxUllRMFplY3dGWjU4bDgyZFdyUWJkUFIKem9DcGRCdTJHQmtmZUtpUjhUb0Q2Q0IxYlFvb0JNSXQ4dG96dzNCL2JDMTRJeEZwUFgxdUNCWWJyK2NzUmNrTQoxOUM3NldqaUFCZ0hpMXNvTkgwQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0tVTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFMWll6VjRkZkF2Y1VPYXVFL0dGT2tvOVRQdWMKZ3h5NHVJczhhOURtTno4S3J0YjRsZG41bTljWnQ0RFRFUTJwbWwwOEY5cVBYRy9EY0FjWWE3ckFlVkZWUWNyMgpXMmI5dEVyMWZjTEhUWlF1aFpzQVVuT3FkcTYreDRkQTBTcTJLbm1hMWlnQmJoUWJCc3cyVGdIOEFtS0NxMXVXClhQVU1XWFE5NlB0eUYvWjEvTUdmcC9CMi9LZHdpVWd0WCtSQnlhWGsxaXJsZzZLb1owcVZDK1lxMTdOZG92QTYKa2habC9XZXlYaEZnaUhaclAwNHR1MzlhVHEyTkhsWk4vWjBvOXVYVUZ2RnVZYURwWGRNakxHNDFhU3p6KzJLVwpwZ0NkVGZXQkpGMzJMV2Npc21lcG5YVENxS2xEeWxDZW13bEg5OWljUGtUd2tOd2xXYlIzOFpKOFlsTT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\n    server: https://cls-hbktlqm5.ccs.tencent-cloud.com\n  name: cls-hbktlqm5\ncontexts:\n- context:\n    cluster: cls-hbktlqm5\n    user: \"86509022\"\n  name: cls-hbktlqm5-86509022-context-default\ncurrent-context: cls-hbktlqm5-86509022-context-default\nkind: Config\npreferences: {}\nusers:\n- name: \"86509022\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURDRENDQWZDZ0F3SUJBZ0lJUnJvbVFheWpCRFF3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TURFeU1qRXhOakEyTVRGYUZ3MDBNREV5TWpFeE5qQTJNVEZhTURJeApFakFRQmdOVkJBb1RDWFJyWlRwMWMyVnljekVjTUJvR0ExVUVBeE1UT0RZMU1Ea3dNakl0TVRZd09EVTJOamMzCk1UQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQUw2NXRiZDdsNjlDTTdsL1dNTHAKSTJSZ3ZXTmtic3BjcGllMlpVUXNIZ0JGcE1jc1JvZFVNUW5LbmZrZHh3NlIrSVAzTWVtVnVEZHYwR0dEalFFVAp6K0V6aU13dmorUkdqZkxrc0E1WWVxRjZnVUJjOFArSkMvVW11SDVNM0taK29IYXVZT2M1VTREeXBscVFpWTdOCmdNdk1SZ29LcmNOaVE1V3BxZ2k2UzNtTnF2Z0I3azFsR2hFeFhoRDZERExmWFRvaGhiMGQ0Q1lZUnlLSkZiWUUKMGFHTSt0TjVTOUdOZXlQZGZORTNFcVVzbHp4emJDajZseWxMQ3NUTGw1L0pLUGRvblBWd2JoV0RlZ1ZtNEp5LwpzbUpoSW81MHdLRDZDdVUxa1JYZXgzcGMrbUZDMUwvV3VMdUVOVWcxSVRPdWh0YjlsdXVqYVlnVUk3V2JvMkVOCi9Wa0NBd0VBQWFNL01EMHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHQVFVRkJ3TUMKQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUFqRFgybQp2c0N4T05VSmhLUDZWRHBFZy9MWTB6Ni9GMHBkblpoVFBFb3pwMDg1R01EbWZCdUxuSmxIbE54czZjc0t5Qk53CmtDMWtUM3Q1NndkWXUxVjBtN0tvZHFudjloWGxBYWZ6S0pmT1F3NlVzZmtqSjJZL0wwT2FMaWhLZ0IzYnpMWVcKZE0veVgwYlpVNzRlSHRraS9ydTRzOGFLamg4UzVMQnpzQ25Yc1dsOXdwdCt6ZXBkVnJVbDZPWmhOOVdTLzU2awpYcWZxMjAzeDMvWk1kdG53R0huOWZxQisyb2Yrc3BwY0dYdGJnanV0LzBXQVVhZWdTcUFqRGxLekcrMkJrTGJmCjB6dFp1UUFBVXBBSVBZeU1PZU1YOUdObmVSWWM5QUQ2ZnpVMWdTaEdIaEY3RDZobDl0NXpjRmZueUdGRmVSR2IKVVBEWVc0MlNJL0I5Vk1NNwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBdnJtMXQzdVhyMEl6dVg5WXd1a2paR0M5WTJSdXlseW1KN1psUkN3ZUFFV2t4eXhHCmgxUXhDY3FkK1IzSERwSDRnL2N4NlpXNE4yL1FZWU9OQVJQUDRUT0l6QytQNUVhTjh1U3dEbGg2b1hxQlFGencKLzRrTDlTYTRma3pjcG42Z2RxNWc1emxUZ1BLbVdwQ0pqczJBeTh4R0NncXR3MkpEbGFtcUNMcExlWTJxK0FIdQpUV1VhRVRGZUVQb01NdDlkT2lHRnZSM2dKaGhISW9rVnRnVFJvWXo2MDNsTDBZMTdJOTE4MFRjU3BTeVhQSE5zCktQcVhLVXNLeE11WG44a285MmljOVhCdUZZTjZCV2JnbkwreVltRWlqblRBb1BvSzVUV1JGZDdIZWx6NllVTFUKdjlhNHU0UTFTRFVoTTY2RzF2Mlc2Nk5waUJRanRadWpZUTM5V1FJREFRQUJBb0lCQUVXVWNMUG9sZlR0UFB1TwpkdTVjcVhuRVJUT09mMUM2TGkvTXZmTDUrVlAyRkdCSlNjMnpMRlM3STVpdmdXQlNab3lXVVJJN2VjSlh1M2puCnlqZzdaeHBzZDVxdU8xdDNWZS9uK0VhemhzR0VkTVRyWTB5R1RlTjQ1ZFBGN0xXYytxTnhpSTZ5ZmtGTHhON0QKWGp2SHd4WVdodkxBNUpXa01xM1dBTjlBUDZxdkh0N0FPVW4zb3lLZklWUTRHVlNXaWVPL3hBR0dMMUVMamZKVwpzYmR4NG92eHdHb1JKMGpod0x3THViMVp1ZnFGQkY0TldKa3NRRG94SW9RNzc0R2VCdWlPQnlBbEZJVW1ydERnCkhyS21Pci92cjVZelNuZjc3UFNrUUlwWnNETG5JLzRMRmxhc0xpQ0U0dVlMYnh5aDNoQVh4dllNL3E1RWUyWHIKakdjUXZvVUNnWUVBK2VTOSt2MitvalRUcFpGMmJLWmtIbzVac3FSSlFGcVB5ZVh1aHdSOWxkeTJvV1M2Z25ibQpnV0RQaVVJUGsvOG4vaGEvU1BVRWpLbjdEUEl3NWp2MEduYU1ucEhZVDdpTWlyajJWMGZEZWF6eUJoZVEvenRqClFBSXZLeXAvcmN6cGttTUpBTWJpamVVQ3dvNE9CQnh1bFpMdUlyc3BMeTBsbCtOUlo5NGs4KzhDZ1lFQXcyTFUKZ2tCR2dqT0ZPR2RCOHg5UHdlcm5VVG00b3huajZ5dGJzY2VacG1wQzZBNG8zcytqc1U3YXRXOVJDOHk4U0ZhOApQR2IrV0haL0cwajlXam9XOFl0d1ZlaUxTMzNsNE9GeFRSM2xxdkZZNDFUbjk0cGYvazdIMnVQTWkwazNhaDh4ClZVUG81OHJRTEpobVlaV2d1b3FjdkZONElkZnRPNVp0SFAveXV6Y0NnWUJLeTh5UmM2RzdxMVF1R252M3lWWHUKVDIxSnF5TEJ3Rm1KZE9rUVFLZldVMW5XdE8rZVhUaGhRVGpkUElpdEk1STAyMW9sM0RDZ2FjQmEvNkxqUnM2cApuUkk1NUMxNnJ5Smg0enJZcFFJOVNTYW43Q1hhUDB4VnZGR2grZlo5YnZmNmVPb1k4VzZlU2cycGJodUQzMzY2CkJtQ0F4TVJ1K25SbUlnUWJzc0ljd1FLQmdRQ2hBakpJMjVxKzlLSFZweEdydmFQR0UwTm5wZjlIT0xDZlBPMmwKQk96VFBFSENaTmk5TTdLYkRIWWlpWWJxQ0Y4bjVZSGM3S3F3VDRYVEFFVDNNMk53elExWFhmaGJ6M1c5NlVtcQoyUFpIOWZiZjd6bnd2WEQ1YWdZN2xQa2Ixc3Y1Z1pidndyU05QbzVxRVhSYytpYW5VazV5eXYzMk5hL2pLTVRsCjN1MHg1UUtCZ1FDVWkyeWVMOTYvcUREZ3p3amdyOWNWZFhGdU8xVFZtNXl2QmMzK015ZFh1dzlYaGFmRGowMXYKUmcvT2loNnptZmRkOEwxb05GNTY2ZmVja3ZKYzRCTWJVeDRRd2o3c1A3SStibGYyUzlLc3B3U00yWTN4SFBkSApsZU02eldnajF1bFBUbWt4SkpWYk4rL2RzZGZHTnpFQU9Qc2dKYkVKbDJFVnBVZWhlOWZ4bVE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\n', 0, 1, '2021-12-24 17:06:43', '2021-12-24 17:06:43');
INSERT INTO `sgr_tenant_cluster` VALUES (4, 1, '', 'microk8s-cluster', 'v1.22.4-3+adc4115d990346', '', 'apiVersion: v1\r\nclusters:\r\n- cluster:\r\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREekNDQWZlZ0F3SUJBZ0lVQzdPVGlodEx1ZWFVYUNOZ3VFQ1F1aWRmV1ljd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0Z6RVZNQk1HQTFVRUF3d01NVEF1TVRVeUxqRTRNeTR4TUI0WERUSXlNREl3TVRFek16RTBOVm9YRFRNeQpNREV6TURFek16RTBOVm93RnpFVk1CTUdBMVVFQXd3TU1UQXVNVFV5TGpFNE15NHhNSUlCSWpBTkJna3Foa2lHCjl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFuWHVTbmVwc3NIRUp0UjlWMUpjQTJaQmxnQk0vTVFXeDAwcU8KQzJuVWk2UCtUZDloOENtam1wS2p5TURwYzBOV2tVQVp3L0JOd1VYN0dHWG5EcjBsaW5oc1lTWkxydkdWOGw3TgpFR3dhd1VsY3FzZXova3dOZDI5Q2pBaTU4TUFaOWtKNUpianJLN2lXU0tPeHdkK0pvcWtSSVh3clJCUzU4VEFlCjVTNXFNK3dlUXhtcy9oSUpnZFNSL3ZSNC8yMERZd3JKT3ZkR0kzb1I1cUxjdm8weldYZ0NvNFZkNUQwYUQzMGMKMW50eVViODNjc2JPcGtvcU9HNnRsejBtNk5tZUNNLy9FVW5ncEtRbkIyUXNwaFhMYjFsT0VEZWF0K3ZyU1RMUApKQzUreG5xK1J0d3VXL0Fod0tJb1pBN2NiT2Y1MHljYlFqclArUTBqWHVZMi9uY0REd0lEQVFBQm8xTXdVVEFkCkJnTlZIUTRFRmdRVWQxbG9jOEpRaVo2R2pLMTZNQ0Z6OSs2bkllOHdId1lEVlIwakJCZ3dGb0FVZDFsb2M4SlEKaVo2R2pLMTZNQ0Z6OSs2bkllOHdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQwpBUUVBVmpBVDFkT1BEN3VMNE9DcXNCTVBTQ2JVeTFBS0wyT3AreFNJQ0NDRUV6V1RhNVk3bkRuYlh0bnVZS1hkCnhOdkdrSStZMFh4TE1oZm1PSklaNUEwTjRwdmRuYzdxSWZ0d01uWHNVU3dMTENkTE5sa1VpcWVRZ1F6UWtUeWcKaHB3aUU4REtlL2xvblMwczh2VUkvdVpEb0xiOFF0d2U2U3ZTcmpoajJsNUJLZUNaYjB6RHl1ZGNRZ1VDYWNJYwpMcWo5SXZ0amZaUG9selAxSTV6aDE0VG5pUFpERGFhYXplNGkxU1VET1JPeFp0TjhFSXlDWWxPN0t0b0txWFNsCmhqb0dzeEVnY1N2NmZBMlhVVGhKd1IvYUZPemxXZW4vOWV2VHlrZjRhSzN0YnFEMWVzL3M4cHVXY1g2ZExwNlYKK1NYc3E4RTl2WjhZRGRlQkRJZS9lSXlnN3c9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\r\n    server: https://47.100.213.49:8325\r\n  name: microk8s-cluster\r\ncontexts:\r\n- context:\r\n    cluster: microk8s-cluster\r\n    user: admin\r\n  name: microk8s\r\ncurrent-context: microk8s\r\nkind: Config\r\npreferences: {}\r\nusers:\r\n- name: admin\r\n  user:\r\n    token: VXJKWEp1ZmJJS09JVklxbDBoNFpoSitkY0g5b3lqcVUvTUxqcFlWbHAxQT0K', 0, 1, '2022-02-07 17:57:37', '2022-02-07 17:57:37');
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='部署发布记录';

-- ----------------------------
-- Records of sgr_tenant_deployment_record
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_deployment_record` VALUES (2, 1, 1, 'docker.io/library/nginx:alpine', 'githook', NULL, '2022-03-02 19:10:01', NULL, NULL, '2022-03-02 19:10:01');
INSERT INTO `sgr_tenant_deployment_record` VALUES (3, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', NULL, '2022-03-02 16:42:59', '成功', '', '2022-03-02 16:43:07');
INSERT INTO `sgr_tenant_deployment_record` VALUES (4, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', NULL, '2022-03-02 16:43:03', '成功', '', '2022-03-02 16:43:07');
INSERT INTO `sgr_tenant_deployment_record` VALUES (5, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', 0, '2022-03-02 08:44:47', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` VALUES (6, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', 0, '2022-03-02 08:45:29', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` VALUES (7, 1, 5, 'yoyofx/prism-desgin:v0.1', 'manual', 0, '2022-03-02 08:57:37', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` VALUES (8, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v36', 'githook', NULL, '2022-03-02 19:09:58', NULL, NULL, '2022-03-02 19:09:58');
INSERT INTO `sgr_tenant_deployment_record` VALUES (9, 2, 10, 'harbor.xiaocui.site/apps/pipeline-4-app-2:v38', 'githook', NULL, '2022-03-02 19:09:59', NULL, NULL, '2022-03-02 19:09:59');
INSERT INTO `sgr_tenant_deployment_record` VALUES (10, 1, 1, 'docker.io/library/nginx:alpine', 'githook', 0, '2022-03-02 11:13:30', NULL, NULL, '2022-03-02 11:13:30');
INSERT INTO `sgr_tenant_deployment_record` VALUES (11, 1, 1, 'docker.io/library/nginx:alpine', 'githook', 0, '2022-03-02 11:13:42', NULL, NULL, '2022-03-02 11:13:42');
INSERT INTO `sgr_tenant_deployment_record` VALUES (12, 1, 1, 'docker.io/library/nginx:alpine', 'githook', 0, '2022-03-02 11:14:41', NULL, NULL, '2022-03-02 11:14:41');
INSERT INTO `sgr_tenant_deployment_record` VALUES (13, 1, 1, 'docker.io/library/nginx:alpine', 'githook', 0, '2022-03-02 11:16:22', NULL, NULL, '2022-03-02 11:16:22');
INSERT INTO `sgr_tenant_deployment_record` VALUES (14, 2, 10, 'yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-03 14:08:30', '成功', '', NULL);
INSERT INTO `sgr_tenant_deployment_record` VALUES (15, 2, 10, 'yoyofx/yoyogo-demo:v0.1', 'manual', 0, '2022-03-03 15:24:13', '成功', '', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_deployments
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_deployments`;
CREATE TABLE `sgr_tenant_deployments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL,
  `name` varchar(50) NOT NULL COMMENT '部署名称(英文唯一)',
  `nickname` varchar(50) NOT NULL COMMENT '部署中文名称',
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
  `service_name` varchar(150) DEFAULT NULL,
  `service_away` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `service_port` int unsigned NOT NULL DEFAULT '0',
  `service_port_type` varchar(8) DEFAULT NULL,
  `last_image` varchar(350) DEFAULT NULL,
  `level` varchar(8) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `levev_idx` (`level`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3 COMMENT='集群部署';

-- ----------------------------
-- Records of sgr_tenant_deployments
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_deployments` VALUES (1, 1, 'dev-nginx-cls-hbktlqm5', 'nginx', 3, 1, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'dev-nginx-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 80, '', '', 'dev');
INSERT INTO `sgr_tenant_deployments` VALUES (2, 1, 'test-nginx-microk8s-cluster', 'nginx', 4, 2, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'svc-test-nginx-microk8s-cluster', 'ClusterPort', 80, '', '', 'test');
INSERT INTO `sgr_tenant_deployments` VALUES (3, 1, 'prod-nginx-cls-hbktlqm5', 'prod-nginx', 3, 1, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'prod-nginx-cls-hbktlqm-svc-cluster-sgr', 'ClusterPort', 80, '', '', 'prod');
INSERT INTO `sgr_tenant_deployments` VALUES (4, 1, 'dev-yoyogodemo-cls-hbktlqm5', 'yoyogo-demo', 3, 1, 2, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'dev-yoyogodemo-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'dev');
INSERT INTO `sgr_tenant_deployments` VALUES (5, 1, 'test-nginx-cls-hbktlqm5', 'prism-desgin', 3, 1, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'test-nginx-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8092, '', '', 'test');
INSERT INTO `sgr_tenant_deployments` VALUES (6, 1, 'prod-yoyogodemo-microk8s-cluster', 'yoyogo-demo正式环境', 4, 2, 2, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'prod-yoyogodemo-microk8s-cluster-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'prod');
INSERT INTO `sgr_tenant_deployments` VALUES (7, 1, 'dev-repo3-cls-hbktlqm5', 'test', 3, 1, 13, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'dev-repo3-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'dev');
INSERT INTO `sgr_tenant_deployments` VALUES (8, 1, 'test-repo3-cls-hbktlqm5', 'test', 3, 1, 13, 1, NULL, NULL, '', '', 'Deployment', 0, 1, 'test-repo3-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'test');
INSERT INTO `sgr_tenant_deployments` VALUES (9, 1, 'prod-repo3-cls-hbktlqm5', 'sfwwf', 3, 1, 13, 1, NULL, NULL, '', '', 'Deployment', 1, 0, 'prod-repo3-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'prod');
INSERT INTO `sgr_tenant_deployments` VALUES (10, 1, 'test-yoyogodemo-cls-hbktlqm5', 'hahaha', 3, 1, 2, 1, NULL, NULL, '', '', 'Deployment', 1, 0, 'test-yoyogodemo-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'test');
INSERT INTO `sgr_tenant_deployments` VALUES (11, 1, 'prod-kubelilin-apiserver-cls-hbktlqm5', '线上部署环境', 3, 1, 15, 1, NULL, NULL, '', '', 'Deployment', 1, 0, 'prod-kubelilin-apiserver-cls-hbktlqm5-svc-cluster-sgr', 'ClusterPort', 8080, '', '', 'prod');
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
  `environments` varchar(255) DEFAULT NULL,
  `workdir` varchar(200) DEFAULT NULL,
  `run_cmd` varchar(200) DEFAULT NULL,
  `run_params` varchar(100) DEFAULT NULL,
  `podstop` varchar(100) DEFAULT NULL,
  `liveness` varchar(300) DEFAULT NULL,
  `readness` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3 COMMENT='应用部署容器配置';

-- ----------------------------
-- Records of sgr_tenant_deployments_containers
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_deployments_containers` VALUES (1, '', 1, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, '[{\"key\":\"test\",\"value\":\"1\"}]', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (2, '', 2, 1, 'docker.io/library/nginx:alpine', '', '', 0.05, 64, 0.07, 128, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (3, '', 3, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (4, '', 4, 1, 'yoyofx/yoyogo-demo:v0.1', '', '', 0.10, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (5, '', 5, 1, 'yoyofx/prism-desgin:v0.1', '', '', 0.25, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (6, '', 6, 1, 'yoyofx/yoyogo-demo:v0.1', '', '', 0.25, 128, 0.25, 256, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (7, '', 7, 1, '', '', '', 0.25, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (8, '', 8, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, '[{\"key\":\"es\",\"value\":\"123\"},{\"key\":\"se\",\"value\":\"133\"}]', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (9, '', 9, 1, '', '', '', 0.25, 128, 0.25, 256, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (10, '', 10, 1, 'yoyofx/yoyogo-demo:v0.1', '', '', 0.25, 128, 0.25, 256, 'null', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (11, '', 11, 1, '', '', '', 0.25, 128, 0.25, 256, 'null', '', '', '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_namespace
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_namespace`;
CREATE TABLE `sgr_tenant_namespace` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户ID',
  `cluster_id` bigint unsigned NOT NULL COMMENT '集群ID',
  `namespace` varchar(50) NOT NULL COMMENT '命名空间名称',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `status` tinyint NOT NULL COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3 COMMENT='集群_命名空间';

-- ----------------------------
-- Records of sgr_tenant_namespace
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_namespace` VALUES (1, 1, 3, 'yoyogo', '2021-12-24 16:24:21', '2021-12-24 16:24:23', 1);
INSERT INTO `sgr_tenant_namespace` VALUES (2, 1, 4, 'sukt-core', '2021-12-24 16:54:47', '2021-12-24 16:54:49', 1);
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_role
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_role`;
CREATE TABLE `sgr_tenant_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role_code` varchar(30) NOT NULL COMMENT '角色编码',
  `role_name` varchar(50) NOT NULL COMMENT '角色名称',
  `description` varchar(50) DEFAULT NULL COMMENT '角色描述',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_role_code_name` (`role_code`,`role_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='租户角色';

-- ----------------------------
-- Records of sgr_tenant_role
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_role` VALUES (1, 'PlatformAdmin', '平台管理员', NULL, 1, 1, '2021-09-24 08:48:09', '2021-09-24 08:48:09');
INSERT INTO `sgr_tenant_role` VALUES (2, 'TenantAdmin', '租户管理员', NULL, 1, 1, '2021-09-24 08:47:16', '2021-09-24 08:47:16');
INSERT INTO `sgr_tenant_role` VALUES (12, 'tuser', '用户', '', 1, 39, '2022-01-11 11:46:09', '2022-01-11 11:46:09');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_user
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_user`;
CREATE TABLE `sgr_tenant_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL COMMENT '租户',
  `user_name` varchar(50) DEFAULT NULL COMMENT '用户名',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `mobile` varchar(20) DEFAULT NULL COMMENT '手机',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息';

-- ----------------------------
-- Records of sgr_tenant_user
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_user` VALUES (1, 1, 'admin', '平台管理员', '123456', '13877668829', 'zl.hxd@hotmail.com', 1, NULL, '2022-01-10 12:53:41');
INSERT INTO `sgr_tenant_user` VALUES (11, 1, 'user1', '111', '1234', '18630535890', '111', 1, NULL, '2022-01-10 13:13:11');
INSERT INTO `sgr_tenant_user` VALUES (15, 1, 'ghost', 'ghost', 'ghost', '18630535890', '86509022@qq.com', 1, NULL, '2022-01-10 13:25:58');
INSERT INTO `sgr_tenant_user` VALUES (17, 39, 'com-dev-admin', '公司研发团队admin', '1234abcd', '', '', 1, '2022-01-11 19:26:55', '2022-01-11 19:29:53');
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
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户角色';

-- ----------------------------
-- Records of sgr_tenant_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_user_role` VALUES (10, 1, 1, '2021-09-13 14:01:59', NULL);
INSERT INTO `sgr_tenant_user_role` VALUES (14, 11, 2, '2022-01-10 21:13:11', NULL);
INSERT INTO `sgr_tenant_user_role` VALUES (17, 0, 2, '2022-01-10 21:21:20', NULL);
INSERT INTO `sgr_tenant_user_role` VALUES (21, 15, 2, '2022-01-10 21:25:58', NULL);
INSERT INTO `sgr_tenant_user_role` VALUES (22, 16, 1, '2022-01-11 19:03:33', NULL);
INSERT INTO `sgr_tenant_user_role` VALUES (23, 17, 2, '2022-01-11 19:26:55', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
