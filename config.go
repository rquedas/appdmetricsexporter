package appdmetricsexporter

import (
	"fmt"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
   config.ExporterSettings `mapstructure:",squash"`
   exporterhelper.TimeoutSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
   MachineAgentHost string `mapstructure:"machineagent_host"`
   MachineAgentPort int `mapstructure:"machineagent_port"`
}


// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {
	if (cfg.MachineAgentHost == ""){
	   return fmt.Errorf("The AppD Machine Agent host is required")
	} 
	return nil
 }