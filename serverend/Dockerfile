#基础镜像
FROM ubuntu:20.04
#把打好包的服务放到镜像里去
COPY user.elf /app/user.elf
#容器中执行命令的默认工作目录
WORKDIR /app
#进到容器里之后执行的命令
ENTRYPOINT ["/app/user.elf"]