# Fasthop


## Getting started

## Add your files

```
cd existing_repo
git remote add origin https://github.com/VINDA-98/Fasthop.git
git remote add origin https://github.com/VINDA-98/Fasthop.git
git branch -M master
git push -uf origin master
```


## 运行前提
```
go version ：1.19+
mysql version：5.7+
redis version：7.0+
```

## 数据库创建
保证数据库已经创建，数据库名与config.yaml中的database.database 属性名一致


## 项目本地运行
```
go mod tidy && go run main.go
```

## 项目在服务器运行
```
nohup ./Fasthop >./out.log 2>&1 &
```