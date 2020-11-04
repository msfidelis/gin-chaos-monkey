# Gin Chaos Monkey - Assalt Middleware for Gin 

:cocktail: :cocktail: :cocktail: Chaos Monkey assalts middleware for Gin Gonic 

## Installation 

```
go get -v github.com/msfidelis/gin-chaos-monkey
```


# Usage 

```go
package main

import (
	chaos_monkey "github.com/mfidelis/gin-chaos-monkey"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

    //Middlewares
    r.Use(gin.Recovery())
    r.Use(chaos_monkey.Load())

	// Healthcheck
	r.GET("/healthcheck", healthcheck.Ok)    

	r.Run()
}
```

# Configuration 

## Enable Chaos Monkey Assalts

```bash
export CHAOS_MONKEY_ENABLED=true
export CHAOS_MONKEY_MODE=soft
export CHAOS_MONKEY_LATENCY=true
```

## Environment Variables Configuration 

| VARIABLE                                  | OPTIONS               | DEFAULT   | 
| ----------------------------------------- | ------------------    | --------- |
| CHAOS_MONKEY_ENABLED                      | true/false            | false     |
| CHAOS_MONKEY_MODE                         | soft/hard/critical    | soft      |
| CHAOS_MONKEY_LATENCY                      | true/false            | false     |
| CHAOS_MONKEY_LATENCY_MAX_TIMEOUT          | miliseconds           | 1000      |
| CHAOS_MONKEY_EXCEPTION                    | true/false            | false     |
| CHAOS_MONKEY_APP_KILLER                   | true/false            | false     |
| CHAOS_MONKEY_MEMORY                       | true/false            | false     |
