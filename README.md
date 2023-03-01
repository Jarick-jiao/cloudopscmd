---
author: cloudopscmd
date: 2022-12-22
title: 'cloudopscmd'
weight: 100 #Defaults 100
bookFlatSection: false    #级联缩进，适用目录
bookCollapseSection: true #折叠目录
disable_comments: true
categories: ['Notepad']  #1.产品
tag: ['']                #2.技术
Series: ['']             #3.
keywords: ['']           #4.
summary: ''              #5.描述
---
  
<!--more-->
[TOC]

## cloudopscmd

```shell
CloudDevKubernetes.com: Display The Kubernetes Messages of Kinds Resources.
More Information At: https://kubernetes.io/docs/reference/kubectl/overview/.

Usage:
  cloudopscmd [flags]
  cloudopscmd [command]

Available Commands:
  help        Help about any command
  kubedeploy  Display Deployment And Statefulset Resources Info.
  kubediags   Kubediags For Diagnostic Kubernetes Issues And Risks.
  kubens      Display Deployment And Statefulset Resources By The Namespaces.
  kuberes     Display Node Resources And Metric For Usage.
  kubetop     Display POD/Container CPU/MEM Resources And Usage Info.
  kubetrace   kubetrace For Diags Networks of Kubernetes Clusters Business Application.
  version     The CloudDevKubernetes.com Information.

Flags:
      --configDir string   The Genconfig Dir Path, If Using The Config File Needs To Be Moved To The Current Path(./config)
      --debug              Debug mode
      --des                Output Parameter Description Info
      --doc                Generated Using The Cloudopscmd Document
      --docDir string      Point The Cloudopscmd Document Dir Path
      --genconfig          Generated Sing The Cloudopscmd Config File
  -h, --help               help for cloudopscmd
      --json               Output In JSON Format, Default Table Format
  -n, --namespace string   If Present, The Namespace Scope For This CLI Request
      --rmitem strings     Remove Data Item. Only One Type Is Allowed To Be Specified

Use "cloudopscmd [command] --help" for more information about a command.
```

### 1. cloudopscmd version

展示版本信息

```shell
Version:        2.0
CommitId:       1a9286d42f7e361c3430de01fcdcb29eaefdffc2
GitBranch:      CloudDevKubernetes.com/dev
GitAuthor:      ******
Build Date:     2022.12.20.180702
Go Version:     go1.18.3
OS/Arch:        linux/amd64
```

### 2. cloudopscmd kuberes

展示节点资源使用情况

```shell
+-----------------+--------+--------+--------+------------+-------------+----------------+-------------+------------+-----------+
|NodeName         |PodCal  |CapMem  |CapCPU  |RequestCPU  |LimitCPU     |RequestMem      |LimitMem     |NodeTaints  |ImageSize  |
+-----------------+--------+--------+--------+------------+-------------+----------------+-------------+------------+-----------+
|192.168.128.115  |43(0)   |62.76   |32      |7673(0.24)  |30350(3.96)  |13150(0.2)      |34136(2.6)   |success     |11.28      |
|192.168.128.122  |67(1)   |62.76   |32      |8012(0.25)  |58718(7.33)  |12345.68(0.19)  |78528(6.36)  |success     |11.41      |
|192.168.128.49   |33(0)   |62.76   |32      |5263(0.16)  |31850(6.05)  |21876(0.34)     |45270(2.07)  |success     |9.64       |
|192.168.129.56   |23(0)   |15.51   |8       |4536(0.57)  |22650(4.99)  |6716(0.42)      |23392(3.48)  |success     |6.97       |
|192.168.254.87   |87(1)   |109.91  |56      |22366(0.4)  |70118(3.14)  |35104.68(0.31)  |98584(2.81)  |success     |14.56      |
+-----------------+--------+--------+--------+------------+-------------+----------------+-------------+------------+-----------+
```

#### 2.1. cloudopscmd kuberes -u

展示节点资源使用情况(resources)和资源使用率（top）

