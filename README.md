# rfidWeb
## 说明
该程序作为插件主要用于服务rfid一键获取及rfid一键录入
## 原理
1. 前端通过fetch不断请求后端（如果有需求可以使用websocket）请求后端返回rfid信息  
2. 如果获取到rfid信息则返回给前端，前端写入rfid到表单或直接提交到微服务addproduct处理中心