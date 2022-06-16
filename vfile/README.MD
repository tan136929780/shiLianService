docker build --rm -t vfile .
docker run -p 8200:8200 -v 挂载日志目录:/app/log vfile