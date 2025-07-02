```
root@master01:/home/jeff/k8s/deployment# kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
serviceaccount/metrics-server created
clusterrole.rbac.authorization.k8s.io/system:aggregated-metrics-reader created
clusterrole.rbac.authorization.k8s.io/system:metrics-server created
rolebinding.rbac.authorization.k8s.io/metrics-server-auth-reader created
clusterrolebinding.rbac.authorization.k8s.io/metrics-server:system:auth-delegator created
clusterrolebinding.rbac.authorization.k8s.io/system:metrics-server created
service/metrics-server created
deployment.apps/metrics-server created
apiservice.apiregistration.k8s.io/v1beta1.metrics.k8s.io created


需要修改配置：
kubectl edit deployment metrics-server -n kube-system

containers:
      - args:
        - --cert-dir=/tmp
        - --secure-port=10250
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s        
在 args: 下添加：
        - --kubelet-insecure-tls

root@master01:/home/jeff/k8s/deployment# kubectl edit deployment metrics-server -n kube-system
deployment.apps/metrics-server edited


root@master01:/home/jeff/k8s/deployment# kubectl get pods -n kube-system | grep metrics
metrics-server-596474b58-v8tvn             1/1     Running   0             49s


root@master01:/home/jeff/k8s/deployment# kubectl top node
NAME       CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%
master01   138m         3%     1513Mi          40%
worker01   33m          0%     599Mi           15%
worker02   38m          0%     627Mi           16%
```
