package bin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/conero/uymas/bin/butil"
	"os/exec"
	"time"
)

const (
	PlgCmdGetTitle   = "plg-cmd-title"
	PlgCmdGetProfile = "plg-cmd--profile"
)

type PlgCProfile struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ExecName    string `json:"execName,omitempty"`
}

func (c *PlgCProfile) GetRunCmd(args ...string) (*exec.Cmd, error) {
	if c.ExecName == "" {
		return nil, errors.New("可执行路径为空(ExecName)！")
	}

	ctx, _ := context.WithTimeout(context.Background(), 300*time.Second)
	//defer cancel()

	cmd := exec.CommandContext(ctx, c.ExecName, args...)
	return cmd, nil
}

func (c *PlgCProfile) Run(args ...string) ([]byte, error) {
	cmd, err := c.GetRunCmd(args...)
	if err != nil {
		return nil, err
	}

	return cmd.CombinedOutput()

}

// PluginCommand Hot swap command (sub command)
type PluginCommand struct {
	Name     string
	Descript string
	*CLI
}

func NewPluginCommand() *PluginCommand {
	pc := &PluginCommand{
		CLI: NewCLI(),
	}

	pc.RegisterFunc(func(arg *Arg) {
		profile := PlgCProfile{
			Name:        pc.Name,
			Description: pc.Descript,
		}
		bys, _ := json.Marshal(profile)
		fmt.Print(string(bys))
	}, PlgCmdGetProfile)

	return pc
}

// The later optimization is for the first detection, and caching is performed after success for quick recall next time
func plgCmdDetect(cc *Arg) *PlgCProfile {
	name := cc.Command
	if name == "" {
		return nil
	}

	// $name
	pBin := butil.RootPath(name)
	cmd := exec.Command(pBin, PlgCmdGetProfile)
	rtBy, err := cmd.CombinedOutput()
	var plg *PlgCProfile
	if err == nil {
		err = json.Unmarshal(rtBy, &plg)
		if err == nil {
			plg.ExecName = pBin
			return plg
		}
	}

	// plg/$name
	pBin = butil.RootPath(fmt.Sprintf("plg/%s", name))
	cmd = exec.Command(pBin, PlgCmdGetProfile)
	rtBy, err = cmd.CombinedOutput()
	if err == nil {
		err = json.Unmarshal(rtBy, &plg)
		if err == nil {
			plg.ExecName = pBin
			return plg
		}
	}

	appName := butil.AppName()

	// name-$name
	pBin = butil.RootPath(fmt.Sprintf("%s-%s", appName, name))
	cmd = exec.Command(pBin, PlgCmdGetProfile)
	rtBy, err = cmd.CombinedOutput()
	if err == nil {
		err = json.Unmarshal(rtBy, &plg)
		if err == nil {
			plg.ExecName = pBin
			return plg
		}
	}

	// name-$name
	pBin = butil.RootPath(fmt.Sprintf("%s_%s", appName, name))
	cmd = exec.Command(pBin, PlgCmdGetProfile)
	rtBy, err = cmd.CombinedOutput()
	if err == nil {
		err = json.Unmarshal(rtBy, &plg)
		if err == nil {
			plg.ExecName = pBin
			return plg
		}
	}

	return plg
}
