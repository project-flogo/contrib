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
curl http://localhost:9096/ping/details
```

You should then see something like:
```
{"Version":"1.1","Appversion":"","Appdescription":""}
```
