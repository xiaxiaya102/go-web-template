-- Go Web Admin 数据库表结构 (PostgreSQL)
-- 创建数据库
-- CREATE DATABASE go_web_admin;
-- \c go_web_admin;

-- 用户表
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(3) WITH TIME ZONE,
    updated_at TIMESTAMP(3) WITH TIME ZONE,
    deleted_at TIMESTAMP(3) WITH TIME ZONE,
    user_id VARCHAR(50) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(20) UNIQUE,
    avatar VARCHAR(255),
    status SMALLINT NOT NULL DEFAULT 1,
    last_login TIMESTAMP(3) WITH TIME ZONE,
    login_count INTEGER NOT NULL DEFAULT 0,
    remark TEXT
);

-- 创建索引
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
CREATE INDEX idx_users_status ON users(status);

-- 角色表
CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(3) WITH TIME ZONE,
    updated_at TIMESTAMP(3) WITH TIME ZONE,
    deleted_at TIMESTAMP(3) WITH TIME ZONE,
    name VARCHAR(50) NOT NULL UNIQUE,
    code VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    status SMALLINT NOT NULL DEFAULT 1,
    sort INTEGER NOT NULL DEFAULT 0
);

-- 创建索引
CREATE INDEX idx_roles_deleted_at ON roles(deleted_at);
CREATE INDEX idx_roles_status ON roles(status);
CREATE INDEX idx_roles_sort ON roles(sort);

-- 权限表
CREATE TABLE permissions (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(3) WITH TIME ZONE,
    updated_at TIMESTAMP(3) WITH TIME ZONE,
    deleted_at TIMESTAMP(3) WITH TIME ZONE,
    name VARCHAR(50) NOT NULL UNIQUE,
    code VARCHAR(50) NOT NULL UNIQUE,
    type VARCHAR(20) NOT NULL,
    parent_id BIGINT DEFAULT 0,
    path VARCHAR(255),
    component VARCHAR(255),
    icon VARCHAR(50),
    sort INTEGER NOT NULL DEFAULT 0,
    status SMALLINT NOT NULL DEFAULT 1,
    description TEXT
);

-- 创建索引
CREATE INDEX idx_permissions_deleted_at ON permissions(deleted_at);
CREATE INDEX idx_permissions_parent_id ON permissions(parent_id);
CREATE INDEX idx_permissions_type ON permissions(type);
CREATE INDEX idx_permissions_status ON permissions(status);
CREATE INDEX idx_permissions_sort ON permissions(sort);

-- 用户角色关联表
CREATE TABLE user_roles (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(3) WITH TIME ZONE,
    updated_at TIMESTAMP(3) WITH TIME ZONE,
    deleted_at TIMESTAMP(3) WITH TIME ZONE,
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    UNIQUE(user_id, role_id)
);

-- 创建索引和外键
CREATE INDEX idx_user_roles_deleted_at ON user_roles(deleted_at);
CREATE INDEX idx_user_roles_user_id ON user_roles(user_id);
CREATE INDEX idx_user_roles_role_id ON user_roles(role_id);
ALTER TABLE user_roles ADD CONSTRAINT fk_user_roles_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE user_roles ADD CONSTRAINT fk_user_roles_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE;

-- 角色权限关联表
CREATE TABLE role_permissions (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(3) WITH TIME ZONE,
    updated_at TIMESTAMP(3) WITH TIME ZONE,
    deleted_at TIMESTAMP(3) WITH TIME ZONE,
    role_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    UNIQUE(role_id, permission_id)
);

-- 创建索引和外键
CREATE INDEX idx_role_permissions_deleted_at ON role_permissions(deleted_at);
CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX idx_role_permissions_permission_id ON role_permissions(permission_id);
ALTER TABLE role_permissions ADD CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE;
ALTER TABLE role_permissions ADD CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE;

-- 菜单表
CREATE TABLE menus (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(3) WITH TIME ZONE,
    updated_at TIMESTAMP(3) WITH TIME ZONE,
    deleted_at TIMESTAMP(3) WITH TIME ZONE,
    name VARCHAR(50) NOT NULL,
    parent_id BIGINT DEFAULT 0,
    path VARCHAR(255),
    component VARCHAR(255),
    icon VARCHAR(50),
    sort INTEGER NOT NULL DEFAULT 0,
    status SMALLINT NOT NULL DEFAULT 1,
    type SMALLINT NOT NULL DEFAULT 1,
    remark TEXT
);

-- 创建索引
CREATE INDEX idx_menus_deleted_at ON menus(deleted_at);
CREATE INDEX idx_menus_parent_id ON menus(parent_id);
CREATE INDEX idx_menus_status ON menus(status);
CREATE INDEX idx_menus_type ON menus(type);
CREATE INDEX idx_menus_sort ON menus(sort);

-- 插入初始数据
-- 插入默认管理员用户（密码：admin123）
INSERT INTO users (user_id, username, password, status, created_at, updated_at) VALUES
('admin', 'admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 1, NOW(), NOW());

-- 插入默认角色
INSERT INTO roles (name, code, description, status, sort, created_at, updated_at) VALUES
('超级管理员', 'super_admin', '系统超级管理员，拥有所有权限', 1, 1, NOW(), NOW()),
('管理员', 'admin', '系统管理员', 1, 2, NOW(), NOW()),
('普通用户', 'user', '普通用户', 1, 3, NOW(), NOW());

-- 插入默认权限
INSERT INTO permissions (name, code, type, parent_id, path, component, icon, sort, status, description, created_at, updated_at) VALUES
('系统管理', 'system', 'menu', 0, '/system', '', 'system', 1, 1, '系统管理菜单', NOW(), NOW()),
('用户管理', 'system:user', 'menu', 1, '/system/user', 'system/user/index', 'user', 1, 1, '用户管理', NOW(), NOW()),
('角色管理', 'system:role', 'menu', 1, '/system/role', 'system/role/index', 'role', 2, 1, '角色管理', NOW(), NOW()),
('权限管理', 'system:permission', 'menu', 1, '/system/permission', 'system/permission/index', 'permission', 3, 1, '权限管理', NOW(), NOW()),
('菜单管理', 'system:menu', 'menu', 1, '/system/menu', 'system/menu/index', 'menu', 4, 1, '菜单管理', NOW(), NOW());

-- 插入默认菜单
INSERT INTO menus (name, parent_id, path, component, icon, sort, status, type, remark, created_at, updated_at) VALUES
('系统管理', 0, '/system', '', 'system', 1, 1, 1, '系统管理菜单', NOW(), NOW()),
('用户管理', 1, '/system/user', 'system/user/index', 'user', 1, 1, 1, '用户管理', NOW(), NOW()),
('角色管理', 1, '/system/role', 'system/role/index', 'role', 2, 1, 1, '角色管理', NOW(), NOW()),
('权限管理', 1, '/system/permission', 'system/permission/index', 'permission', 3, 1, 1, '权限管理', NOW(), NOW()),
('菜单管理', 1, '/system/menu', 'system/menu/index', 'menu', 4, 1, 1, '菜单管理', NOW(), NOW());

-- 为超级管理员分配角色
INSERT INTO user_roles (user_id, role_id, created_at, updated_at) VALUES
(1, 1, NOW(), NOW());

-- 为超级管理员角色分配所有权限
INSERT INTO role_permissions (role_id, permission_id, created_at, updated_at) VALUES
(1, 1, NOW(), NOW()),
(1, 2, NOW(), NOW()),
(1, 3, NOW(), NOW()),
(1, 4, NOW(), NOW()),
(1, 5, NOW(), NOW());
