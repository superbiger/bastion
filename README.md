### 开发

#### 起步
```shell script
git clone git@github.com:dev-escort/bastion.git
cd bastion

# 根据 secret_example.toml 模板里面的内容 创建配置文件
# 创建出 secret_dev.toml 配置文件, 填写上对应的 mysql 地址和 redis 配置
touch secret_dev.toml

# 本地开发 mysql 和 redis 需要自己搭建  
# 对应的表结构 sql 在 tools/sql 文件夹下

# 运行项目
go run main.go

```
#### 生产文档
```shell script
swag init
```

#### 实用工具
http://stming.cn/tool/sql2go.html

### 问题

#### 运行 xxx.test.go 时候找不到配置文件？
```shell script
# goland 编辑器中运行测试文件的时候，可能会抛出找不到配置的错误
# 需要在把你项目位置 作为环境变量设置到全局 
# 重启 goland  点击绿色小三角按钮，运行测试文件就能找到配置了

vi ~/.zshrc 
export BASTION_WORKSPACE="你的项目地址"
source ~/.zshrc
```
