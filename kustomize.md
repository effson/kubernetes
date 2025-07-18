# 1.什么是 kustomize
- Kustomize 是 Kubernetes 官方推荐的原生配置管理工具，它允许你在不修改原始 YAML 文件的基础上，对资源进行个性化定制（比如：修改镜像、端口、标签、命名空间等）
- Kustomize 是 Kubernetes 原生支持的 YAML 模板定制工具，支持参数覆盖、继承组合、资源拼接，无需使用 Helm 之类的模板引擎

# 2.kustomize 使用
## 2.1 kubectl使用kustomization文件管理k8s
> kubectl 原生支持 kustomization.yaml，可以直接用 kubectl 来应用、查看或管理 Kustomize 管理的 Kubernetes 配置资源
## 2.2 命令格式
```
kubectl kustomize <kustomization_directory>
```
```
kubectl apply -k <目录路径>
```
## 2.3 命令使用
### 2.3.1 生成资源
#### 2.3.1.1 configMapGenerator生成configMap
##### 2.3.1.1.1 基于属性文件生成configMap
```
root@master01:/home/kustomize/k1# cat <<EOF >application.properties
> FOO=Bar
> EOF
```
```
root@master01:/home/kustomize/k1# ls
application.properties
```
```
root@master01:/home/kustomize/k1# cat <<EOF >./kustomization.yaml
> configMapGenerator:
> - name: example-configmap-v1
>   files:
>   - application.properties
> EOF
```
```
root@master01:/home/kustomize/k1# ls
application.properties  kustomization.yaml
```
```
root@master01:/home/kustomize/k1# kubectl kustomize ./
apiVersion: v1
data:
  application.properties: |
    FOO=Bar
kind: ConfigMap
metadata:
  name: example-configmap-v1-g4hk9g2ff8
```
##### 2.3.1.1.2 env文件生成configMap
```
root@master01:/home/kustomize/k1# cat <<EOF >.env
> FOO=Bar
> EOF
```
```
root@master01:/home/kustomize/k1# cat <<EOF >./kustomization.yaml
> configMapGenerator:
> - name: example-configmap-env
>   envs:
>   - .env
> EOF
```
```
root@master01:/home/kustomize/k1# kubectl kustomize ./
apiVersion: v1
data:
  FOO: Bar
kind: ConfigMap
metadata:
  name: example-configmap-env-42cfbf598f
```
#### 2.3.1.2 使用生成的configMap
> base：deployment.yaml
```
apiVersion: v1
kind: Deployment
metadata:
  name: my-app
  labels:
    app: my-app
spec:
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
      spec:
        containers:
        - name: app
          image: my-app
          volumeMounts:
          - name: config
            mountPath: /config
        volumes:
        - name: config
          configMap:
            name: example-configmap-env-42cfbf598f
```
> kustomization.yaml
```
resources:
- deployment.yaml
configMapGenerator:
- name: example-configmap-env
  envs:
  - .env
```
> kubectl kustomize ./
```
root@master01:/home/kustomize/k2# kubectl kustomize ./
apiVersion: v1
data:
  FOO: Bar
kind: ConfigMap
metadata:
  name: example-configmap-env-42cfbf598f
---
apiVersion: v1
kind: Deployment
metadata:
  labels:
    app: my-app
  name: my-app
spec:
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - image: my-app
        name: app
        volumeMounts:
        - mountPath: /config
          name: config
      volumes:
      - configMap:
          name: example-configmap-env-42cfbf598f
        name: config
```
#### 2.3.1.3 设置贯穿性字段
贯穿性字段在 Kustomize 中的术语叫做：全局变更字段（Global fields），也称为统一注入字段：使用场景如下：
- 设置所有资源所属命名空间
- 为所有对象添加相同前缀/后缀
- 为对象添加相同的标签(Label)集合
- 为对象添加相同的注解(Annotations)集合
> deployment.yaml
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
      spec:
        containers:
        - name: nginx
          image: nginx
```
> kustomization.yaml
```
namespace: my-namespace
namePrefix: dev-
nameSuffix: "-001"
commonLabels:
  app: bingo
commonAnnotations:
  oncallPager: 400-500-600
resources:
- deployment.yaml
```
建议使用新版:
```
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

namespace: my-namespace
namePrefix: dev-
nameSuffix: "-001"

labels:
  - pairs:
      app: bingo

commonAnnotations:
  oncallPager: "400-500-600"

resources:
  - deployment.yaml
```
> kubectl kustomize ./
```
root@master01:/home/kustomize/k3# kubectl kustomize ./
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    oncallPager: 400-500-600
  labels:
    app: bingo
  name: dev-nginx-deployment-001
  namespace: my-namespace
spec:
  selector:
    matchLabels:
      app: bingo
  template:
    metadata:
      annotations:
        oncallPager: 400-500-600
      labels:
        app: bingo
    spec:
      containers:
      - image: nginx
        name: nginx
