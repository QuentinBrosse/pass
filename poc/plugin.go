package poc

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"

	yaml "gopkg.in/yaml.v2"
)

type PluginConstructor func(*YamlPlugin) Plugin

var PluginTypeNameToConstructor = map[string]PluginConstructor{
	"FLAG_VALUE": NewFlagValuePlugin,
	// "FLAG_PATH":  NewFlagPath,
	"VAR_ENV": NewVarEnvPlugin,
}

type YamlPlugin struct {
	BinaryName  string            `yaml:"binary_name"`
	Constructor PluginConstructor `yaml:"type"`
	FlagName    string            `yaml:"flag_name"`
	VarEnvName  string            `yaml:"var_env_name"`
}

type Plugin interface {
	Prepare() error
	InjectPassword(*exec.Cmd, string) error
	CleanUp() error
}

func (c *PluginConstructor) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw string

	err := unmarshal(&raw)
	if err != nil {
		return err
	}

	constructor, ok := PluginTypeNameToConstructor[raw]
	if !ok {
		return fmt.Errorf("unknown plugin type name: %s", raw)
	}

	*c = constructor
	return nil
}

func NewPluginFromConfig(name string, file io.Reader) (Plugin, error) {
	yamlPlugin := &YamlPlugin{}

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("fail to read yamlPlugin file %s: %s", name, err)
	}

	err = yaml.Unmarshal(raw, yamlPlugin)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal yamlPlugin file %s: %s", name, err)
	}

	yamlPlugin.BinaryName = name
	plugin := yamlPlugin.Constructor(yamlPlugin)

	return plugin, nil
}

//
// FlagValuePlugin
//

type FlagValuePlugin struct {
	flagName string
}

func NewFlagValuePlugin(plugin *YamlPlugin) Plugin {
	return &FlagValuePlugin{
		flagName: plugin.FlagName,
	}
}

func (fv *FlagValuePlugin) Prepare() error {
	return nil
}

func (fv *FlagValuePlugin) InjectPassword(cmd *exec.Cmd, password string) error {
	begin := []string{cmd.Args[0], fv.flagName, password}
	cmd.Args = append(begin, cmd.Args[1:]...)
	return nil
}

func (fv *FlagValuePlugin) CleanUp() error {
	return nil
}

//
// VarEnvPlugin
//

type VarEnvPlugin struct {
	varEnvName string
}

func NewVarEnvPlugin(plugin *YamlPlugin) Plugin {
	return &VarEnvPlugin{
		varEnvName: plugin.VarEnvName,
	}
}

func (ve *VarEnvPlugin) Prepare() error {
	return nil
}

func (ve *VarEnvPlugin) InjectPassword(cmd *exec.Cmd, password string) error {
	cmd.Env = append(cmd.Env, ve.varEnvName+"="+password)
	return nil
}

func (ve *VarEnvPlugin) CleanUp() error {
	return nil
}
