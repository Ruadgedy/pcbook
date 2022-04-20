# GRPC Go 使用
##课程目标：构建一个定制PC的APP，其中客户端与服务端通信使用GRPC
# GRPC 4种模式
1. 客户端unary  服务端unary
这种模式很像HTTP

2. 客户端stream  服务端unary

3. 客户端unary  服务端stream

4. 客户端、服务端 bistream

# GRPC vs REST
![GRPC vs REST](./img/QQ20220416-212016@2x.png)

# unary GRPC & What will we do
1. Define proto service
2. Implement server in Go
> Write server-side code to handle the unary RPC request:save laptop to an in-memory store
3. Implement client in Go
> Write client-side code to call the unary RPC on server.Also write unit test for the interaction between client and server
4. Handle errors and deadline 
> Learn how to set deadline for the request,check for cancellation,handle errors,and return suitable status code to the client

# server-streaming gRPC & What will we do 
1. Define RPC in proto file
> Add the server-streaming RPC to search laptops with filtering
2. Implement server in Go
> Write server-side code to handle the server-streaming RPC request
3. Implement client in Go
> Write client-side code to call the server-streaming RPC on server
4. Write unit test

# client-streaming gRPC & What will we do 
1. Define RPC in proto file
> Add client-streaming RPC to upload a laptop image
2. Implement server in Go
> Write server-side code to handle the client-streaming RPC
3. Implement client in Go
> Write client-side code to call the client-streaming RPC
4. Write unit test

# bidirectional gRPC & What will we do 
1. Define RPC in proto file
> Add bidirectional-streaming RPC to rate a laptop
2. Implement server in Go
> Write server-side code to handle the bidirectional-streaming RPC
3. Implement client in Go
> Write client-side code to call the bidirectional-streaming RPC
4. Write unit test

# What is GRPC interceptor?
![interceptor](./img/QQ20220419-132058@2x.png)
## What will we do?
1. Implement server interceptors
 - Authenticate users with JWT
 - Authorize access by roles
2. Implement client interceptors
 - Login user to get JWT access token
 - Attach token to RPC request
> 有两种类型的interceptor，一种是unary拦截器，另一种是stream拦截器