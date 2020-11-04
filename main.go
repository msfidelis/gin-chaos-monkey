package chaos

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Load() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if IsEnabled() && IsGonnaAssault() {
			assaults := GetAssaltsEnabled()
			assault_type := getAssaultType(assaults)
			assaultAction(assault_type, ctx)
		}
	}
}

func IsEnabled() bool {
	enabled := os.Getenv("CHAOS_MONKEY_ENABLED")
	if enabled == "true" {
		return true
	}
	return false
}

func GetAssaltsEnabled() []string {
	var enabled []string

	assalt_available := []string{
		"CHAOS_MONKEY_LATENCY",
		"CHAOS_MONKEY_EXCEPTION",
		"CHAOS_MONKEY_APP_KILLER",
		"CHAOS_MONKEY_MEMORY",
	}

	for i := 0; i < len(assalt_available); i++ {
		assalt := assalt_available[i]
		if strings.ToLower(os.Getenv(strings.ToUpper(assalt))) == "true" {
			enabled = append(enabled, assalt)
		}
	}

	return enabled
}

func IsGonnaAssault() bool {
	rand.Seed(time.Now().Unix())
	modes := map[string]int{
		"":         100,
		"soft":     100,
		"hard":     50,
		"critical": 10,
	}

	quorum := MakeRange(0, modes[os.Getenv("CHAOS_MONKEY_MODE")])

	r := quorum[rand.Intn(len(quorum))]
	if r == 0 {
		return true
	}
	return false
}

func getAssaultType(assaults []string) string {
	rand.Seed(time.Now().Unix())
	return assaults[rand.Intn(len(assaults))]
}

func assaultAction(assault string, ctx *gin.Context) {
	switch assault {
	case "CHAOS_MONKEY_LATENCY":
		latencyAssault(ctx)
		break
	case "CHAOS_MONKEY_EXCEPTION":
		exceptionAssault(ctx)
		break
	case "CHAOS_MONKEY_APP_KILLER":
		appKillerAssault(ctx)
		break
	case "CHAOS_MONKEY_MEMORY":
		memoryAssault(ctx)
		break
	default:
		latencyAssault(ctx)
		break
	}
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func latencyAssault(ctx *gin.Context) {
	fmt.Println("[CHAOS MONKEY] - Latency Assault")
	time.Sleep(time.Second)
	ctx.Next()
}

func exceptionAssault(ctx *gin.Context) {
	fmt.Println("[CHAOS MONKEY] - Exception Assault")
	ctx.JSON(http.StatusServiceUnavailable, "Service Unavailable")
	ctx.Abort()
}

func appKillerAssault(ctx *gin.Context) {
	panic("[CHAOS MONKEY] - App Killer Assault")
}

func memoryAssault(ctx *gin.Context) {
	fmt.Println("[CHAOS MONKEY] - Memory Assault")
	var s []int
	sum := 1
	for sum < 9999998 {
		sum += 1
		s = append(s, sum)
	}
	ctx.Next()
}
