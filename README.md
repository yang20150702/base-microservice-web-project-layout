# 开发文档

基于微服务架构思想来设计web项目的代码结构，对外提供restful-api和graphql接口

base-microservice-web-project-layout

组件：gin、gorm、graphql-go

## API管理

API归属于各个业务方。
+ 不同业务的API必须分组
+ 通过init函数显式注册到GlobalEngine（全局路由引擎）
+ 中间件要统一管理，采用标准化组件



## 开发建议

1. git commit 提交规范化
2. git flow工作流程，保证产品开发、迭代以及发布流程清晰明了
3. 充分利用golang的作用域以及短变量命名
4. 通过golint、等工具对golang的语法错误进行检测以及代码格式化
5. 单元测试的粒度要足够小，不应该有外部依赖。如果有条件，单元测试统一接入gitlab-ci
6. 代码合并统一走merge request。并进行代码review

