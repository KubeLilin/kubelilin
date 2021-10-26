/******************************************/
/*   DatabaseName = sgr_paas   */
/*   TableName = sgr_role_menu_map   */
/******************************************/
CREATE TABLE `sgr_role_menu_map` (
                                     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                     `role_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
                                     `menu_id` bigint(20) unsigned NOT NULL COMMENT '菜单ID',
                                     `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                     `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单权限影射'
;

/******************************************/
/*   DatabaseName = sgr_paas   */
/*   TableName = sgr_sys_menu   */
/******************************************/
CREATE TABLE `sgr_sys_menu` (
                                `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                `tenant_id` bigint(11) NOT NULL COMMENT '租户',
                                `menu_code` varchar(100) NOT NULL COMMENT '编码',
                                `menu_name` varchar(50) NOT NULL COMMENT '目录名称',
                                `icon` varchar(50) DEFAULT NULL COMMENT '图标',
                                `path` varchar(100) NOT NULL COMMENT '路由路径',
                                `component` varchar(100) DEFAULT NULL COMMENT 'react组件路径',
                                `is_root` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否是根目录',
                                `parent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父层级id',
                                `sort` int(11) NOT NULL DEFAULT '0' COMMENT '权重，正序排序',
                                `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
                                `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COMMENT='菜单'
;

/******************************************/
/*   DatabaseName = sgr_paas   */
/*   TableName = sgr_tenant   */
/******************************************/
CREATE TABLE `sgr_tenant` (
                              `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                              `t_name` varchar(50) NOT NULL COMMENT '租户名称',
                              `t_code` varchar(16) NOT NULL COMMENT '租户编码',
                              `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态',
                              `creation_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
                              `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `un_code` (`t_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COMMENT='租户'
;

/******************************************/
/*   DatabaseName = sgr_paas   */
/*   TableName = sgr_tenant_cluster   */
/******************************************/
CREATE TABLE `sgr_tenant_cluster` (
                                      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
                                      `tenant_id` bigint(20) unsigned DEFAULT NULL COMMENT '租户ID',
                                      `nickname` varchar(50) NOT NULL COMMENT '别名',
                                      `name` varchar(50) NOT NULL COMMENT '集群名称',
                                      `version` varchar(50) DEFAULT NULL COMMENT 'k8s 版本号',
                                      `distribution` varchar(30) DEFAULT NULL COMMENT '来源',
                                      `config` text NOT NULL COMMENT 'k8s config text',
                                      `sort` int(11) DEFAULT NULL COMMENT '排序',
                                      `status` tinyint(4) NOT NULL COMMENT '状态',
                                      `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
                                      `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                      PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='集群信息'
;

/******************************************/
/*   DatabaseName = sgr_paas   */
/*   TableName = sgr_tenant_namespace   */
/******************************************/
CREATE TABLE `sgr_tenant_namespace` (
                                        `id` bigint(20) unsigned NOT NULL,
                                        `tenant_id` bigint(20) unsigned DEFAULT NULL COMMENT '租户ID',
                                        `cluster_id` bigint(20) unsigned DEFAULT NULL COMMENT '集群ID',
                                        `namespace` varchar(50) NOT NULL COMMENT '命名空间名称',
                                        `create_time` datetime NOT NULL COMMENT '创建时间',
                                        `update_time` datetime NOT NULL COMMENT '更新时间',
                                        `status` tinyint(4) NOT NULL COMMENT '状态',
                                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='集群_命名空间'
;

/******************************************/
/*   DatabaseName = sgr_paas   */
/*   TableName = sgr_tenant_role   */
/******************************************/
CREATE TABLE `sgr_tenant_role` (
                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                   `role_code` varchar(30) NOT NULL COMMENT '角色编码',
                                   `role_name` varchar(50) NOT NULL COMMENT '角色名称',
                                   `description` varchar(50) DEFAULT NULL COMMENT '角色描述',
                                   `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态',
                                   `tenant_id` bigint(11) NOT NULL COMMENT '租户',
                                   `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                   `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `un_role_code_name` (`role_code`,`role_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='租户角色'
;

/******************************************/
/*   DatabaseName = sgr_paas   */
/*   TableName = sgr_tenant_user   */
/******************************************/
CREATE TABLE `sgr_tenant_user` (
                                   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                   `tenant_id` bigint(11) NOT NULL COMMENT '租户',
                                   `user_name` varchar(50) DEFAULT NULL COMMENT '用户名',
                                   `account` varchar(50) NOT NULL COMMENT '账号',
                                   `password` varchar(255) NOT NULL COMMENT '密码',
                                   `mobile` varchar(20) DEFAULT NULL COMMENT '手机',
                                   `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
                                   `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态',
                                   `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                   `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息'
;

/******************************************/
/*   DatabaseName = sgr_paas   */
/*   TableName = sgr_tenant_user_role   */
/******************************************/
CREATE TABLE `sgr_tenant_user_role` (
                                        `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                                        `user_id` bigint(20) NOT NULL COMMENT '用户id',
                                        `role_id` bigint(20) NOT NULL COMMENT '角色id',
                                        `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
                                        `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='用户角色'
;
