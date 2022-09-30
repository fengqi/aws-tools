# AWS EC2 小工具

# 配置
1. 参考 [AWS access keys](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#get-your-aws-access-keys) 获取 access key id 和 secret access key
2. 参考 [Creating the Credentials File](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#creating-the-credentials-file) 将配置写入到 ~/.aws/credentials 内

```
[default]
region = <YOUR_REGION>
aws_access_key_id = <YOUR_ACCESS_KEY_ID>
aws_secret_access_key = <YOUR_SECRET_ACCESS_KEY>
```

其中 region 可以留空，通过运行时参数指定
# 开机、关机

./aws-ec2-power -i i-05aff47edc3b4f -a start
./aws-ec2-power -r us-west-1 -i i-05aff47edc3b4f -a start

./aws-ec2-power -i i-05aff47edc3b4f -a stop
./aws-ec2-power -r us-west-1 -i i-05aff47edc3b4f -a stop

# 获取公网 ip

/aws-ec2-public-ip -i i-05aff47edc3b4f
/aws-ec2-public-ip -r us-west-1 -i i-05aff47edc3b4f