# Gin Chaos Monkey - Assault Middleware for Gin :cocktail: :cocktail: :cocktail:

<div align=>
	<img align="right" width="150px" src="/.github/assets/img/color.png">
</div> 


<br><br>

## Contents

- [Gin Web Framework](https://github.com/gin-gonic/gin)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Assault types](#assault-types)
  - [Environment variables configuration](#environment-variables-configuration)
  - [Development](#development)

## Installation 

```
go get -v github.com/msfidelis/gin-chaos-monkey
```

<br><br>

# Usage 

## Basic Usage

```go
package main

import (
	chaos "github.com/msfidelis/gin-chaos-monkey"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

        //Middlewares
        r.Use(gin.Recovery())
        r.Use(chaos.Load())

	// Healthcheck
	r.GET("/healthcheck", healthcheck.Ok)    

	r.Run()
}
```

## Specific Routes

```golang

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	chaos "github.com/msfidelis/gin-chaos-monkey"
)

func main() {
	router := gin.Default()

	//Enable Chaos Monkey in Specific Route
	router.GET("/healthcheck/chaos", chaos.Load(), healthcheck.Ok)

	router.GET("/healthcheck", healthcheck.Ok)

	router.Run()
}
```

# ASSAULT TYPES


### LATENCY ASSAULT 

This assault increase latency on response time for web requests. You can set `CHAOS_MONKEY_LATENCY_MAX_TIMEOUT` environment variable to customize a max time to increase in requests. 


### EXCEPTION ASSAULT

This assault randomly returns 5xx errors for HTTP requests. You can set `CHAOS_MONKEY_EXCEPTION_HTTP_STATUS_CODE` to customize status codes to return in HTTP exception. Default: `503`


### APP KILLER ASSAULT

This assault randomly inject an `panic` exception on application runtime

### MEMORY ASSAULT 

Increases the RAM consumption of the application

### CPU ASSAULT 

Increases the CPU consumption of the application

# CONFIGURATION 

## Enable Chaos Monkey Assalts

```bash
export CHAOS_MONKEY_ENABLED=true
export CHAOS_MONKEY_MODE=soft
export CHAOS_MONKEY_LATENCY=true
export CHAOS_MONKEY_LATENCY_MIN_TIME=5000
export CHAOS_MONKEY_LATENCY_MAX_TIME=10000
```

## Environment Variables Configuration 

| VARIABLE                                  | OPTIONS               	| DEFAULT   | 
| ----------------------------------------- | --------------------------| --------- |
| CHAOS_MONKEY_ENABLED                      | true/false            	| false     |
| CHAOS_MONKEY_MODE                         | soft/hard/critical/hell	| soft      |
| CHAOS_MONKEY_LATENCY                      | true/false            	| false     |
| CHAOS_MONKEY_LATENCY_MIN_TIME       		| miliseconds           	| max_time  |
| CHAOS_MONKEY_LATENCY_MAX_TIME       		| miliseconds           	| 1000      |
| CHAOS_MONKEY_EXCEPTION                    | true/false            	| false     |
| CHAOS_MONKEY_EXCEPTION_HTTP_STATUS_CODE   | 5xx                   	| 503       |
| CHAOS_MONKEY_APP_KILLER                   | true/false            	| false     |
| CHAOS_MONKEY_MEMORY                       | true/false            	| false     |
| CHAOS_MONKEY_CPU                       	| true/false            	| false     |

## Development 

### Running Tests 

```
go test -v
```
