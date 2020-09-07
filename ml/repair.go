package ml

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type ContentDistrict struct {
	ContentId int    `gorm:"primary_key"`
	Title string
	Content string
	CrawlSourceId string 
}

//func (ContentDistrict) TableName() string  {
//	return "tb_content_district"
//}

var db *gorm.DB
var err error

func init()  {
	// https://www.liwenzhou.com/posts/Go/gorm_crud/ GORM CRUD指南
	// https://www.liwenzhou.com/posts/Go/gorm/  GORM 入门指南
	fmt.Println("开始修复数据")

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tb_"+defaultTableName
	}

	db, err = gorm.Open("mysql","root:root@tcp(192.168.33.10:3306)/test?charset=utf8&parseTime=True&loc=Local")

	dealError(err)
	//设置全局表名禁用复数
	db.SingularTable(true)

	fmt.Println("connection succedssed")

	//defer db.Close()
}

func RepairData()  {
	//根据标题查询到采集ID  然后让中间表的district_id都对应到那个contentId
	var titles = make([]string,0)
	titles = append(titles, "hello","world","nihao")

	for _,v := range titles {
		fmt.Println(v)
		res := findByTitle(v)
		fmt.Println(res)
	}

	defer db.Close()


	//contentDistrict := TbContentDistrict{Title: "hello",Content: "content",CrawlSourceId: "123456"}
	//db.Create(&contentDistrict)
}

func findByTitle(title string) []int  {

	contentDistrict := new (ContentDistrict)
	//db.Create(&contentDistrict)

	dt := db.Debug().Where("title = ?", title).First(&contentDistrict)
	fmt.Printf("%v",dt)


	//db.Delete(&contentDistrict)
	//db.Debug().Unscoped().Delete(&contentDistrict)
	//db.Model(&contentDistrict).Where("title = ?", true).Update("name", "hello")


	res := []int{1,2}
	return res
}

func dealError(err error)  {
	fmt.Println("error",err)
}
