# Installation
## Command Line
### Installing
```
sudo apt install gcc g++ sqlite3
go get github.com/mattn/go-sqlite3
go get github.com/ritewhose/df
cd $GOPATH/src/github.com/
mv ritewhose shppr
cd shppr/df
git checkout cmdpack
go build -o gtb7 cmd/bot/main.go

export dftoken="<discord_api_token_here>"
export dfprefix="." # Or some other symbol:wq
```
### Running
```
./gtb7 2>> gtb7.log &
```

## Docker
(Doesn't really work yet)
```
docker build df -t ritewhose/df
docker run --name="df" -d -e dfprefix="$dfprefix" -e dftoken="$dftoken" ritewhose/df
```
