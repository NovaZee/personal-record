## 1：kind

官方文档：[kind – Configuration](https://kind.sigs.k8s.io/docs/user/configuration/)

- 启动集群模板：kind create cluster --config my-cluster-multi-node.yaml --name k8s-cluster

```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
# One control plane node and three "workers".
#
# While these will not add more real compute capacity and
# have limited isolation, this can be useful for testing
# rolling updates etc.
#
# The API-server and other control plane components will be
# on the control-plane node.
#
# You probably don't need this unless you are testing Kubernetes itself.
nodes:
- role: control-plane
- role: worker //节点1，2，3
- role: worker
- role: worker
```

- 删除：kind delete cluster --name k8s-cluster 快速删除指定集群
- 删除所有： kind delete clusters --all

## 2：dashboard

- ***apply dashboard.yaml： kubectl apply -f .\dashboard.yaml***

```
PS F:\k8s> kubectl apply -f .\dashboard.yaml
namespace/kubernetes-dashboard created
serviceaccount/kubernetes-dashboard created
service/kubernetes-dashboard created
secret/kubernetes-dashboard-certs created
secret/kubernetes-dashboard-csrf created
secret/kubernetes-dashboard-key-holder created
configmap/kubernetes-dashboard-settings created
role.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrole.rbac.authorization.k8s.io/kubernetes-dashboard created
rolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
deployment.apps/kubernetes-dashboard created
service/dashboard-metrics-scraper created
Warning: spec.template.metadata.annotations[seccomp.security.alpha.kubernetes.io/pod]: non-functional in v1.27+; use the "seccompProfile" field instead
deployment.apps/dashboard-metrics-scraper created  
```

访问地址：**http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/**

- ***控制台访问令牌：kubectl apply -f kube-system-default.yaml***

- ***导出令牌***：

  ```
  $TOKEN=((kubectl -n kube-system describe secret default | Select-String "token:") -split " +")[1]
  
  kubectl config set-credentials xxxx --token="${TOKEN}"  //xxx为 cluster，可以查看./kube/config
  
  echo $TOKEN
  
  ```

- ***kubectl proxy***

## 3：metric-server

容器args添加--kubelet-insecure-tls 跳过tis证书验证：

```
containers:
- args:
  - --kubelet-insecure-tls
```

修改与kubelet通信优先级

```
 - --kubelet-preferred-address-types=InternalIP,Hostname,InternalDNS,ExternalDNS,ExternalIP
```

观察pod启动

```
kubeclt top pod -A
kubectl top pod --sort-by=cpu -A
```

