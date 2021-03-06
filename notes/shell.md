NGINX日志统计UV


```shell
awk '{print $1}' access.log | sort | uniq -c | wc -l
```





