/*
 Navicat Premium Data Transfer

 Source Server         : centos8
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : centos8.cn:3306
 Source Schema         : gin-vben-admin

 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : 65001

 Date: 12/05/2021 20:09:26
*/

CREATE DATABASE IF NOT EXISTS go_vben_admin
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_general_ci;
USE go_vben_admin;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/role/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/menu/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/log/del_batch', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/log/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/password', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/role/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/dashboard/analysis', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/dept/edit', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/user/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/dept/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/user/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/menu/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/menu/edit', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/user/edit', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/user/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/dashboard', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/dept', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/dept/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/log/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/log', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/dept/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/user', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/role/edit', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/menu/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/user/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '/system/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/dashboard', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/dashboard/analysis', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/user/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/user', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/menu/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/dept/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/dept', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/log/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/log', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/user/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '49', '/system/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/user', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/role/edit', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/user/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/menu/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/dept/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/role/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/dept/edit', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/user/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/dept/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/log/del_batch', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/log/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/dept', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/log/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/log', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/menu/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/dashboard/analysis', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/menu/edit', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/user/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/dashboard', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/menu/del', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/password', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/dept/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/user/edit', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/role/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/user/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '55', '/system/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/dashboard/analysis', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/dept/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/log', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/menu/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/user/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/password', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/log/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/dashboard', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/dept', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/user', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/user/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '56', '/system/user/info', 'GET', '', '', '');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `pid` int(11) NULL DEFAULT NULL,
  `dept_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `order_no` int(11) NULL DEFAULT NULL,
  `level` int(11) NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dept_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (1, '2021-05-12 00:19:16', '2021-05-12 00:19:16', NULL, 0, '账号管理组', 1, 1, 1, '测试');
INSERT INTO `sys_dept` VALUES (3, '2021-05-10 21:06:35', '2021-05-10 21:06:35', NULL, 8, '财务部', 1, 2, 1, '');
INSERT INTO `sys_dept` VALUES (5, '2021-04-24 02:35:40', '2021-04-24 02:35:40', '2021-05-12 00:18:49', 1, '研发部', 1, 2, 1, '');
INSERT INTO `sys_dept` VALUES (6, '2021-04-24 02:36:02', '2021-04-24 02:36:02', '2021-05-10 03:29:16', 1, '市场部', 1, 2, 1, '');
INSERT INTO `sys_dept` VALUES (7, '2021-05-12 01:41:12', '2021-05-12 01:41:12', NULL, 0, '华南分部', 3, 1, 1, '');
INSERT INTO `sys_dept` VALUES (8, '2021-05-12 00:18:58', '2021-05-12 00:18:58', NULL, 0, '华中分部', 2, 1, 1, '');
INSERT INTO `sys_dept` VALUES (9, '2021-04-24 02:36:32', '2021-04-24 02:36:32', NULL, 7, '研发部', 1, 2, 1, '');
INSERT INTO `sys_dept` VALUES (10, '2021-04-24 02:36:57', '2021-04-24 02:36:57', '2021-05-10 03:29:19', 7, '市场部', 1, 2, 1, '');
INSERT INTO `sys_dept` VALUES (11, '2021-04-26 06:53:46', '2021-04-26 06:53:46', '2021-05-10 03:29:23', 7, '财务部', 1, 2, 1, '');
INSERT INTO `sys_dept` VALUES (12, '2021-04-28 01:00:16', '2021-04-28 01:00:16', '2021-05-10 18:38:29', 0, '总部', 0, 1, 1, '');
INSERT INTO `sys_dept` VALUES (13, '2021-04-28 00:47:46', '2021-04-28 00:47:46', NULL, 13, '系统管理', 1, 2, 1, '');
INSERT INTO `sys_dept` VALUES (14, '2021-04-28 01:01:23', '2021-04-28 01:01:23', '2021-05-10 18:38:29', 12, '系统管理', 1, 2, 1, '');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sort` int(11) NULL DEFAULT NULL,
  `flag` int(11) NULL DEFAULT NULL,
  `affix` tinyint(1) NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `en_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `cn_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `type` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '0代表目录1代表菜单2代表按钮',
  `parent_id` int(10) NULL DEFAULT 0,
  `status` int(11) NULL DEFAULT NULL,
  `api_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `keepalive` tinyint(10) NULL DEFAULT NULL COMMENT '是否缓存',
  `is_ext` int(255) NULL DEFAULT NULL COMMENT '是否外链',
  `level` int(11) NULL DEFAULT NULL,
  `api_method` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_menu_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '2021-05-11 21:35:24', '2021-05-11 21:35:24', NULL, '', '/dashboard', 0, 0, 0, NULL, '', '仪表大盘', 'LAYOUT', 0, 0, 1, '/dashboard', 1, 0, 1, 1);
INSERT INTO `sys_menu` VALUES (2, '2021-05-12 00:18:04', '2021-05-12 00:18:04', NULL, '', 'analysis', 4, 0, 0, 'ant-design:apple-filled', '', '今日数据', '/dashboard/analysis/index', 1, 1, 1, '/dashboard/analysis', 1, 0, 2, 3);
INSERT INTO `sys_menu` VALUES (3, '2021-04-28 06:37:14', '2021-04-28 06:37:14', NULL, '', '/system', 3, 0, 0, NULL, '', '系统管理', 'LAYOUT', 0, 0, 1, '/system', 0, 0, 1, 1);
INSERT INTO `sys_menu` VALUES (4, '2021-05-11 23:37:47', '2021-05-11 23:37:47', NULL, '', 'role', 3, 0, 0, 'ant-design:team-outlined', '', '角色管理', '/system/role/index', 1, 3, 1, '/system/role', 0, 0, 2, 1);
INSERT INTO `sys_menu` VALUES (5, '2021-05-11 23:41:04', '2021-05-11 23:41:04', NULL, '', 'user', 5, 0, 0, 'ant-design:qq-circle-filled', '', '账号管理', '/system/user/index', 1, 3, 1, '/system/user', 0, 0, 2, 1);
INSERT INTO `sys_menu` VALUES (8, '2021-05-12 00:17:10', '2021-05-12 00:17:10', NULL, '', 'menu', 7, 0, 0, 'ant-design:appstore-twotone', '', '菜单管理', '/system/menu/index', 1, 3, 1, '/system/menu', 0, 0, 2, 1);
INSERT INTO `sys_menu` VALUES (10, '2021-05-10 03:34:28', '2021-05-10 03:34:28', NULL, '', 'del', 6, 0, 0, NULL, '', '删除角色', '', 2, 4, 1, '/system/role/del', 0, 0, 3, 4);
INSERT INTO `sys_menu` VALUES (11, '2021-05-10 15:17:22', '2021-05-10 15:17:22', NULL, '', 'edit', 4, 0, 0, NULL, '', '编辑角色', '', 2, 4, 1, '/system/role/edit', 0, 0, 3, 3);
INSERT INTO `sys_menu` VALUES (26, '2021-05-10 03:35:44', '2021-05-10 03:35:44', NULL, '', 'edit', 3, 0, 0, NULL, '', '编辑菜单', '', 2, 8, 1, '/system/menu/edit', 1, 0, 3, 3);
INSERT INTO `sys_menu` VALUES (32, '2021-05-10 03:35:51', '2021-05-10 03:35:51', NULL, '', 'del', 6, 0, 0, NULL, '', '删除菜单', '', 2, 8, 1, '/system/menu/del', 1, 0, 3, 4);
INSERT INTO `sys_menu` VALUES (35, '2021-05-11 22:27:11', '2021-05-11 22:27:11', NULL, '', 'dept', 2, 0, 0, 'ant-design:apartment-outlined', '', '部门管理', '/system/dept/index', 1, 3, 1, '/system/dept', 1, 0, 2, 1);
INSERT INTO `sys_menu` VALUES (36, '2021-05-12 00:17:40', '2021-05-12 00:17:40', NULL, '', 'password', 8, 0, 0, 'ant-design:lock-outlined', '', '修改密码', '/system/password/index', 1, 3, 1, '/system/password', 1, 0, 2, 2);
INSERT INTO `sys_menu` VALUES (37, '2021-05-11 20:52:26', '2021-05-11 20:52:26', NULL, '', 'edit', 2, 0, 0, NULL, '', '编辑部门', '', 2, 35, 1, '/system/dept/edit', 1, 0, 3, 3);
INSERT INTO `sys_menu` VALUES (38, '2021-05-10 03:38:38', '2021-05-10 03:38:38', NULL, '', 'edit', 4, 0, 0, NULL, '', '编辑账号', '', 2, 5, 1, '/system/user/edit', 1, 0, 3, 3);
INSERT INTO `sys_menu` VALUES (40, '2021-05-10 03:38:48', '2021-05-10 03:38:48', NULL, '', 'del', 6, 0, 0, NULL, '', '删除账号', '', 2, 5, 1, '/system/user/del', 1, 0, 3, 4);
INSERT INTO `sys_menu` VALUES (42, '2021-05-11 22:38:46', '2021-05-11 22:38:46', NULL, '', 'log', 7, 0, 0, 'ant-design:read-filled', '', '日志管理', '/system/log/index', 1, 3, 1, '/system/log', 1, 0, 2, 1);
INSERT INTO `sys_menu` VALUES (51, '2021-05-11 21:01:46', '2021-05-11 21:01:46', NULL, '', 'add', 2, 0, 0, NULL, '', '新增菜单', '', 2, 8, 1, '/system/menu/add', 1, 0, 3, 2);
INSERT INTO `sys_menu` VALUES (52, '2021-05-10 15:05:41', '2021-05-10 15:05:41', NULL, '', 'add', 3, 0, 0, NULL, '', '新增部门', '', 2, 35, 1, '/system/dept/add', 1, 0, 3, 2);
INSERT INTO `sys_menu` VALUES (53, '2021-05-10 15:33:55', '2021-05-10 15:33:55', NULL, '', 'add', 2, 0, 0, NULL, '', '新增角色', '', 2, 4, 1, '/system/role/add', 1, 0, 3, 2);
INSERT INTO `sys_menu` VALUES (54, '2021-05-11 21:01:30', '2021-05-11 21:01:30', NULL, '', 'add', 2, 0, 0, NULL, '', '新增账号', '', 2, 5, 1, '/system/user/add', 1, 0, 3, 2);
INSERT INTO `sys_menu` VALUES (55, '2021-05-11 20:52:38', '2021-05-11 20:52:38', NULL, '', 'del', 5, 0, 0, NULL, '', '删除部门', '', 2, 35, 1, '/system/dept/del', 1, 0, 3, 4);
INSERT INTO `sys_menu` VALUES (59, '2021-05-11 20:59:48', '2021-05-11 20:59:48', NULL, '', 'del', 2, 0, 0, NULL, '', '删除日志', '', 2, 42, 1, '/system/log/del', 1, 0, 3, 4);
INSERT INTO `sys_menu` VALUES (60, '2021-05-11 19:42:49', '2021-05-11 19:42:49', NULL, '', 'del_batch', 2, 0, 0, NULL, '', '批量删除日志', '', 2, 42, 1, '/system/log/del_batch', 1, 0, 3, 4);
INSERT INTO `sys_menu` VALUES (61, '2021-05-11 20:57:31', '2021-05-11 20:57:31', NULL, '', '', 1, 0, 0, NULL, '', '部门列表', '', 2, 35, 1, '/system/dept/list', 1, 0, 3, 1);
INSERT INTO `sys_menu` VALUES (62, '2021-05-11 20:58:53', '2021-05-11 20:58:53', NULL, '', '', 1, 0, 0, NULL, '', '角色列表', '', 2, 4, 1, '/system/role/list', 1, 0, 3, 1);
INSERT INTO `sys_menu` VALUES (63, '2021-05-11 20:59:33', '2021-05-11 20:59:33', NULL, '', '', 1, 0, 0, NULL, '', '日志列表', '', 2, 42, 1, '/system/log/list', 1, 0, 3, 1);
INSERT INTO `sys_menu` VALUES (64, '2021-05-11 21:00:41', '2021-05-11 21:00:41', NULL, '', '', 1, 0, 0, NULL, '', '账号列表', '', 2, 5, 1, '/system/user/list', 1, 0, 3, 1);
INSERT INTO `sys_menu` VALUES (65, '2021-05-11 21:01:12', '2021-05-11 21:01:12', NULL, '', '', 1, 0, 0, NULL, '', '菜单列表', '', 2, 8, 1, '/system/menu/list', 1, 0, 3, 1);

-- ----------------------------
-- Table structure for sys_op_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_op_log`;
CREATE TABLE `sys_op_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `business_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `request_method` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `operator_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `oper_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `oper_location` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `oper_param` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` int(1) NULL DEFAULT NULL,
  `oper_time` datetime NULL DEFAULT NULL,
  `json_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `update_by` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `browser` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `os` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `platform` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `latency_time` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_op_log_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_op_log
