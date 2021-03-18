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

package timer

import (
	"time"

	"github.com/godbus/dbus"

	"github.com/lscheidler/check-systemd/util/dbus/unit"
)

// Timer struct providing some properties from dbus interface org.freedesktop.systemd1.Timer
type Timer struct {
	Unit                    *unit.Unit
	NextElapseUSecRealtime  time.Time
	NextElapseUSecMonotonic int
	Result                  string
}

// New gets the dbus properties for unitName and returns a struct
func New(conn *dbus.Conn, unitName *string) (*Timer, error) {
	su := &Timer{}
	var err error

	if su.Unit, err = unit.New(conn, unitName); err != nil {
		return nil, err
	}

	if su.NextElapseUSecRealtime, err = su.Unit.GetPropertyUInt64ToUnix("org.freedesktop.systemd1.Timer.NextElapseUSecRealtime"); err != nil {
		return nil, err
	}

	if su.NextElapseUSecMonotonic, err = su.Unit.GetPropertyUInt64ToInt("org.freedesktop.systemd1.Timer.NextElapseUSecMonotonic"); err != nil {
		return nil, err
	}

	if su.Result, err = su.Unit.GetPropertyString("org.freedesktop.systemd1.Timer.Result"); err != nil {
		return nil, err
	}

	return su, nil
}
