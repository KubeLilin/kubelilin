/*
 Navicat Premium Data Transfer

 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Schema         : sgr_pass

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 29/12/2021 10:33:12
*/

CREATE DATABASE IF NOT EXISTS sgr_pass
     DEFAULT CHARACTER SET utf8
     DEFAULT COLLATE utf8_general_ci;

use sgr_pass;

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
) ENGINE=InnoDB AUTO_INCREMENT=156 DEFAULT CHARSET=utf8mb4  COMMENT='角色菜单权限影射';

-- ----------------------------
-- Records of sgr_role_menu_map
-- ----------------------------
BEGIN;
INSERT INTO `sgr_role_menu_map` VALUES (46, 2, 7, '2021-09-23 17:58:14', '2021-09-23 17:58:14');
INSERT INTO `sgr_role_menu_map` VALUES (47, 2, 5, '2021-09-23 17:58:14', '2021-09-23 17:58:14');
INSERT INTO `sgr_role_menu_map` VALUES (48, 2, 4, '2021-09-23 17:58:14', '2021-09-23 17:58:14');
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
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4  COMMENT='菜单';

-- ----------------------------
-- Records of sgr_sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sgr_sys_menu` VALUES (1, 1, '001', '概览', 'dashboard', '/dashboard/analysis', './dashboard/analysis', 1, 0, 0, 1, NULL, '2021-12-24 07:53:29');
INSERT INTO `sgr_sys_menu` VALUES (3, 1, '002', '管理中心', 'SettingOutlined', '/account', './account', 1, 0, 99, 1, '2021-09-17 14:15:51', '2021-12-24 07:23:17');
INSERT INTO `sgr_sys_menu` VALUES (4, 1, '003', '菜单管理', 'smile', '/account/route', './account/route', 0, 3, 2, 1, '2021-09-17 14:16:30', '2021-12-24 07:00:45');
INSERT INTO `sgr_sys_menu` VALUES (5, 1, '004', '用户管理', 'smile', '/account/manage', './account/manage', 0, 3, 4, 1, '2021-09-17 14:16:43', '2021-09-29 09:44:05');
INSERT INTO `sgr_sys_menu` VALUES (6, 1, '005', '租户管理', 'smile', '/account/tenant', './account/tenant', 0, 3, 5, 1, '2021-09-17 14:16:50', '2021-09-29 09:44:08');
INSERT INTO `sgr_sys_menu` VALUES (7, 1, '006', '角色管理', 'user', '/account/role', './account/role', 0, 3, 3, 1, '2021-09-17 14:17:01', '2021-09-29 09:44:01');
INSERT INTO `sgr_sys_menu` VALUES (24, 0, '', '资源中心', 'ClusterOutlined', '/resources', './resources', 1, 0, 0, 1, '2021-12-24 06:32:51', '2021-12-24 07:21:09');
INSERT INTO `sgr_sys_menu` VALUES (25, 0, '', '集群管理', '', '/resources/clusters', './resources/clusters', 0, 24, 0, 1, '2021-12-24 06:43:14', '2021-12-24 06:43:14');
INSERT INTO `sgr_sys_menu` VALUES (26, 0, '', '命名空间', '', '/resources/namespaces', './resources/namespaces', 0, 24, 0, 1, '2021-12-24 06:43:44', '2021-12-24 06:43:44');
INSERT INTO `sgr_sys_menu` VALUES (27, 0, '', '应用中心', 'SendOutlined', '/applications', './applications', 1, 0, 0, 1, '2021-12-24 06:47:01', '2021-12-24 07:24:01');
INSERT INTO `sgr_sys_menu` VALUES (28, 0, '', '应用管理', '', '/applications/apps', './applications/apps', 0, 27, 10, 1, '2021-12-24 06:48:52', '2021-12-28 03:48:22');
INSERT INTO `sgr_sys_menu` VALUES (29, 0, '', '配置管理', '', '/applications/configs', './applications/configs', 0, 27, 0, 1, '2021-12-24 06:50:20', '2021-12-24 06:50:20');
INSERT INTO `sgr_sys_menu` VALUES (30, 0, '', '服务与路由', '', '/resources/services', './resources/services', 0, 24, 0, 1, '2021-12-24 07:01:27', '2021-12-24 07:01:27');
INSERT INTO `sgr_sys_menu` VALUES (31, 0, '', '服务治理', '', '/applications/serviceconfig', './applications/serviceconfig', 0, 27, 0, 1, '2021-12-24 07:03:05', '2021-12-24 07:03:05');
INSERT INTO `sgr_sys_menu` VALUES (32, 0, '', 'DevOps', 'ProjectOutlined', '/devops', './devops', 1, 0, 0, 1, '2021-12-24 07:30:43', '2021-12-24 07:30:43');
INSERT INTO `sgr_sys_menu` VALUES (33, 0, '', '流水线管理', '', '/devops', './devops', 0, 32, 0, 1, '2021-12-24 07:31:08', '2021-12-24 07:31:08');
INSERT INTO `sgr_sys_menu` VALUES (34, 0, '', '组件中心', 'RocketOutlined', '/components', './components', 1, 0, 0, 1, '2021-12-24 07:33:40', '2021-12-24 07:33:40');
INSERT INTO `sgr_sys_menu` VALUES (35, 0, '', '微服务网关', '', '/components/apigateway', './components/apigateway', 0, 34, 0, 1, '2021-12-24 07:44:50', '2021-12-24 07:44:50');
INSERT INTO `sgr_sys_menu` VALUES (36, 0, '', '监控中心', 'RadarChartOutlined', '/monitor', './monitor', 1, 0, 0, 1, '2021-12-24 07:51:59', '2021-12-24 07:51:59');
INSERT INTO `sgr_sys_menu` VALUES (37, 0, '', '部署环境', '', '/applications/info/deployments', './applications/info/deployments', 0, 27, 0, 1, '2021-12-27 13:07:42', '2021-12-27 13:07:42');
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
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COMMENT='租户';

-- ----------------------------
-- Records of sgr_tenant
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant` VALUES (1, 'admin', 'admin', 2, '2021-10-14 15:11:18', '2021-10-14 15:11:18');
INSERT INTO `sgr_tenant` VALUES (35, 'adfasdf', '44444', 1, NULL, NULL);
INSERT INTO `sgr_tenant` VALUES (36, '第二个租户', 'SECOND', 1, NULL, NULL);
INSERT INTO `sgr_tenant` VALUES (37, '第三个用户', 'Third', 1, NULL, NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 COMMENT='集群应用';

-- ----------------------------
-- Records of sgr_tenant_application
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_application` VALUES (1, 0, 'nginx', NULL, '', 'nginx', NULL, 5, 7, 1, NULL, NULL, '');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_cluster
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_cluster`;
CREATE TABLE `sgr_tenant_cluster` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `tenant_id` bigint unsigned DEFAULT NULL COMMENT '租户ID',
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
  `replicas` int unsigned DEFAULT '1',
  `service_enable` tinyint unsigned DEFAULT NULL,
  `service_name` varchar(150) DEFAULT NULL,
  `service_away` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `service_port` int unsigned DEFAULT NULL,
  `service_port_type` varchar(8) DEFAULT NULL,
  `last_image` varchar(350) DEFAULT NULL,
  `level` varchar(8) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3 COMMENT='集群部署';

-- ----------------------------
-- Records of sgr_tenant_deployments
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_deployments` VALUES (1, 1, 'dev-nginx-cls-hbktlqm5', 'nginx', 3, 1, 1, 1, NULL, NULL, '', '', 'Deployment', 2, 1, 'dev-nginx-cls-hbktlqm5.svc.cluster.sgr', 'ClusterPort', 80, '', '', 'dev');
INSERT INTO `sgr_tenant_deployments` VALUES (2, 1, 'test-nginx-microk8s-cluster', 'nginx', 4, 2, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'test-nginx-microk8s-cluster.svc.cluster.sgr', 'ClusterPort', 80, '', '', 'test');
INSERT INTO `sgr_tenant_deployments` VALUES (3, 1, 'prod-nginx-cls-hbktlqm5', 'prod-nginx', 3, 1, 1, 1, NULL, NULL, '', '', 'Deployment', 1, 1, 'prod-nginx-cls-hbktlqm-svc-cluster-sgr', 'ClusterPort', 80, '', '', 'prod');
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of sgr_tenant_deployments_containers
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_deployments_containers` VALUES (1, '', 1, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (2, '', 2, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, '', '', '', '', '', '', '');
INSERT INTO `sgr_tenant_deployments_containers` VALUES (3, '', 3, 1, 'docker.io/library/nginx:alpine', '', '', 0.25, 128, 0.25, 256, '', '', '', '', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_namespace
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_namespace`;
CREATE TABLE `sgr_tenant_namespace` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned DEFAULT NULL COMMENT '租户ID',
  `cluster_id` bigint unsigned DEFAULT NULL COMMENT '集群ID',
  `namespace` varchar(50) NOT NULL COMMENT '命名空间名称',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  `status` tinyint NOT NULL COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3 COMMENT='集群_命名空间';

-- ----------------------------
-- Records of sgr_tenant_namespace
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_namespace` VALUES (1, 1, 3, 'yoyogo', '2021-12-24 16:24:21', '2021-12-24 16:24:23', 1);
INSERT INTO `sgr_tenant_namespace` VALUES (2, 1, 4, 'sukt-core', '2021-12-24 16:54:47', '2021-12-24 16:54:49', 1);
INSERT INTO `sgr_tenant_namespace` VALUES (5, 1, 3, 'sgr-test', '2021-12-24 09:22:06', '2021-12-24 09:22:06', 1);
INSERT INTO `sgr_tenant_namespace` VALUES (6, 1, 4, 'sgr-dev', '2021-12-27 10:24:41', '2021-12-27 10:24:41', 1);
INSERT INTO `sgr_tenant_namespace` VALUES (7, 1, 4, 'sgr-test', '2021-12-27 11:20:34', '2021-12-27 11:20:34', 1);
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4  COMMENT='租户角色';

-- ----------------------------
-- Records of sgr_tenant_role
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_role` VALUES (1, 'PlatformAdmin', '平台管理员', NULL, 1, 1, '2021-09-24 08:48:09', '2021-09-24 08:48:09');
INSERT INTO `sgr_tenant_role` VALUES (2, 'TenantAdmin', '租户管理员', NULL, 1, 1, '2021-09-24 08:47:16', '2021-09-24 08:47:16');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_user
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_user`;
CREATE TABLE `sgr_tenant_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint NOT NULL COMMENT '租户',
  `user_name` varchar(50) DEFAULT NULL COMMENT '用户名',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `mobile` varchar(20) DEFAULT NULL COMMENT '手机',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4  COMMENT='用户信息';

-- ----------------------------
-- Records of sgr_tenant_user
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_user` VALUES (1, 1, 'admin', 'admin', '123456', '13877668829', 'zl.hxd@hotmail.com', 1, '2021-08-30 17:33:09', '2021-09-14 16:23:43');
COMMIT;

-- ----------------------------
-- Table structure for sgr_tenant_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sgr_tenant_user_role`;
CREATE TABLE `sgr_tenant_user_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户id',
  `role_id` bigint NOT NULL COMMENT '角色id',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4  COMMENT='用户角色';

-- ----------------------------
-- Records of sgr_tenant_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sgr_tenant_user_role` VALUES (5, 0, 0, '2021-09-10 11:47:50', NULL);
INSERT INTO `sgr_tenant_user_role` VALUES (6, 0, 0, '2021-09-10 13:54:36', NULL);
INSERT INTO `sgr_tenant_user_role` VALUES (8, 5, 1, '2021-09-10 15:25:03', NULL);
INSERT INTO `sgr_tenant_user_role` VALUES (10, 1, 1, '2021-09-13 14:01:59', NULL);
INSERT INTO `sgr_tenant_user_role` VALUES (11, 10, 1, '2021-10-14 16:57:07', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
