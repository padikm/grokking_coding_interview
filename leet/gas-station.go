package leet

func CanCompleteCircuit(gas []int, cost []int) int {
	return canCompleteCircuit(gas, cost)
}

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	for i := 0; i < l; i++ {
		cGas := gas[i]
		cPos := i % l
		cCost := cost[i]
		j := (i + 1) % l
		if cGas < cCost {
			continue
		}
		for j != cPos {
			if cGas >= cCost {
				cGas = cGas - cCost + gas[j]
				cCost = cost[j]
			}

			if cGas < cCost {
				break
			}
			j = (j + 1) % l
		}

		if j == cPos {
			return cPos
		}

	}
	return -1
}

func CanCompleteCircuit1(gas []int, cost []int) int {
	if sum(cost) > sum(gas) {
		return -1
	}

	toTalSum := 0
	res := 0
	for i := 0; i < len(gas); i++ {
		toTalSum = toTalSum + gas[i] - cost[i]
		if toTalSum < 0 {
			toTalSum = 0
			res = i + 1
		}
	}
	return res

}

func sum(a []int) int {
	res := 0
	for i := 0; i < len(a); i++ {
		res += a[i]
	}
	return res
}
