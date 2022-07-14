package util

/*
* Product - 多个数组的排列组合
*
* PARAMS:
*   sets - 二维数组
*
* RETURNS:
*   排列组合数组
 */
func Product(sets [][]any) [][]any {
    lens := func(i int) int { return len(sets[i]) }
    product := [][]any{}
    for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
        var r []any
        for j, k := range ix {
            r = append(r, sets[j][k])
        }
        product = append(product, r)
    }
    return product
}

func nextIndex(ix []int, lens func(i int) int) {
    for j := len(ix) - 1; j >= 0; j-- {
        ix[j]++
        if j == 0 || ix[j] < lens(j) {
            return
        }
        ix[j] = 0
    }
}
