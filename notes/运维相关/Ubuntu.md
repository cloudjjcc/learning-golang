# 安装Ubuntu Server 20.04





## 网络

Netplan 是一个简便的Linux网络配置工具，其会从 `/etc/netplan/*.yaml` 读取网络设置，并在 `/run` 生成特定的后端配置文件，用于网络设备进程控制。

Netplan 当前版本支持两种网络管理接口:

- NetworkManager
- Systemd-networkd

```shell
ip -a #查看网卡信息
sudo iw dev wlp2s0b1 scan # 扫描WiFi

```

进入 /etc/netplan/ 目录，编辑已存在的yaml 格式文件，文件名一般为 50-cloud-init.yaml 。下面以静态IP地址为例，介绍配置文件格式：

```yaml
network:
  version: 2
  renderer: networkd
  wifis:
    wlp2s0b1:
      dhcp4: no
      dhcp6: no
      addresses: [192.168.0.21/24]
      gateway4: 192.168.0.1
      nameservers:
        addresses: [192.168.0.1, 114.114.114.114]
      access-points:
        "network_ssid_name":
          password: "**********"
```

说明：

- 第三行：networkd，即使用 Systemd-networkd 接口。
- 第五行：wlp2s0b1，即为上一步通过ip a命令，得到的无线网卡设备名称。
- 最后两行：无线网络ssid名称与密码。
- 若 /etc/netplan/ 目录下存在多个yaml文件，那 netplan 将全部读取，此时需注意不同文件配置冲突。
- yaml 文件使用行首空格来区分不同部分，据传，每个部分的行首空格数量应一致，另传，冒号后与设置内容之间，应有一个空格。

键入 sudo netplan apply ，生成后端配置文件。



注意：

```shell
sudo apt-get install wpasupplicant # 支持wpa 加密
sudo apt-get install network-manager
```



## 蓝牙







安装 bluez

```shell
sudo apt-get -y install bluetooth bluez bluez-tools rfkill
```



修改蓝牙名称

```shell
sudo vim /etc/bluetooth/main.conf # 修改Name
```



确保未禁用蓝牙



```shell
rfkill
rfkill unblock
```





查看蓝牙设备

```shell
hcitool dev
hciconfig
```



开启和关闭蓝牙

```shell
#设备打开
sudo hciconfig hci0 up
#设备关闭
sudo hciconfig hci0 down
```



使用 Bluetoothctl  进入交互模式操作蓝牙



```shell
bluetoothctl
```







## 查看温度信息

```shell
sudo apt-get install lm-sensors
sensors
```



## ZSH



## Brew

```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh --proxy localhost:7890)"
```



## Docker

```shell
sudo apt install docker.io
sudo apt install dokcer-compose
sudo groupadd docker
sudo usermod -aG docker cloudjjcc
newgrp docker
sudo systemctl restart docker

docker info 
```





## mysql

### 安装

```shell
sudo apt install mysql-server mysql-client
sudo mysql_secure_installation
sudo vim  /etc/mysql/mysql.conf.d/mysqld.cnf #修改 bind_addr
sudo service mysql restart
```



### 设置用户权限

```sql
CREATE USER 'work'@'%' IDENTIFIED WITH mysql_native_password BY 'work@dev';
GRANT ALL PRIVILEGES ON *.* TO 'work'@'%' WITH GRANT OPTION;
```





### 数据备份与导入

利用mysqldump 工具导出数据：

--column-satistics 选项关闭列统计（mysql8.0的mysqldump工具导出mysql5的数据时需要加上）

```shell
mysqldump -uroot -p123456 -h 127.0.0.1  --column-statistics=0 --all-databases > sqlfile.sql
```





## SMB

```shell
sudo apt-get install samba samba-common
```

