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

package service

import (
	"time"

	"github.com/godbus/dbus"

	"github.com/lscheidler/check-systemd/util/dbus/unit"
)

// Service struct providing some properties from dbus interface org.freedesktop.systemd1.Service
type Service struct {
	Unit *unit.Unit

	Type                            string
	RemainAfterExit                 bool
	ExecMainStartTimestamp          time.Time
	ExecMainStartTimestampMonotonic int
	ExecMainExitTimestamp           time.Time
	ExecMainExitTimestampMonotonic  int
	ExecMainCode                    int
	ExecMainStatus                  int
	StatusText                      string
	Result                          string
}

// New gets the dbus properties for unitName and returns a struct
func New(conn *dbus.Conn, unitName *string) (*Service, error) {
	su := &Service{}

	var err error

	if su.Unit, err = unit.New(conn, unitName); err != nil {
		return nil, err
	}

	if su.Type, err = su.Unit.GetPropertyString("org.freedesktop.systemd1.Service.Type"); err != nil {
		return nil, err
	}

	if su.RemainAfterExit, err = su.Unit.GetPropertyBool("org.freedesktop.systemd1.Service.RemainAfterExit"); err != nil {
		return nil, err
	}

	if su.ExecMainStartTimestamp, err = su.Unit.GetPropertyUInt64ToUnix("org.freedesktop.systemd1.Service.ExecMainStartTimestamp"); err != nil {
		return nil, err
	}

	if su.ExecMainStartTimestampMonotonic, err = su.Unit.GetPropertyUInt64ToInt("org.freedesktop.systemd1.Service.ExecMainStartTimestampMonotonic"); err != nil {
		return nil, err
	}

	if su.ExecMainExitTimestamp, err = su.Unit.GetPropertyUInt64ToUnix("org.freedesktop.systemd1.Service.ExecMainExitTimestamp"); err != nil {
		return nil, err
	}

	if su.ExecMainExitTimestampMonotonic, err = su.Unit.GetPropertyUInt64ToInt("org.freedesktop.systemd1.Service.ExecMainExitTimestampMonotonic"); err != nil {
		return nil, err
	}

	if su.ExecMainCode, err = su.Unit.GetPropertyInt32ToInt("org.freedesktop.systemd1.Service.ExecMainCode"); err != nil {
		return nil, err
	}

	if su.ExecMainStatus, err = su.Unit.GetPropertyInt32ToInt("org.freedesktop.systemd1.Service.ExecMainStatus"); err != nil {
		return nil, err
	}

	if su.StatusText, err = su.Unit.GetPropertyString("org.freedesktop.systemd1.Service.StatusText"); err != nil {
		return nil, err
	}

	if su.Result, err = su.Unit.GetPropertyString("org.freedesktop.systemd1.Service.Result"); err != nil {
		return nil, err
	}

	return su, nil
}
