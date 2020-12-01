# st
中文简体与繁体转换工具, 可用于命令行以及IDE file watch

### 安裝
```
   go get -u github.com/xzp51/st
```

### 使用
```
  st -st=s2t path ...
```
### GOLAND
```
1. 找到file watchers
2. 配置哪些文件:
   file type: 如 GO
   scope: 如 project files
3. 配置工具:
   Program: st
   Arguments: -st=s2t $FilePath$
   Output paths: $FilePath$
```
