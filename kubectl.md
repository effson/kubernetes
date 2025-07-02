### 1. kubectl api-resources
列出 当前 Kubernetes 集群中支持的所有资源类型（Kind），包括：<br>
Pod, Service, Deployment, StatefulSet, ConfigMap 等常见资源<br>
各种 CRD（CustomResourceDefinition）资源（如果有安装）,每个资源的缩写（SHORTNAMES）、是否命名空间级（NAMESPACED）、所在 API 组（APIGROUP）
```
root@master01:/home/jeff/k8s/deployment# kubectl api-resources
NAME                              SHORTNAMES   APIVERSION                        NAMESPACED   KIND
bindings                                       v1                                true         Binding
componentstatuses                 cs           v1                                false        ComponentStatus
configmaps                        cm           v1                                true         ConfigMap
endpoints                         ep           v1                                true         Endpoints
events                            ev           v1                                true         Event
limitranges                       limits       v1                                true         LimitRange
namespaces                        ns           v1                                false        Namespace
nodes                             no           v1                                false        Node
persistentvolumeclaims            pvc          v1                                true         PersistentVolumeClaim
persistentvolumes                 pv           v1                                false        PersistentVolume
pods                              po           v1                                true         Pod
podtemplates                                   v1                                true         PodTemplate
replicationcontrollers            rc           v1                                true         ReplicationController
resourcequotas                    quota        v1                                true         ResourceQuota
secrets                                        v1                                true         Secret
serviceaccounts                   sa           v1                                true         ServiceAccount
services                          svc          v1                                true         Service
mutatingwebhookconfigurations                  admissionregistration.k8s.io/v1   false        MutatingWebhookConfiguration
validatingwebhookconfigurations                admissionregistration.k8s.io/v1   false        ValidatingWebhookConfiguration
customresourcedefinitions         crd,crds     apiextensions.k8s.io/v1           false        CustomResourceDefinition
apiservices                                    apiregistration.k8s.io/v1         false        APIService
controllerrevisions                            apps/v1                           true         ControllerRevision
daemonsets                        ds           apps/v1                           true         DaemonSet
deployments                       deploy       apps/v1                           true         Deployment
replicasets                       rs           apps/v1                           true         ReplicaSet
statefulsets                      sts          apps/v1                           true         StatefulSet
selfsubjectreviews                             authentication.k8s.io/v1          false        SelfSubjectReview
tokenreviews                                   authentication.k8s.io/v1          false        TokenReview
localsubjectaccessreviews                      authorization.k8s.io/v1           true         LocalSubjectAccessReview
selfsubjectaccessreviews                       authorization.k8s.io/v1           false        SelfSubjectAccessReview
selfsubjectrulesreviews                        authorization.k8s.io/v1           false        SelfSubjectRulesReview
subjectaccessreviews                           authorization.k8s.io/v1           false        SubjectAccessReview
horizontalpodautoscalers          hpa          autoscaling/v2                    true         HorizontalPodAutoscaler
cronjobs                          cj           batch/v1                          true         CronJob
jobs                                           batch/v1                          true         Job
certificatesigningrequests        csr          certificates.k8s.io/v1            false        CertificateSigningRequest
leases                                         coordination.k8s.io/v1            true         Lease
bgpconfigurations                              crd.projectcalico.org/v1          false        BGPConfiguration
bgpfilters                                     crd.projectcalico.org/v1          false        BGPFilter
bgppeers                                       crd.projectcalico.org/v1          false        BGPPeer
blockaffinities                                crd.projectcalico.org/v1          false        BlockAffinity
caliconodestatuses                             crd.projectcalico.org/v1          false        CalicoNodeStatus
clusterinformations                            crd.projectcalico.org/v1          false        ClusterInformation
felixconfigurations                            crd.projectcalico.org/v1          false        FelixConfiguration
globalnetworkpolicies                          crd.projectcalico.org/v1          false        GlobalNetworkPolicy
globalnetworksets                              crd.projectcalico.org/v1          false        GlobalNetworkSet
hostendpoints                                  crd.projectcalico.org/v1          false        HostEndpoint
ipamblocks                                     crd.projectcalico.org/v1          false        IPAMBlock
ipamconfigs                                    crd.projectcalico.org/v1          false        IPAMConfig
ipamhandles                                    crd.projectcalico.org/v1          false        IPAMHandle
ippools                                        crd.projectcalico.org/v1          false        IPPool
ipreservations                                 crd.projectcalico.org/v1          false        IPReservation
kubecontrollersconfigurations                  crd.projectcalico.org/v1          false        KubeControllersConfiguration
networkpolicies                                crd.projectcalico.org/v1          true         NetworkPolicy
networksets                                    crd.projectcalico.org/v1          true         NetworkSet
endpointslices                                 discovery.k8s.io/v1               true         EndpointSlice
events                            ev           events.k8s.io/v1                  true         Event
flowschemas                                    flowcontrol.apiserver.k8s.io/v1   false        FlowSchema
prioritylevelconfigurations                    flowcontrol.apiserver.k8s.io/v1   false        PriorityLevelConfiguration
ingressclasses                                 networking.k8s.io/v1              false        IngressClass
ingresses                         ing          networking.k8s.io/v1              true         Ingress
networkpolicies                   netpol       networking.k8s.io/v1              true         NetworkPolicy
runtimeclasses                                 node.k8s.io/v1                    false        RuntimeClass
poddisruptionbudgets              pdb          policy/v1                         true         PodDisruptionBudget
clusterrolebindings                            rbac.authorization.k8s.io/v1      false        ClusterRoleBinding
clusterroles                                   rbac.authorization.k8s.io/v1      false        ClusterRole
rolebindings                                   rbac.authorization.k8s.io/v1      true         RoleBinding
roles                                          rbac.authorization.k8s.io/v1      true         Role
priorityclasses                   pc           scheduling.k8s.io/v1              false        PriorityClass
csidrivers                                     storage.k8s.io/v1                 false        CSIDriver
csinodes                                       storage.k8s.io/v1                 false        CSINode
csistoragecapacities                           storage.k8s.io/v1                 true         CSIStorageCapacity
storageclasses                    sc           storage.k8s.io/v1                 false        StorageClass
volumeattachments                              storage.k8s.io/v1                 false        VolumeAttachment
```
### 2. kubectl api-versions
列出 当前 Kubernetes 集群中支持的所有 API 版本
```
root@master01:/home/jeff/k8s/deployment# kubectl api-versions
admissionregistration.k8s.io/v1
apiextensions.k8s.io/v1
apiregistration.k8s.io/v1
apps/v1
authentication.k8s.io/v1
authorization.k8s.io/v1
autoscaling/v1
autoscaling/v2
batch/v1
certificates.k8s.io/v1
coordination.k8s.io/v1
crd.projectcalico.org/v1
discovery.k8s.io/v1
events.k8s.io/v1
flowcontrol.apiserver.k8s.io/v1
flowcontrol.apiserver.k8s.io/v1beta3
networking.k8s.io/v1
node.k8s.io/v1
policy/v1
rbac.authorization.k8s.io/v1
scheduling.k8s.io/v1
storage.k8s.io/v1
v1
```
### 3. kubectl apply 、kubectl create
kubectl apply 与 kubectl create 都能创建资源
```
✅ 核心区别总结
功能点	                 kubectl create	                  kubectl apply
用途	                 初次创建资源	                    创建或更新资源（推荐方式）
是否幂等	               否（重复执行会报错）	              是（多次执行不会报错，且支持自动合并）
更新资源	               ❌ 不能更新已存在资源	            ✅ 可以更新已有资源
推荐程度	               一般仅用于简单/临时资源	          ✅ 推荐用于部署、配置管理、CI/CD
使用机制	               直接创建资源（POST）	            使用 kubectl.kubernetes.io/last-applied-configuration 做对比和合并
对 YAML 的需求	         每次都应是完整的资源定义	          可以只包含修改的字段（建议完整）
```
### 4. kubectl expose
kubectl expose 是一个快捷命令，用于基于已有的 Pod、Deployment、ReplicaSet、Service 等资源 快速创建一个 Service（服务资源），让其他组件可以访问这些资源
```
kubectl expose RESOURCE NAME [--port=PORT] [--target-port=TARGET_PORT] [--type=TYPE] [flags]

RESOURCE	要暴露的资源类型（如 deployment、pod、rc）
NAME	要暴露的资源名称
--port	Service 对外暴露的端口（port 字段）
--target-port	容器中实际监听的端口（targetPort 字段）
--type	Service 类型，常见有 ClusterIP（默认）、NodePort、LoadBalancer

kubectl expose deployment nginx-deployment --port=80 --target-port=80
kubectl expose deployment nginx-deployment --port=80 --target-port=80 --type=NodePort
```
