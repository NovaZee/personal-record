

# kubectl 
[Kubectl命令参考](https://kubernetes.io/zh-cn/docs/reference/kubectl/generated/kubectl/)  

 kubectl is a command line interface for running commands against Kubernetes clusters. kubectl looks for a file named config in the $HOME/.kube directory. You can specify other kubeconfig files by setting the KUBECONFIG environment variable or by setting the --kubeconfig flag.
### 常用命令
- yaml查看 `kubectl get pods -o yaml`
- 特定资源的 YAML 配置  `kubectl get pod my-pod -o yaml`
- 通过标签选择器查找资源 `kubectl get pods -l app=nginx`
- 创建资源 `kubectl create -f ./my-manifest.yaml`
- 更新资源 `kubectl apply -f ./my-manifest.yaml` 
- 删除资源 `kubectl delete -f ./my-manifest.yaml`
- 强制删除 `kubectl delete pod my-pod --force`
- 优雅删除 `kubectl delete pod my-pod --grace-period=30`
- 删除指定标签的资源 `kubectl delete pods -l app=nginx`
- 日志查看 `kubectl logs my-pod`
- 进入容器 `kubectl exec -ti my-pod -- /bin/bash`
- 监控资源状态 `kubectl get pods -w`  -w：watch
---
- 查看某类资源的标签 `kubectl get pods --show-labels`
- 按标签筛选资源 ：使用 -l 或 --selector 按标签筛选资源。 `kubectl get pods -l env=production`
---
- 标记不可调度 `kubectl cordon my-node`
- 标记可调度 `kubectl uncordon my-node`
---
- 回滚重启 `kubectl rollout undo deployment/my-deployment`
- 重启 `kubectl rollout restart deployment/my-deployment`
- 查看历史版本 `kubectl rollout history deployment/my-deployment`
- 查看历史版本详细信息 `kubectl rollout history deployment/my-deployment --revision=2`
---
- 新增污点 `kubectl taint nodes <node-name> <key>=<value>:<effect>`
- 删除污点 `kubectl taint nodes <node-name> <key>:<effect>-`
- 污点 `kubectl taint nodes node1 key=value:NoSchedule-`

### 格式化输出 
-o 后面可选格式：输出格式。可选值为： json、yaml、name、go-template、go-template-file、template、templatefile、jsonpath、jsonpath-as-json、jsonpath-file。
- 扩展输出`kubectl get pods -o wide`
- 自定义列输出 `kubectl get pods -o custom-columns=NAME:.metadata.name,RS:.metadata.labels.app`
- yaml输出 `kubectl get pods -o yaml`
- json输出 `kubectl get pods -o json`
- jsonpath输出 `kubectl get pods -o jsonpath='{.items[*].metadata.name}'`

### Events
- 查看资源事件 `kubectl get events`
- 列举 default 命名空间中近期的事件
  `kubectl events`

- 列举所有命名空间中近期的事件
  `kubectl events --all-namespaces`

- 列举指定 Pod 的近期事件，然后等待更多事件，并在出现新事件时列举出来
  `kubectl events --for pod/web-pod-13je7 --watch`

- 以 YAML 格式列举近期的事件
  `kubectl events -o yaml`

- 仅列举类别为 “Warning” 或 “Normal” 的近期事件
  `kubectl events --types=Warning,Normal`