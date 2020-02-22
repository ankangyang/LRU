package main

import (
	"fmt"
	"lru"
	"math/rand"
)

func main(){
	fmt.Printf("test lru \n")

	var sumVistor int = 1000 //访问次数
	var countSuccess int = 0 //命中次数
	var pageArr []int = make([]int, 0, sumVistor)
	var inst lru.Lru
	inst.Init(3)
	//初始化所有访问的页面
	//var allPage = []int{1,2,3,4,5,6,7,8,9,10} //作为于IO慢的存储
	//随机访问页面
	for i := 0; i < sumVistor; i++ {
		//随机访问10个页面
		pageIndex := rand.Intn(10) + 1 //即将要加载的页面
		//尝试缓存中获取
		_,success := inst.GetCache(pageIndex)
		pageArr = append(pageArr, pageIndex)
		if 1 == success{
			countSuccess++
		}else{
			//IO慢存储中访问
			//...
			//访问的页面放入缓存中
			inst.SetCache(pageIndex)
		}
	}

	//打印历史访问页面
	fmt.Println("访问历史页面:")
	for _,page := range pageArr {
		fmt.Printf(" %d ",page)
	}
	//命中次数
	fmt.Printf("\n访问次数：%d  命中次数：%d \n", sumVistor, countSuccess )
	inst.Fini()

	return
}


