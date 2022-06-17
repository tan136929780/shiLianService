docker build --rm -t vfile .

docker run -p 8200:8200 vfile

项目首次运行时启动后需要执行（执行模型创建） 
        vfile/preload/init
