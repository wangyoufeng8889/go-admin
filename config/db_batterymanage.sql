-- 开始初始化数据 ;
BEGIN;

INSERT INTO sys_setting VALUES (1, '电池监控管理系统', 'https://gitee.com/mydearzwj/image/raw/master/img/go-admin.png', '2020-08-05 14:27:34', '2020-08-07 10:13:31' ,NULL);



INSERT INTO sys_dict_type VALUES (11, '电池状态', 'sys_charge_status', 0, '1', '', '电池状态列表', '2020-04-11 15:52:48', NULL, NULL);
INSERT INTO sys_dict_type VALUES (12, '网络状态', 'sys_net_status', 0, '1', '', '网络状态列表', '2020-04-11 15:52:48', NULL, NULL);
INSERT INTO sys_dict_type VALUES (13, '电池类型', 'sys_pkg_type', 0, '1', '', '电池类型列表', '2020-04-11 15:52:48', NULL, NULL);
INSERT INTO sys_dict_type VALUES (14, 'DTU类型', 'sys_dtu_type', 0, '1', '', 'DTU类型列表', '2020-04-11 15:52:48', NULL, NULL);


INSERT INTO sys_dict_data VALUES (32, 0, '搁置', '0', 'sys_charge_status', '', '', '', 0, '', '1', '', '电池搁置', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (33, 0, '充电', '1', 'sys_charge_status', '', '', '', 0, '', '1', '', '电池充电', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (34, 0, '放电', '2', 'sys_charge_status', '', '', '', 0, '', '1', '', '电池放电', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (35, 0, '离线', '0', 'sys_net_status', '', '', '', 0, '', '1', '', '电池离线', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (36, 0, '在线', '1', 'sys_net_status', '', '', '', 0, '', '1', '', '电池在线', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (37, 0, '三元', '1', 'sys_pkg_type', '', '', '', 0, '', '1', '', '三元锂电池', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (38, 0, '磷酸铁锂', '2', 'sys_pkg_type', '', '', '', 0, '', '1', '', '磷酸铁锂电池', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (39, 0, '2G', '2', 'sys_dtu_type', '', '', '', 0, '', '1', '', 'DTU_2G', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (40, 0, '4G-CAT4', '4', 'sys_dtu_type', '', '', '', 0, '', '1', '', 'DTU_CAT1', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (41, 0, '4G-CAT1', '6', 'sys_dtu_type', '', '', '', 0, '', '1', '', 'DTU_CAT4', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);

INSERT INTO casbin_rule VALUES (100,'p', 'common', '/api/v1/logout', 'POST', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (101,'p', 'tester', '/api/v1/logout', 'POST', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (102,'p', 'common', '/api/v1/getinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (103,'p', 'tester', '/api/v1/getinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (104,'p', 'common', '/api/v1/menurole', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (105,'p', 'tester', '/api/v1/menurole', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (106,'p', 'admin', '/api/bm1/battery/dashboard', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (107,'p', 'common', '/api/bm1/battery/dashboard', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (108,'p', 'tester', '/api/bm1/battery/dashboard', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (109,'p', 'admin', '/api/bm1/battery/batterylist', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (110,'p', 'common', '/api/bm1/battery/batterylist', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (111,'p', 'tester', '/api/bm1/battery/batterylist', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (112,'p', 'admin', '/api/bm1/battery/batterylist/:bms_specinfoId', 'DELETE', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (113,'p', 'common', '/api/bm1/battery/batterylist/:bms_specinfoId', 'DELETE', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (114,'p', 'tester', '/api/bm1/battery/batterylist/:bms_specinfoId', 'DELETE', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (115,'p', 'admin', '/api/bm1/battery/batterydetail', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (116,'p', 'common', '/api/bm1/battery/batterydetail', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (117,'p', 'tester', '/api/bm1/battery/batterydetail', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (118,'p', 'admin', '/api/bm1/battery/batterydetail/batterysoc', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (119,'p', 'common', '/api/bm1/battery/batterydetail/batterysoc', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (120,'p', 'tester', '/api/bm1/battery/batterydetail/batterysoc', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (121,'p', 'admin', '/api/bm1/battery/dtulist', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (122,'p', 'common', '/api/bm1/battery/dtulist', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (123,'p', 'tester', '/api/bm1/battery/dtulist', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (124,'p', 'admin', '/api/bm1/battery/dtulist/:dtu_specInfoId', 'DELETE', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (125,'p', 'common', '/api/bm1/battery/dtulist/:dtu_specInfoId', 'DELETE', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (126,'p', 'tester', '/api/bm1/battery/dtulist/:dtu_specInfoId', 'DELETE', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (127,'p', 'admin', '/api/bm1/battery/batterymove', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (128,'p', 'common', '/api/bm1/battery/batterymove', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (129,'p', 'tester', '/api/bm1/battery/batterymove', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (130,'p', 'admin', '/api/bm1/battery/batterymove/location', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (131,'p', 'common', '/api/bm1/battery/batterymove/location', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (132,'p', 'tester', '/api/bm1/battery/batterymove/location', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (133,'p', 'admin', '/api/bm1/battery/dtudetail', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (134,'p', 'common', '/api/bm1/battery/dtudetail', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (135,'p', 'tester', '/api/bm1/battery/dtudetail', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (136,'p', 'admin', '/api/bm1/battery/dtudetail/dtu_statusinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (137,'p', 'common', '/api/bm1/battery/dtudetail/dtu_statusinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (138,'p', 'tester', '/api/bm1/battery/dtudetail/dtu_statusinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (139,'p', 'admin', '/api/bm1/battery/dtudetail/dtu_specinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (140,'p', 'common', '/api/bm1/battery/dtudetail/dtu_specinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (141,'p', 'tester', '/api/bm1/battery/dtudetail/dtu_specinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (142,'p', 'admin', '/api/bm1/battery/dtudetail/dtucsq', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (143,'p', 'common', '/api/bm1/battery/dtudetail/dtucsq', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES (144,'p', 'tester', '/api/bm1/battery/dtudetail/dtucsq', 'GET', NULL, NULL, NULL);

INSERT INTO `sys_menu` VALUES (4, 'battery', '电池信息', 'example', '/api/bm1/battery', '/0/4', 'M', '', '', 0, 1, '', 'Layout', 4, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL);
INSERT INTO `sys_menu` VALUES (5, 'batterylist', '电池列表', 'component', '/api/bm1/battery/batterylist', '/0/4/5', 'C', '', 'system:batterylist:list', 4, 1, '', '/batterylist/index', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:10:42', NULL);
INSERT INTO `sys_menu` VALUES (6, 'batterydetail', '电池详情', 'date', '/api/bm1/battery/batterydetail', '/0/4/6', 'C', '', '', 4, 1, '', '/batterydetail/index', 2, '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:10:42', NULL);
INSERT INTO `sys_menu` VALUES (7, 'dtulist', 'DTU列表', 'component', '/api/bm1/battery/dtulist', '/0/4/7', 'C', '', '', 4, 1, '', '/dtulist/index', 3, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:10:42', NULL);
INSERT INTO `sys_menu` VALUES (8, 'dtudetail', 'DTU详情', 'date', '/api/bm1/battery/dtudetail', '/0/4/8', 'C', '', '', 4, 1, '', '/dtudetail/index', 4, '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:10:42', NULL);

INSERT INTO sys_menu VALUES (500, '', '查看电池', '', '', '/0/2/5/500', 'F', '', 'system:batterylist:query', 5, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL);
INSERT INTO sys_menu VALUES (501, '', '删除电池', '', '', '/0/2/5/501', 'F', '', 'system:batterylist:remove', 5, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL);

INSERT INTO sys_menu VALUES (502, '', '查看DTU', '', '', '/0/2/7/502', 'F', '', 'system:dtulist:query', 7, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL);
INSERT INTO sys_menu VALUES (503, '', '删除DTU', '', '', '/0/2/7/503', 'F', '', 'system:dtulist:remove', 7, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL);


INSERT INTO sys_role_menu VALUES (2, 4, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (2, 5, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (2, 6, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (2, 7, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (2, 8, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (2, 500, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (2, 502, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 4, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 5, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 6, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 7, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 8, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 500, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 501, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 502, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 503, 'admin', NULL, NULL);

INSERT INTO sys_user VALUES (3, 'john', '13688888888', 2, '', '', '0', 'ww@126.com', 1, 1, '1', '', 'dd', '0', '2020-09-17 15:10:57', '2020-09-17 15:10:57', NULL, 'john', '$2a$10$mAA7kc1nvxs5E3A7do9AZuhMMVNdnMNFPqNUjpgnukOgLM5gwFL6a');

COMMIT;

-- 数据完成 ;