-- ----------------------------
INSERT INTO `sys_op_log` VALUES (1, '用户登录', '10', 'POST', '', 'chris', '', '/api/login', '127.0.0.1', '内部IP', '', 1, '2021-05-12 01:22:43', '', 'chris', '', 'Chrome 89.0.4389.82', 'Android 6.0', 'Linux', '登录成功', '1ns', '2021-05-12 01:22:43', NULL);
INSERT INTO `sys_op_log` VALUES (2, '用户登录', '10', 'POST', '', 'chris', '', '/api/login', '127.0.0.1', '内部IP', '', 1, '2021-05-12 01:22:45', '', 'chris', '', 'Chrome 89.0.4389.82', 'Android 6.0', 'Linux', '登录成功', '1ns', '2021-05-12 01:22:45', NULL);
INSERT INTO `sys_op_log` VALUES (3, '用户登录', '10', 'POST', '', 'chris', '', '/api/login', '127.0.0.1', '内部IP', '', 1, '2021-05-12 01:22:47', '', 'chris', '', 'Chrome 89.0.4389.82', 'Android 6.0', 'Linux', '登录成功', '1ns', '2021-05-12 01:22:47', NULL);
INSERT INTO `sys_op_log` VALUES (4, '用户登录', '10', 'POST', '', 'chris', '', '/api/login', '127.0.0.1', '内部IP', '', 1, '2021-05-12 01:22:49', '', 'chris', '', 'Chrome 89.0.4389.82', 'Android 6.0', 'Linux', '登录成功', '1ns', '2021-05-12 01:22:49', NULL);
INSERT INTO `sys_op_log` VALUES (5, '用户登录', '10', 'POST', '', 'chris', '', '/api/login', '127.0.0.1', '内部IP', '', 1, '2021-05-12 01:24:40', '', 'chris', '', 'Chrome 89.0.4389.82', 'Android 6.0', 'Linux', '登录成功', '1ns', '2021-05-12 01:24:40', NULL);
INSERT INTO `sys_op_log` VALUES (6, '', '3', 'PUT', '', 'chris', '', '/api/system/role/edit', '127.0.0.1', '内部IP', '{\"id\":\"49\",\"roleName\":\"测试账号\",\"roleValue\":\"test\",\"status\":\"1\",\"remark\":\"\",\"menu\":[\"1\",\"2\",\"61\",\"62\",\"63\",\"64\",\"65\"]}', 1, '2021-05-12 01:25:10', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '47.0441ms', '2021-05-12 01:25:10', NULL);
INSERT INTO `sys_op_log` VALUES (7, '', '2', 'POST', '', 'chris', '', '/api/system/role/add', '127.0.0.1', '内部IP', '{\"roleName\":\"超级管理员\",\"roleValue\":\"super\",\"status\":\"1\",\"menu\":[\"1\",\"2\",\"3\",\"35\",\"61\",\"37\",\"52\",\"55\",\"4\",\"62\",\"53\",\"11\",\"10\",\"5\",\"64\",\"54\",\"38\",\"40\",\"42\",\"63\",\"60\",\"59\",\"8\",\"65\",\"51\",\"26\",\"32\",\"36\"]}', 1, '2021-05-12 01:25:32', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '2.9973ms', '2021-05-12 01:25:32', NULL);
INSERT INTO `sys_op_log` VALUES (8, '', '2', 'POST', '', 'chris', '', '/api/system/role/add', '127.0.0.1', '内部IP', '{\"roleName\":\"超级管理员\",\"roleValue\":\"super\",\"status\":\"1\",\"menu\":[\"1\",\"2\",\"3\",\"35\",\"61\",\"37\",\"52\",\"55\",\"4\",\"62\",\"53\",\"11\",\"10\",\"5\",\"64\",\"54\",\"38\",\"40\",\"42\",\"63\",\"60\",\"59\",\"8\",\"65\",\"51\",\"26\",\"32\",\"36\"]}', 1, '2021-05-12 01:25:37', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '1.9992ms', '2021-05-12 01:25:37', NULL);
INSERT INTO `sys_op_log` VALUES (9, '', '2', 'POST', '', 'chris', '', '/api/system/role/add', '127.0.0.1', '内部IP', '{\"roleName\":\"超级管理员\",\"roleValue\":\"super\",\"status\":\"1\",\"menu\":[\"1\",\"2\",\"3\",\"35\",\"61\",\"37\",\"52\",\"55\",\"4\",\"62\",\"53\",\"11\",\"10\",\"5\",\"64\",\"54\",\"38\",\"40\",\"42\",\"63\",\"60\",\"59\",\"8\",\"65\",\"51\",\"26\",\"32\",\"36\"]}', 1, '2021-05-12 01:25:40', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '3.0023ms', '2021-05-12 01:25:40', NULL);
INSERT INTO `sys_op_log` VALUES (10, '', '2', 'POST', '', 'chris', '', '/api/system/role/add', '127.0.0.1', '内部IP', '{\"roleName\":\"超级管理员\",\"roleValue\":\"super\",\"status\":\"1\",\"menu\":[\"1\",\"2\",\"3\",\"35\",\"61\",\"37\",\"52\",\"55\",\"4\",\"62\",\"53\",\"11\",\"10\",\"5\",\"64\",\"54\",\"38\",\"40\",\"42\",\"63\",\"60\",\"59\",\"8\",\"65\",\"51\",\"26\",\"32\",\"36\"]}', 1, '2021-05-12 01:25:44', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '1.9995ms', '2021-05-12 01:25:44', NULL);
INSERT INTO `sys_op_log` VALUES (11, '', '2', 'POST', '', 'chris', '', '/api/system/role/add', '127.0.0.1', '内部IP', '{\"roleName\":\"超级管理员\",\"roleValue\":\"super\",\"status\":\"1\",\"menu\":[\"1\",\"2\",\"3\",\"35\",\"61\",\"37\",\"52\",\"55\",\"4\",\"62\",\"53\",\"11\",\"10\",\"5\",\"64\",\"54\",\"38\",\"40\",\"42\",\"63\",\"60\",\"59\",\"8\",\"65\",\"51\",\"26\",\"32\",\"36\"]}', 1, '2021-05-12 01:25:48', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '1.9986ms', '2021-05-12 01:25:48', NULL);
INSERT INTO `sys_op_log` VALUES (12, '', '2', 'POST', '', 'chris', '', '/api/system/role/add', '127.0.0.1', '内部IP', '{\"roleName\":\"超级管理员\",\"roleValue\":\"super\",\"status\":\"1\",\"menu\":[\"1\",\"2\",\"3\",\"35\",\"61\",\"37\",\"52\",\"55\",\"4\",\"62\",\"53\",\"11\",\"10\",\"5\",\"64\",\"54\",\"38\",\"40\",\"42\",\"63\",\"60\",\"59\",\"8\",\"65\",\"51\",\"26\",\"32\",\"36\"]}', 1, '2021-05-12 01:26:48', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '121.1005ms', '2021-05-12 01:26:48', NULL);
INSERT INTO `sys_op_log` VALUES (13, '', '2', 'POST', '', 'chris', '', '/api/system/role/add', '127.0.0.1', '内部IP', '{\"roleName\":\"测试账号\",\"roleValue\":\"test\",\"status\":\"1\",\"menu\":[\"36\",\"1\",\"2\",\"61\",\"62\",\"64\",\"63\",\"65\"]}', 1, '2021-05-12 01:27:52', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '49.9707ms', '2021-05-12 01:27:52', NULL);
INSERT INTO `sys_op_log` VALUES (14, '', '3', 'PUT', '', 'chris', '', '/api/system/user/edit', '127.0.0.1', '内部IP', '{\"id\":\"1\",\"username\":\"chris\",\"role_id\":\"1\",\"roleName\":\"55\",\"deptName\":\"1\",\"phone\":\"123123\",\"email\":\"2501170033@qq.com\",\"remark\":\"\"}', 1, '2021-05-12 01:28:19', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '9.487ms', '2021-05-12 01:28:19', NULL);
INSERT INTO `sys_op_log` VALUES (15, '', '2', 'POST', '', 'chris', '', '/api/system/user/add', '127.0.0.1', '内部IP', '{\"username\":\"guest\",\"pwd\":\"123456\",\"roleName\":\"56\",\"deptName\":\"1\",\"phone\":\"23111\",\"email\":\"2501170033@qq.com\"}', 1, '2021-05-12 01:28:48', '', 'chris', '', 'Chrome 89.0.4389.82', 'Windows 10', 'Windows', 'Operation log', '5.0156ms', '2021-05-12 01:28:48', NULL);
INSERT INTO `sys_op_log` VALUES (16, '', '3', 'PUT', '', 'chris', '', '/api/system/dept/edit', '127.0.0.1', '内部IP', '{\"id\":\"7\",\"pid\":\"0\",\"deptName\":\"华南分部\",\"parentDept\":\"顶级目录\",\"orderNo\":2,\"status\":\"1\",\"remark\":\"\"}', 1, '2021-05-12 01:41:05', '', 'chris', '', 'Chrome 89.0.4389.82', 'Android 6.0', 'Linux', 'Operation log', '3.9975ms', '2021-05-12 01:41:05', NULL);
INSERT INTO `sys_op_log` VALUES (17, '', '3', 'PUT', '', 'chris', '', '/api/system/dept/edit', '127.0.0.1', '内部IP', '{\"id\":\"7\",\"pid\":\"0\",\"deptName\":\"华南分部\",\"parentDept\":\"顶级目录\",\"orderNo\":3,\"status\":\"1\",\"remark\":\"\"}', 1, '2021-05-12 01:41:12', '', 'chris', '', 'Chrome 89.0.4389.82', 'Android 6.0', 'Linux', 'Operation log', '4.012ms', '2021-05-12 01:41:12', NULL);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `role_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `order_no` int(11) NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `uix_sys_role_role_name`(`role_name`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 57 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (55, '超级管理员', 'super', '', '1', 0, '2021-05-12 01:26:47', '2021-05-12 01:26:47', NULL);
INSERT INTO `sys_role` VALUES (56, '测试账号', 'test', '', '1', 0, '2021-05-12 01:27:52', '2021-05-12 01:27:52', NULL);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `sys_role_id` int(11) NOT NULL,
  `sys_menu_id` int(11) NOT NULL,
  PRIMARY KEY (`sys_role_id`, `sys_menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (1, 1);
