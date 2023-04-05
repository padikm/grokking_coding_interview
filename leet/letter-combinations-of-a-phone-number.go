package leet

func LetterCombinations(digits string) []string {
	m := make(map[string]string)
	m["1"] = ""
	m["2"] = "abc"
	m["3"] = "def"
	m["4"] = "ghi"
	m["5"] = "jkl"
	m["6"] = "mno"
	m["7"] = "pqrs"
	m["8"] = "tuv"
	m["9"] = "wxyz"

	fres := make([]string, 0)
	if len(digits) == 0 {
		return fres
	}

	letterCombinations1(digits, m, 0, "", &fres)
	return fres
}

func letterCombinations1(digits string, m map[string]string, i int, res string, fres *[]string) {
	if i == len(digits) {
		*fres = append(*fres, res)
		//fmt.Println(res)
		return
	}
	s := m[string(digits[i])]
	for j := 0; j < len(s); j++ {
		res = res + string(s[j])
		letterCombinations1(digits, m, i+1, res, fres)
		res = res[0 : len(res)-1]
	}
}
