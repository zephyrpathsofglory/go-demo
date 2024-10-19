/*
od 科目一 第一次考试第三题
无线网络设备有 ssid 编号，每个设备有一个坐标和一个信号覆盖范围半径r，如果设备B在设备A 的覆盖范围内（B的坐标在以 A 为圆心，半径 Ar 的圆中，
那么设备B的将被设置成设备A的 ssid，如果将2个设备用网线连接，那么2个设备可以使用同一个ssid。给定一组设备，允许使用 1 根不限制长度的网线，
连接2个设备。使得相同 ssid 设备的数量最大。给出最大设备数量。
*/

// 记录题目，解答不对
package od

import (
	"math"
	"sort"
)

type Device struct {
	X uint
	Y uint
	R uint
}

type Devices []Device

func (d Devices) Len() int {
	return len(d)
}

func (d Devices) Swap(x, y int) {
	d[x], d[y] = d[y], d[x]
}

func (d Devices) Less(x, y int) bool {
	return d[x].R < d[y].R || (d[x].R == d[y].R && d[x].X > d[y].X) ||
		(d[x].R == d[y].R && d[x].X == d[y].X && d[x].Y > d[y].Y)
}

func maxDevices(devices Devices) uint {
	sort.Sort(devices)

	deviceStack := make(Devices, 0, len(devices))
	deviceStackIdx := make([]int, 0, len(devices))
	seen := make(map[int]bool, len(devices))
	for idx, d := range devices {
		if seen[idx] {
			continue
		}

		deviceStack = append(deviceStack, d)
		deviceStackIdx = append(deviceStackIdx, idx)

		for len(deviceStack) > 0 {
			dev := deviceStack[0]
			devIdx := deviceStackIdx[0]
			for i, dd := range devices {
				if seen[i] || devIdx == i {
					continue
				}

				if math.Pow(
					float64(dev.X-dd.X),
					2,
				)+math.Pow(
					float64(dev.X-dd.Y),
					2,
				) <= math.Pow(
					float64(dev.R),
					2,
				) {
					deviceStack = append(deviceStack, dd)
					deviceStackIdx = append(deviceStackIdx, i)
				}
			}
		}

		for i := idx + 1; i < len(devices); i++ {
		}

		seen[idx] = true
	}

	return 0
}