```shell
+-----------------+--------+--------+--------+------------+-------------+-------------+----------------+-------------+----------------+
|NodeName         |PodCal  |CapMem  |CapCPU  |RequestCPU  |LimitCPU     |UsageCPU     |RequestMem      |LimitMem     |UsageMem        |
+-----------------+--------+--------+--------+------------+-------------+-------------+----------------+-------------+----------------+
|192.168.128.115  |43(0)   |62.76   |32      |7673(0.24)  |30350(3.96)  |2319(0.07)   |13150(0.2)      |34136(2.6)   |9721.29(0.15)   |
|192.168.128.122  |67(1)   |62.76   |32      |8012(0.25)  |58718(7.33)  |3940(0.12)   |12345.68(0.19)  |78528(6.36)  |13386.42(0.21)  |
|192.168.128.49   |33(0)   |62.76   |32      |5263(0.16)  |31850(6.05)  |11389(0.36)  |21876(0.34)     |45270(2.07)  |12304.2(0.19)   |
|192.168.129.56   |23(0)   |15.51   |8       |4536(0.57)  |22650(4.99)  |665(0.08)    |6716(0.42)      |23392(3.48)  |4950.65(0.31)   |
|192.168.254.87   |87(1)   |109.91  |56      |22366(0.4)  |70118(3.14)  |3272(0.06)   |35104.68(0.31)  |98584(2.81)  |32118.27(0.29)  |
+-----------------+--------+--------+--------+------------+-------------+-------------+----------------+-------------+----------------+
```

#### 2.2. cloudopscmd kuberes -e

仅展示错误记录

```shell
+-----------------+--------+--------+--------+------------+-------------+----------------+-------------+------------+-----------+
|NodeName         |PodCal  |CapMem  |CapCPU  |RequestCPU  |LimitCPU     |RequestMem      |LimitMem     |NodeTaints  |ImageSize  |
+-----------------+--------+--------+--------+------------+-------------+----------------+-------------+------------+-----------+
|192.168.128.122  |67(1)   |62.76   |32      |8012(0.25)  |58718(7.33)  |12345.68(0.19)  |78528(6.36)  |success     |11.41      |
|192.168.254.87   |87(1)   |109.91  |56      |22366(0.4)  |70118(3.14)  |35104.68(0.31)  |98584(2.81)  |success     |14.56      |
+-----------------+--------+--------+--------+------------+-------------+----------------+-------------+------------+-----------+
```

### 3. cloudopscmd kubens

展示命名空间基本资源使用情况

```shell
+-------------------+-----------+-----+----------+------------+------------+----------+----------+
|Namespace          |Container  |Pod  |ErrCount  |RequestCPU  |RequestMem  |LimitCPU  |LimitMem  |
+-------------------+-----------+-----+----------+------------+------------+----------+----------+
|cert-manager       |3          |3    |          |70          |166         |1500      |3072      |
|cosi               |2          |2    |          |100         |20          |100       |30        |
|xxxx-system       |91         |79   |          |8846        |25725       |88636     |152952    |
|demo               |29         |22   |          |8600        |12544       |8600      |12544     |
|kube-system        |7          |7    |          |900         |740         |4200      |5660      |
|kubevirt           |9          |9    |          |80          |1194        |0         |0         |
|nativestor-system  |19         |5    |          |4150        |3066        |7100      |7830      |
|operators          |18         |14   |          |3274        |3622        |7200      |12364     |
|rook-ceph          |42         |22   |          |4500        |13824       |2000      |3584      |
+-------------------+-----------+-----+----------+------------+------------+----------+----------+
Display Deployment And Statefulset Resources By The Namespaces.
[Note]: Unit: cpu(m), mem(Mi), ImageSize(G), CapMem(Gi), CapCPU(c), Replices[replices(err)], PodCal[pod(err)]
[PrintTime]:2022-12-22 11:23:35
```

#### 3.1. cloudopscmd kubens -d

展示命名空间基本资源使用情况（Deployment）

