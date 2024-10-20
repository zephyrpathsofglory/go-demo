/*
双遥控电路：初始时3个开关均为断开，给定操作员操作 records，每个record 为 [time, switchID], 表示更改了 switch 状态，设备在 limit 时间内
没有操作会自动断开，如果自动断开时恰好收到操作指令，则该switch闭合，返回最后一次操作完成后，且所有switch均断开后，设备的累计工作时长。

switch 1 闭合 切 switch 2 或者 switch 3 中的一个闭合，则设备工作

注意： 每个时间点是可以操作多个设备的
*/

package od

var (
	switchState = map[int]bool{
		1: false,
		2: false,
		3: false,
	}

	switchOnMap = map[int]int{}
	deviceOn    = false
)

func DeviceWorkingTime(records [][]int, limit int) int {
	var tm int
	switchOnTime := 0

	for _, record := range records {
		t := record[0]
		operateSwitch := record[1]

		if t != tm {
			// 过了 1 个或者 多个 时间单位
			for i := tm + 1; i <= t; i++ {
				tm = i
				if deviceOn {
					switchOnTime += 1
				}
				checkState(tm - limit)
				checkDeviceOn()
			}
		}

		switchState[operateSwitch] = !switchState[operateSwitch]
		if switchState[operateSwitch] {
			switchOnMap[operateSwitch] = tm
		} else {
			delete(switchOnMap, operateSwitch)
		}
		checkDeviceOn()
	}

	if deviceOn {
		for i := 1; i <= limit; i++ {
			switchOnTime += 1
			checkState(tm + i - limit)
			checkDeviceOn()
			if !deviceOn {
				break
			}
		}
	}

	return switchOnTime
}

func checkDeviceOn() {
	if switchState[1] && (switchState[2] || switchState[3]) {
		deviceOn = true
		return
	}

	deviceOn = false
}

func checkState(thre int) {
	for k, v := range switchOnMap {
		if v == thre {
			delete(switchOnMap, k)
			switchState[k] = false
		}
	}
}
