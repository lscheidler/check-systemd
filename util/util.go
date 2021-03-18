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

package util

import (
	"github.com/godbus/dbus"

	"github.com/lscheidler/check-systemd/util/dbus/service"
	"github.com/lscheidler/check-systemd/util/dbus/timer"
	"github.com/lscheidler/check-systemd/util/dbus/unit"
)

// Util holds the dbus Connection
type Util struct {
	conn *dbus.Conn
}

func New() (*Util, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	c := Util{conn: conn}
	return &c, nil
}

func (u *Util) Close() {
	u.conn.Close()
}

func (u *Util) Unit(unitName *string) (*unit.Unit, error) {
	return unit.New(u.conn, unitName)
}

func (u *Util) Service(unitName *string) (*service.Service, error) {
	return service.New(u.conn, unitName)
}

func (u *Util) Timer(unitName *string) (*timer.Timer, error) {
	return timer.New(u.conn, unitName)
}
