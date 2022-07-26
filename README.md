# password-tool

### 为什么
因为总是忘记ssh密码，而且有多个设备同时办公的需求，所以做了这个工具。
这个工具通过将ssh的登录配置写入文件，并可以将配置同步至git代码仓库的方式，来达到多个设备之间的配置同步


### 使用方式
- 默认使用的是本地仓库，无法使用 git push 同步到远程仓库
- 使用可以使用下面的命令添加一个默认的远程仓库。
```shell
password-tool repo add --name main --url git@github.com:xxx/password-tool-repo.git
```
- 添加ssh配置
```shell
password-tool add --name test --host 10.167.3.219 --port 22 --user xxx --password xxx
```
- 列出ssh配置
```shell
password-tool ls [keyword]  # 如果增加keyword，那么就可以安装keyword进行搜索
```

