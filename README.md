# NFLY

通用被动消息接口框架服务端

# 分部件概览

- [x] 数据库互动
- [ ] 各部件缓存
- [ ] 用户管理
- [ ] 通知接口
- [ ] 传输协议
- [ ] Worker 池
- [ ] 异步调度
- [ ] 推送健康度检查
- [ ] 推送测试检查
- [ ] 通用插件接口
- [x] 外部 API 接口
- [ ] 安全过滤
- [x] CLI 接口
- [ ] GUI 接口

## 用户管理

- [x] 用户组
- [ ] 用户权限
- [x] 用户注册 (密码使用哈希算法存储)
- [x] 用户登录
- [ ] 批量管理
- [ ] 用户限制
- [ ] 用户屏蔽
- [ ] 客户端绑定

## 通知接口

被动式数据接口

- [x] 通用型
- [ ] 特定型

## 传输协议

- [x] HTTP
- [ ] TCP-based 自研协议

## Session 

session 在使用 API 的登录接口时由服务端生成返回。  
每次 API 请求动作必须含有 session 字段(uuid) 以识别用户。   

## Token 更新算法流程

1. 登录API
2. 服务端返回 token (根据通知关联)
3. 服务端生成token 定义session过期period
4. 返回token

## Worker

Worker 作为自动任务执行器，具有状态、无状态，有创建或关闭通知的权限。使用 Worker 需要调用插件，并配置相关参数。

## Worker 池

作为 Worker 的容器执行池，队列可配置为有序或无序。

## 异步调度

根据特殊算法以获取相关调度值，默认使用 Worker 健康度与插件优先级。 一般情况只批量执行 Worker Pool，也可单独执行某 Pool 中的 Worker。

## 推送健康度检查

检测通知的可连通性，稳定性，客户端回应延迟等，并根据相关算法得出健康度。

## 通用插件实现

通过重写插件类方法实现。

### 插件结构定义字段

- 名称
- 描述
- 版本号
- 维护者
- 网站
- OTA地址
- 能力标签
- 优先级

### 插件结构定义重写方法

- 注册插件
- 初始化
- 互通 channel
- 作用域
- 读取配置
- 插件调用前
- 插件调用
- 请求更新推送 <=> ↑↓
- 插件调用后
- 插件权限申请
- 插件终止前

### 流程

1. 导入插件到指定位置
2. 插件绑定到 Worker 
3. Worker 运行
4. 调用插件
5. 插件推送消息
6. 计划任务
7. 停止

## 推送测试

发送 "ping pong"，测试与被动链接的客户端是否通信正常。

## safeguards

- 蜜罐
- 抗重放 (时间戳)
- 验证token
- 写数据库
- 屏蔽

## cli接口

- [x] 服务端控制

## gui接口

- 可视化服务端控制
- 被动客户端