```
#### 2.3.1.4 组合和定制资源
##### 2.3.1.4.1 组合
> deployment.yaml
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  selector:
    matchLabels:
      run: my-nginx
  replicas: 2
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - name: my-nginx
        image: nginx
        ports:
        - containerPort: 80
```
> service.yaml
```
apiVersion: v1
kind: Service
metadata:
  name: my-nginx
  labels:
    run: my-nginx
spec:
  ports:
  - port: 80
    protocol: TCP
  selector:
    run: my-nginx
```
> kustomization.yaml
```
resources:
- deployment.yaml
- service.yaml
```
> kubectl kustomize ./<br>
结果是将两部分组合起来<br>
> kubectl apply -k ./
```
root@master01:/home/kustomize/k4# kubectl apply -k ./
service/my-nginx unchanged
deployment.apps/my-nginx created
```
```
root@master01:/home/kustomize/k4# kubectl get pods
NAME                                                          READY   STATUS              RESTARTS   AGE
my-nginx-684dd4dcd4-gxn7x                                     1/1     Running             0          30s
my-nginx-684dd4dcd4-lk562                                     1/1     Running             0          30s
```
> kubectl get -k ./
```
root@master01:/home/kustomize/k4# kubectl get -k ./
NAME               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
service/my-nginx   ClusterIP   10.96.126.170   <none>        80/TCP    5m7s

NAME                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/my-nginx   2/2     2            2           2m8s
```
> kubectl delete -k ./
```
root@master01:/home/kustomize/k4# kubectl delete -k ./
service "my-nginx" deleted
deployment.apps "my-nginx" deleted
```
##### 2.3.1.4.2 定制
###### 2.3.1.4.2.1 patchesStrategicMerge\pathes
patchesStrategicMerge，它是 Kustomize 中用来局部修改资源字段的重要功能,可以使用它修改 base 目录中的 Deployment、Service 等资源:
- 修改 replicas
- 替换镜像
- 添加或修改 label/annotation
- 更改 port、env、volume 等配置
> deployment.yaml 同 2.3.1.4.1
> increase-replica.yaml
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  replicas: 3
```
> set-rmemory.yaml
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  template:
    spec:
      containers:
      - name: my-nginx
        resources:
          limits:
            memory: 512Mi
```
> kustomization.yaml
```
resources:
- deployment.yaml
- service.yaml
patchesStrategicMerge:
- increase-replica.yaml
- set-rmemory.yaml
```
新版：
```
resources:
- deployment.yaml
patches:
  - path: increase-replica.yaml
    target:
      kind: Deployment
      name: my-nginx
  - path: set-memory.yaml
    target:
      kind: Deployment
      name: my-nginx
```
> kubectl kustomize ./
```
root@master01:/home/kustomize/k5# kubectl kustomize ./
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      run: my-nginx
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - image: nginx
        name: my-nginx
        ports:
        - containerPort: 80
        resources:
          limits:
            memory: 512Mi
```
###### 2.3.1.4.2.2 patchesJson6902
使用 RFC 6902 JSON Patch 语法，来对资源对象进行精确字段修改<br>
> deployment.yaml
> patch.yaml or patch.json
```
- op: replace
  path: /spec/replicas
  value: 3

- op: add
  path: /spec/template/spec/containers/0/resources
  value:
    limits:
      memory: 512Mi
```
or:
```
[
  {
    "op": "replace",
    "path": "/spec/replicas",
    "value": 3
  },
  {
    "op": "add",
    "path": "/spec/template/spec/containers/0/resources",
    "value": {
      "limits": {
        "memory": "512Mi"
      }
    }
  }
]
```
> kustomization.yaml
```
resources:
- deployment.yaml
patchesJson6902:
  - target:
      group: apps
      version: v1
      kind: Deployment
      name: my-nginx
    path: patch.json
```
or
```
resources:
- deployment.yaml
patchesJson6902:
  - target:
      group: apps
      version: v1
      kind: Deployment
      name: my-nginx
    path: patch.yaml
```
> kustomization.yaml 新版写法同样是应用patches：
```
resources:
- deployment.yaml
patches:
  - path: patch.json
    target:
      kind: Deployment
      name: my-nginx
```
一样的作用
###### 2.3.1.4.2.3 kustomization.yaml文件images字段注入
> deployment.yaml同上
> kustomization.yaml
```
resources:
- deployment.yaml
images:
- name: nginx
  newName: nginx
  newTag: 1.25.0
```
> kubectl kustomize ./
```
root@master01:/home/kustomize/k7# kubectl kustomize ./
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  replicas: 2
  selector:
    matchLabels:
      run: my-nginx
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - image: nginx:1.25.0
        name: my-nginx
        ports:
        - containerPort: 80
```
###### 2.3.1.4.2.4 通过变量注入字段
> deployment.yaml
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  selector:
    matchLabels:
      run: my-nginx
  replicas: 2
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - name: my-nginx
        image: nginx
        command: ["start","--host","$(MY_SERVICE_NAME)"]
```
> service.yaml
```
apiVersion: v1
kind: Service
metadata:
  name: my-nginx
  labels:
    run: my-nginx
spec:
  ports:
  - port: 80
    protocol: TCP
  selector:
    run: my-nginx
```
> kustomization.yaml
```
namePrefix: dev-
nameSuffix: "-001"

