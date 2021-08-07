package plugin

import (
	"github.com/spiral/errors"
	"github.com/spiral/roadrunner/v2/plugins/config"
	"github.com/spiral/roadrunner/v2/plugins/logger"
)

// PluginName is the name for the plugin as found inside the container
const PluginName = "plugin"

// Plugin is the structure that will be initiated as per endure.Container interface
type Plugin struct {
	logger logger.Logger
	cfg    Config
}

// Init initiates the plugin with any injected services implementing endure.Container, returning an error if the
// plugin fails to start, if the error is of type errors.Disabled then the plugin will not be active
func (p *Plugin) Init(cfg config.Configurer, logger logger.Logger) error {
	const op = errors.Op("custom_plugin_init")

	// if the config does not have a section matching PluginName section
	// then return an error
	if !cfg.Has(PluginName) {
		return errors.E(op, errors.Disabled)
	}

	// read in the section of the config by the plugin name
	// if it cannot be read then return an error
	err := cfg.UnmarshalKey(PluginName, &p.cfg)
	if err != nil {
		return errors.E(op, errors.Disabled, err)
	}

	p.logger = logger

	p.cfg.InitDefaults()

	return nil
}

// Name returns endure.Named interface implementation
func (p *Plugin) Name() string {
	return PluginName
}

// Action is purely an example for the interface
func (p *Plugin) Action() error {
	p.logger.Debug("action was called")

	return nil
}

// RPC provides a struct instance that allows for methods to be called via the RPC server
// in roadrunner
func (p *Plugin) RPC() interface{} {
	return &rpc{srv: p}
}