INSERT INTO `sys_role_menu` VALUES (1, 2);
INSERT INTO `sys_role_menu` VALUES (1, 3);
INSERT INTO `sys_role_menu` VALUES (1, 4);
INSERT INTO `sys_role_menu` VALUES (1, 5);
INSERT INTO `sys_role_menu` VALUES (1, 8);
INSERT INTO `sys_role_menu` VALUES (1, 10);
INSERT INTO `sys_role_menu` VALUES (1, 11);
INSERT INTO `sys_role_menu` VALUES (1, 26);
INSERT INTO `sys_role_menu` VALUES (1, 32);
INSERT INTO `sys_role_menu` VALUES (1, 35);
INSERT INTO `sys_role_menu` VALUES (1, 36);
INSERT INTO `sys_role_menu` VALUES (1, 37);
INSERT INTO `sys_role_menu` VALUES (1, 38);
INSERT INTO `sys_role_menu` VALUES (1, 40);
INSERT INTO `sys_role_menu` VALUES (1, 42);
INSERT INTO `sys_role_menu` VALUES (1, 51);
INSERT INTO `sys_role_menu` VALUES (1, 52);
INSERT INTO `sys_role_menu` VALUES (1, 53);
INSERT INTO `sys_role_menu` VALUES (1, 54);
INSERT INTO `sys_role_menu` VALUES (1, 55);
INSERT INTO `sys_role_menu` VALUES (1, 59);
INSERT INTO `sys_role_menu` VALUES (1, 60);
INSERT INTO `sys_role_menu` VALUES (1, 61);
INSERT INTO `sys_role_menu` VALUES (1, 62);
INSERT INTO `sys_role_menu` VALUES (1, 63);
INSERT INTO `sys_role_menu` VALUES (1, 64);
INSERT INTO `sys_role_menu` VALUES (1, 65);
INSERT INTO `sys_role_menu` VALUES (1, 66);
INSERT INTO `sys_role_menu` VALUES (49, 1);
INSERT INTO `sys_role_menu` VALUES (49, 2);
INSERT INTO `sys_role_menu` VALUES (49, 61);
INSERT INTO `sys_role_menu` VALUES (49, 62);
INSERT INTO `sys_role_menu` VALUES (49, 63);
INSERT INTO `sys_role_menu` VALUES (49, 64);
INSERT INTO `sys_role_menu` VALUES (49, 65);
INSERT INTO `sys_role_menu` VALUES (55, 1);
INSERT INTO `sys_role_menu` VALUES (55, 2);
INSERT INTO `sys_role_menu` VALUES (55, 3);
INSERT INTO `sys_role_menu` VALUES (55, 4);
INSERT INTO `sys_role_menu` VALUES (55, 5);
INSERT INTO `sys_role_menu` VALUES (55, 8);
INSERT INTO `sys_role_menu` VALUES (55, 10);
INSERT INTO `sys_role_menu` VALUES (55, 11);
INSERT INTO `sys_role_menu` VALUES (55, 26);
INSERT INTO `sys_role_menu` VALUES (55, 32);
INSERT INTO `sys_role_menu` VALUES (55, 35);
INSERT INTO `sys_role_menu` VALUES (55, 36);
INSERT INTO `sys_role_menu` VALUES (55, 37);
INSERT INTO `sys_role_menu` VALUES (55, 38);
INSERT INTO `sys_role_menu` VALUES (55, 40);
INSERT INTO `sys_role_menu` VALUES (55, 42);
INSERT INTO `sys_role_menu` VALUES (55, 51);
INSERT INTO `sys_role_menu` VALUES (55, 52);
INSERT INTO `sys_role_menu` VALUES (55, 53);
INSERT INTO `sys_role_menu` VALUES (55, 54);
INSERT INTO `sys_role_menu` VALUES (55, 55);
INSERT INTO `sys_role_menu` VALUES (55, 59);
INSERT INTO `sys_role_menu` VALUES (55, 60);
INSERT INTO `sys_role_menu` VALUES (55, 61);
INSERT INTO `sys_role_menu` VALUES (55, 62);
INSERT INTO `sys_role_menu` VALUES (55, 63);
INSERT INTO `sys_role_menu` VALUES (55, 64);
INSERT INTO `sys_role_menu` VALUES (55, 65);
INSERT INTO `sys_role_menu` VALUES (56, 1);
INSERT INTO `sys_role_menu` VALUES (56, 2);
INSERT INTO `sys_role_menu` VALUES (56, 36);
INSERT INTO `sys_role_menu` VALUES (56, 61);
INSERT INTO `sys_role_menu` VALUES (56, 62);
INSERT INTO `sys_role_menu` VALUES (56, 63);
INSERT INTO `sys_role_menu` VALUES (56, 64);
INSERT INTO `sys_role_menu` VALUES (56, 65);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'static/upload/avatar/default.png',
  `role_id` int(11) NULL DEFAULT NULL,
  `status` int(11) NULL DEFAULT 1,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `dept` int(11) NULL DEFAULT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uix_sys_user_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'chris', NULL, '7edf620a4a49f509b33e8472e8ca71ca', 'static/upload/avatar/default.png', 55, 1, '2021-05-12 01:28:19', '2021-05-12 01:28:19', 1, '123123', '2501170033@qq.com', NULL);
INSERT INTO `sys_user` VALUES (2, 'guest', '', '71de42a0f9bf70b948f6627d61f92801', 'static/upload/avatar/default.png', 56, 1, '2021-05-12 01:28:48', '2021-05-12 01:28:48', 1, '23111', '2501170033@qq.com', '');

SET FOREIGN_KEY_CHECKS = 1;
