package miraapi

import (
	"errors"
	"fmt"
	"github.com/bearsh/hid"
	"testing"
)

func TestConnect(t *testing.T) {
	if !hid.Supported() {
		t.Fatal(errors.New("hid not supported on this platform"))
	}

	deviceInfos := hid.Enumerate(Boox_Mira_VID, Boox_Mira_PID)
	if len(deviceInfos) == 0 {
		t.Fatal(errors.New("no connected device found"))
	}

	device, err := deviceInfos[0].Open()
	if err != nil {
		t.Fatal(fmt.Errorf("%s: open failed: %w", deviceInfos[0].Path, err))
	}
	t.Logf("%s: opened", deviceInfos[0].Path)

	defer func() {
		if err := device.Close(); err != nil {
			t.Logf("%s: close failed: %v", deviceInfos[0].Path, err)
		} else {
			t.Logf("%s: closed", deviceInfos[0].Path)
		}
	}()

	if err := Refresh(device); err != nil {
		t.Fatal(fmt.Errorf("%s: Refresh failed: %w", deviceInfos[0].Path, err))
	}
}
