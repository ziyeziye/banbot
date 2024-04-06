package core

import (
	"fmt"
	"github.com/banbox/banexg/log"
	"regexp"
	"slices"
	"strings"
)

/*
GroupByPairQuotes
将[key]:pairs...输出为下面字符串
【key】
Quote: Base1 Base2 ...
*/
func GroupByPairQuotes(items map[string][]string) string {
	res := make(map[string]map[string][]string)
	for key, arr := range items {
		slices.Sort(arr)
		quoteMap := make(map[string][]string)
		for _, pair := range arr {
			baseCode, quoteCode, _, _ := SplitSymbol(pair)
			baseList, _ := quoteMap[quoteCode]
			quoteMap[quoteCode] = append(baseList, baseCode)
		}
		for quote, baseList := range quoteMap {
			slices.Sort(baseList)
			quoteMap[quote] = baseList
		}
		res[key] = quoteMap
	}
	var b strings.Builder
	for key, quoteMap := range res {
		b.WriteString(fmt.Sprintf("【%s】\n", key))
		for quoteCode, arr := range quoteMap {
			baseStr := strings.Join(arr, " ")
			b.WriteString(fmt.Sprintf("%s(%d): %s\n", quoteCode, len(arr), baseStr))
		}
	}
	return b.String()
}

/*
PrintStagyGroups
从core.StgPairTfs输出策略+时间周期的币种信息到控制台
*/
func PrintStagyGroups() {
	groups := make(map[string][]string)
	for stagy, pairMap := range StgPairTfs {
		for pair, tf := range pairMap {
			key := fmt.Sprintf("%s_%s", stagy, tf)
			arr, _ := groups[key]
			groups[key] = append(arr, pair)
		}
	}
	text := GroupByPairQuotes(groups)
	log.Info("group pairs by stagy_tf:\n" + text)
}

var (
	reCoinSplit = regexp.MustCompile("[/:-]")
	splitCache  = map[string][4]string{}
)

/*
SplitSymbol
返回：Base，Quote，Settle，Identifier
*/
func SplitSymbol(pair string) (string, string, string, string) {
	if cache, ok := splitCache[pair]; ok {
		return cache[0], cache[1], cache[2], cache[3]
	}
	parts := reCoinSplit.Split(pair, -1)
	settle, ident := "", ""
	if len(parts) > 2 {
		settle = parts[2]
	}
	if len(parts) > 3 {
		ident = parts[3]
	}
	splitCache[pair] = [4]string{parts[0], parts[1], settle, ident}
	return parts[0], parts[1], settle, ident
}
