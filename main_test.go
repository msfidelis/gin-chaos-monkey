package chaos_test

import (
	"os"
	"reflect"
	"testing"

	chaos "github.com/msfidelis/gin-chaos-monkey"
)

func TestEnabledIsFalseByDefault(t *testing.T) {
	enabled := chaos.IsEnabled()
	if enabled != false {
		t.Errorf("Expected %s, got %v", "false", enabled)
	}
}

func TestEnabledIsTrueByEnvironmentVariable(t *testing.T) {
	err := os.Setenv("CHAOS_MONKEY_ENABLED", "true")
	if err != nil {
		t.Errorf("Error to set CHAOS_MONKEY_ENABLED environment variable to true")
	}
	enabled := chaos.IsEnabled()
	if enabled != true {
		t.Errorf("Expected %s, got %v", "true", enabled)
	}
}

func TestEnabledIsFalseByEnvironmentVariable(t *testing.T) {
	err := os.Setenv("CHAOS_MONKEY_ENABLED", "false")
	if err != nil {
		t.Errorf("Error to set CHAOS_MONKEY_ENABLED environment variable to false")
	}
	enabled := chaos.IsEnabled()
	if enabled != false {
		t.Errorf("Expected %s, got %v", "false", enabled)
	}
}

func TestEnabledWithInvalidString(t *testing.T) {
	err := os.Setenv("CHAOS_MONKEY_ENABLED", "doutorequemtemdoutorado")
	if err != nil {
		t.Errorf("Error to set CHAOS_MONKEY_ENABLED environment variable to doutorequemtemdoutorado")
	}
	enabled := chaos.IsEnabled()
	if enabled != false {
		t.Errorf("Expected %s, got %v", "false", enabled)
	}
}

func TestAssaultsEnabledsWithNoAssaults(t *testing.T) {
	assaults := chaos.GetAssaltsEnabled()
	if len(assaults) > 0 {
		t.Errorf("Expected %v, got %v", 0, len(assaults))
	}
}

func TestAssaultsEnabledsWithOnlyLatency(t *testing.T) {
	os.Setenv("CHAOS_MONKEY_LATENCY", "true")
	os.Setenv("CHAOS_MONKEY_EXCEPTION", "false")
	os.Setenv("CHAOS_MONKEY_APP_KILLER", "false")
	os.Setenv("CHAOS_MONKEY_MEMORY", "false")

	assaults := chaos.GetAssaltsEnabled()
	if len(assaults) != 1 {
		t.Errorf("Expected %v, got %v", 1, len(assaults))
	}
	if assaults[0] != "CHAOS_MONKEY_LATENCY" {
		t.Errorf("Expected %v, got %v", "CHAOS_MONKEY_LATENCY", assaults[0])
	}
}

func TestAssaultsEnabledsWithOnlyException(t *testing.T) {
	os.Setenv("CHAOS_MONKEY_LATENCY", "false")
	os.Setenv("CHAOS_MONKEY_EXCEPTION", "true")
	os.Setenv("CHAOS_MONKEY_APP_KILLER", "false")
	os.Setenv("CHAOS_MONKEY_MEMORY", "false")

	assaults := chaos.GetAssaltsEnabled()
	if len(assaults) != 1 {
		t.Errorf("Expected %v, got %v", 1, len(assaults))
	}
	if assaults[0] != "CHAOS_MONKEY_EXCEPTION" {
		t.Errorf("Expected %v, got %v", "CHAOS_MONKEY_EXCEPTION", assaults[0])
	}
}

func TestAssaultsEnabledsWithOnlyAppKiller(t *testing.T) {
	os.Setenv("CHAOS_MONKEY_LATENCY", "false")
	os.Setenv("CHAOS_MONKEY_EXCEPTION", "false")
	os.Setenv("CHAOS_MONKEY_APP_KILLER", "true")
	os.Setenv("CHAOS_MONKEY_MEMORY", "false")

	assaults := chaos.GetAssaltsEnabled()
	if len(assaults) != 1 {
		t.Errorf("Expected %v, got %v", 1, len(assaults))
	}
	if assaults[0] != "CHAOS_MONKEY_APP_KILLER" {
		t.Errorf("Expected %v, got %v", "CHAOS_MONKEY_APP_KILLER", assaults[0])
	}
}

func TestAssaultsEnabledsWithOnlyMemory(t *testing.T) {
	os.Setenv("CHAOS_MONKEY_LATENCY", "false")
	os.Setenv("CHAOS_MONKEY_EXCEPTION", "false")
	os.Setenv("CHAOS_MONKEY_APP_KILLER", "false")
	os.Setenv("CHAOS_MONKEY_MEMORY", "true")

	assaults := chaos.GetAssaltsEnabled()
	if len(assaults) != 1 {
		t.Errorf("Expected %v, got %v", 1, len(assaults))
	}
	if assaults[0] != "CHAOS_MONKEY_MEMORY" {
		t.Errorf("Expected %v, got %v", "CHAOS_MONKEY_MEMORY", assaults[0])
	}
}

func TestAssaultsEnabledsWithAllEnabled(t *testing.T) {
	os.Setenv("CHAOS_MONKEY_LATENCY", "true")
	os.Setenv("CHAOS_MONKEY_EXCEPTION", "true")
	os.Setenv("CHAOS_MONKEY_APP_KILLER", "true")
	os.Setenv("CHAOS_MONKEY_MEMORY", "true")

	assaults := chaos.GetAssaltsEnabled()
	if len(assaults) != 4 {
		t.Errorf("Expected %v, got %v", 1, len(assaults))
	}
}

func TestIsGonnaAssaultIsBool(t *testing.T) {
	assault := chaos.IsGonnaAssault()

	if reflect.TypeOf(assault).String() != "bool" {
		t.Errorf("Expected value is type bool %v, got %v", 1, assault)
	}
}
