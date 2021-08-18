/******************************************/
/*   DatabaseName = sgr_platform   */
/*   TableName = sgr_context   */
/******************************************/
CREATE TABLE `sgr_sys_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint(11) NOT NULL COMMENT '租户',
  `menu_code` varchar(100) NOT NULL COMMENT '编码',
  `menu_name` varchar(50) NOT NULL COMMENT '目录名称',
  `is_root` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否是根目录',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父层级id',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '权重，正序排序',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `createtion_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_c_code` (`context_code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单'
;

/******************************************/
/*   DatabaseName = sgr_platform   */
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='租户'
;

/******************************************/
/*   DatabaseName = sgr_platform   */
/*   TableName = tenant_role   */
/******************************************/
CREATE TABLE `sgr_tenant_role` (
  `id` bigint(20) unsigned NOT NULL,
  `role_code` varchar(30) NOT NULL COMMENT '角色编码',
  `role_name` varchar(50) NOT NULL COMMENT '角色名称',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态',
  `tenant_id` bigint(11) NOT NULL COMMENT '租户',
  `createtion_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_role_code_name` (`role_code`,`role_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='租户角色'
;

/******************************************/
/*   DatabaseName = sgr_platform   */
/*   TableName = tenant_user   */
/******************************************/
CREATE TABLE `sgr_tenant_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint(11) NOT NULL COMMENT '租户',
  `user_name` varchar(50) DEFAULT NULL COMMENT '用户名',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `mobile` varchar(10) DEFAULT NULL COMMENT '手机',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态',
  `createtion_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息'
;

/******************************************/
/*   DatabaseName = sgr_platform   */
/*   TableName = tenant_user_role   */
/******************************************/
CREATE TABLE `tenant_user_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `creation_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色'
;
