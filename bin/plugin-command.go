package bin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/conero/uymas/v2/bin/butil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
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

	// name_$name
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

// Dependent on suffix, like exe, bat, cmd.
func plgCmdListWindows() []string {
	var plugs []string
	rootName := butil.Basedir()
	appName := butil.AppName()

	// to scan file
	toScan := func(vDir string) {
		entrys, err := os.ReadDir(vDir)
		if err != nil {
			return
		}

		for _, entry := range entrys {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			ext := path.Ext(name)
			extLower := strings.ToLower(ext)
			if extLower == ".exe" {
				plugName := strings.ReplaceAll(name, ext, "")
				plugName = strings.ReplaceAll(plugName, appName+"-", "")
				plugName = strings.ReplaceAll(plugName, appName+"_", "")
				if plugName == appName {
					continue
				}
				plugs = append(plugs, plugName)
			}

		}
	}

	toScan(rootName)
	toScan(butil.RootPath("plg/"))

	return plugs
}

// Dependent on `file` that use to detect.
func plgCmdListLinux() []string {
	var plugs []string
	rootName := butil.Basedir()
	appName := butil.AppName()

	// to scan file
	toScan := func(vDir string) {
		entrys, err := os.ReadDir(vDir)
		if err != nil {
			return
		}

		for _, entry := range entrys {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			pathName := path.Join(vDir, name)

			// ELF 64-bit LSB executable
			cmd := exec.Command("file", pathName)
			rtBy, er := cmd.CombinedOutput()
			if er != nil {
				continue
			}
			rtStr := string(rtBy)
			if strings.Index(rtStr, "ELF") > -1 && strings.Index(rtStr, "LSB executable") > -1 {
				plugName := name
				plugName = strings.ReplaceAll(plugName, appName+"-", "")
				plugName = strings.ReplaceAll(plugName, appName+"_", "")
				if plugName == appName {
					continue
				}
				plugs = append(plugs, plugName)
			}

		}
	}

	toScan(rootName)
	toScan(butil.RootPath("plg/"))

	return plugs
}

// PlgCmdList Get Plugin Sub-Command list, support windows/linux/darwin.
func PlgCmdList() []string {
	var plugs []string
	switch runtime.GOOS {
	case "windows":
		return plgCmdListWindows()
	case "linux", "darwin":
		return plgCmdListLinux()

	}
	return plugs
}
