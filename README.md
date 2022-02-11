# 物联大师

[![Go](https://github.com/zgwit/iot-master/actions/workflows/go.yml/badge.svg)](https://github.com/zgwit/iot-master/actions/workflows/go.yml)

[iot-master.com](https://iot-master.com)

物联大师是[真格智能实验室](https://labs.zgwit.com)推出的一款通用的数据采集和自动控制系统，
集成了Modbus和一些主流PLC的通讯协议，适用于大部分物联网或工业互联网数据应用场景。
该产品提供数据采集，历史存储，自动控制等功能，可以从一定程度上取代PLC或MCU。
物联大师平台提供丰富的元件库和在线模板，可以直接用于大部分物联网项目后端，快速，方便，高效。

> 作者曾经接触多个物联网实际项目的后端，需求大同小异， 因为团队不同，实现方式就千奇百怪了，大家其实都在重复地造轮子。
> 痛定思痛，于是决定提取共同的部分，做成了通用的物联大师， 并且通过开源的方式免费分享给小伙伴儿们使用。

## 给谁用？

1. 物联网企业，比如：智慧养老、智慧小区、智慧农业、智慧养殖、智慧厂房、智慧仓库等
2. 设备制造商，比如：锅炉、液压、锻造、成型、清洗、机床（暂不支持CNC）等
3. 政府单位，比如：智慧交通、环境监控、水利设施、灾害监测、物联网小镇等
4. 其他


## 典型的应用案例

- 项目部署在云服务器，使用数据透传连接设备（支持大部分DTU和移动通讯模块）
- 使用485总线连接标准的Modbus设备，比如：传感器、继电器（开关）
- 配置定时采集，合法检查
- 创建工程，配置定时任务，自动控制，异常告警等
- 开放接口对接APP或小程序，实现远程操控，定时，自动控制，查看历史曲线等
- 使用大数据屏展示实时数据

> 如果以上正是您所需的，请聊聊合作（联系方式在底部）


## 项目架构图

![结构图](https://github.com/zgwit/iot-master/raw/main/docs/frame.svg)

## 技术栈

项目使用Golang进行开发，普通桌面机实测5w并发无压力，云端未实测，主要看带宽。

> PS：项目曾经使用Nodejs开发后端，但是Nodejs的单线程模型，并不适合物联网程序开发，有兴趣可以查看js分支。


| 模块        | 选型    |  说明  |
| --------   | -----   | ---- |
| 后端框架     | gin    | 简单好用，灵活高效   |
| 前端框架     | Angular和ZORRO    |  Angular集成度高，学习成本虽高，但使用方便  |
| 关系数据库   | storm(boltdb)    |  内嵌数据库，可以省去单独部署关系数据库的麻烦，而且存储结构化数据方便  |
| 历史数据库   | tstorage | 内嵌时序数据库 |

## 开发目标

- [x] 数据通道
    - [x] TCP通道，以及注册包和心跳包支持
    - [ ] UDP通道，以及注册包和心跳包支持
    - [ ] 串口通道
- [x] 协议支持
    - [x] Modbus RTU、TCP（ASCII不常用，暂无必要）（**推荐**RTU转TCP的网关，可以加速远程控制）
    - [x] Omron PLC（hostlink, fins）
    - [ ] Mitsubishi PLC (melsec)
    - [ ] Siemens PLC (S7)
- [x] 设备 & 采集 & 控制
    - [x] 定时轮询
    - [ ] 滤波（均值，中值，最大，最小等）
    - [x] 变量映射
    - [x] 控制指令
    - [x] 定时任务
    - [x] 自动控制
    - [x] 存入历史数据库
    - [x] 报警器
- [ ] 远程控制中心（商业版高阶功能，收费）
    - [ ] 统一管理
    - [ ] 短信报警，电话报警
    - [ ] 数据透传，虚拟串口，远程调试
    - [ ] API服务，对接APP和小程序
    
## 其他

- 项目的早期和支线版本已经在实际的养猪物联网和养鱼物联网项目中使用，效果良好
- 项目主线还在待续开发中，有兴趣的小伙伴可以加入进来
- 开源版本限制单机单核，有5W+连接需求请使用商业版（支持多机多核）

## 联系方式

- 邮箱：[jason@zgwit.com](mailto:jason@zgwit.com)
- 手机：[15161515197](tel:15161515197)(微信同号)

![微信号](https://labs.zgwit.com/qrcode.jpg)

[![真格智能实验室](https://labs.zgwit.com/logo.png)](https://labs.zgwit.com)
