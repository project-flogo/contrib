# Crypto functions

This package adds function for encrypting and decrypting using AES GCM.

## encrypt()
Encrypt given byte array with provide 128bit or 256 bit key.

Example usage
```
coerce.toString(crypto.encrypt(coerce.toBytes("AES256Key-32Characters1234567890"), coerce.toBytes("encrypt-me-text")))
```


### Input Args

| Arg        | Type      | Description
|:---        | :---      | :---    
| key        | byte array| 128bit or 256 bit key
| plaintext  | byte array| Plaintext that will be encrypted

### Output

| Arg     | Type   | Description
|:---      | :---   | :---    
| returnType | byte array| Encrypted `plaintext` 


## decrypt()
Decrypt given encrypted byte array with provide 128bit or 256 bit key.

Example usage
```
coerce.toString(crypto.decrypt(coerce.toBytes("AES256Key-32Characters1234567890"), coerce.toBytes($activity[Encrypt].output.name)))
```


### Input Args

| Arg        | Type      | Description
|:---        | :---      | :---    
| key        | byte array| 128bit or 256 bit key
| ciphertext | byte array| Encrypted payload that will be decrypted

### Output

| Arg        | Type      | Description
|:---        | :---      | :---    
| returnType | byte array| Decrypted `ciphertext` 
