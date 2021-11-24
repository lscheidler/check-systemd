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

package unit

import (
	"time"

	"github.com/godbus/dbus"
)

// Unit struct providing some properties from dbus interface org.freedesktop.systemd1.Unit
type Unit struct {
	UnitPath   dbus.ObjectPath
	DbusObject dbus.BusObject

	ID                     string
	LoadState              string
	ActiveState            string
	SubState               string
	UnitFileState          string
	ActiveEnterTimestamp   time.Time
	ActiveExitTimestamp    time.Time
	InactiveEnterTimestamp time.Time
	InactiveExitTimestamp  time.Time
}

// New gets the dbus properties for unitName and returns a struct
func New(conn *dbus.Conn, unitName *string) (*Unit, error) {
	u := &Unit{}

	obj := conn.Object("org.freedesktop.systemd1", "/org/freedesktop/systemd1")

	var unitPath dbus.ObjectPath
	var err error

	if err := obj.Call("org.freedesktop.systemd1.Manager.GetUnit", 0, unitName).Store(&unitPath); err != nil {
		return nil, err
	}
	u.UnitPath = unitPath

	u.DbusObject = conn.Object("org.freedesktop.systemd1", unitPath)
	if u.ID, err = u.GetPropertyString("org.freedesktop.systemd1.Unit.Id"); err != nil {
		return nil, err
	}
	if u.LoadState, err = u.GetPropertyString("org.freedesktop.systemd1.Unit.LoadState"); err != nil {
		return nil, err
	}
	if u.ActiveState, err = u.GetPropertyString("org.freedesktop.systemd1.Unit.ActiveState"); err != nil {
		return nil, err
	}
	if u.SubState, err = u.GetPropertyString("org.freedesktop.systemd1.Unit.SubState"); err != nil {
		return nil, err
	}
	if u.UnitFileState, err = u.GetPropertyString("org.freedesktop.systemd1.Unit.UnitFileState"); err != nil {
		return nil, err
	}
	if u.InactiveEnterTimestamp, err = u.GetPropertyUInt64ToUnix("org.freedesktop.systemd1.Unit.InactiveEnterTimestamp"); err != nil {
		return nil, err
	}
	if u.InactiveExitTimestamp, err = u.GetPropertyUInt64ToUnix("org.freedesktop.systemd1.Unit.InactiveExitTimestamp"); err != nil {
		return nil, err
	}
	if u.ActiveEnterTimestamp, err = u.GetPropertyUInt64ToUnix("org.freedesktop.systemd1.Unit.ActiveEnterTimestamp"); err != nil {
		return nil, err
	}
	if u.ActiveExitTimestamp, err = u.GetPropertyUInt64ToUnix("org.freedesktop.systemd1.Unit.ActiveExitTimestamp"); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *Unit) GetProperty(path string) (dbus.Variant, error) {
	return u.DbusObject.GetProperty(path)
}

func (u *Unit) GetPropertyString(path string) (string, error) {
	property, err := u.DbusObject.GetProperty(path)
	if err != nil {
		return "", err
	}
	return property.Value().(string), nil
}

func (u *Unit) GetPropertyBool(path string) (bool, error) {
	property, err := u.DbusObject.GetProperty(path)
	if err != nil {
		return false, err
	}
	return property.Value().(bool), nil
}

func (u *Unit) GetPropertyUInt64ToInt(path string) (int, error) {
	property, err := u.DbusObject.GetProperty(path)
	if err != nil {
		return 0, err
	}
	return int(property.Value().(uint64)), nil
}

func (u *Unit) GetPropertyInt32ToInt(path string) (int, error) {
	property, err := u.DbusObject.GetProperty(path)
	if err != nil {
		return 0, err
	}
	return int(property.Value().(int32)), nil
}

func (u *Unit) GetPropertyUInt64ToUnix(path string) (time.Time, error) {
	property, err := u.DbusObject.GetProperty(path)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(property.Value().(uint64)), 0), nil
}
