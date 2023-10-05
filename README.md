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
The responses from these APIs are as follows:
```
{
	"message": "ok",
}
```
Among these, 8081 and 8082 are set to wait 4 seconds before responding. This represents servers with slow responses.
```go
router.GET("/", func(c *gin.Context) {
		time.Sleep(4 * time.Second)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
```

### loadbalancer
There are two types of load balancers in this repository, representing different algorithms.
Specifically, there are two types: lc (Least Connections) and lr (Least Response Time).The two types can be changed by modifying the config.json.

```json:config.json
{
    "type": "lr"
}
```
To launch the load balancer, enter the following command. (Don't forget to start the test server beforehand.)
```
make run
```
### Results
Let's take a look at the actual execution results from here.
First, let's take a look at the results for Least Connections.
```
make run
{"level":"info","msg":"access to endpoint","url":"http://localhost:8081/","connections":0}
{"level":"info","msg":"access to endpoint","url":"http://localhost:8081/","connections":0}
{"level":"info","msg":"access to endpoint","url":"http://localhost:8082/","connections":0}
{"level":"info","msg":"access to endpoint","url":"http://localhost:8083/","connections":0}
{"level":"info","msg":"access to endpoint","url":"http://localhost:8083/","connections":0}
{"level":"info","msg":"access to endpoint","url":"http://localhost:8081/","connections":0}
```
I think you can see that the server is actually being changed so that the connections become 0.


Next, let's take a look at the results for Least Response Time.
```
make run
{"level":"info","msg":"access to endpoint","url":"http://localhost:8085/","connections":0}
{"level":"info","msg":"access to endpoint","url":"http://localhost:8085/","connections":0}
{"level":"info","msg":"access to endpoint","url":"http://localhost:8085/","connections":0}
{"level":"info","msg":"access to endpoint","url":"http://localhost:8085/","connections":0}
```
Now, as explained earlier, you can see that there is no access to 8081 and 8082, which have slower responses.

## Reference
 [Golangでロードバランサーを実装する](https://bmf-tech.com/posts/Golang%E3%81%A7%E3%83%AD%E3%83%BC%E3%83%89%E3%83%90%E3%83%A9%E3%83%B3%E3%82%B5%E3%83%BC%E3%82%92%E5%AE%9F%E8%A3%85%E3%81%99%E3%82%8B)

  [Creating a Load Balancer in GO](https://medium.com/@leonardo5621_66451/building-a-load-balancer-in-go-1c68131dc0ef)

 https://github.com/kasvith/simplelb

 https://github.com/leonardo5621/golang-load-balancer



 ## License
Code licensed under 
[the MIT License](https://github.com/seipan/mylb/blob/main/LICENSE).
