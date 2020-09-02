## go-redis


* 官方文档


```
The Conn interface is the primary interface for working with Redis. Applications create connections by calling the Dial, DialWithTimeout or NewConn functions. In the future, functions will be added for creating sharded and other types of connections.

The application must call the connection Close method when the application is done with the connection.
```