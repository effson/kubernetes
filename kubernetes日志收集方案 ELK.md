# kubernetes日志收集方案 ELK



# 一、为什么收集日志

收集日志可以用于：

- 分析用户行为
- 监控服务器状态
- 增强系统或应用安全性等。





# 二、收集哪些日志

- kubernetes集群节点系统日志
- kubernetes集群节点应用程序日志
- kubernetes集群中部署的应用程序日志



# 三、日志收集方案

## 3.1 日志收集技术栈ELK(ELKB)  +  Filebeat
![image-20200106175502929](https://github.com/user-attachments/assets/ff557e7d-b081-439d-8986-97982fcb0048)




## 3.2 日志收集技术栈  EK(EFK)  + fluentd

![image-20220408005109209](https://github.com/user-attachments/assets/d7e4c89f-ce6d-47ee-8308-29f8951b9589)






# 四、ELK集群部署

> 为了增加ELK集群的运行效率，一般建议在k8s集群之外使用物理机部署ELK集群，当然也可以直接在k8s集群内部署。



## 4.1 主机准备

| 主机     | 软件          | 版本   | 配置 | IP             |
| -------- | ------------- | ------ | ---- | -------------- |
| kibana   | kibana        | 7.17.2 | 2C2G | 192.168.10.200 |
| elastic  | elasticsearch | 7.17.2 | 2C4G | 192.168.10.201 |
| logstash | logstash      | 7.17.2 | 2C4G | 192.168.10.202 |



~~~powershell
# hostname set-hostname xxx
~~~



~~~powershell
# cat /etc/hosts
192.168.10.200 kibana
192.168.10.201 elastic
192.168.10.202 logstash
~~~





## 4.2 软件安装

> 由于软件下载较慢，请提前准备好以下软件。



### 4.2.1 安装jdk

> 所有主机全部安装，可考虑使用openjdk也可以使用oracle jdk。



~~~powershell
[root@kibana ~]# yum -y install java-11-openjdk

[root@elastic ~]# yum -y install java-11-openjdk

[root@logstash ~]# yum -y install java-11-openjdk
~~~





### 4.2.2 安装kibana


![image-20220407154620497](https://github.com/user-attachments/assets/716dee19-ee01-45c7-b77a-7687a702c670)





或



![image-20220407154735245](https://github.com/user-attachments/assets/5245ff93-f3d8-4002-81f2-44afe3f316be)




![image-20220407154902092](https://github.com/user-attachments/assets/68f9fd1b-b123-4f59-ae0a-e2df10af9e4f)





![image-20220407154953016](https://github.com/user-attachments/assets/f8cbff4d-bc25-4438-9c7e-490624a22355)






![image-20220407155103366](https://github.com/user-attachments/assets/61f6dd61-df9e-43e7-b99f-0d8d48b3734a)







![image-20220407155556894](https://github.com/user-attachments/assets/7be91f95-baa6-4cc3-ac98-104f263078d5)




![image-20220407155702199](https://github.com/user-attachments/assets/406e7e39-752f-4605-86f7-11db40f951cc)





~~~powershell
# wget https://artifacts.elastic.co/downloads/kibana/kibana-7.17.2-x86_64.rpm
~~~



~~~powershell
# yum -y install kibana-7.17.2-x86_64.rpm
~~~





### 4.2.3 安装elasticsearch




![image-20220407160045282](https://github.com/user-attachments/assets/4735e635-27fd-4955-a328-74eb8792c93c)




![image-20220407160127340](https://github.com/user-attachments/assets/7d9cf97c-cadf-49fd-937f-997f0edb044f)







~~~powershell
# wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.17.2-x86_64.rpm
~~~



~~~powershell
# yum -y install elasticsearch-7.17.2-x86_64.rpm
~~~





### 4.2.4 安装logstash



![image-20220407160405048](kubernetes日志收集方案 ELK.assets/image-20220407160405048.png)
![image-20220407160405048](https://github.com/user-attachments/assets/7fc565ce-d939-4ce4-b626-f0b0fcb197b7)



![image-20220407160443555](kubernetes日志收集方案 ELK.assets/image-20220407160443555.png)
![image-20220407160443555](https://github.com/user-attachments/assets/3bcd6187-ee6f-4574-b7cb-9938feae81ce)





~~~powershell
# wget https://artifacts.elastic.co/downloads/logstash/logstash-7.17.2-x86_64.rpm
~~~



~~~powershell
# yum -y install logstash-7.17.2-x86_64.rpm
~~~





## 4.3 软件配置及启动



### 4.3.1 kibana软件配置及启动



~~~powershell
[root@kibana ~]# cat -n /etc/kibana/kibana.yml | grep -v "#" | grep -v "^$"
     2  server.port: 5601
     7  server.host: "192.168.10.200"
    32  elasticsearch.hosts: ["http://192.168.10.201:9200"]
   115  i18n.locale: "zh-CN"
~~~



~~~powershell
说明：
server.port 是开启kibana监听端口
server.host 设置远程连接主机IP地址，用于远程访问使用
elasticsearch.hosts 设置elasticsearch.hosts主机IP，用于连接elasticsearch主机，可以为多个值
i18n.locale 设置语言支持，不需要再汉化，直接修改后即可支持中文
~~~



~~~powershell
[root@kibana ~]# systemctl enable kibana
[root@kibana ~]# systemctl start kibana
~~~



~~~powershell
[root@kibana ~]# ss -anput | grep ":5601"
tcp    LISTEN     0      128    192.168.10.200:5601                  *:*                   users:(("node",pid=2571,fd=71))
~~~





### 4.3.2 elasticsearch软件配置及启动



~~~powershell
修改配置文件
[root@elastic ~]#  cat -n /etc/elasticsearch/elasticsearch.yml | grep -v "#" | grep -v "^$"
    17  cluster.name: k8s-elastic
    23  node.name: elastic
    33  path.data: /var/lib/elasticsearch
    37  path.logs: /var/log/elasticsearch
    56  network.host: 192.168.10.201
    61  http.port: 9200
    70  discovery.seed_hosts: ["192.168.10.201"]
    74  cluster.initial_master_nodes: ["192.168.10.201"]
~~~



~~~powershell
说明
cluster.name 集群名称
node.name 节点名称
path.data 数据目录
path.logs 日志目录
network.host 主机IP
http.port 监听端口
discovery.seed_hosts 主机发现列表
cluster.initial_master_nodes 集群master节点
~~~





~~~powershell
启动服务并验证
[root@elastic ~]# systemctl enable elasticsearch
[root@elastic ~]# systemctl start elasticsearch
~~~



~~~powershell
[root@elastic ~]# ss -anput | grep ":9200"
tcp    LISTEN     0      128    [::ffff:192.168.10.201]:9200               [::]:*                   users:(("java",pid=9726,fd=219))
~~~





~~~powershell
[root@elastic ~]# curl http://192.168.10.201:9200
{
  "name" : "elastic",
  "cluster_name" : "k8s-elastic",
  "cluster_uuid" : "cW78ZkrhS4OV41DV5CtWWQ",
  "version" : {
    "number" : "7.17.2",
    "build_flavor" : "default",
    "build_type" : "rpm",
    "build_hash" : "de7261de50d90919ae53b0eff9413fd7e5307301",
    "build_date" : "2022-03-28T15:12:21.446567561Z",
    "build_snapshot" : false,
    "lucene_version" : "8.11.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
~~~





### 4.3.3 logstash软件配置及启动



#### 4.3.3.1 修改配置文件



~~~powershell
[root@logstash ~]# cat -n /etc/logstash/logstash.yml | grep -v "#" | grep -v "^$"
    19  node.name: logstash
    28  path.data: /var/lib/logstash
   133  api.http.host: 192.168.10.202
   139  api.http.port: 9600-9700
   280  path.logs: /var/log/logstash
   
分布式架构中 api.http.host一定要配置为logstash主机IP，不然无法远程访问。
~~~





#### 4.3.3.2 启动服务

**logstash进程不用预先启动，使用时启动即可**



#### 4.3.3.3 验证logstash可用性



~~~powershell
标准输入及标准输出验证
[root@logstash ~]# /usr/share/logstash/bin/logstash -e 'input {stdin{} } output {stdout {} }'
Using bundled JDK: /usr/share/logstash/jdk
OpenJDK 64-Bit Server VM warning: Option UseConcMarkSweepGC was deprecated in version 9.0 and will likely be removed in a future release.
WARNING: Could not find logstash.yml which is typically located in $LS_HOME/config or /etc/logstash. You can specify the path using --path.settings. Continuing using the defaults
Could not find log4j2 configuration at path /usr/share/logstash/config/log4j2.properties. Using default config which logs errors to the console
[INFO ] 2022-04-07 16:34:16.332 [main] runner - Starting Logstash {"logstash.version"=>"7.17.2", "jruby.version"=>"jruby 9.2.20.1 (2.5.8) 2021-11-30 2a2962fbd1 OpenJDK 64-Bit Server VM 11.0.14.1+1 on 11.0.14.1+1 +indy +jit [linux-x86_64]"}
[INFO ] 2022-04-07 16:34:16.339 [main] runner - JVM bootstrap flags: [-Xms1g, -Xmx1g, -XX:+UseConcMarkSweepGC, -XX:CMSInitiatingOccupancyFraction=75, -XX:+UseCMSInitiatingOccupancyOnly, -Djava.awt.headless=true, -Dfile.encoding=UTF-8, -Djruby.compile.invokedynamic=true, -Djruby.jit.threshold=0, -Djruby.regexp.interruptible=true, -XX:+HeapDumpOnOutOfMemoryError, -Djava.security.egd=file:/dev/urandom, -Dlog4j2.isThreadContextMapInheritable=true]
[INFO ] 2022-04-07 16:34:16.385 [main] settings - Creating directory {:setting=>"path.queue", :path=>"/usr/share/logstash/data/queue"}
[INFO ] 2022-04-07 16:34:16.423 [main] settings - Creating directory {:setting=>"path.dead_letter_queue", :path=>"/usr/share/logstash/data/dead_letter_queue"}
[WARN ] 2022-04-07 16:34:16.890 [LogStash::Runner] multilocal - Ignoring the 'pipelines.yml' file because modules or command line options are specified
[INFO ] 2022-04-07 16:34:16.956 [LogStash::Runner] agent - No persistent UUID file found. Generating new UUID {:uuid=>"608c9b46-8138-44b6-8cd8-c42d0fb08a90", :path=>"/usr/share/logstash/data/uuid"}
[INFO ] 2022-04-07 16:34:18.587 [Api Webserver] agent - Successfully started Logstash API endpoint {:port=>9600, :ssl_enabled=>false}
[INFO ] 2022-04-07 16:34:19.018 [Converge PipelineAction::Create<main>] Reflections - Reflections took 76 ms to scan 1 urls, producing 119 keys and 419 values
[WARN ] 2022-04-07 16:34:19.571 [Converge PipelineAction::Create<main>] line - Relying on default value of `pipeline.ecs_compatibility`, which may change in a future major release of Logstash. To avoid unexpected changes when upgrading Logstash, please explicitly declare your desired ECS Compatibility mode.
[WARN ] 2022-04-07 16:34:19.590 [Converge PipelineAction::Create<main>] stdin - Relying on default value of `pipeline.ecs_compatibility`, which may change in a future major release of Logstash. To avoid unexpected changes when upgrading Logstash, please explicitly declare your desired ECS Compatibility mode.
[INFO ] 2022-04-07 16:34:19.839 [[main]-pipeline-manager] javapipeline - Starting pipeline {:pipeline_id=>"main", "pipeline.workers"=>2, "pipeline.batch.size"=>125, "pipeline.batch.delay"=>50, "pipeline.max_inflight"=>250, "pipeline.sources"=>["config string"], :thread=>"#<Thread:0xc108bf4 run>"}
[INFO ] 2022-04-07 16:34:20.536 [[main]-pipeline-manager] javapipeline - Pipeline Java execution initialization time {"seconds"=>0.69}
WARNING: An illegal reflective access operation has occurred
WARNING: Illegal reflective access by com.jrubystdinchannel.StdinChannelLibrary$Reader (file:/usr/share/logstash/vendor/bundle/jruby/2.5.0/gems/jruby-stdin-channel-0.2.0-java/lib/jruby_stdin_channel/jruby_stdin_channel.jar) to field java.io.FilterInputStream.in
WARNING: Please consider reporting this to the maintainers of com.jrubystdinchannel.StdinChannelLibrary$Reader
WARNING: Use --illegal-access=warn to enable warnings of further illegal reflective access operations
WARNING: All illegal access operations will be denied in a future release
[INFO ] 2022-04-07 16:34:20.602 [[main]-pipeline-manager] javapipeline - Pipeline started {"pipeline.id"=>"main"}
The stdin plugin is now waiting for input:
[INFO ] 2022-04-07 16:34:20.662 [Agent thread] agent - Pipelines running {:count=>1, :running_pipelines=>[:main], :non_running_pipelines=>[]}
abc 输入abc字符，查看其输出
{
    "@timestamp" => 2022-04-07T08:35:24.663Z,
          "host" => "logstash",
       "message" => "abc",
      "@version" => "1"
}
以json格式输出abc内容
~~~





~~~powershell
使用logstash输入内容到elasticsearch验证

[root@logstash ~]# /usr/share/logstash/bin/logstash -e 'input { stdin{} } output { elasticsearch { hosts => ["192.168.10.201:9200"] index => "logstash-%{+YYYY.MM.dd}" } }'

hello elasticsearch

此内容将会通过kibana页面中的索引看到，但是需要在kibana页面中添加索引
~~~









## 4.4 kibana访问



![image-20220407163953277](https://github.com/user-attachments/assets/6d3ea6cb-dd19-4390-a26a-b73be95696ac)





![image-20220407165920127](https://github.com/user-attachments/assets/7662e903-4824-4696-a640-ad502207ebb4)





![image-20220407165959162](https://github.com/user-attachments/assets/0c55478f-6e47-401d-9c98-ea67fd0865dd)






![image-20220407170046439](https://github.com/user-attachments/assets/2a248d5c-93e2-4af9-bf4b-05b46b736951)



![image-20220407170142024](https://github.com/user-attachments/assets/bea20c50-4f53-455b-ba89-3a043edc635d)







![image-20220407172812976](https://github.com/user-attachments/assets/c434b3e8-c302-42fd-bfef-dd2243ff8379)






![image-20220407172535568](https://github.com/user-attachments/assets/ada59954-0532-48be-9fe8-43a9c9e6e899)







## 4.5 编写logstash用于收集日志配置文件

> 通过filebeat进行收集



~~~powershell
[root@logstash ~]# cat /etc/logstash/conf.d/logstash-to-elastic.conf
input {
  beats {
    host => "0.0.0.0"
    port => "5044"
  }
}

filter {

}


output {
    elasticsearch {
      hosts => "192.168.10.201:9200"
      index => "k8s-%{+YYYY.MM.dd}"
    }
}
~~~



## 4.6 运行logstash

> 如果不涉及多个配置文件，可以直接使用systemctl start logstash;如果有多个配置文件，只想启动一个配置文件，可以使用如下方法。



### 4.6.1 直接在后台运行

~~~powershell
[root@logstash ~]# /usr/share/logstash/bin/logstash -f /etc/logstash/conf.d/logstash-to-elastic.conf --path.data /usr/share/logstash/data1 &
~~~



### 4.6.2 通过rc.local设置自动后台运行



~~~powershell
[root@logstash ~]# cat /etc/rc.local
...
/usr/share/logstash/bin/logstash -f /etc/logstash/conf.d/logstash-to-elastic.conf &

查看文件默认权限
[root@logstash ~]# ls -l /etc/rc.d/rc.local
-rw-r--r-- 1 root root 562 1月   9 13:40 /etc/rc.d/rc.local

修改文件权限
[root@logstash ~]# chmod +x  /etc/rc.d/rc.local

查看修改后文件权限
[root@logstash ~]# ls -l /etc/rc.d/rc.local
-rwxr-xr-x 1 root root 562 1月   9 13:40 /etc/rc.d/rc.local
~~~



# 五、收集k8s集群节点系统日志

>通过在work节点以DaemonSet方法运行filebeat应用实现

## 5.1 下载filebeat镜像



> 所有work节点




![image-20220407175113144](https://github.com/user-attachments/assets/4a965d31-5f9b-4315-a40b-f0b0bf3ca91d)





![image-20220407175159994](https://github.com/user-attachments/assets/1ddc4d86-b456-4c59-9372-40fa8720dc70)



~~~powershell
下载filebeat镜像

[root@k8s-work1 ~]# docker pull elastic/filebeat:7.17.2
~~~



~~~powershell
[root@k8s-work1 ~]# docker images
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
docker.elastic/filebeat   7.17.2               00c5b17745d1        3 weeks ago         359MB

~~~



或



~~~powershell
使用containerd时使用
# crictl pull elastic/filebeat:7.17.2
~~~



~~~powershell
# crictl images
IMAGE                                           TAG                 IMAGE ID            SIZE
docker.io/elastic/filebeat                      7.17.2              2314640a78873       107MB
~~~





## 5.2 创建filebeat资源清单文件



~~~powershell
[root@k8s-master1 ~]# cat filebeat-to-logstash.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: k8s-logs-filebeat-config
  namespace: kube-system

data:
  filebeat.yml: |
    filebeat.inputs:
      - type: log
        paths:
          - /var/log/messages
        fields:
          app: k8s
          type: module
        fields_under_root: true

    setup.ilm.enabled: false
    setup.template.name: "k8s-module"
    setup.template.pattern: "k8s-module-*"

    output.logstash:
      hosts: ['192.168.10.202:5044']
      index: "k8s-module-%{+yyyy.MM.dd}"

---

apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: k8s-logs
  namespace: kube-system
spec:
  selector:
    matchLabels:
      project: k8s
      app: filebeat
  template:
    metadata:
      labels:
        project: k8s
        app: filebeat
    spec:
      containers:
      - name: filebeat
        image: docker.io/elastic/filebeat:7.17.2
        args: [
          "-c", "/etc/filebeat.yml",
          "-e",
        ]
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
          limits:
            cpu: 500m
            memory: 500Mi
        securityContext:
          runAsUser: 0
        volumeMounts:
        - name: filebeat-config
          mountPath: /etc/filebeat.yml
          subPath: filebeat.yml
        - name: k8s-logs
          mountPath: /var/log/messages
      volumes:
      - name: k8s-logs
        hostPath:
          path: /var/log/messages
      - name: filebeat-config
        configMap:
          name: k8s-logs-filebeat-config
~~~





## 5.3 应用filebeat资源清单文件

~~~powershell
[root@k8s-master1 ~]# kubectl apply -f filebeat-to-logstash.yaml
~~~



## 5.4 验证结果



~~~powershell
查看pod
# kubectl get pods -n kube-system -o wide
NAME                                       READY   STATUS    RESTARTS   AGE   IP               NODE         
k8s-logs-s8qw6                             1/1     Running   0          15s   10.244.194.83    k8s-worker1   <none>           <none>
~~~



~~~powershell
查看pod输出日志
[root@k8s-master1 ~]# kubectl logs k8s-logs-6mqq5 -n kube-system
~~~



## 5.5 在kibana中添加索引



![image-20220407182931857](https://github.com/user-attachments/assets/9f3f3fe0-6399-4d98-b3d8-39b846c2c3d9)



![image-20220407203219708](https://github.com/user-attachments/assets/75a97182-d667-4f60-a95d-a6619a6c1c77)



![image-20220407203312600](https://github.com/user-attachments/assets/a20acb1d-4859-496f-8f30-15c97b97d03a)



![image-20220407203403595](https://github.com/user-attachments/assets/7ec29c60-9868-4c35-a3d1-ca8196fed916)



![image-20220407203454587](https://github.com/user-attachments/assets/9afe36b6-29ee-42ec-9dad-bf66f3e91836)



![image-20220407203546362](https://github.com/user-attachments/assets/2ab7a9ac-18a7-4e26-9729-7d65ecdf372c)





# 六、收集kubernetes节点应用程序日志

> 本案例在k8s-worker1主机上安装nginx并收集其日志



## 6.1 安装nginx应用



~~~powershell
# wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo
~~~



~~~powershell
# yum -y install nginx
~~~



~~~powershell
# cd /usr/share/nginx/html/

# ls
404.html  50x.html  en-US  icons  img  index.html  nginx-logo.png  poweredby.png

# echo "work1 web page" > index.html
~~~





~~~powershell

# systemctl enable nginx
# systemctl start nginx
~~~



~~~powershell
# curl http://192.168.10.15
work1 web page
~~~





## 6.2 编写filebeat资源清单文件



~~~powershell
[root@k8s-master1 ~]# cat filebeat-to-logstash-nginx.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: k8s-filebeat-config-nginx-logs
  namespace: default

data:
  filebeat.yml: |
    filebeat.inputs:
      - type: log
        paths:
          - /var/log/nginx/access.log
        fields:
          app: k8s
          type: module
        fields_under_root: true

      - type: log
        paths:
          - /var/log/nginx/error.log
        fields:
          app: k8s
          type: module
        fields_under_root: true

    setup.ilm.enabled: false
    setup.template.name: "k8s-module"
    setup.template.pattern: "k8s-module-*"

    output.logstash:
      hosts: ['192.168.10.202:5055']

---

apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: k8s-logs
  namespace: default
spec:
  selector:
    matchLabels:
      project: k8s
      app: filebeat
  template:
    metadata:
      labels:
        project: k8s
        app: filebeat
    spec:
      nodeName: k8s-worker1
      containers:
      - name: filebeat
        image: docker.io/elastic/filebeat:7.17.2
        imagePullPolicy: IfNotPresent
        args: [
          "-c", "/etc/filebeat.yml",
          "-e",
        ]
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
          limits:
            cpu: 500m
            memory: 500Mi
        securityContext:
          runAsUser: 0
        volumeMounts:
        - name: filebeat-config
          mountPath: /etc/filebeat.yml
          subPath: filebeat.yml
        - name: nginx-access
          mountPath: /var/log/nginx/access.log
        - name: nginx-error
          mountPath: /var/log/nginx/error.log
      volumes:
      - name: nginx-access
        hostPath:
          path: /var/log/nginx/access.log
      - name: nginx-error
        hostPath:
          path: /var/log/nginx/error.log
      - name: filebeat-config
        configMap:
          name: k8s-filebeat-config-nginx-logs
~~~





## 6.3 编写logstash配置文件



~~~powershell
[root@logstash ~]# cat /etc/logstash/conf.d/nginx-logstash-to-elastic.conf
input {
  beats {
    host => "0.0.0.0"
    port => "5055"
  }
}

filter {

}


output {
    elasticsearch {
      hosts => "192.168.10.201:9200"
      index => "nginx-%{+YYYY.MM.dd}"
    }
}
~~~



## 6.4 重启logstash



~~~powershell
[root@logstash ~]# /usr/share/logstash/bin/logstash -f /etc/logstash/conf.d/nginx-logstash-to-elastic.conf --path.data /usr/share/logstash/data2 &
~~~



~~~powershell
[root@logstash ~]# ss -anput | grep ":5055"
tcp    LISTEN     0      128    [::]:5055               [::]:*                   users:(("java",pid=14296,fd=106))
~~~





## 6.5 应用filebeat资源清单文件



~~~powershell
[root@k8s-master1 ~]# kubectl apply -f filebeat-to-logstash-nginx.yaml
configmap/k8s-filebeat-config-nginx-logs created
daemonset.apps/k8s-logs created
~~~





~~~powershell
[root@k8s-master1 ~]# kubectl get pods -o wide
NAME             READY   STATUS    RESTARTS   AGE   IP              NODE          NOMINATED NODE   READINESS GATES
k8s-logs-ndznb   1/1     Running   0          14s   10.244.194.84   k8s-worker1   <none>           <none>
~~~





~~~powershell
[root@k8s-master1 ~]# # kubectl logs k8s-logs-ndznb
~~~





## 6.6 在kibana界面添加索引



![image-20220407223459697](kubernetes日志收集方案 ELK.assets/image-20220407223459697.png)
![image-20220407223459697](https://github.com/user-attachments/assets/ffb2cd37-9ccc-41e9-8634-ea9ad324aa5a)



![image-20220407223538269](kubernetes日志收集方案 ELK.assets/image-20220407223538269.png)
![image-20220407223538269](https://github.com/user-attachments/assets/ebc01de4-8710-4053-b305-1471a7f4d21c)





![image-20220407223613198](kubernetes日志收集方案 ELK.assets/image-20220407223613198.png)
![image-20220407223613198](https://github.com/user-attachments/assets/c7665045-c5a0-4759-98d2-e458bb01f022)



![image-20220407223651714](kubernetes日志收集方案 ELK.assets/image-20220407223651714.png)
![image-20220407223651714](https://github.com/user-attachments/assets/02ca8e93-d753-4585-8ef2-4d2f17a23803)





![image-20220407223727384](kubernetes日志收集方案 ELK.assets/image-20220407223727384.png)
![image-20220407223727384](https://github.com/user-attachments/assets/c755ab05-8661-41d0-baa1-1ad713fe8f46)



![image-20220407223838150](kubernetes日志收集方案 ELK.assets/image-20220407223838150.png)
![image-20220407223838150](https://github.com/user-attachments/assets/b1172c4b-32bb-45cf-a9d3-beece54f26e1)







# 七、收集kubernetes集群中以Pod方式运行的应用日志

> 通过在应用程序Pod中运行filebeat(sidecar)实现,本次将以tomcat为例进行说明。



## 7.1 准备tomcat数据目录

> 默认tomcat容器中没有网站首页文件，不添加会导致pod中容器无法正常运行。



~~~powershell
[root@k8s-worker1 ~]# mkdir /opt/tomcatwebroot
~~~



~~~powershell
[root@k8s-worker1 ~]# echo "tomcat running" > /opt/tomcatwebroot/index.html
~~~







## 7.2 编写tomcat应用资源清单文件



~~~powershell
[root@k8s-master1 ~]# cat tomcat-logs.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tomcat-demo
  namespace: default
spec:
  replicas: 2
  selector:
    matchLabels:
      project: www
      app: tomcat-demo
  template:
    metadata:
      labels:
        project: www
        app: tomcat-demo
    spec:
      nodeName: k8s-worker1
      containers:
      - name: tomcat
        image: tomcat:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8080
          name: web
          protocol: TCP
        resources:
          requests:
            cpu: 0.5
            memory: 1Gi
          limits:
            cpu: 1
            memory: 2Gi
        livenessProbe:
          httpGet:
            path: /
            port: 8080
          initialDelaySeconds: 60
          timeoutSeconds: 20
        readinessProbe:
          httpGet:
            path: /
            port: 8080
          initialDelaySeconds: 60
          timeoutSeconds: 20
        volumeMounts:
        - name: tomcat-logs
          mountPath: /usr/local/tomcat/logs
        - name: tomcatwebroot
          mountPath: /usr/local/tomcat/webapps/ROOT

      - name: filebeat
        image: docker.io/elastic/filebeat:7.17.2
        imagePullPolicy: IfNotPresent
        args: [
          "-c", "/etc/filebeat.yml",
          "-e",
        ]
        resources:
          limits:
            memory: 500Mi
          requests:
            cpu: 100m
            memory: 100Mi
        securityContext:
          runAsUser: 0
        volumeMounts:
        - name: filebeat-config
          mountPath: /etc/filebeat.yml
          subPath: filebeat.yml
        - name: tomcat-logs
          mountPath: /usr/local/tomcat/logs
      volumes:
      - name: tomcat-logs
        emptyDir: {}
      - name: tomcatwebroot
        hostPath:
          path: /opt/tomcatwebroot
          type: Directory
      - name: filebeat-config
        configMap:
          name: filebeat-config
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: filebeat-config
  namespace: default

data:
  filebeat.yml: |-
    filebeat.inputs:
    - type: log
      paths:
        - /usr/local/tomcat/logs/catalina.*

      fields:
        app: www
        type: tomcat-catalina
      fields_under_root: true
      multiline:
        pattern: '^\['
        negate: true
        match: after

    setup.ilm.enabled: false
    setup.template.name: "tomcat-catalina"
    setup.template.pattern: "tomcat-catalina-*"

    output.logstash:
      hosts: ['192.168.10.202:5056']
~~~







## 7.3  编写logstash配置文件



~~~powershell
编写logstash配置文件，不影响以往配置文件
[root@logstash ~]# cat /etc/logstash/conf.d/tomcat-logstash-to-elastic.conf
input {
  beats {
    host => "0.0.0.0"
    port => "5056"
  }
}

filter {

}


output {
    elasticsearch {
      hosts => "192.168.10.201:9200"
      index => "tomcat-catalina-%{+yyyy.MM.dd}"
    }
}
~~~



~~~powershell
[root@logstash ~]# /usr/share/logstash/bin/logstash -f /etc/logstash/conf.d/tomcat-logstash-to-elastic.conf --path.data /usr/share/logstash/data3 &
~~~



~~~powershell
验证端口是否启动
[root@logstash ~]# ss -anput | grep ":5056"
tcp    LISTEN     0      128    [::]:5056               [::]:*                   users:(("java",pid=14144,fd=106))
~~~





## 7.4 应用tomcat应用资源清单文件



~~~powershell
[root@k8s-master1 ~]# kubectl apply -f tomcat-logs.yaml

[root@k8s-master1 ~]# kubectl get deployment.apps
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
tomcat-demo   2/2     2            2           5m26s
[root@k8s-master1 ~]# kubectl get pods
NAME                           READY   STATUS    RESTARTS   AGE
tomcat-demo-664584f857-k8whd   2/2     Running   0          5m33s
tomcat-demo-664584f857-xncpk   2/2     Running   0          5m33s
~~~





## 7.5 验证Pod 中tomcat及filebeat是否正常



~~~powershell
查看tomcat产生日志
[root@k8s-master1 ~]# kubectl logs tomcat-demo-664584f857-k8whd -c tomcat

查看filebeat收集日志
[root@k8s-master1 ~]# kubectl logs tomcat-demo-664584f857-k8whd -c filebeat
~~~







## 7.6 在kibana页面中添加索引



![image-20220408002236127](https://github.com/user-attachments/assets/dd75e2b5-1bdf-4822-9e3d-c9e377c1b31d)


![image-20220408002721389](https://github.com/user-attachments/assets/f72f0a26-c0cb-4028-a2d2-1f1201dfd12e)




![image-20220408003250140](https://github.com/user-attachments/assets/d13a5f76-3c24-42b8-84e5-1b45ccc8bfcd)








![image-20220408003550821](https://github.com/user-attachments/assets/db4e5528-6f99-43ce-a7f0-e254a762e4db)


![image-20220408003610470](https://github.com/user-attachments/assets/b8daf800-d119-4b72-8119-ba567260bc85)




![image-20220408003632515](https://github.com/user-attachments/assets/9dd2ae76-195a-4598-8f03-ad6b99e2a3c0)



![image-20220408003733569](https://github.com/user-attachments/assets/5f0ff8c3-3bdc-4cfc-b1d8-b525c308541e)





