#用于说明go后台程序制作和前端制作及安装部署
##制作go可执行程序
set GOOS=linu
set CGO_ENABLED=0
go build main.go
生成文件为改名goamdin_linux

##制作VUE包
npm run build:prod

##docker运行
docker exec -it cb46fff4e389 bash
cd usr/share/nginx/html/
docker cp dist/ cb46fff4e389:/usr/share/nginx/html/
exit
docker restart cb46fff4e389
