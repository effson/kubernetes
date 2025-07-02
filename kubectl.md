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
### 5. kubectl explain
kubectl explain 是 Kubernetes 提供的命令，用来查看资源（Kind）及其字段的含义、结构和 API 说明
```
kubectl explain RESOURCE[.FIELD...]


root@master01:/home/jeff/k8s/deployment# kubectl explain deployment
GROUP:      apps
KIND:       Deployment
VERSION:    v1

DESCRIPTION:
    Deployment enables declarative updates for Pods and ReplicaSets.

FIELDS:
  apiVersion    <string>
    APIVersion defines the versioned schema of this representation of an object.
    Servers should convert recognized schemas to the latest internal value, and
    may reject unrecognized values. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources

  kind  <string>
    Kind is a string value representing the REST resource this object
    represents. Servers may infer this from the endpoint the client submits
    requests to. Cannot be updated. In CamelCase. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds

  metadata      <ObjectMeta>
    Standard object's metadata. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata

  spec  <DeploymentSpec>
    Specification of the desired behavior of the Deployment.

  status        <DeploymentStatus>
    Most recently observed status of the Deployment.


root@master01:/home/jeff/k8s/deployment# kubectl explain deployment.metadata
GROUP:      apps
KIND:       Deployment
VERSION:    v1

FIELD: metadata <ObjectMeta>

DESCRIPTION:
    Standard object's metadata. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    ObjectMeta is metadata that all persisted resources must have, which
    includes all objects users must create.

FIELDS:
  annotations   <map[string]string>
    Annotations is an unstructured key value map stored with a resource that may
    be set by external tools to store and retrieve arbitrary metadata. They are
    not queryable and should be preserved when modifying objects. More info:
    https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations

  creationTimestamp     <string>
    CreationTimestamp is a timestamp representing the server time when this
    object was created. It is not guaranteed to be set in happens-before order
    across separate operations. Clients may not set this value. It is
    represented in RFC3339 form and is in UTC.

    Populated by the system. Read-only. Null for lists. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata

  deletionGracePeriodSeconds    <integer>
    Number of seconds allowed for this object to gracefully terminate before it
    will be removed from the system. Only set when deletionTimestamp is also
    set. May only be shortened. Read-only.

  deletionTimestamp     <string>
    DeletionTimestamp is RFC 3339 date and time at which this resource will be
    deleted. This field is set by the server when a graceful deletion is
    requested by the user, and is not directly settable by a client. The
    resource is expected to be deleted (no longer visible from resource lists,
    and not reachable by name) after the time in this field, once the finalizers
    list is empty. As long as the finalizers list contains items, deletion is
    blocked. Once the deletionTimestamp is set, this value may not be unset or
    be set further into the future, although it may be shortened or the resource
    may be deleted prior to this time. For example, a user may request that a
    pod is deleted in 30 seconds. The Kubelet will react by sending a graceful
    termination signal to the containers in the pod. After that 30 seconds, the
    Kubelet will send a hard termination signal (SIGKILL) to the container and
    after cleanup, remove the pod from the API. In the presence of network
    partitions, this object may still exist after this timestamp, until an
    administrator or automated process can determine the resource is fully
    terminated. If not set, graceful deletion of the object has not been
    requested.

    Populated by the system when a graceful deletion is requested. Read-only.
    More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata

  finalizers    <[]string>
    Must be empty before the object is deleted from the registry. Each entry is
    an identifier for the responsible component that will remove the entry from
    the list. If the deletionTimestamp of the object is non-nil, entries in this
    list can only be removed. Finalizers may be processed and removed in any
    order.  Order is NOT enforced because it introduces significant risk of
    stuck finalizers. finalizers is a shared field, any actor with permission
    can reorder it. If the finalizer list is processed in order, then this can
    lead to a situation in which the component responsible for the first
    finalizer in the list is waiting for a signal (field value, external system,
    or other) produced by a component responsible for a finalizer later in the
    list, resulting in a deadlock. Without enforced ordering finalizers are free
    to order amongst themselves and are not vulnerable to ordering changes in
    the list.

  generateName  <string>
    GenerateName is an optional prefix, used by the server, to generate a unique
    name ONLY IF the Name field has not been provided. If this field is used,
    the name returned to the client will be different than the name passed. This
    value will also be combined with a unique suffix. The provided value has the
    same validation rules as the Name field, and may be truncated by the length
    of the suffix required to make the value unique on the server.

    If this field is specified and the generated name exists, the server will
    return a 409.

    Applied only if Name is not specified. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency

  generation    <integer>
    A sequence number representing a specific generation of the desired state.
    Populated by the system. Read-only.

  labels        <map[string]string>
    Map of string keys and values that can be used to organize and categorize
    (scope and select) objects. May match selectors of replication controllers
    and services. More info:
    https://kubernetes.io/docs/concepts/overview/working-with-objects/labels

  managedFields <[]ManagedFieldsEntry>
    ManagedFields maps workflow-id and version to the set of fields that are
    managed by that workflow. This is mostly for internal housekeeping, and
    users typically shouldn't need to set or understand this field. A workflow
    can be the user's name, a controller's name, or the name of a specific apply
    path like "ci-cd". The set of fields is always in the version that the
    workflow used when modifying the object.

  name  <string>
    Name must be unique within a namespace. Is required when creating resources,
    although some resources may allow a client to request the generation of an
    appropriate name automatically. Name is primarily intended for creation
    idempotence and configuration definition. Cannot be updated. More info:
    https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names

  namespace     <string>
    Namespace defines the space within which each name must be unique. An empty
    namespace is equivalent to the "default" namespace, but "default" is the
    canonical representation. Not all objects are required to be scoped to a
    namespace - the value of this field for those objects will be empty.

    Must be a DNS_LABEL. Cannot be updated. More info:
    https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces

  ownerReferences       <[]OwnerReference>
    List of objects depended by this object. If ALL objects in the list have
    been deleted, this object will be garbage collected. If this object is
    managed by a controller, then an entry in this list will point to this
    controller, with the controller field set to true. There cannot be more than
    one managing controller.

  resourceVersion       <string>
    An opaque value that represents the internal version of this object that can
    be used by clients to determine when objects have changed. May be used for
    optimistic concurrency, change detection, and the watch operation on a
    resource or set of resources. Clients must treat these values as opaque and
    passed unmodified back to the server. They may only be valid for a
    particular resource or set of resources.

    Populated by the system. Read-only. Value must be treated as opaque by
    clients and . More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency

  selfLink      <string>
    Deprecated: selfLink is a legacy read-only field that is no longer populated
    by the system.

  uid   <string>
    UID is the unique in time and space value for this object. It is typically
    generated by the server on successful creation of a resource and is not
    allowed to change on PUT operations.

    Populated by the system. Read-only. More info:
    https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
```
### 6. kubectl get
获取资源
```
root@master01:/home/jeff/k8s/deployment# kubectl get ns
NAME              STATUS   AGE
default           Active   12h
kube-node-lease   Active   12h
kube-public       Active   12h
kube-system       Active   12h

root@master01:/home/jeff/k8s/deployment# kubectl get pods -n kube-system
NAME                                       READY   STATUS    RESTARTS      AGE
calico-kube-controllers-5fc7d6cf67-z5wks   1/1     Running   0             12h
calico-node-5m2p9                          1/1     Running   0             12h
calico-node-lznhz                          1/1     Running   0             12h
calico-node-r796x                          1/1     Running   0             12h
coredns-76f75df574-gc2wd                   1/1     Running   0             12h
coredns-76f75df574-gv4qz                   1/1     Running   0             12h
etcd-master01                              1/1     Running   0             12h
kube-apiserver-master01                    1/1     Running   0             12h
kube-controller-manager-master01           1/1     Running   1 (11h ago)   12h
kube-proxy-7xmgr                           1/1     Running   0             12h
kube-proxy-rp26m                           1/1     Running   0             12h
kube-proxy-xjqz2                           1/1     Running   0             12h
kube-scheduler-master01                    1/1     Running   1 (11h ago)   12h

```
