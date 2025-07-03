```
apiVersion: v1
kind: Pod
metadata:
  name: pod1
  namespace: default
  labels:
    app: app_nginx
spec:
  tolerations:
  - key: "key1"
    operator: "Exists"
    effect: "NoSchedule"
  nodeSelector:
    kubernetes.io/hostname: "worker01"
  restartPolicy: Always
  containers:
  - name: app-container
    image: nginx:latest
    imagePullPolicy: IfNotPresent

    ports:
    - containerPort: 80

    resources:
      requests:
        cpu: "200m"
        memory: "256Mi"
      limits:
        cpu: "300m"
        memory: "512Mi"

    env:
      - name: APP_ENV
        value: "production"
      - name: CONFIG_PATH
        value: "/etc/demo/config/app.conf"
      - name: SECRET_TOKEN
        valueFrom:
          secretKeyRef:
            name: demo-secret
            key: token
```
### initContainers:<br>
initContainer 是 Kubernetes Pod 中的一种特殊容器，在主容器（即 containers 中的容器）启动之前 先执行初始化任务。
可以有一个或多个 initContainer，按顺序串行执行，全部成功后才会启动主容器.
#### 🔧 作用：
做初始化工作，比如：<br>
下载配置文件<br>
检查依赖服务是否就绪<br>
拷贝文件或设置权限<br>
等待数据库、Redis 等服务启动完毕<br>
避免在主容器中写很多复杂的启动逻辑脚本，让职责更清晰<br>
```
apiVersion: v1
kind: Pod
metadata:
  name: with-init
spec:
  initContainers:
  - name: init-myservice
    image: busybox
    command: ['sh', '-c', 'until nslookup myservice; do echo waiting; sleep 2; done']

  containers:
  - name: app
    image: nginx
    ports:
    - containerPort: 80
```
myservice 是在 Kubernetes 集群中定义的 Service 的名字，不断执行 <mark>nslookup myservice</mark>，直到它能解析成功（说明 myservice 服务已创建并在 DNS 中可解析），否则每隔 2 秒打印一次 waiting<br>
创建了以下一个名为 myservice 的 Service：
```
apiVersion: v1
kind: Service
metadata:
  name: myservice
spec:
  selector:
    app: app_nginx
  ports:
    - port: 80
      targetPort: 80
```
