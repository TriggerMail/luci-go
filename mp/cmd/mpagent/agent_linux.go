// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/common/logging"
)

type LinuxStrategy struct {
}

// chown changes ownership of a path.
func (LinuxStrategy) chown(ctx context.Context, username, path string) error {
	user, err := user.Lookup(username)
	if err != nil {
		return err
	}
	uid, err := strconv.Atoi(user.Uid)
	if err != nil {
		return err
	}
	gid, err := strconv.Atoi(user.Gid)
	if err != nil {
		return err
	}
	return os.Chown(path, uid, gid)
}

// configureAutoMount mounts the specified disk and configures mount on startup.
//
// Assumes the disk is already formatted as ext4.
func (LinuxStrategy) configureAutoMount(ctx context.Context, disk string) error {
	// Wait for the disk to be attached.
	for {
		err := exec.Command("blkid", disk).Run()
		if err == nil {
			logging.Infof(ctx, "Attached: %s.", disk)
			break
		}
		s := err.(*exec.ExitError).Sys().(syscall.WaitStatus)
		if s.ExitStatus() != 2 {
			// 2 means the specified device wasn't found.
			// Keep waiting if 2, otherwise return an error.
			return err
		}
		time.Sleep(60 * time.Second)
		logging.Infof(ctx, "Waiting for disk to be attached...")
	}
	// Configure auto-mount using fstab.
	f, err := os.OpenFile("/etc/fstab", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), disk) {
			logging.Infof(ctx, "Already mounted: %s.", disk)
			return nil
		}
	}
	logging.Infof(ctx, "Mounting: %s.", disk)
	// Ensure the cursor is at the end so the new line is appended.
	f.Seek(0, io.SeekEnd)
	line := []byte(fmt.Sprintf("%s /b ext4 defaults,nobootwait,nofail 0 2\n", disk))
	if _, err := f.Write(line); err != nil {
		return err
	}
	// Ensure the file is closed before mount reads from it.
	f.Close()
	return exec.Command("/bin/mount", "--all").Run()
}

// reboot reboots the machine.
func (LinuxStrategy) reboot(ctx context.Context) error {
	return exec.Command("/sbin/shutdown", "-r", "now").Run()
}

type SystemdStrategy struct {
	LinuxStrategy
}

// enableSwarming enables installed service.
func (SystemdStrategy) enableSwarming(ctx context.Context) error {
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		return err
	}
	return exec.Command("systemctl", "enable", "swarming-start-bot").Run()
}

// start starts the agent.
func (SystemdStrategy) start(ctx context.Context, _ string) error {
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		return err
	}
	if err := exec.Command("systemctl", "enable", "machine-provider-agent").Run(); err != nil {
		return err
	}
	return exec.Command("systemctl", "start", "machine-provider-agent").Run()
}

// stop stops all instances of the agent.
func (SystemdStrategy) stop(ctx context.Context) error {
	return exec.Command("systemctl", "stop", "machine-provider-agent").Run()
}

type UpstartStrategy struct {
	LinuxStrategy
}

// enableSwarming enables installed service.
func (UpstartStrategy) enableSwarming(ctx context.Context) error {
	return nil
}

// start starts the agent.
func (UpstartStrategy) start(ctx context.Context, _ string) error {
	if err := exec.Command("initctl", "reload-configuration").Run(); err != nil {
		return err
	}
	return exec.Command("start", "machine-provider-agent").Run()
}

// stop stops all instances of the agent.
func (UpstartStrategy) stop(ctx context.Context) error {
	return exec.Command("stop", "machine-provider-agent").Run()
}

// systemdFound returns true iff systemd is supported on this machine.
func systemdFound() bool {
	return exec.Command("which", "systemctl").Run() == nil
}

// upstartFound returns true iff upstart is supported on this machine.
func upstartFound() bool {
	return exec.Command("init", "--version").Run() == nil
}

// getAgent returns an agent which runs on Linux, depending on supported init systems.
func getAgent(ctx context.Context) (*Agent, error) {
	if systemdFound() {
		agent := Agent{
			agentAutoStartPath:        "/etc/systemd/system/machine-provider-agent.service",
			agentAutoStartTemplate:    "machine-provider-agent.service.tmpl",
			logsDir:                   "/var/log/machine-provider-agent",
			swarmingAutoStartPath:     "/etc/systemd/system/swarming-start-bot.service",
			swarmingAutoStartTemplate: "swarming-start-bot.service.tmpl",
			swarmingBotDir:            "/b/s",
			strategy:                  SystemdStrategy{},
		}
		logging.Infof(ctx, "Using systemd Linux agent.")
		return &agent, nil
	}

	if upstartFound() {
		agent := Agent{
			agentAutoStartPath:        "/etc/init/machine-provider-agent.conf",
			agentAutoStartTemplate:    "machine-provider-agent.conf.tmpl",
			logsDir:                   "/var/log/messages/machine-provider-agent",
			swarmingAutoStartPath:     "/etc/init/swarming-start-bot.conf",
			swarmingAutoStartTemplate: "swarming-start-bot.conf.tmpl",
			swarmingBotDir:            "/b/s",
			strategy:                  UpstartStrategy{},
		}
		logging.Infof(ctx, "Using Upstart Linux agent.")
		return &agent, nil
	}

	return nil, errors.New("unsupported init system, expected systemd or Upstart")
}
