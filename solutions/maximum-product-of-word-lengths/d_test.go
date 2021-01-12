/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package maximum_product_of_word_lengths

import "testing"

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  int
	}{
		{
			words: []string{
				"bdcecbcadca",
				"caafd",
				"bcadc",
				"eaedfcd",
				"fcdecf",
				"dee",
				"bfedd",
				"ffafd",
				"eceaffa",
				"caabe",
				"fbdb",
				"acafbccaa",
				"cdc",
				"ecfdebaafde",
				"cddbabf",
				"adc",
				"cccce",
				"cbbe",
				"beedf",
				"fafbfdcb",
				"ceecfabedbd",
				"aadbedeaf",
				"cffdcfde",
				"fbbdfdccce",
				"ccada",
				"fb",
				"fa",
				"ec",
				"dddafded",
				"accdda",
				"acaad",
				"ba",
				"dabe",
				"cdfcaa",
				"caadfedd",
				"dcdcab",
				"fadbabace",
				"edfdb",
				"dbaaffdfa",
				"efdffceeeb",
				"aefdf",
				"fbadcfcc",
				"dcaeddd",
				"baeb",
				"beddeed",
				"fbfdffa",
				"eecacbbd",
				"fcde",
				"fcdb",
				"eac",
				"aceffea",
				"ebabfffdaab",
				"eedbd",
				"fdeed",
				"aeb",
				"fbb",
				"ad",
				"bcafdabfbdc",
				"cfcdf",
				"deadfed",
				"acdadbdcdb",
				"fcbdbeeb",
				"cbeb",
				"acbcafca",
				"abbcbcbaef",
				"aadcafddf",
				"bd",
				"edcebadec",
				"cdcbabbdacc",
				"adabaea",
				"dcebf",
				"ffacdaeaeeb",
				"afedfcadbb",
				"aecccdfbaff",
				"dfcfda",
				"febb",
				"bfffcaa",
				"dffdbcbeacf",
				"cfa",
				"eedeadfafd",
				"fcaa",
				"addbcad",
				"eeaaa",
				"af",
				"fafc",
				"bedbbbdfae",
				"adfecadcabe",
				"efffdaa",
				"bafbcbcbe",
				"fcafabcc",
				"ec",
				"dbddd",
				"edfaeabecee",
				"fcbedad",
				"abcddfbc",
				"afdafb",
				"afe",
				"cdad",
				"abdffbc",
				"dbdbebdbb"},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
