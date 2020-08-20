-- 开始初始化数据 ;
BEGIN;
INSERT INTO sys_dict_type VALUES (11, '电池状态', 'sys_charge_status', 0, '1', '1', '电池状态列表', '2020-04-11 15:52:48', NULL, NULL);
INSERT INTO sys_dict_type VALUES (12, '网络状态', 'sys_net_status', 0, '1', '1', '网络状态列表', '2020-04-11 15:52:48', NULL, NULL);

INSERT INTO sys_dict_data VALUES (32, 0, '搁置', '0', 'sys_charge_status', NULL, NULL, NULL, 0, NULL, '1', NULL, '电池搁置', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (33, 0, '充电', '1', 'sys_charge_status', NULL, NULL, NULL, 0, NULL, '1', NULL, '电池充电', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (34, 0, '放电', '2', 'sys_charge_status', NULL, NULL, NULL, 0, NULL, '1', NULL, '电池放电', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (35, 0, '离线', '0', 'sys_net_status', NULL, NULL, NULL, 0, NULL, '1', NULL, '电池离线', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);
INSERT INTO sys_dict_data VALUES (36, 0, '在线', '1', 'sys_net_status', NULL, NULL, NULL, 0, NULL, '1', NULL, '电池在线', '2020-03-15 18:38:42', '2020-03-15 18:38:42', NULL);

INSERT INTO casbin_rule VALUES ('p', 'admin', '/api/bm1/batterylist', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'common', '/api/bm1/batterylist', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'tester', '/api/bm1/batterylist', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'admin', '/api/bm1/batterylist/:pkg_id', 'DELETE', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'common', '/api/bm1/batterylist/:pkg_id', 'DELETE', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'tester', '/api/bm1/batterylist/:pkg_id', 'DELETE', NULL, NULL, NULL);

INSERT INTO casbin_rule VALUES ('p', 'admin', '/api/bm1/batterydetail', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'common', '/api/bm1/batterydetail', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'tester', '/api/bm1/batterydetail', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'admin', '/api/bm1/batterydetail/dtu_statusinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'common', '/api/bm1/batterydetail/dtu_statusinfo', 'GET', NULL, NULL, NULL);
INSERT INTO casbin_rule VALUES ('p', 'tester', '/api/bm1/batterydetail/dtu_statusinfo', 'GET', NULL, NULL, NULL);

INSERT INTO `sys_menu` VALUES (4, 'battery', '电池信息', 'example', '/api/bm1/battery', '/0/4', 'M', '', '', 0, 1, '', 'Layout', 4, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL);
INSERT INTO `sys_menu` VALUES (5, 'batterylist', '电池列表', 'component', '/api/bm1/batterylist', '/0/4/5', 'C', '', '', 4, 1, '', '/batterylist/index', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:10:42', NULL);
INSERT INTO `sys_menu` VALUES (6, 'batterydetail', '电池详情', 'date', '/api/bm1/batterydetail', '/0/4/6', 'C', '', '', 4, 1, '', '/batterydetail/index', 50, '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:10:42', NULL);
INSERT INTO sys_role_menu VALUES (2, 4, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (2, 5, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (2, 6, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (2, 7, 'common', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 4, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 5, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 6, 'admin', NULL, NULL);
INSERT INTO sys_role_menu VALUES (1, 7, 'admin', NULL, NULL);


COMMIT;

-- 数据完成 ;