```shell
+-------------------+-----------+-----+----------+------------+------------+----------+----------+
|Namespace          |Container  |Pod  |ErrCount  |RequestCPU  |RequestMem  |LimitCPU  |LimitMem  |
+-------------------+-----------+-----+----------+------------+------------+----------+----------+
|cert-manager       |3          |3    |          |70          |166         |1500      |3072      |
|cosi               |2          |2    |          |100         |20          |100       |30        |
|xxxx-system       |78         |71   |          |8134        |12105       |84736     |139952    |
|demo               |13         |12   |          |4300        |3380        |4300      |3380      |
|kube-system        |5          |5    |          |900         |740         |4200      |5660      |
|kubevirt           |9          |9    |          |80          |1194        |0         |0         |
|nativestor-system  |19         |5    |          |4150        |3066        |7100      |7830      |
|operators          |18         |14   |          |3274        |3622        |7200      |12364     |
|rook-ceph          |42         |22   |          |4500        |13824       |2000      |3584      |
+-------------------+-----------+-----+----------+------------+------------+----------+----------+
```

#### 3.2. cloudopscmd kubens -f

展示命名空间基本资源使用情况（Statefulset）

```shell
+--------------+-----------+-----+----------+------------+------------+----------+----------+
|Namespace     |Container  |Pod  |ErrCount  |RequestCPU  |RequestMem  |LimitCPU  |LimitMem  |
+--------------+-----------+-----+----------+------------+------------+----------+----------+
|xxxx-system  |13         |8    |          |712         |13620       |3900      |13000     |
|demo          |16         |10   |          |4300        |9164        |4300      |9164      |
|kube-system   |2          |2    |          |0           |0           |0         |0         |
+--------------+-----------+-----+----------+------------+------------+----------+----------+
```

### 4. cloudopscmd kubedeploy --des

展示部署组,POD,Container资源使用情况

```shell
+-------------+-------------+---------------------+---------------------+----------+------------+------------+----------+-------------+------------------------------+
|NameSpace    |TYPE         |ResourceName         |ContainerName        |Replices  |RequestCPU  |RequestMem  |LimitCPU  |LimitMem     |IMAGE                         |
+-------------+-------------+---------------------+---------------------+----------+------------+------------+----------+-------------+------------------------------+
|kube-system  |deployment   |coredns              |coredns              |2(0)      |100         |70          |0(0)      |170(2.43)    |coredns:1.7.0                 |
|kube-system  |deployment   |kube-ovn-controller  |kube-ovn-controller  |1(0)      |200         |200         |1000(5)   |1024(5.12)   |kube-ovn:v1.8.9               |
|kube-system  |deployment   |kube-ovn-monitor     |kube-ovn-monitor     |1(0)      |200         |200         |200(1)    |200(1)       |kube-ovn:v1.8.9               |
|kube-system  |deployment   |ovn-central          |ovn-central          |1(0)      |300         |200         |3000(10)  |4096(20.48)  |kube-ovn:v1.8.9               |
|kube-system  |statefulset  |snapshot-controller  |snapshot-controller  |2(0)      |0           |0           |0(NaN)    |0(NaN)       |snapshot-controller:v3.0.2.1  |
+-------------+-------------+---------------------+---------------------+----------+------------+------------+----------+-------------+------------------------------+
Display Deployment And Statefulset Resources Info.
[Note]: Unit: cpu(m), mem(Mi), ImageSize(G), CapMem(Gi), CapCPU(c), Replices[replices(err)], PodCal[pod(err)]
```

#### 4.1. cloudopscmd kubedeploy  -n cert-manager  -d

指定命名空间展示Deployment部署资源,POD,Container资源使用情况

```shell
+--------------+------------+-------------------------+---------------+----------+------------+------------+------------+-------------+--------------------------------+
|NameSpace     |TYPE        |ResourceName             |ContainerName  |Replices  |RequestCPU  |RequestMem  |LimitCPU    |LimitMem     |IMAGE                           |
+--------------+------------+-------------------------+---------------+----------+------------+------------+------------+-------------+--------------------------------+
|cert-manager  |deployment  |cert-manager             |cert-manager   |1(0)      |40          |64          |500(12.5)   |1024(16)     |cert-manager-controller:v1.4.0  |
|cert-manager  |deployment  |cert-manager-cainjector  |cainjector     |1(0)      |15          |82          |500(33.33)  |1024(12.49)  |cert-manager-cainjector:v1.4.0  |
|cert-manager  |deployment  |cert-manager-webhook     |webhook        |1(0)      |15          |20          |500(33.33)  |1024(51.2)   |cert-manager-webhook:v1.4.0     |
+--------------+------------+-------------------------+---------------+----------+------------+------------+------------+-------------+--------------------------------+
```

