/*
某店铺将用于组成套餐的商品记作字符串 goods，其中 goods[i] 表示对应商品。请返回该套餐内所含商品的 全部排列方式 。
goods 中可能有重复商品

返回结果 无顺序要求，但不能含有重复的元素。



示例 1:

输入：goods = "agew"
输出：["aegw","aewg","agew","agwe","aweg","awge","eagw","eawg","egaw","egwa","ewag","ewga","gaew","gawe","geaw","gewa",
"gwae","gwea","waeg","wage","weag","wega","wgae","wgea"]


提示：

1 <= goods.length <= 8
*/

package permute

func goodsOrder(goods string) []string {
	var dfs func(string, string)

	var res []string

	dfs = func(preset string, remain string) {
		if len(remain) == 1 {
			res = append(res, preset+remain)
			return
		}

		appearance := make(map[string]bool)

		for idx, c := range remain {
			if !appearance[string(c)] {
				dfs(preset+string(c), remain[0:idx]+remain[idx+1:])
				appearance[string(c)] = true
			}
		}
	}

	dfs("", goods)

	return res
}
