package poc

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	yaml "gopkg.in/yaml.v2"
)

type PluginConstructor func(*YamlPlugin) Plugin

var PluginTypeNameToConstructor = map[string]PluginConstructor{
	"FLAG_VALUE": NewFlagValuePlugin,
	"FLAG_PATH":  NewFlagPathPlugin,
	"VAR_ENV":    NewVarEnvPlugin,
}

type YamlPlugin struct {
	BinaryName  string            `yaml:"binary_name"`
	Constructor PluginConstructor `yaml:"type"`
	FlagName    string            `yaml:"flag_name"`
	VarEnvName  string            `yaml:"var_env_name"`
}

type Plugin interface {
	SetPassword(string)
	Prepare() error
	InjectPassword(*exec.Cmd) error
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
	password string
}

func NewFlagValuePlugin(plugin *YamlPlugin) Plugin {
	return &FlagValuePlugin{
		flagName: plugin.FlagName,
	}
}

func (fp *FlagValuePlugin) SetPassword(password string) {
	fp.password = password
}

func (fp *FlagValuePlugin) Prepare() error {
	return nil
}

func (fp *FlagValuePlugin) InjectPassword(cmd *exec.Cmd) error {
	begin := []string{cmd.Args[0], fp.flagName, fp.password}
	cmd.Args = append(begin, cmd.Args[1:]...)
	return nil
}

func (fp *FlagValuePlugin) CleanUp() error {
	return nil
}

//
// FlagPathPlugin
//

type FlagPathPlugin struct {
	flagName         string
	password         string
	passwordFileName string
}

func NewFlagPathPlugin(plugin *YamlPlugin) Plugin {
	return &FlagPathPlugin{
		flagName: plugin.FlagName,
	}
}

func (fp *FlagPathPlugin) SetPassword(password string) {
	fp.password = password
}

func (fp *FlagPathPlugin) Prepare() error {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		return fmt.Errorf("fail to create password tmp file: %s", err)
	}

	// TODO: Maybe add a chown here (configurable in plugin settings ?)

	_, err = file.WriteString(fp.password)
	if err != nil {
		return fmt.Errorf("fail to write in password tmp file: %s", err)
	}

	_ = file.Close()

	fp.passwordFileName = file.Name()

	return nil
}

func (fp *FlagPathPlugin) InjectPassword(cmd *exec.Cmd) error {
	begin := []string{cmd.Args[0], fp.flagName, fp.passwordFileName}
	cmd.Args = append(begin, cmd.Args[1:]...)
	return nil
}

func (fp *FlagPathPlugin) CleanUp() error {
	err := os.Remove(fp.passwordFileName)
	if err != nil {
		return fmt.Errorf("fail to remove the password tmp file: %s", err)
	}
	return nil
}

//
// VarEnvPlugin
//

type VarEnvPlugin struct {
	varEnvName string
	password   string
}

func NewVarEnvPlugin(plugin *YamlPlugin) Plugin {
	return &VarEnvPlugin{
		varEnvName: plugin.VarEnvName,
	}
}

func (ve *VarEnvPlugin) SetPassword(password string) {
	ve.password = password
}

func (ve *VarEnvPlugin) Prepare() error {
	return nil
}

func (ve *VarEnvPlugin) InjectPassword(cmd *exec.Cmd) error {
	cmd.Env = append(cmd.Env, ve.varEnvName+"="+ve.password)
	return nil
}

func (ve *VarEnvPlugin) CleanUp() error {
	return nil
}
