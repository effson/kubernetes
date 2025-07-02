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
```
