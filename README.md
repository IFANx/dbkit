## Environment
- MySQL # store test statistic
- go

## Dependency
```cmd
# before the following command, create database dbkit in MySQL (*)
cd deploy
mysql -h{host} -P{port} -u{username} -p dbkit < dbkit.sql 
-- {host},{port},{username},{password} should be the same as config.json 
```

## Run
```cmd
go build
./dbkit
```