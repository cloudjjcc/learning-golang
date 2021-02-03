protobuffer 基本语法



指定版本信息

```protobuf
syntax = "proto3";//指定版本信息
```



消息类型



```protobuf
message Prod{
	string name=1;
}
```



枚举



```protobuf
enum Type{
	Male=0;
	Female=1;
}
```



数组

```protobuf
message Person //message为关键字，作用为定义一种消息类型
{
    string name = 1;    //姓名
    int32 id = 2;       //id
    string email = 3;   //邮件
}

message AddressBook
{
    repeated Person people = 1;
}
```

