package repo

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql驱动
	"github.com/jinzhu/gorm"
	"zkzj_resource_project/pkg/question/model"
)

type Repo struct{
	db *gorm.DB
}


//初始化repo
func NewRepo() *Repo{
	db := initDb()
	return &Repo{db:db}
}

//获取题目数量
func (r *Repo) GetQuestionNum(questionId int32)int32{

	var count int32

	r.db.Model(&model.QuestionSingleRelT{}).Where("ques_id = ?",questionId).Count(&count)

	fmt.Println(count)

	return count
}


//初始化数据库
func initDb()*gorm.DB{
	//配置MySQL连接参数
	username := "root"           //账号
	password := "Root2020@@"     //密码
	host := "47.119.140.240"     //数据库地址，可以是Ip或者域名
	port := 3306                //数据库端口
	Dbname := "zk_resource_db"  //数据库名
	timeout := "20s"            //连接超时，10秒
	var err error
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	Db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if Db.Error != nil {
		fmt.Printf("database error %v", Db.Error)
	}
	return Db
}

