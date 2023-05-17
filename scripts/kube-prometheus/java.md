## 采集Java程序JVM信息

## 创建 Spring Boot Application 应用程序
进行 https://start.spring.io 使用版本 Spring Boot v2.7.11和JDK 17，并创建一个具有以下依赖项的简单JAVA应用程序。 
* Spring Boot Actuator (Ops)
* Prometheus (Observability)
* Spring Web (Optional: only to create a simple REST controller.)

Maven POM 会生成以下依赖:
```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-actuator</artifactId>
</dependency>

<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web</artifactId>
</dependency>

<dependency>
    <groupId>io.micrometer</groupId>
    <artifactId>micrometer-registry-prometheus</artifactId>
    <scope>runtime</scope>
</dependency>
```

接下来，我们需要公开一个执行器端点，Prometheus将通过该端点以Prometheus能够理解的格式收集指标数据。为此，我们需要添加以下属性。 
```properties
management.endpoints.web.exposure.include=prometheus
```
接下来，让我们添加一个简单的控制器和一个简单的接口端点。
```java
@RestController
@SpringBootApplication
public class MonitorApplication {

	public static void main(String[] args) {
		SpringApplication.run(MonitorApplication.class, args);
	}
	
	@GetMapping("/hello")
	public String hello() {
		return "Hello World!";
	}
}
```
现在，让我们启动应用程序并打开以下URL。
```url
http://localhost:8080/actuator/prometheus
```
打开上述端点后，您将发现以下格式的一些指标数据,例如：
```prometheus
jvm_memory_used_bytes{area="heap",id="G1 Survivor Space",} 1005592.0
```
