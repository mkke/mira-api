## HID API for Boox Mira Monitor

This uses the HID command structure information from 
https://github.com/ipodnerd3019/mira-js/blob/master/src/mira.js

### Usage

```golang
import (
    "github.com/bearsh/hid"
    mira "github.com/mkke/mira-api"
)

func main() {
    deviceInfos := hid.Enumerate(mira.BooxMiraVID, mira.BooxMiraPID)
    device, _ := deviceInfos[0].Open()
    defer device.Close()
    mira.Refresh(device)
}
```