#### 4.2. cloudopscmd kubedeploy  -n xxxx-system  -f

指定命名空间展示Statefulset部署资源,POD,Container资源使用情况

```shell
+--------------+-------------+------------------------------+-----------------+----------+------------+------------+----------+-------------+--------------------------------------------+
|NameSpace     |TYPE         |ResourceName                  |ContainerName    |Replices  |RequestCPU  |RequestMem  |LimitCPU  |LimitMem     |IMAGE                                       |
+--------------+-------------+------------------------------+-----------------+----------+------------+------------+----------+-------------+--------------------------------------------+
|xxxx-system  |statefulset  |alertmanager-kube-prometheus  |alertmanager     |1(0)      |10          |64          |500(50)   |500(7.81)    |alertmanager:v0.23.0-v3.8.24                |
|xxxx-system  |statefulset  |alertmanager-kube-prometheus  |config-reloader  |1(0)      |100         |50          |100(1)    |50(1)        |prometheus-config-reloader:v0.52.0-v3.8.24  |
|xxxx-system  |statefulset  |alertmanager-kube-prometheus  |proxy            |1(0)      |1           |20          |100(100)  |200(10)      |oauth2-proxy:v7.1.3-v3.8.24                 |
|xxxx-system  |statefulset  |minio                         |minio            |6(0)      |0           |2048        |0(NaN)    |0(0)         |minio:v3.8.12                               |
|xxxx-system  |statefulset  |prometheus-kube-prometheus-0  |prometheus       |1(0)      |400         |1000        |2000(5)   |10000(10)    |prometheus:v2.29.2-v3.8.24                  |
|xxxx-system  |statefulset  |prometheus-kube-prometheus-0  |config-reloader  |1(0)      |100         |50          |100(1)    |50(1)        |prometheus-config-reloader:v0.52.0-v3.8.24  |
|xxxx-system  |statefulset  |prometheus-kube-prometheus-0  |thanos-sidecar   |1(0)      |100         |128         |1000(10)  |2000(15.62)  |thanos:v0.17.1-v3.8.24                      |
|xxxx-system  |statefulset  |prometheus-kube-prometheus-0  |proxy            |1(0)      |1           |20          |100(100)  |200(10)      |oauth2-proxy:v7.1.3-v3.8.24                 |
+--------------+-------------+------------------------------+-----------------+----------+------------+------------+----------+-------------+--------------------------------------------+
```

#### 4.3. cloudopscmd kubedeploy -e

仅展示错误记录

```shell
+--------------+------------+------------------------+------------------------+----------+------------+------------+----------+----------+-------------------------------+
|NameSpace     |TYPE        |ResourceName            |ContainerName           |Replices  |RequestCPU  |RequestMem  |LimitCPU  |LimitMem  |IMAGE                          |
+--------------+------------+------------------------+------------------------+----------+------------+------------+----------+----------+-------------------------------+
|xxxx-system  |deployment  |nfs-client-provisioner  |nfs-client-provisioner  |1(1)      |0           |0           |0(NaN)    |0(NaN)    |nfs-client-provisioner:v3.8.3  |
+--------------+------------+------------------------+------------------------+----------+------------+------------+----------+----------+-------------------------------+
```

### 5. cloudopscmd kubetop  --des

展示组件资源(POD)使用量和使用率（top）

