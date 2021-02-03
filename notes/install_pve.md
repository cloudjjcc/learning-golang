闲置的笔记本可以用来做家庭服务器，目前主流的虚拟机平台方案有两种：

+ ESXI

+ PVE

  

  ESXI 在个人电脑上安装有以下问题：

  + 4代inter CPU 需要关闭MMIO

  +  需要自己打包网卡驱动

  

所以我选择了PVE。

用Rufus 制作U盘安装盘的时候注意选择DD模式





PVE基于Debian系统，可以通过以下步骤更新软件源：

```sh
apt install git
git clone https://github.com/gzzchh/pve_knife
cd pve_knife
sh pve_knife.sh
```





关闭笔记本合上后盖休眠：

systemd 处理某些电源相关的 ACPI事件，可以通过从 /etc/system/logind.conf以下选项进行配置：

> HandlePowerKey按下电源键后的行为，默认power off
>
> HandleSleepKey 按下挂起键后的行为，默认suspend
>
> HandleHibernateKey 按下休眠键后的行为，默认hibernate
>
> HandleLidSwitch 合上笔记本盖后的行为，默认suspend
>
> 

触发的行为可以有

> ignore、power off、reboot、halt、suspend、hibernate、hybrid-sleep、lock 或 exec。

如果要合盖不休眠只需要把HandleLidSwitch选项设置为如下即可：

```
HandleLidSwitch=ignore
```

注意：设置完成保存后运行下列命令才生效。

```
systemctl restart systemd-logind
```

