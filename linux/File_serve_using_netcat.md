# Serving a file on a specific port via netcat!


It excites me every time. Not netcat itself. What quirks are possible with some standard-linux-calls.

Here today i learned, how to serve a file with netcat on the local network/internet. Plus: closing connection after first download! For this, all what i need here is `cat` and, obv., `netcat`. (Just assuming, the GNU Version one.)

```sh
cat file.txt | nc -u -c 127.0.0.1 8080
```
Streams the content of `file.txt`, on Port `8080`, as long as the first download finishes and then terminates the connections.
The `-c` flag will terminate the connection on `EOF`. With the "original" version, we could do the same with a `-q 0` which translates to "quit in 0 seconds after download".
The `-u` flag just enforces to use UDP over TCP.

(Sources are here the manpages served. Just use for ex. this one: https://linux.die.net/man/1/nc)