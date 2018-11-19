# EFTL Example

## Install

Step 1 : Install FTL Server
FTL download; Follow the installation instructions for your platform

Step 2: Install eFTL Server
EFTL download; Follow the installation instructions for your platform here

Step 3:
Get Client-Server files
"git clone github.com/project-flogo/contrib/trigger/eftl"

Step 4:
To start the EFTL server run:
go run helper/main.go -ftl

Then in another terminal run:
go run helper/main.go -eftl

Step 5:
Start Listener Server
go run server/server.go


Step 6:
To install run the following commands:
```
flogo create -f flogo.json
cd eftl
flogo build
```


## Testing

Run:
```
bin/eftl
```

Then open another terminal and run client:
```
go run client/client.go
```

You should then see something like on server screen:
```
2018/11/19 14:12:59 Request URI : /a
2018/11/19 14:12:59 application/json; charset=UTF-8
2018/11/19 14:12:59 {"message":"hello world"}
```

