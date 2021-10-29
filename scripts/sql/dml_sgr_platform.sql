INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('31','1','3','2021-09-17 14:29:29','2021-09-17 14:29:29');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('32','1','4','2021-09-17 14:29:29','2021-09-17 14:29:29');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('33','1','5','2021-09-17 14:29:29','2021-09-17 14:29:29');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('34','1','6','2021-09-17 14:29:29','2021-09-17 14:29:29');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('35','1','7','2021-09-17 14:29:29','2021-09-17 14:29:29');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('46','2','7','2021-09-23 17:58:14','2021-09-23 17:58:14');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('47','2','5','2021-09-23 17:58:14','2021-09-23 17:58:14');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('48','2','4','2021-09-23 17:58:14','2021-09-23 17:58:14');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('49','10','1','2021-09-24 16:45:46','2021-09-24 16:45:46');
INSERT INTO `sgr_role_menu_map` (`id`, `role_id`, `menu_id`, `creation_time`, `update_time`) VALUES ('50','11','1','2021-09-24 16:51:16','2021-09-24 16:51:16');

INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES ('1','1','001','仪表盘','dashboard','/dashboard','','1','0','0','1',null,'2021-09-27 14:33:40');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES ('3','1','002','系统管理','user','/account','','1','0','0','1','2021-09-17 14:15:51','2021-09-27 15:13:25');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES ('4','1','003','路由管理','smile','/account/route','./account/route','0','3','2','1','2021-09-17 14:16:30','2021-09-29 09:43:54');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES ('5','1','004','用户管理','smile','/account/manage','./account/manage','0','3','4','1','2021-09-17 14:16:43','2021-09-29 09:44:05');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES ('6','1','005','租户管理','smile','/account/tenant','./account/tenant','0','3','5','1','2021-09-17 14:16:50','2021-09-29 09:44:08');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES ('7','1','006','角色管理','user','/account/role','./account/role','0','3','3','1','2021-09-17 14:17:01','2021-09-29 09:44:01');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES ('8','1','007','分析页','smile','/dashboard/analysis','./dashboard/analysis','0','1','1','1','2021-09-27 12:00:00','2021-09-29 06:26:46');
INSERT INTO `sgr_sys_menu` (`id`, `tenant_id`, `menu_code`, `menu_name`, `icon`, `path`, `component`, `is_root`, `parent_id`, `sort`, `status`, `creation_time`, `update_time`) VALUES ('9','0','008','监控','smile','/dashboard/monitor','./dashboard/monitor','0','1','2','1','2021-09-28 12:57:49','2021-09-29 09:43:39');

INSERT INTO `sgr_tenant` (`id`, `t_name`, `t_code`, `status`, `creation_time`, `update_time`) VALUES ('1','admin','admin','2','2021-10-14 15:11:18','2021-10-14 15:11:18');
INSERT INTO `sgr_tenant` (`id`, `t_name`, `t_code`, `status`, `creation_time`, `update_time`) VALUES ('35','adfasdf','44444','1',null,null);
INSERT INTO `sgr_tenant` (`id`, `t_name`, `t_code`, `status`, `creation_time`, `update_time`) VALUES ('36','第二个租户','SECOND','1',null,null);
INSERT INTO `sgr_tenant` (`id`, `t_name`, `t_code`, `status`, `creation_time`, `update_time`) VALUES ('37','第三个用户','Third','1',null,null);


INSERT INTO `sgr_tenant_role` (`id`, `role_code`, `role_name`, `description`, `status`, `tenant_id`, `creation_time`, `update_time`) VALUES ('1','PlatformAdmin','平台管理员',null,'1','1','2021-09-24 08:48:09','2021-09-24 08:48:09');
INSERT INTO `sgr_tenant_role` (`id`, `role_code`, `role_name`, `description`, `status`, `tenant_id`, `creation_time`, `update_time`) VALUES ('2','TenantAdmin','租户管理员',null,'1','1','2021-09-24 08:47:16','2021-09-24 08:47:16');
INSERT INTO `sgr_tenant_role` (`id`, `role_code`, `role_name`, `description`, `status`, `tenant_id`, `creation_time`, `update_time`) VALUES ('10','guest','guest',null,'1','1','2021-09-24 08:49:15','2021-09-24 08:49:15');
INSERT INTO `sgr_tenant_role` (`id`, `role_code`, `role_name`, `description`, `status`, `tenant_id`, `creation_time`, `update_time`) VALUES ('11','g2','g2',null,'1','1','2021-09-24 08:49:21','2021-09-24 08:49:21');

INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES ('1','1','admin','admin','123456','13877668829','zl.hxd@hotmail.com','1','2021-08-30 17:33:09','2021-09-14 16:23:43');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES ('2','1','xx','xx','123456','13877668821','xx@hotmail.com','1','2021-08-31 10:01:54','2021-09-16 16:45:28');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES ('3','1','xx2','xx2','123456','13877668811','xx2@hotmail.com','1','2021-08-31 10:53:08','2021-09-14 16:25:17');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES ('4','1','xx1','xx1','123456','13872368829','xx1@hotmail.com','1','2021-09-10 11:17:02','2021-09-14 16:23:45');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES ('5','1','小崔','xiaocui','123456','13347668829','xiaocui@hotmail.com','1','2021-09-10 15:22:11','2021-09-16 16:47:54');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES ('8','1','dxx1','大崔x1','123456','18623535890','86509022@qq.com','1',null,'2021-10-14 14:25:43');
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES ('9','36','第二个租户Administrator','SECONDadmin','123456','','','0','2021-10-14 16:47:28',null);
INSERT INTO `sgr_tenant_user` (`id`, `tenant_id`, `user_name`, `account`, `password`, `mobile`, `email`, `status`, `creation_time`, `update_time`) VALUES ('10','37','第三个用户Administrator','Thirdadmin','123456','','','0','2021-10-14 16:57:07',null);

INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES ('5','0','0','2021-09-10 11:47:50',null);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES ('6','0','0','2021-09-10 13:54:36',null);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES ('8','5','1','2021-09-10 15:25:03',null);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES ('10','1','1','2021-09-13 14:01:59',null);
INSERT INTO `sgr_tenant_user_role` (`id`, `user_id`, `role_id`, `creation_time`, `update_time`) VALUES ('11','10','1','2021-10-14 16:57:07',null);

