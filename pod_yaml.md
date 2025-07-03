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
