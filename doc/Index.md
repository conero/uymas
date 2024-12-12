## 文档信息

> 2024年12月12日 
>
> Joshua Conero









### secure

基于go原始库封装及优化aes不同模式的数据加密

#### aes 加密

- ECB       电子密码本模式（<span style="color:red;">不安全，不再推荐</span>），Electronic Codebook Book
- CBC      密码分组链接模式，Cipher Block Chaining。相同的明文块不会生成相同的密文块，适合加密大文件；适用于大多数通用加密场景，尤其是需要较高安全性的场景。
- OFB      输出反馈模式，Output FeedBack
- CTR       计算器模式，Counter。适用于需要高并发处理的场景，如云计算、分布式系统等。
- CFB        密码反馈模式，Cipher FeedBack 。适用于流式数据加密，如实时通信、视频流等。
- OFB        输出反馈模式，Cipher FeedBack。适用于对错误传播敏感的场景，如卫星通信、广播等。
- GCM       服务对称密码模式（CTR变种）， Galois Counter Mode。适用于需要同时加密和认证的场景，如 TLS 1.2+、IPsec 等。



​	

