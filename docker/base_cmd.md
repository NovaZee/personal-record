#### 1:镜像导入导出

```
导出：docker save -o xxx-image.tar xxx:latest
上传：scp file user@remote-host:/path/to/destination
导入：docker load -i xxx-image.tar
```
