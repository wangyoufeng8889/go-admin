package batterymanage
import (
	orm "go-admin/global"
	"go-admin/models"
	"go-admin/tools"
	"time"
)
//dtu为主，dtu pkg脱离后只保留 dtu
type Dtu_aliyun struct {
	Dtu_aliyunId     int    `json:"dtu_aliyunId" gorm:"size:10;primary_key;AUTO_INCREMENT"`
	Dtu_uptime time.Time  `json:"dtu_uptime"`
	Dtu_id      string `json:"dtu_id" gorm:"size:20;primary_key;"`
	Dtu_aliyunStatus uint8    `json:"dtu_aliyunStatus" gorm:"Type：uint8"`
	DataScope  string `json:"dataScope" gorm:"-"`
	UpdateBy  string `gorm:"size:128;" json:"updateBy"`
	models.BaseModel
}
func (Dtu_aliyun) TableName() string {
	return "user_dtu_aliyun"
}
func (e *Dtu_aliyun) GetDtu_aliyun(is_oneList string) ([]Dtu_aliyun,int, error) {
	var doc []Dtu_aliyun

	table := orm.Eloquent.Select("*").Table(e.TableName())
	if e.Dtu_aliyunId != 0 {
		table = table.Where("dtu_aliyun_id = ?", e.Dtu_aliyunId)
	}
	if e.Dtu_id != "" {
		table = table.Where("dtu_id = ?", e.Dtu_id)
	}
	// 数据权限控制
	dataPermission := new(models.DataPermission)
	dataPermission.UserId, _ = tools.StringToInt(e.DataScope)
	table, err := dataPermission.GetDataScope(e.TableName(), table)
	if err != nil {
		return nil, 0, err
	}
	var count int
	if is_oneList == "YES" {
		if err := table.Order("dtu_uptime desc").First(&doc).Error; err != nil {
			return nil, 0, err
		}
		table.Where("`deleted_at` IS NULL").Count(&count)
	}else{
		if err := table.Order("dtu_aliyun_id").Find(&doc).Error; err != nil {
			return nil, 0, err
		}
		table.Where("`deleted_at` IS NULL").Count(&count)
	}
	return doc, count, nil
}