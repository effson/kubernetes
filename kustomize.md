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