```shell
+----+------+------------------------------------------+--------------+-----------------+------------+----------+------------+----------+----------+--------------+
|ID  |Type  |ResourceName                              |Namespace     |NodeName         |RequestCPU  |LimitCPU  |RequestMem  |LimitMem  |UsageCPU  |UsageMem      |
+----+------+------------------------------------------+--------------+-----------------+------------+----------+------------+----------+----------+--------------+
|0   |POD   |cert-manager-747584bd86-2drtq             |cert-manager  |192.168.128.122  |40          |500       |64          |1024      |2(-0.08)  |45.89(-0.02)  |
|1   |POD   |cert-manager-cainjector-6f6b878cbb-f2f56  |cert-manager  |192.168.128.122  |15          |500       |82          |1024      |9(-0.01)  |135.8(0.06)   |
|2   |POD   |cert-manager-webhook-7d9d78689d-5jx9l     |cert-manager  |192.168.128.122  |15          |500       |20          |1024      |3(-0.02)  |22.73(0)      |
+----+------+------------------------------------------+--------------+-----------------+------------+----------+------------+----------+----------+--------------+
Display POD/Container CPU/MEM Resources And Usage Info.
[Note]: metric(%)=(UsageCPU - RequestCPU）/（LimitCPU - RequestCPU), (Limit == 0 && Request == 0 || Request > Limit) = 100
[Note]: Unit: cpu(m), mem(Mi), ImageSize(G), CapMem(Gi), CapCPU(c), Replices[replices(err)], PodCal[pod(err)]
```

#### 5.1. cloudopscmd kubetop -n cert-manager --containers

展示组件资源(Container)使用量和使用率（top）

```shell
+----+-----------+------------------------------------------+--------------+-----------------+---------------+------------+----------+------------+----------+----------+--------------+
|ID  |Type       |ResourceName                              |Namespace     |NodeName         |ContainerName  |RequestCPU  |LimitCPU  |RequestMem  |LimitMem  |UsageCPU  |UsageMem      |
+----+-----------+------------------------------------------+--------------+-----------------+---------------+------------+----------+------------+----------+----------+--------------+
|0   |container  |cert-manager-747584bd86-2drtq             |cert-manager  |192.168.128.122  |cert-manager   |40          |500       |64          |1024      |1(-0.08)  |45.89(-0.02)  |
|1   |container  |cert-manager-cainjector-6f6b878cbb-f2f56  |cert-manager  |192.168.128.122  |cainjector     |15          |500       |82          |1024      |8(-0.01)  |135.8(0.06)   |
|2   |container  |cert-manager-webhook-7d9d78689d-5jx9l     |cert-manager  |192.168.128.122  |webhook        |15          |500       |20          |1024      |3(-0.02)  |22.66(0)      |
+----+-----------+------------------------------------------+--------------+-----------------+---------------+------------+----------+------------+----------+----------+--------------+
```

### 6. cloudopscmd kubediags

定义规则做巡检

```shell
+----+---------------------------+-----------------------------------------------------------------------------+--------------+----------+--------------------+
|ID  |Case                       |Point                                                                        |CurrentState  |Severity  |Describe            |
+----+---------------------------+-----------------------------------------------------------------------------+--------------+----------+--------------------+
|0   |CPULimitInfo_3             |deployment/nativestor-system/topolvm-controller/liveness-prometheus          |100(4)        |提示      |容器CPU超售高于:3   |
|1   |MEMLimitInfo_3             |deployment/nativestor-system/topolvm-controller/liveness-prometheus          |256(5.12)     |提示      |容器内存超售高于:3  |
|2   |MEMLimitWarning_5          |deployment/nativestor-system/topolvm-controller/liveness-prometheus          |256(5.12)     |预警      |容器内存超售高于:5  |
|33  |MEMOversoldNodeInfo_2      |node/192.168.128.49                                                          |45270(2.07)   |提示      |节点内存超售高于:2  |
|34  |MEMOversoldNodeInfo_2      |node/192.168.129.56                                                          |23392(3.48)   |提示      |节点内存超售高于:2  |
|35  |MEMOversoldNodeInfo_2      |node/192.168.254.87                                                          |98584(2.81)   |提示      |节点内存超售高于:2  |
|36  |PODCapacityNodeWarning_80  |node/192.168.254.87                                                          |87(1)         |预警      |节点POD数高于:80    |
+----+---------------------------+-----------------------------------------------------------------------------+--------------+----------+--------------------+
```

