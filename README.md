### /user/register
Service 层 RegisterUser 检查数据库中是否已有该用户名。
如果已存在 → 返回错误。
否则对密码进行 bcrypt 哈希，存入数据库。
### /user/login
Service 层 Authenticate 检查用户名和密码是否正确。
如果正确 → 生成一个 JWT token。
把 token 保存到 Redis并设置过期时间
### /user/admin(保护接口)
通过 AuthMiddleware 中间件进行鉴权。
