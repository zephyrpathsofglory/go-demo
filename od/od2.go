package od

import "sort"

/*
某文件系统中有N个目录，每个目录都有一个独一无二的ID。每个目录只有一个父目录，但每个父目录下可以有零个或者多个子目录，目录结构呈树状结构。

假设，根目录的ID为0，且根目录没有父目录，其他所有目录的ID用唯一的正整数表示，并统一编号。

现给定目录ID和其父目录ID的对应父子关系表[子目录ID，父目录ID]，以及一个待删除的目录ID，请计算并返回一个ID序列，表示因为删除指定目录后剩下的所有目录，返回的ID序列以递增序输出。

注意：

被删除的目录或文件编号一定在输入的ID序列中。
当一个目录删除时，它所有的子目录都会被删除。

二、输入描述
输入的第一行为父子关系表的长度m;

接下来的m行为m个父子关系对;

最后一行为待删除的ID。序列中的元素以空格分割。

三、输出描述
输出一个序列，表示因为删除指定目录后，剩余的目录ID。
*/

func deleteDirectory(dirTree [][2]int, dirName int) []int {
	treeMap := make(map[int][]int, len(dirTree)+1)
	for _, pair := range dirTree {
		child, parent := pair[0], pair[1]
		if len(treeMap[parent]) == 0 {
			treeMap[parent] = []int{child}
		} else {
			treeMap[parent] = append(treeMap[parent], child)
		}
	}

	var removeDir func(int)
	removeDir = func(d int) {
		for _, i := range treeMap[d] {
			removeDir(i)
		}
		delete(treeMap, d)
	}

	removeDir(dirName)

	dirs := make(map[int]bool, len(dirTree)+1)
	for k, v := range treeMap {
		dirs[k] = true
		for _, i := range v {
			dirs[i] = true
		}
	}

	delete(dirs, dirName)

	sortDir := make([]int, 0, len(dirs))

	for k := range dirs {
		sortDir = append(sortDir, k)
	}

	sort.Ints(sortDir)

	return sortDir
}
