# OpenFaas_Demo
OpenFaas GoLang Demo

## godemo
给指定手机号发送短信

### 部署
UI界面部署，http://10.11.11.115:31112/ui/, 输入用户名密码  
`Deploy New Function` -> `CUSTOM`, 填上godemo.yaml文件中的镜像名称、函数名称自定义(函数名称为URL地址访问中的名称);  
例如:  
镜像名: harbor.chengdd.cn/dev/godemo:v4  
函数名：smsv4  

### 调用
```shell
curl --location --request GET 'http://10.11.11.115:31112/function/smsv4' \
    --header 'Content-Type: text/plain' \
    --data '12356332824'
```
