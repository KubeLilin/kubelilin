# ASP.NET Core Prometheus 集成&采集&展示

## 创建 ASP.NET Core Web API
此项目在.NET 7 环境创建
```bash
dotnet new webapi
```
## 添加包
```bash
dotnet add package prometheus-net.AspNetCore --version 8.0.0
```
添加指标端点代码
### program.cs
```csharp
...
...
app.UseMetricServer();
...
...
```
完整文件内容如下:
```csharp
using Prometheus;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddControllers();

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseMetricServer();

app.UseAuthorization();

app.MapControllers();

app.Run();

```

运行 http://localhost:5069/metrics 端口号替换成运行端口

端点指标内容如下：

```
# HELP process_private_memory_bytes Process private memory size
# TYPE process_private_memory_bytes gauge
process_private_memory_bytes 0
# HELP system_runtime_alloc_total (B) Allocation Rate
# TYPE system_runtime_alloc_total counter
system_runtime_alloc_total 1230088
# HELP system_runtime_gen_0_size (B) Gen 0 Size
# TYPE system_runtime_gen_0_size gauge
system_runtime_gen_0_size 0
# HELP system_runtime_poh_size (B) POH (Pinned Object Heap) Size
# TYPE system_runtime_poh_size gauge
system_runtime_poh_size 0
# HELP prometheus_net_meteradapter_instruments_connected Number of instruments that are currently connected to the adapter.
# TYPE prometheus_net_meteradapter_instruments_connected gauge
prometheus_net_meteradapter_instruments_connected 0
# HELP system_runtime_gc_committed (MB) GC Committed Bytes
# TYPE system_runtime_gc_committed gauge
system_runtime_gc_committed 0
# HELP process_working_set_bytes Process working set
# TYPE process_working_set_bytes gauge
process_working_set_bytes 118489088
# HELP system_runtime_threadpool_thread_count ThreadPool Thread Count
# TYPE system_runtime_threadpool_thread_count gauge
system_runtime_threadpool_thread_count 5
# HELP system_net_sockets_bytes_received Bytes Received
# TYPE system_net_sockets_bytes_received gauge
system_net_sockets_bytes_received 10322
# HELP system_runtime_gen_1_gc_count_total Gen 1 GC Count
# TYPE system_runtime_gen_1_gc_count_total counter
system_runtime_gen_1_gc_count_total 0
# HELP system_runtime_il_bytes_jitted (B) IL Bytes Jitted
# TYPE system_runtime_il_bytes_jitted gauge
system_runtime_il_bytes_jitted 283857
# HELP microsoft_aspnetcore_hosting_current_requests Current Requests
# TYPE microsoft_aspnetcore_hosting_current_requests gauge
microsoft_aspnetcore_hosting_current_requests 0
# HELP microsoft_aspnetcore_server_kestrel_current_tls_handshakes Current TLS Handshakes
# TYPE microsoft_aspnetcore_server_kestrel_current_tls_handshakes gauge
microsoft_aspnetcore_server_kestrel_current_tls_handshakes 0
# HELP system_net_sockets_datagrams_received Datagrams Received
# TYPE system_net_sockets_datagrams_received gauge
system_net_sockets_datagrams_received 0
# HELP microsoft_aspnetcore_server_kestrel_total_connections Total Connections
# TYPE microsoft_aspnetcore_server_kestrel_total_connections gauge
microsoft_aspnetcore_server_kestrel_total_connections 5
# HELP system_net_sockets_incoming_connections_established Incoming Connections Established
# TYPE system_net_sockets_incoming_connections_established gauge
system_net_sockets_incoming_connections_established 5
```
