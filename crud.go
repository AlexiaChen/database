package database

import (
	"database/sql"
	"gorm.io/gorm"
)

type Crud struct {
	Db *gorm.DB
}
type Sort struct {
	FieldName string `json:"fieldName"`
	Order     string `json:"order"`
}
type Search struct {
	FieldName string `json:"fieldName"`
	Value     string `json:"value"`
}

//var (
//	DefaultSort = []Sort{
//		{
//			FieldName: "created_at",
//			Order:     "desc",
//		},
//	}
//)

func (c *Crud) Create(model interface{}) error {
	return c.Db.Create(model).Error
}
func (c *Crud) QueryOne(filter interface{}) error {
	return c.Db.Where(*&filter).First(filter).Error
}
func (c *Crud) Delete(filter interface{}) error {
	return c.Db.Delete(filter).Error
}
func (c *Crud) Save(model interface{}) error {
	return c.Db.Save(model).Error
}

//func (c *Crud) QueryList(page, size int, count *int64, sort []Sort, searchFieldName []string, searchValue string, filter, result interface{}) {
//	db := c.Db
//	if sort == nil {
//		db = c.setSort(db, DefaultSort)
//	} else {
//		db = c.setSort(db, sort)
//	}
//	if searchValue != "" {
//		db = c.setSearch(db, searchFieldName, searchValue)
//	}
//	if filter != nil {
//		db = db.Where(filter)
//	}
//	db.Find(result).Count(count)
//	db.Offset((page - 1) * size).Limit(size).Find(result)
//}
//func (c *Crud) setSort(db *gorm.DB, sort []Sort) *gorm.DB {
//	for _, v := range sort {
//		db = db.Order(v.FieldName + " " + v.Order)
//	}
//	return db
//}
//func (c *Crud) setSearch(db *gorm.DB, fieldName []string, value string) *gorm.DB {
//	for k, v := range fieldName {
//		if k == 0 {
//			db = db.Where(v+" LIKE ?", "%"+value+"%")
//		} else {
//			db = db.Or(v+" LIKE ?", "%"+value+"%")
//		}
//	}
//	return db
//}

func GetResult(rows *sql.Rows) []map[string]interface{} {
	var result []map[string]interface{}
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			panic(err.Error())
		}

		// 创建一个切片，用于存储每一行的列值
		values := make([]interface{}, len(columns))
		for i, _ := range columns {
			values[i] = new(interface{})
		}

		// 将每一行的列值存储到 values 中
		if err := rows.Scan(values...); err != nil {
			panic(err.Error())
		}

		// 创建一个 map，将列名和对应的值存储到 map 中
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			rowMap[colName] = *(values[i].(*interface{}))
		}

		// 将每一行的 map 存储到 result 中
		result = append(result, rowMap)
	}
	return result
}
