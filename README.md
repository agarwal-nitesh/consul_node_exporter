# consul_node_exporter
Add service registry for node exporter.

### Config
Specify the consul address and token with service write permission. Following policy maybe used.
```
service_prefix "" {
   policy = "write"
}
```
### Run
go run register_consul.go
