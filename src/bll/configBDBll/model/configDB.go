package model

// config数据库对象
type ConfigDB struct {
	// 名称
	ConfigKey string `gorm:"column:ConfigKey"`

	// 值
	ConfigValue string  `gorm:"column:ConfigValue"`

	// 描述
	ConfigDesc string  `gorm:"column:ConfigDesc"`
}

// TableName 获取表名
func (thisObj *ConfigDB) TableName() string {
	return "b_config"
}

// CheckData 检查对象是否正确
func (thisObj *ConfigDB) CheckData() []string {
	return nil
}