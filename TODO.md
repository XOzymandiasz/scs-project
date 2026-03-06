# Client
    Before communication with Server must register/login to the TTP
    Encrypt ID with TTP
    Generate two pait public key and send the public keys to the ttp
    Send request to server

# Server
    Before communication with Client must register/login to the TTP
    Encrypt ID with TTP
    Generate two pait public key and send the public keys to the ttp
    Forward Client request to ttp and send request for server authentication

# TTP 
    Encrypt ID
    Returns X.509 public key certificate to Client and Server
    TTP validates Server's request
    Sends to server and client information about positive authentication
    Sends to Client request for user authentication
