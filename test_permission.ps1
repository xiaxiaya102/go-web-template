# 权限系统测试脚本

Write-Host "=== 权限系统测试 ===" -ForegroundColor Green

# 1. 登录获取token
Write-Host "`n1. 登录测试..." -ForegroundColor Yellow
$loginBody = @{
    userName = "admin"
    password = "admin123"
} | ConvertTo-Json

try {
    $loginResponse = Invoke-RestMethod -Uri 'http://localhost:8089/api/login' -Method POST -ContentType 'application/json' -Body $loginBody
    $token = $loginResponse.result.token
    Write-Host "✅ 登录成功" -ForegroundColor Green
    Write-Host "Token: $($token.Substring(0,50))..." -ForegroundColor Gray
} catch {
    Write-Host "❌ 登录失败: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

# 设置认证头
$headers = @{
    Authorization = "Bearer $token"
}

# 2. 获取用户信息
Write-Host "`n2. 获取用户信息..." -ForegroundColor Yellow
try {
    $userInfo = Invoke-RestMethod -Uri 'http://localhost:8089/api/user/info' -Method GET -Headers $headers
    Write-Host "✅ 用户信息获取成功" -ForegroundColor Green
    Write-Host "用户ID: $($userInfo.result.user_id)" -ForegroundColor Gray
    Write-Host "用户名: $($userInfo.result.username)" -ForegroundColor Gray
    Write-Host "是否超级管理员: $($userInfo.result.is_super_admin)" -ForegroundColor Gray
    Write-Host "角色数量: $($userInfo.result.roles.Count)" -ForegroundColor Gray
    Write-Host "权限数量: $($userInfo.result.permissions.Count)" -ForegroundColor Gray
} catch {
    Write-Host "❌ 获取用户信息失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 3. 获取用户权限
Write-Host "`n3. 获取用户权限..." -ForegroundColor Yellow
try {
    $permissions = Invoke-RestMethod -Uri 'http://localhost:8089/api/user/permissions' -Method GET -Headers $headers
    Write-Host "✅ 权限列表获取成功" -ForegroundColor Green
    Write-Host "权限总数: $($permissions.result.Count)" -ForegroundColor Gray
    Write-Host "前5个权限: $($permissions.result[0..4] -join ', ')" -ForegroundColor Gray
} catch {
    Write-Host "❌ 获取权限失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 4. 获取用户角色
Write-Host "`n4. 获取用户角色..." -ForegroundColor Yellow
try {
    $roles = Invoke-RestMethod -Uri 'http://localhost:8089/api/user/roles' -Method GET -Headers $headers
    Write-Host "✅ 角色列表获取成功" -ForegroundColor Green
    Write-Host "角色: $($roles.result -join ', ')" -ForegroundColor Gray
} catch {
    Write-Host "❌ 获取角色失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 5. 测试权限检查
Write-Host "`n5. 测试权限检查..." -ForegroundColor Yellow
$testPermissions = @("user:list", "user:create", "user:delete", "permission:create")

foreach ($permission in $testPermissions) {
    try {
        $checkResult = Invoke-RestMethod -Uri "http://localhost:8089/api/user/check-permission?permission=$permission" -Method GET -Headers $headers
        $hasPermission = $checkResult.result
        $status = if ($hasPermission) { "✅" } else { "❌" }
        Write-Host "$status $permission : $hasPermission" -ForegroundColor $(if ($hasPermission) { "Green" } else { "Red" })
    } catch {
        Write-Host "❌ 检查权限 $permission 失败: $($_.Exception.Message)" -ForegroundColor Red
    }
}

# 6. 测试需要权限的接口
Write-Host "`n6. 测试需要权限的接口..." -ForegroundColor Yellow

# 测试用户列表接口（需要user:list权限）
try {
    $userList = Invoke-RestMethod -Uri 'http://localhost:8089/api/user' -Method GET -Headers $headers
    Write-Host "✅ 用户列表接口访问成功" -ForegroundColor Green
    Write-Host "用户总数: $($userList.result.total)" -ForegroundColor Gray
} catch {
    Write-Host "❌ 用户列表接口访问失败: $($_.Exception.Message)" -ForegroundColor Red
}

# 测试角色列表接口（需要role:list权限）
try {
    $roleList = Invoke-RestMethod -Uri 'http://localhost:8089/api/role' -Method GET -Headers $headers
    Write-Host "✅ 角色列表接口访问成功" -ForegroundColor Green
    Write-Host "角色总数: $($roleList.result.total)" -ForegroundColor Gray
} catch {
    Write-Host "❌ 角色列表接口访问失败: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n=== 测试完成 ===" -ForegroundColor Green
