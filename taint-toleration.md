## taint和toleration
Kubernetes 中，taint（污点）和 toleration（容忍）是节点调度控制机制的一部分，一起用来控制哪些 Pod 可以或不能 被调度到某些节点上
#### Taint 是节点上的标签，表示“不欢迎谁”；
#### Toleration 是 Pod 上的声明，表示“可以接受某种不欢迎”

### 1. taint
```
kubectl taint nodes <node-name> <key>=<value>:<effect>
```
```
effect:
NoSchedule       :  不能调度到该节点，已经在该节点上运行的pod不受影响
PreferNoSchedule :  尽量避免调度到该节点，但非强制
NoExecute        :  不会调度到该节点并且会将该节点上的pod驱逐

去掉 Kubernetes 节点上的 taint（污点）：
kubectl taint nodes <节点名> <key>[=<value>]:<effect>-


root@master01:/home/jeff/k8s/deployment# kubectl taint node worker01 key=forbiddenkey:NoSchedule
node/worker01 tainted


root@master01:/home/jeff/k8s/deployment# kubectl describe node worker01
Name:               worker01
Roles:              <none>
Labels:             beta.kubernetes.io/arch=amd64
                    beta.kubernetes.io/os=linux
                    kubernetes.io/arch=amd64
                    kubernetes.io/hostname=worker01
                    kubernetes.io/os=linux
Annotations:        kubeadm.alpha.kubernetes.io/cri-socket: unix:///var/run/containerd/containerd.sock
                    node.alpha.kubernetes.io/ttl: 0
                    projectcalico.org/IPv4Address: 192.168.23.170/24
                    projectcalico.org/IPv4IPIPTunnelAddr: 192.244.5.0
                    volumes.kubernetes.io/controller-managed-attach-detach: true
CreationTimestamp:  Wed, 02 Jul 2025 00:09:17 +0000
Taints:             key=forbiddenkey:NoSchedule
Unschedulable:      false
Lease:
  HolderIdentity:  worker01
  AcquireTime:     <unset>
  RenewTime:       Wed, 02 Jul 2025 13:27:20 +0000
Conditions:
  Type                 Status  LastHeartbeatTime                 LastTransitionTime                Reason                       Message
  ----                 ------  -----------------                 ------------------                ------                       -------
  NetworkUnavailable   False   Wed, 02 Jul 2025 00:11:10 +0000   Wed, 02 Jul 2025 00:11:10 +0000   CalicoIsUp                   Calico is running on this node
  MemoryPressure       False   Wed, 02 Jul 2025 13:27:12 +0000   Wed, 02 Jul 2025 00:09:17 +0000   KubeletHasSufficientMemory   kubelet has sufficient memory available
  DiskPressure         False   Wed, 02 Jul 2025 13:27:12 +0000   Wed, 02 Jul 2025 00:09:17 +0000   KubeletHasNoDiskPressure     kubelet has no disk pressure
  PIDPressure          False   Wed, 02 Jul 2025 13:27:12 +0000   Wed, 02 Jul 2025 00:09:17 +0000   KubeletHasSufficientPID      kubelet has sufficient PID available
  Ready                True    Wed, 02 Jul 2025 13:27:12 +0000   Wed, 02 Jul 2025 00:10:30 +0000   KubeletReady                 kubelet is posting ready status. AppArmor enabled
Addresses:
  InternalIP:  192.168.23.170
  Hostname:    worker01
Capacity:
  cpu:                4
  ephemeral-storage:  29751268Ki
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             3960972Ki
  pods:               110
Allocatable:
  cpu:                4
  ephemeral-storage:  27418768544
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             3858572Ki
  pods:               110
System Info:
  Machine ID:                 7757c7090bab473cb231af78e19d571f
  System UUID:                84ae4d56-58e5-22af-f77a-76b52ea8ea3c
  Boot ID:                    6de153e4-38c7-43d5-bbf8-7ac38c933a23
  Kernel Version:             6.8.0-62-generic
  OS Image:                   Ubuntu 24.04.2 LTS
  Operating System:           linux
  Architecture:               amd64
  Container Runtime Version:  containerd://1.7.27
  Kubelet Version:            v1.29.2
  Kube-Proxy Version:         v1.29.2
PodCIDR:                      10.244.1.0/24
PodCIDRs:                     10.244.1.0/24
Non-terminated Pods:          (2 in total)
  Namespace                   Name                 CPU Requests  CPU Limits  Memory Requests  Memory Limits  Age
  ---------                   ----                 ------------  ----------  ---------------  -------------  ---
  kube-system                 calico-node-r796x    250m (6%)     0 (0%)      0 (0%)           0 (0%)         13h
  kube-system                 kube-proxy-xjqz2     0 (0%)        0 (0%)      0 (0%)           0 (0%)         13h
Allocated resources:
  (Total limits may be over 100 percent, i.e., overcommitted.)
  Resource           Requests   Limits
  --------           --------   ------
  cpu                250m (6%)  0 (0%)
  memory             0 (0%)     0 (0%)
  ephemeral-storage  0 (0%)     0 (0%)
  hugepages-1Gi      0 (0%)     0 (0%)
  hugepages-2Mi      0 (0%)     0 (0%)
Events:              <none>
```
可以看到
```
Taints:             key=forbiddenkey:NoSchedule
```
```
root@master01:/home/jeff/k8s/deployment# kubectl taint node worker01 key=forbiddenkey:NoSchedule-
node/worker01 untainted
```
### 2. toleration
toleration 是加在 Pod 上的，表示“可以容忍某些 taint”<br>
YAML：
```
tolerations:
- key: "type"
  operator: "Equal"
  value: "highcpu"
  effect: "NoSchedule"

operator：
Equal  ：默认值。表示 key 和 value 都要匹配 才容忍
Exists ：表示 只匹配 key 就行，value 不重要
```
