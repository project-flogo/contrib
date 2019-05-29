# Utils Functions
This package adds a set of functions that can be used in Flogo versions >= 0.9.0.


## Functions

| Name         | Decription             | Sample                                                |
|:-------------|:-----------------------|:------------------------------------------------------|
| decodestring | Decodestring returns the string represented by the base 64 encoded input string. | utils.decodestring(\"SGVsbG8gV29ybGQ=\") |
| encodestring | Encodestring returns a base 64 encoded copy of the input string. | utils.encodestring("Hello World") |
| uuid         | UUID generates a random UUID according to RFC 4122. | utils.uuid() |