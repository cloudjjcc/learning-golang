mac系统下制作黑苹果启动盘



### 下载系统镜像

从黑果小兵处下载：https://blog.daliansky.net/



### 制作U盘启动盘

利用balenaEtcher 制作U盘启动盘



### 修改OC驱动

```shell
#列出所有分区
~ diskutil list
/dev/disk0 (internal, physical):
   #:                       TYPE NAME                    SIZE       IDENTIFIER
   0:      GUID_partition_scheme                        *251.0 GB   disk0
   1:                        EFI EFI                     314.6 MB   disk0s1
   2:                 Apple_APFS Container disk1         250.7 GB   disk0s2

/dev/disk1 (synthesized):
   #:                       TYPE NAME                    SIZE       IDENTIFIER
   0:      APFS Container Scheme -                      +250.7 GB   disk1
                                 Physical Store disk0s2
   1:                APFS Volume MacOS - Data            219.8 GB   disk1s1
   2:                APFS Volume Preboot                 573.5 MB   disk1s2
   3:                APFS Volume Recovery                632.1 MB   disk1s3
   4:                APFS Volume VM                      5.4 GB     disk1s4
   5:                APFS Volume MacOS                   15.4 GB    disk1s5
   6:              APFS Snapshot com.apple.os.update-... 15.4 GB    disk1s5s1

/dev/disk2 (external, physical):
   #:                       TYPE NAME                    SIZE       IDENTIFIER
   0:      GUID_partition_scheme                        *30.9 GB    disk2
   1:                        EFI EFI                     209.7 MB   disk2s1
   2:       Microsoft Basic Data WEPE                    716.8 MB   disk2s2
   3:                  Apple_HFS Install macOS Big Sur   13.9 GB    disk2s3

/dev/disk3 (disk image):
   #:                       TYPE NAME                    SIZE       IDENTIFIER
   0:      GUID_partition_scheme                        +1.0 TB     disk3
   1:                        EFI EFI                     209.7 MB   disk3s1
   2:                  Apple_HFS Time Machine Backups    1.0 TB     disk3s2
   
# 加载分区（Mac系统默认隐藏EFI分区）
sudo diskutil mount disk2s1
```



