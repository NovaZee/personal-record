# istio

官方网站是最好老师：[Istio / 入门](https://istio.io/latest/zh/docs/ambient/getting-started/)

What's a service mesh? And why do I need one

### 背景：

##### **Service Mesh**：

- **概述：**Service Mesh 实际上就是处于 TCP/IP 之上的一个抽象层，它假设底层的 L3/L4 网络能够点对点地传输字节（当然，它也假设网络环境是不可靠的，所以 Service Mesh 也必须具备处理网络故障的能力）。

  从某种程度上说，Service Mesh 有点类似 TCP/IP 。TCP 对网络端点间传输字节的机制进行了抽象，而Service Mesh则是对服务节点间请求的路由机制进行了抽象。Service Mesh 不关心消息体是什么，也不关心它们是如何编码的。应用程序的目标是“将某些东西从A传送到B”，而 Service Mesh 所要做的就是实现这个目标，并处理传送过程中可能出现的任何故障。

- **目标：**应用运行时提供统一的、应用层面的可见性和可控性。

##### **类SDN（拓展）：**

解耦网络设备的控制层和数据层，使网络更具灵活性和可管理性。传统网络架构中，网络设备（如路由器和交换机）的控制层和数据层是紧密耦合的，设备具有独立的控制逻辑，彼此难以协调。而在SDN中，控制层和数据层被分离，控制逻辑集中在一个集中式控制器上，从而简化了网络管理和配置。

1. **控制平面与数据平面分离**：SDN将网络设备的控制功能从设备中剥离出来，放在集中控制器上；设备仅负责转发数据。
2. **集中控制**：网络管理员可以通过SDN控制器管理网络的所有流量，实现集中化管理与自动化配置。
3. **编程化网络**：SDN允许使用编程接口（例如OpenFlow）与控制器通信，从而实现网络设备的动态配置和资源分配。
4. **可视化和动态调整**：管理员可以实时监控网络流量，进行网络拓扑的动态调整，更好地优化资源利用

##### **istio**：
简单来说：为微服务应用提供了一套**完整的解决方案**，将现在化基建抽离出来组合，而不是与应用代码强耦合。
因为只是解决方案，所以不止一种！

- 阿里基于 golang 重写了 Envoy 并在这基础上构建了 Sofa-Mosn/Pilot， 下沉了 Istio 中颇受诟病的 Mixer 的限流能力，
- 腾讯基于 Envoy 进行改造整合内部的 TSF 服务框架，
- 美团基于 Envoy 改造整合内部的 Octo 服务框架，
- 微博基于 Motan-Go 研制出 Motan-Mesh，整合了自己的服务治理体系，
- 华为的 ServiceComb 也是类似的做法， Mixer完全下沉，
- Twitter推出 Conduit，基于Rust，也将Mixer完成下沉，
- 猫眼参考了 Envoy/Sofa-Mosn 的架构，自研了Maoyan-Mesh。
- 等！
### 概述：

- 轻量级的网络代理
- 三大作用
  - **可靠性功能**：请求重试、超时、金丝雀（流量拆分/转移），降级，熔断，故障注入等
  - **可观测性功能**：每个服务或单个路由的成功率、延迟和请求量的聚合;绘制服务拓扑图;等。
  - **安全功能**：双向 TLS、访问控制等  ？

### 部署安装：

个人实践helm为例：官方教程，无坑！

1：首先安装helm
2：配置 Helm 存储库

```
helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update
```

3：为 Istio 组件，创建命名空间 istio-system

```
kubectl create namespace istio-system
```

4：安装 Istio Base Chart

**作用**：`istio-base` 主要是安装 Istio 的基础 CRDs，这些定义是 Istio 控制平面和数据平面之间进行配置和通信的基础。

**主要功能**：Istio 的 CRDs 是控制平面与数据平面之间进行管理的重要定义，包括 VirtualService、DestinationRule、Gateway 等资源的定义

```
helm install istio-base istio/base -n istio-system --set defaultRevision=default
helm ls -n istio-system ## 验证
##验证结果
NAME       NAMESPACE    REVISION UPDATED                                 STATUS   CHART        APP VERSION
istio-base istio-system 1        2024-04-17 22:14:45.964722028 +0000 UTC deployed base-1.23.2  1.23.2
```

5：安装 Istio Discovery Chart

```
helm install istiod istio/istiod -n istio-system --wait
helm ls -n istio-system ## 验证
##验证结果：
NAME       NAMESPACE    REVISION UPDATED                                 STATUS   CHART         APP VERSION
istio-base istio-system 1        2024-04-17 22:14:45.964722028 +0000 UTC deployed base-1.23.2   1.23.2
istiod     istio-system 1        2024-04-17 22:14:45.964722028 +0000 UTC deployed istiod-1.23.2 1.23.2
```

6：检查 `istiod` 服务是否安装成功，确认其 Pod 是否正在运行

**作用**：`istiod` 是 Istio 控制平面的核心组件，整合了 Pilot、Citadel 和 Galley 等组件的功能，提供服务发现、流量管理、安全性和配置管理等功能，是服务网格正常运作的关键部分。

```
kubectl get deployments -n istio-system --output wide
```

6：（可选）入站/出站 网关

