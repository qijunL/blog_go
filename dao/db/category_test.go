package db

//测试文件
import "testing"

func init(){
	//parseTime=true将mysql中的时间类型自动解析为go结构体中的时间类型
	dns := "root:admin@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

//获取单个分类的信息
func TestGetCategoryById(t *testing.T) {
	category,err := GetCategoryById(1)
	if err!=nil {
		panic(err)
	}
	t.Logf("category:%#v",category)
}

//获取多个分类的信息

func TestGetCategoryList(t *testing.T) {
	var categoryIds [] int64
	categoryIds = append(categoryIds,1,2,3)
	categoryList,err := GetCategoryList(categoryIds)
	if err!=nil {
		panic(err)
	}
	for _,v := range categoryList{
		t.Logf("id:%d category:%#v\n",v.CategoryId,v)
	}

}

func TestGetAllCategoryList(t *testing.T) {
	categoryList,err := GetAllCategoryList()
	if err!=nil {
		panic(err)
	}
	for _,v := range categoryList{
		t.Logf("id:%d category:%#v\n",v.CategoryId,v)
	}
}
