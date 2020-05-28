// Copyright 2016-2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package link

import (
	"fmt"

	"os"
	"log"

	"github.com/vishvananda/netlink"
)

// DeleteByName deletes the interface with the name ifName.
func DeleteByName(ifName string) error {
	iface, err := netlink.LinkByName(ifName)
	if err != nil {
		return fmt.Errorf("failed to lookup %q: %v", ifName, err)
	}

	if err = netlink.LinkDel(iface); err != nil {
		return fmt.Errorf("failed to delete %q: %v", ifName, err)
	}

	return nil
}

// Rename renames a network link
func Rename(curName, newName string) error {
	logFileName := "/users/sqi009/cilium-start-time.log"
	logFile, _  := os.OpenFile(logFileName,os.O_RDWR|os.O_APPEND|os.O_CREATE,0644)
	defer logFile.Close()
	debugLog := clog.New(logFile,"[Info: link.go]",clog.Lmicroseconds)
	debugLog.Println("[cilium] netlink.LinkByName(curName) start")

	link, err := netlink.LinkByName(curName)
	if err != nil {
		return err
	}
	debugLog.Println("[cilium] netlink.LinkSetName(link, newName) start")
	return netlink.LinkSetName(link, newName)
}
