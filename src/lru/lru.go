package lru

//缓存数据结构
type  Cache struct{
	PageIndex int //缓存页面内容(为了简化，将页面索引和页面认为是一样的)
	LastTime int //未使用时间
}

type  Lru struct{
	cacheSize int //缓存器大小
	curLen  int //当前缓存个数
	curLRUPage  int //当前最长时间未使用页面
	cacheArr []*Cache //缓存存储空间
}

func (inst *Lru) Init(cacheSize int)(){
	inst.cacheSize = cacheSize
	inst.curLen = 0
	inst.curLRUPage = -1
	inst.cacheArr = make([]*Cache, cacheSize)
	for i:=0; i<cacheSize;i++{
		inst.cacheArr[i] = nil
	}
	return
}
//新放入的缓存页面，时间设置为1
func (inst *Lru) SetCache(pageIndex int)(){
	////查看是否存在
	//exist := false
	//for i := 0; i < inst.curLen; i++{
	//	if(inst.cacheArr[i].PageIndex == pageIndex){
	//		exist = true
	//	}
	//}
	//if exist{
	//
	//}

	//有空余缓存
	if(inst.curLen < inst.cacheSize){
		inst.cacheArr[inst.curLen] = &Cache {pageIndex,0}
		inst.curLen++
	}else{
		//exchange
		inst.cacheArr[inst.curLRUPage].PageIndex =  pageIndex
		inst.cacheArr[inst.curLRUPage].LastTime = 0
	}
	//update time
	inst.updateLRUPage()
	inst.updateLastTime()
	return
}

//1 命中 0：未命中
func (inst *Lru) GetCache(pageIndex int)(page int, success int ){
	success = 0
	for i := 0; i < inst.curLen; i++{
		if(inst.cacheArr[i].PageIndex == pageIndex){
			inst.cacheArr[i].LastTime = 0
			success = 1
		}
	}
	if(1 == success){
		inst.updateLRUPage()
		inst.updateLastTime()
	}
	return page,success
}
//更新页面最长未使用时间
func (inst *Lru) updateLastTime(){
	for i := 0; i < inst.curLen; i++{
		inst.cacheArr[i].LastTime += 1
	}
}

//并同时设置最长未使用页面的索引
func (inst *Lru) updateLRUPage(){
	if(inst.curLRUPage == -1){
		inst.curLRUPage = 0
	}
	for i := 0; i < inst.curLen; i++{
		if inst.cacheArr[i].LastTime > inst.cacheArr[inst.curLRUPage].LastTime {
			inst.curLRUPage = i
		}
	}
	return
}


func (inst *Lru) Fini(){
	inst.curLen = 0
	inst.curLRUPage = -1
}
