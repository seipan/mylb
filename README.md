<div align="center">

![Last commit](https://img.shields.io/github/last-commit/seipan/mylb?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/seipan/mylb?style=flat-square)
![Issues](https://img.shields.io/github/issues/seipan/mylb?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/seipan/mylb?style=flat-square)
[![go](https://github.com/seipan/mylb/actions/workflows/go.yml/badge.svg)](https://github.com/seipan/loghook/actions/workflows/go.yml)
[![testserver-e2e](https://github.com/seipan/mylb/actions/workflows/e2e-testserver.yml/badge.svg)](https://github.com/seipan/mylb/actions/workflows/e2e-testserver.yml)

<img src="https://cdn-icons-png.flaticon.com/512/5880/5880629.png" alt="eyecatch" height="200">

# mylb

⭐ Implementing a Load Balancer in Golang(just a toy)  ⭐

<br>
<br>


</div>

## Usage
### testserver
You can launch servers for load balancer testing. There are a total of four servers.
```
cd testserver
make up
```
If you want to restart the server, use the following command.
```
cd testserver
make re
```
The four servers to be launched are as follows.
```
http://localhost:8081
http://localhost:8082
http://localhost:8083
http://localhost:8085
http://localhost:8086
http://localhost:8087
http://localhost:8088
http://localhost:8089
```
If you access these endpoints, you will be redirected to /health.

## Reference
 [Golangでロードバランサーを実装する](https://bmf-tech.com/posts/Golang%E3%81%A7%E3%83%AD%E3%83%BC%E3%83%89%E3%83%90%E3%83%A9%E3%83%B3%E3%82%B5%E3%83%BC%E3%82%92%E5%AE%9F%E8%A3%85%E3%81%99%E3%82%8B)

 ## License
Code licensed under 
[the MIT License](https://github.com/seipan/bulma/blob/main/LICENSE).
