/*
Copyright 2021 Lars Eric Scheidler

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/lscheidler/check-systemd/util"
	"github.com/lscheidler/go-nagios"
)

func main() {
	var unitName = flag.String("u", "", "unit")
	flag.Parse()

	nagios := nagios.New()
	defer nagios.Exit()

	if strings.Compare(*unitName, "") == 0 {
		nagios.Unknown("Option -u must be set.")
		return
	}

	nagios.SetName(*unitName)

	u, err := util.New()
	if err != nil {
		nagios.Critical(fmt.Sprintf("%#v\n", err))
		panic(err)
	}
	defer u.Close()

	isTimer := strings.HasSuffix(*unitName, ".timer")

	if !isTimer {
		if u, err := u.Service(unitName); err == nil {
			checkState(nagios, "Loaded", u.Unit.LoadState, "loaded")
			checkState(nagios, "Active", u.Unit.ActiveState, "active")
			checkState(nagios, "Sub", u.Unit.SubState, "running")
			checkState(nagios, "Enabled", u.Unit.UnitFileState, "enabled")
		} else {
			nagios.Critical(fmt.Sprintf("%#v\n", err))
			panic(err)
		}
	} else {
		if u, err := u.Timer(unitName); err == nil {
			checkState(nagios, "Loaded", u.Unit.LoadState, "loaded")
			checkState(nagios, "Active", u.Unit.ActiveState, "active")
			checkStates(nagios, "Sub", u.Unit.SubState, []string{"waiting", "running"})
			checkState(nagios, "Enabled", u.Unit.UnitFileState, "enabled")
		} else {
			nagios.Critical(fmt.Sprintf("%#v\n", err))
			panic(err)
		}

		if isTimer {
			serviceUnitName := strings.Replace(*unitName, ".timer", ".service", 1)
			if u, err := u.Service(&serviceUnitName); err == nil {
				checkState(nagios, "Service Loaded", u.Unit.LoadState, "loaded")
				if u.ExecMainStatus == 0 {
					nagios.Ok(fmt.Sprintf("Service Status: %d", u.ExecMainStatus))
				} else {
					nagios.Critical(fmt.Sprintf("Service Status: %d (%s)", u.ExecMainStatus, u.StatusText))
				}
			} else {
				nagios.Critical(fmt.Sprintf("%#v\n", err))
				panic(err)
			}
		}
	}

	// for .timer, check and additional .service
	//    ActiveEnterTimestamp: 1613589472948566
	//    ActiveExitTimestamp: 0
	//    InactiveEnterTimestamp: 0
	//    InactiveExitTimestamp:1613589472898947
}

func checkState(nagios *nagios.Nagios, name string, currentState string, desiredState string) {
	if strings.Compare(currentState, desiredState) == 0 {
		nagios.Ok(name + ": " + currentState)
	} else {
		nagios.Critical(name + ": " + currentState)
	}
}

func checkStates(nagios *nagios.Nagios, name string, currentState string, desiredStates []string) {
	for _, desiredState := range desiredStates {
		if strings.Compare(currentState, desiredState) == 0 {
			nagios.Ok(name + ": " + currentState)
			return
		}
	}
	nagios.Critical(name + ": " + currentState)
}
