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
