-- Go Web Admin 数据库表结构 (MySQL)
-- 创建数据库
CREATE DATABASE IF NOT EXISTS `go_web_admin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `go_web_admin`;

-- 用户表
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` varchar(50) NOT NULL COMMENT '用户ID',
  `username` varchar(50) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0-禁用，1-启用',
  `last_login` datetime(3) DEFAULT NULL COMMENT '最后登录时间',
  `login_count` int NOT NULL DEFAULT '0' COMMENT '登录次数',
  `remark` text COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_user_id` (`user_id`),
  UNIQUE KEY `idx_users_username` (`username`),
  UNIQUE KEY `idx_users_email` (`email`),
  UNIQUE KEY `idx_users_phone` (`phone`),
  KEY `idx_users_deleted_at` (`deleted_at`),
  KEY `idx_users_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 角色表
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `code` varchar(50) NOT NULL COMMENT '角色编码',
  `description` text COMMENT '角色描述',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0-禁用，1-启用',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_roles_name` (`name`),
  UNIQUE KEY `idx_roles_code` (`code`),
  KEY `idx_roles_deleted_at` (`deleted_at`),
  KEY `idx_roles_status` (`status`),
  KEY `idx_roles_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 权限表
CREATE TABLE `permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) NOT NULL COMMENT '权限名称',
  `code` varchar(50) NOT NULL COMMENT '权限编码',
  `type` varchar(20) NOT NULL COMMENT '权限类型：menu-菜单，button-按钮，api-接口',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT '父级ID',
  `path` varchar(255) DEFAULT NULL COMMENT '路由路径',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `icon` varchar(50) DEFAULT NULL COMMENT '图标',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0-禁用，1-启用',
  `description` text COMMENT '权限描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_permissions_name` (`name`),
  UNIQUE KEY `idx_permissions_code` (`code`),
  KEY `idx_permissions_deleted_at` (`deleted_at`),
  KEY `idx_permissions_parent_id` (`parent_id`),
  KEY `idx_permissions_type` (`type`),
  KEY `idx_permissions_status` (`status`),
  KEY `idx_permissions_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- 用户角色关联表
CREATE TABLE `user_roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_roles_user_role` (`user_id`,`role_id`),
  KEY `idx_user_roles_deleted_at` (`deleted_at`),
  KEY `idx_user_roles_user_id` (`user_id`),
  KEY `idx_user_roles_role_id` (`role_id`),
  CONSTRAINT `fk_user_roles_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_user_roles_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';

-- 角色权限关联表
CREATE TABLE `role_permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `permission_id` bigint unsigned NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_role_permissions_role_permission` (`role_id`,`permission_id`),
  KEY `idx_role_permissions_deleted_at` (`deleted_at`),
  KEY `idx_role_permissions_role_id` (`role_id`),
  KEY `idx_role_permissions_permission_id` (`permission_id`),
  CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_role_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

-- 菜单表
CREATE TABLE `menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) NOT NULL COMMENT '菜单名称',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT '父级菜单ID',
  `path` varchar(255) DEFAULT NULL COMMENT '路由路径',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `icon` varchar(50) DEFAULT NULL COMMENT '图标',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0-隐藏，1-显示',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型：1-菜单，2-按钮',
  `remark` text COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_menus_deleted_at` (`deleted_at`),
  KEY `idx_menus_parent_id` (`parent_id`),
  KEY `idx_menus_status` (`status`),
  KEY `idx_menus_type` (`type`),
  KEY `idx_menus_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单表';

-- 插入初始数据
-- 插入默认管理员用户（密码：admin123）
INSERT INTO `users` (`user_id`, `username`, `password`, `status`, `created_at`, `updated_at`) VALUES
('admin', 'admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 1, NOW(), NOW());

-- 插入默认角色
INSERT INTO `roles` (`name`, `code`, `description`, `status`, `sort`, `created_at`, `updated_at`) VALUES
('超级管理员', 'super_admin', '系统超级管理员，拥有所有权限', 1, 1, NOW(), NOW()),
('管理员', 'admin', '系统管理员', 1, 2, NOW(), NOW()),
('普通用户', 'user', '普通用户', 1, 3, NOW(), NOW());

-- 插入默认权限
INSERT INTO `permissions` (`name`, `code`, `type`, `parent_id`, `path`, `component`, `icon`, `sort`, `status`, `description`, `created_at`, `updated_at`) VALUES
('系统管理', 'system', 'menu', 0, '/system', '', 'system', 1, 1, '系统管理菜单', NOW(), NOW()),
('用户管理', 'system:user', 'menu', 1, '/system/user', 'system/user/index', 'user', 1, 1, '用户管理', NOW(), NOW()),
('角色管理', 'system:role', 'menu', 1, '/system/role', 'system/role/index', 'role', 2, 1, '角色管理', NOW(), NOW()),
('权限管理', 'system:permission', 'menu', 1, '/system/permission', 'system/permission/index', 'permission', 3, 1, '权限管理', NOW(), NOW()),
('菜单管理', 'system:menu', 'menu', 1, '/system/menu', 'system/menu/index', 'menu', 4, 1, '菜单管理', NOW(), NOW());

-- 插入默认菜单
INSERT INTO `menus` (`name`, `parent_id`, `path`, `component`, `icon`, `sort`, `status`, `type`, `remark`, `created_at`, `updated_at`) VALUES
('系统管理', 0, '/system', '', 'system', 1, 1, 1, '系统管理菜单', NOW(), NOW()),
('用户管理', 1, '/system/user', 'system/user/index', 'user', 1, 1, 1, '用户管理', NOW(), NOW()),
('角色管理', 1, '/system/role', 'system/role/index', 'role', 2, 1, 1, '角色管理', NOW(), NOW()),
('权限管理', 1, '/system/permission', 'system/permission/index', 'permission', 3, 1, 1, '权限管理', NOW(), NOW()),
('菜单管理', 1, '/system/menu', 'system/menu/index', 'menu', 4, 1, 1, '菜单管理', NOW(), NOW());

-- 为超级管理员分配角色
INSERT INTO `user_roles` (`user_id`, `role_id`, `created_at`, `updated_at`) VALUES
(1, 1, NOW(), NOW());

-- 为超级管理员角色分配所有权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`, `created_at`, `updated_at`) VALUES
(1, 1, NOW(), NOW()),
(1, 2, NOW(), NOW()),
(1, 3, NOW(), NOW()),
(1, 4, NOW(), NOW()),
(1, 5, NOW(), NOW());
