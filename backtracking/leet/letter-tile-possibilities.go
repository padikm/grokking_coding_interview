package leet

func NumTilePossibilities(tiles string) int {
	final := make(map[string]bool)
	used := make(map[int]bool)
	numTilePossibilities1(tiles, []uint8{}, &final, &used)
	return len(final)
}

func numTilePossibilities1(tiles string, res []uint8, final *map[string]bool, used *map[int]bool) {
	if len(res) != 0 {
		(*final)[string(res)] = true
	}
	for i := 0; i < len(tiles); i++ {
		if (*used)[i] {
			continue
		}
		(*used)[i] = true
		res = append(res, tiles[i])
		numTilePossibilities1(tiles, res, final, used)
		(*used)[i] = false
		res = res[0 : len(res)-1]
	}
}
