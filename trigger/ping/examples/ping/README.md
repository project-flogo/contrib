# Flogo Ping Example

## Install

To install run the following commands:
```
flogo create -f flogo.json
cd Ping
flogo build
```

## Testing

Run:
```
bin/Ping
```

Then open another terminal and run:
```
curl --request GET http://localhost:9096/ping/details -H "Authorization: Bearer <Token>"
```

You should then see something like:
```
{
   "Version":"1.1",
   "Appversion":"",
   "Appdescription":""
}Details:{
   "NumGoroutine":2,
   "Alloc":762472,
   "TotalAlloc":762472,
   "Sys":69926912,
   "Mallocs":1078,
   "Frees":101,
   "LiveObjects":977,
   "NumGC":0
}
```
#####
```
curl --request GET http://localhost:9096/ping -H "Authorization: Bearer <Token>"
```

You should then see something like:
```
{"response":"Ping successful"}
```
