Java key store是Java语言常用的秘钥管理工具，会对应生成jks格式的文件。本文将简要介绍如何将jks文件转换为Golang使用的文件格式。

## 依赖

keytool 和 openssl需要预先安装

## jks to pkcs12

指令为
```
keytool -v -importkeystore -srckeystore <jks文件> -srcstoretype jks -srcstorepass <jks文件访问密码> -destkeystore <pkcs12文件名> -deststoretype pkcs12 -deststorepass <pkcs12文件访问密码>
```

## pkcs12 to client private key

指令为
```
openssl pkcs12 -in <client pkcs12文件名> -nocerts -nodes|openssl rsa -out <私钥文件名>
```

## pkcs12 to client cert

指令为
```
openssl pkcs12 -in <client pkcs12文件名> -clcerts -nokeys |openssl x509 -out <client证书文件名>
```

## pkcs12 to ca cert

指令为
```
openssl pkcs12 -in <ca pkcs12文件名> -cacerts -nokeys |openssl x509 -out <ca证书文件名>
```
