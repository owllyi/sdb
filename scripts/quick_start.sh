ps aux | grep "sdb" | grep "main.go" | awk '{print "kill -9 " $2}' | sh -x
go build ./cmd/sdb/main.go
nohup ./main -config ./configs/config.yml &