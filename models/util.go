package models

import "strconv"

func convertUrlValuesToInt(values []string) []int {
	ret := make([]int, 0)

	for _, str := range values {
		if val, err := strconv.Atoi(str); err != nil {
			ret = append(ret, val)
		}
	}

	return ret
}

func convertUrlValuesToLangArr(values []string) []SubmissionLang {
	langs := make([]SubmissionLang, len(values))

	for ind, val := range values {
		langs[ind] = SubmissionLang(val)
	}

	return langs
}

func convertUrlValuesToStatusArr(values []string) []SubmissionStatus {
	statuses := make([]SubmissionStatus, len(values))

	for ind, val := range values {
		statuses[ind] = SubmissionStatus(val)
	}

	return statuses
}
