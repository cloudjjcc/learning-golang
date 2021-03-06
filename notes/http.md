# http请求结构

状态行、请求头、消息主体。类似于下面这样：

```
<method> <request-URL> <version>
<headers>

<entity-body>
```



1. GET 用于信息获取，而且应该是安全的 和 幂等的。

   所谓安全的意味着该操作用于获取信息而非修改信息。换句话说，GET 请求一般不应产生副作用。就是说，它仅仅是获取资源信息，就像数据库查询一样，不会修改，增加数据，不会影响资源的状态。

   幂等的意味着对同一 URL 的多个请求应该返回同样的结果。

   GET 请求报文示例：

   ```
    GET /books/?sex=man&name=Professional HTTP/1.1
    Host: www.example.com
    User-Agent: Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7.6)
    Gecko/20050225 Firefox/1.0.1
    Connection: Keep-Alive
   ```

2. POST 表示可能修改变服务器上的资源的请求。

   ```
    POST / HTTP/1.1
    Host: www.example.com
    User-Agent: Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7.6)
    Gecko/20050225 Firefox/1.0.1
    Content-Type: application/x-www-form-urlencoded
    Content-Length: 40
    Connection: Keep-Alive
   
    sex=man&name=Professional  
   ```



# http1.1

默认Connection:keep-alive,得到http响应并不断开连接，除非指明Connection:close

同一个tcp连接只有在上一个http请求完成后才能发送下一个请求









# http2.0

http2.0是一种安全高效的下一代http传输协议。

安全是因为http2.0建立在https协议的基础上，高效是因为它是通过二进制分帧来进行数据传输。

- Stream： 一个双向流，一条连接可以有多个 streams。
- Message： 也就是逻辑上面的 request，response。
- Frame:：数据传输的最小单位。每个 Frame 都属于一个特定的 stream 或者整个连接。一个 message 可能有多个 frame 组成。





## 二进制分帧（Binary Format）

在二进制分帧层上，http2.0会将所有传输信息分割为更小的消息和帧，并对它们采用二进制格式的编码将其封装，新增的二进制分帧层同时也能够保证http的各种动词，方法，首部都不受影响，兼容上一代http标准。

`Frame Format`
 Frame 是 HTTP/2 里面最小的数据传输单位，一个 Frame 定义如下:



```ruby
+-----------------------------------------------+
|                 Length (24)                   |
+---------------+---------------+---------------+
|   Type (8)    |   Flags (8)   |
+-+-------------+---------------+-------------------------------+
|R|                 Stream Identifier (31)                      |
+=+=============================================================+
|                   Frame Payload (0...)                      ...
+---------------------------------------------------------------+
```

- Length：也就是 Frame 的长度，默认最大长度是 16KB，如果要发送更大的 Frame，需要显式的设置 max frame size。
- Type：Frame 的类型，譬如有 DATA，HEADERS，PRIORITY 等。
- Flag 和 R：保留位，可以先不管。
- Stream Identifier：标识所属的 stream，如果为 0，则表示这个 frame 属于整条连接。
- Frame Payload：根据不同 Type 有不同的格式。





## 多路复用（MultiPlexing）

多路复用允许同时通过单一的http/2 连接发起多重的请求-响应消息。有了新的分帧机制后，http/2 不再依赖多个TCP连接去实现多流并行了。每个数据流都拆分成很多互不依赖的帧，而这些帧可以交错（乱序发送），还可以分优先级，最后再在另一端把它们重新组合起来。

HTTP/2 通过 stream 支持了连接的多路复用，提高了连接的利用率。Stream 有很多重要特性：

- 一条连接可以包含多个 streams，多个 streams 发送的数据互相不影响。
- Stream 可以被 client 和 server 单方面或者共享使用。
- Stream 可以被任意一段关闭。
- Stream 会确定好发送 frame 的顺序，另一端会按照接受到的顺序来处理
- Stream 用一个唯一 ID 来标识。

这里在说一下 Stream ID，如果是 client 创建的 stream，ID 就是奇数，如果是 server 创建的，ID 就是偶数。ID 0x00 和 0x01 都有特定的使用场景。Stream ID 不可能被重复使用，如果一条连接上面 ID 分配完了，client 会新建一条连接。而 server 则会给 client 发送一个 GOAWAY frame 强制让 client 新建一条连接。







## **头部压缩（Header Compression）**

http/2使用encoder来减少需要传输的header大小，通讯双方各自缓存一份头部字段表，既避免了重复header的传输，又减小了需要传输的大小。

http/2使用的是专门为首部压缩而设计的HPACK②算法。

![img](https://upload-images.jianshu.io/upload_images/11345047-3505579e020e4e5b.png?imageMogr2/auto-orient/strip|imageView2/2/w/600/format/webp)

其中，http1.X中的首部信息header封装到Headers帧中，而request body将被封装到Data帧中。



![img](https://user-gold-cdn.xitu.io/2019/10/31/16e208ee1d0caab8?imageView2/0/w/1280/h/960/format/webp/ignore-error/1)



## **服务端推送（Server Push）**

http2.0能通过push的方式将客户端需要的内容预先推送过去，所以也叫“cache push”。



## 重置连接



很多app客户端都有取消图片下载的功能场景，对于http1.x来说，是通过设置tcp segment里的reset flag来通知对端关闭连接的。这种方式会直接断开连接，下次再发请求就必须重新建立连接。http2.0引入RST_STREAM类型的frame，可以在不断开连接的前提下取消某个request的stream，表现更好。



## 流量控制

每个 http2 流都拥有自己的公示的流量窗口，它可以限制另一端发送数据。