### 7. cloudopscmd kubetrace -n cert-manager

组件端口探测，指定命名空间

```shell
+-----------+--------------------------+--------------+---------------------------------------+------------------+-----------------+-------------+
|TraceName  |Name                      |Namespace     |PodName                                |PortList          |NodeName         |Status       |
+-----------+--------------------------+--------------+---------------------------------------+------------------+-----------------+-------------+
|service    |cert-manager-new          |cert-manager  |null                                   |10.4.207.49:9402  |null             |pass[0.00s]  |
|service    |cert-manager-webhook-new  |cert-manager  |null                                   |10.4.221.58:443   |null             |pass[0.00s]  |
|endpoint   |cert-manager-new          |cert-manager  |cert-manager-747584bd86-2drtq          |10.3.0.10:9402    |192.168.128.122  |pass[0.00s]  |
|endpoint   |cert-manager-webhook-new  |cert-manager  |cert-manager-webhook-7d9d78689d-5jx9l  |10.3.0.12:10250   |192.168.128.122  |pass[0.00s]  |
+-----------+--------------------------+--------------+---------------------------------------+------------------+-----------------+-------------+
```

#### 7.1. cloudopscmd kubetrace -n xxxx-system -e

组件端口探测，仅展示错误记录

```shell
+-----------+-----------------------------+--------------+--------------------------------+-------------------+----------------+---------------------------------------------------------------------+
|TraceName  |Name                         |Namespace     |PodName                         |PortList           |NodeName        |Status                                                               |
+-----------+-----------------------------+--------------+--------------------------------+-------------------+----------------+---------------------------------------------------------------------+
|service    |devops-next-webhook-service  |xxxx-system  |null                            |10.4.212.38:443    |null            |fail[0.00s]:dial tcp 10.4.212.38:443: connect: connection refused    |
|service    |minio-console                |xxxx-system  |null                            |10.4.182.238:9001  |null            |fail[0.00s]:dial tcp 10.4.182.238:9001: connect: connection refused  |
|endpoint   |alertmanager-operated        |xxxx-system  |alertmanager-kube-prometheus-0  |10.3.1.151:9094    |192.168.128.49  |fail[0.00s]:dial tcp 10.3.1.151:9094: connect: connection refused    |
|endpoint   |devops-next-webhook-service  |xxxx-system  |rds-api-75f7b844d4-g24zz        |10.3.0.153:9443    |192.168.254.87  |fail[0.00s]:dial tcp 10.3.0.153:9443: connect: connection refused    |
|endpoint   |minio-console                |xxxx-system  |minio-0                         |10.3.0.168:9001    |192.168.128.49  |fail[0.00s]:dial tcp 10.3.0.168:9001: connect: connection refused    |
|endpoint   |minio-console                |xxxx-system  |minio-1                         |10.3.0.174:9001    |192.168.128.49  |fail[0.00s]:dial tcp 10.3.0.174:9001: connect: connection refused    |
|endpoint   |minio-console                |xxxx-system  |minio-2                         |10.3.0.176:9001    |192.168.128.49  |fail[0.00s]:dial tcp 10.3.0.176:9001: connect: connection refused    |
|endpoint   |minio-console                |xxxx-system  |minio-3                         |10.3.0.180:9001    |192.168.128.49  |fail[0.00s]:dial tcp 10.3.0.180:9001: connect: connection refused    |
|endpoint   |minio-console                |xxxx-system  |minio-4                         |10.3.0.190:9001    |192.168.128.49  |fail[0.00s]:dial tcp 10.3.0.190:9001: connect: connection refused    |
|endpoint   |minio-console                |xxxx-system  |minio-5                         |10.3.0.195:9001    |192.168.128.49  |fail[0.00s]:dial tcp 10.3.0.195:9001: connect: connection refused    |
+-----------+-----------------------------+--------------+--------------------------------+-------------------+----------------+---------------------------------------------------------------------+
```
