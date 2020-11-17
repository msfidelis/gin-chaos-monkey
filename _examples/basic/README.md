## Basic Usage 

```bash
export CHAOS_MONKEY_ENABLED=true ; export CHAOS_MONKEY_MODE=critical; export CHAOS_MONKEY_LATENCY=true; export CHAOS_MONKEY_EXCEPTION=true; go run main.go
```

## Requests

### Basic Request

```bash
while true; do curl 0.0.0.0:8080/healthcheck ; echo ; done
```