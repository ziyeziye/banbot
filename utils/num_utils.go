package utils

import (
	"fmt"
	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
	"math"
	"slices"
)

const thresFloat64Eq = 1e-9

/*
NumSign 获取数字的方向；1，-1或0
*/
func NumSign(obj interface{}) int {
	if val, ok := obj.(int); ok {
		if val > 0 {
			return 1
		} else if val < 0 {
			return -1
		} else {
			return 0
		}
	} else if val, ok := obj.(float32); ok {
		if val > 0 {
			return 1
		} else if val < 0 {
			return -1
		} else {
			return 0
		}
	} else if val, ok := obj.(float64); ok {
		if val > 0 {
			return 1
		} else if val < 0 {
			return -1
		} else {
			return 0
		}
	} else {
		panic(fmt.Errorf("invalid type for NumSign: %t", obj))
	}
}

/*
EqualNearly 判断两个float是否近似相等，解决浮点精读导致不等
*/
func EqualNearly(a, b float64) bool {
	return EqualIn(a, b, thresFloat64Eq)
}

/*
EqualIn 判断两个float是否在一定范围内近似相等
*/
func EqualIn(a, b, thres float64) bool {
	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}
	return math.Abs(a-b) <= thres
}

func ConvertFloat64(i interface{}) float64 {
	switch v := i.(type) {
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		return 0
	}
}

func ConvertInt64(i interface{}) int64 {
	switch v := i.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	default:
		return 0
	}
}

// 计算两个数的最大公约数（GCD）的函数，使用欧几里得算法
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

/*
GcdMultiple
计算一个切片中所有元素的最大公约数
*/
func GcdInts(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // 没有数时返回0
	}

	// 从第一个数字开始，逐个将其与后面的数字进行最大公约数计算
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = gcd(result, numbers[i])
	}
	return result
}

type Cluster struct {
	Center float64
	Items  []float64
}

type ClusterRes struct {
	Clusters []Cluster
	RowGIds  []int
}

func KMeansVals(vals []float64, num int) *ClusterRes {
	if len(vals) == 0 {
		return nil
	}
	if num == 1 {
		sumVal := float64(0)
		for _, v := range vals {
			sumVal += v
		}
		avgVal := sumVal / float64(len(vals))
		return &ClusterRes{
			Clusters: []Cluster{{Center: avgVal, Items: vals}},
			RowGIds:  make([]int, len(vals)),
		}
	}
	// 输入值域在0~1之间
	minVal := slices.Min(vals)
	scale := 1 / (slices.Max(vals) - minVal)
	if len(vals) == 1 {
		scale = 1 / minVal
	}
	offset := 0 - minVal*scale
	var d clusters.Observations
	for _, val := range vals {
		d = append(d, clusters.Coordinates{val*scale + offset})
	}
	// 进行聚类
	km := kmeans.New()
	groups, err_ := km.Partition(d, num)
	if err_ != nil {
		return nil
	}
	slices.SortFunc(groups, func(a, b clusters.Cluster) int {
		return int((a.Center[0] - b.Center[0]) * 1000)
	})
	// 生成返回结果
	resList := make([]Cluster, 0, len(groups))
	seps := make([]float64, 0, len(groups))
	for i, group := range groups {
		var center = (group.Center[0] - offset) / scale
		var items = make([]float64, 0, len(group.Observations))
		for _, it := range group.Observations {
			coords := it.Coordinates()
			items = append(items, (coords[0]-offset)/scale)
		}
		resList = append(resList, Cluster{
			Center: center,
			Items:  items,
		})
		curMax := slices.Max(items)
		curMin := slices.Min(items)
		if len(seps) > 0 {
			seps[i-1] = (seps[i-1] + curMin) / 2
		}
		seps = append(seps, curMax)
	}
	// 计算每个项所属的分组
	rowGids := make([]int, 0, len(vals))
	for _, v := range vals {
		gid := len(groups) - 1
		for i, end := range seps {
			if v < end {
				gid = i
				break
			}
		}
		rowGids = append(rowGids, gid)
	}
	return &ClusterRes{
		Clusters: resList,
		RowGIds:  rowGids,
	}
}
