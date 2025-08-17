-- 修复MySQL数据库表结构的SQL脚本
-- 如果表已存在，需要先修改字段类型以避免GORM迁移错误

-- 修复permissions表的type字段
ALTER TABLE `permissions` MODIFY COLUMN `type` varchar(20) NOT NULL COMMENT '权限类型：menu-菜单，button-按钮，api-接口';
ALTER TABLE `permissions` MODIFY COLUMN `name` varchar(50) NOT NULL COMMENT '权限名称';
ALTER TABLE `permissions` MODIFY COLUMN `code` varchar(50) NOT NULL COMMENT '权限编码';
ALTER TABLE `permissions` MODIFY COLUMN `path` varchar(255) DEFAULT NULL COMMENT '路由路径';
ALTER TABLE `permissions` MODIFY COLUMN `component` varchar(255) DEFAULT NULL COMMENT '组件路径';
ALTER TABLE `permissions` MODIFY COLUMN `icon` varchar(50) DEFAULT NULL COMMENT '图标';
ALTER TABLE `permissions` MODIFY COLUMN `description` text COMMENT '权限描述';

-- 修复users表字段
ALTER TABLE `users` MODIFY COLUMN `user_id` varchar(50) NOT NULL COMMENT '用户ID';
ALTER TABLE `users` MODIFY COLUMN `username` varchar(50) NOT NULL COMMENT '用户名';
ALTER TABLE `users` MODIFY COLUMN `password` varchar(255) NOT NULL COMMENT '密码';
ALTER TABLE `users` MODIFY COLUMN `email` varchar(100) DEFAULT NULL COMMENT '邮箱';
ALTER TABLE `users` MODIFY COLUMN `phone` varchar(20) DEFAULT NULL COMMENT '手机号';
ALTER TABLE `users` MODIFY COLUMN `avatar` varchar(255) DEFAULT NULL COMMENT '头像';
ALTER TABLE `users` MODIFY COLUMN `remark` text COMMENT '备注';

-- 修复roles表字段
ALTER TABLE `roles` MODIFY COLUMN `name` varchar(50) NOT NULL COMMENT '角色名称';
ALTER TABLE `roles` MODIFY COLUMN `code` varchar(50) NOT NULL COMMENT '角色编码';
ALTER TABLE `roles` MODIFY COLUMN `description` text COMMENT '角色描述';

-- 修复menus表字段
ALTER TABLE `menus` MODIFY COLUMN `name` varchar(50) NOT NULL COMMENT '菜单名称';
ALTER TABLE `menus` MODIFY COLUMN `path` varchar(255) DEFAULT NULL COMMENT '路由路径';
ALTER TABLE `menus` MODIFY COLUMN `component` varchar(255) DEFAULT NULL COMMENT '组件路径';
ALTER TABLE `menus` MODIFY COLUMN `icon` varchar(50) DEFAULT NULL COMMENT '图标';
ALTER TABLE `menus` MODIFY COLUMN `remark` text COMMENT '备注';