resources:
- deployment.yaml
- service.yaml
vars:
  - name: MY_SERVICE_NAME
    objref:
      kind: Service
      name: my-nginx
      apiVersion: v1
```
新版推荐使用replacement:
```
namePrefix: dev-
nameSuffix: "-001"

resources:
- deployment.yaml
- service.yaml

replacements:
  - source:
      kind: Service
      name: my-nginx
      fieldPath: metadata.name
    targets:
      - select:
          kind: Deployment
          name: my-nginx
        fieldPaths:
          - spec.template.spec.containers.0.command.2
```
### 2.3.2 基准Bases与覆盖overlay
```
root@master01:/home/kustomize/kd_multi# mkdir base
root@master01:/home/kustomize/kd_multi# mkdir dev
root@master01:/home/kustomize/kd_multi# mkdir prod
root@master01:/home/kustomize/kd_multi# ls
base  dev  prod
```
在base目录创建deployment.yaml、service.yaml和kustomization.yaml:
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
spec:
  selector:
    matchLabels:
      run: my-nginx
  replicas: 2
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - name: my-nginx
        image: nginx
```
```
apiVersion: v1
kind: Service
metadata:
  name: my-nginx
  labels:
    run: my-nginx
spec:
  ports:
  - port: 80
    protocol: TCP
  selector:
    run: my-nginx
```
```
resources:
- deployment.yaml
- service.yaml
```
进入dev目录，创建kustomization.yaml：
```
bases:
- ../base
namePrefix: dev-
```
进入prod目录，创建kustomization.yaml：
```
bases:
- ../base
namePrefix: prod-
```
推荐使用新版：
```
resources:
- ../base
namePrefix: prod-
```
> kubectl kustomize ./
```
root@master01:/home/kustomize/kd_multi# cd prod/

root@master01:/home/kustomize/kd_multi/prod# ls
kustomization.yaml

root@master01:/home/kustomize/kd_multi/prod# kubectl kustomize ./
apiVersion: v1
kind: Service
metadata:
  labels:
    run: my-nginx
  name: prod-my-nginx
spec:
  ports:
  - port: 80
    protocol: TCP
  selector:
    run: my-nginx
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: prod-my-nginx
spec:
  replicas: 2
  selector:
    matchLabels:
      run: my-nginx
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
      - image: nginx
        name: my-nginx
```

### 2.3.3 kustomize 应用、查看和删除对象
```
root@master01:/home/kustomize/kd_multi/prod# kubectl apply -k .
service/prod-my-nginx created
deployment.apps/prod-my-nginx created
root@master01:/home/kustomize/kd_multi/prod# kubectl get -k .
NAME                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
service/prod-my-nginx   ClusterIP   10.109.16.235   <none>        80/TCP    13s

NAME                            READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/prod-my-nginx   2/2     2            2           13s
root@master01:/home/kustomize/kd_multi/prod# kubectl delete -k .
service "prod-my-nginx" deleted
deployment.apps "prod-my-nginx" deleted
```
# 3. kustomize客户端应用
## 3.1 kustomize客户端下载
~~~
root@master01:/home/kustomize# curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash
v5.7.0
kustomize installed to /home/kustomize/kustomize
~~~
~~~
root@master01:/home/kustomize# mv kustomize /usr/local/bin/
root@master01:/home/kustomize# kustomize version
v5.3.0
~~~

## 3.2 kustomize创建资源
> deployment.yaml文件同上
> kustomization.yaml
```
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

namespace: my-namespace
namePrefix: dev-
nameSuffix: "-001"

labels:
  - pairs:
      app: bingo

commonAnnotations:
  oncallPager: "400-500-600"

resources:
  - deployment.yaml
```
> kustomize build ./
```
root@master01:/home/kustomize/k9# kustomize build ./
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    oncallPager: 400-500-600
  labels:
    app: bingo
  name: dev-prod-my-nginx-001
  namespace: my-namespace
spec:
  replicas: 2
  selector:
    matchLabels:
      run: my-nginx
  template:
    metadata:
      annotations:
        oncallPager: 400-500-600
      labels:
        run: my-nginx
    spec:
      containers:
      - image: nginx
        name: my-nginx
```
## 3.3 kustomize修改资源

```
root@master01:/home/kustomize/k9# kustomize edit set image nginx=nginx:1.25.0

root@master01:/home/kustomize/k9# vim kustomization.yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

namespace: my-namespace
namePrefix: dev-
nameSuffix: "-001"

labels:
- pairs:
    app: bingo

commonAnnotations:
  oncallPager: 400-500-600

resources:
- deployment.yaml
images:
- name: nginx
  newName: nginx
  newTag: 1.25.0
```
```
root@master01:/home/kustomize/k9# kustomize build ./
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    oncallPager: 400-500-600
  labels:
    app: bingo
  name: dev-prod-my-nginx-001
  namespace: my-namespace
spec:
  replicas: 2
  selector:
    matchLabels:
      run: my-nginx
  template:
    metadata:
      annotations:
        oncallPager: 400-500-600
      labels:
        run: my-nginx
    spec:
      containers:
      - image: nginx:1.25.0
        name: my-nginx
```
