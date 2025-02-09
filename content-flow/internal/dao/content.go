package dao

import (
	"content-flow/internal/model"
	"fmt"
	"gorm.io/gorm"
	"hash/fnv"
	"log"
	"math/big"
)

const contentNumTables = 4

type ContentDao struct {
	db *gorm.DB
}

func NewContentDao(db *gorm.DB) *ContentDao {
	return &ContentDao{db: db}
}

func (c *ContentDao) First(contentID string) (*model.ContentDetail, error) {
	var detail model.ContentDetail
	if err := c.db.
		Table(getContentDetailsTable(contentID)).
		Where("content_id = ?", contentID).First(&detail).Error; err != nil {
		log.Printf("content first error = %v", err)
		return &detail, nil
	}
	return &detail, nil
}

func (c *ContentDao) UpdateByID(contentID string, column string, value interface{}) error {
	if err := c.db.Table(getContentDetailsTable(contentID)).
		Where("content_id = ?", contentID).
		Update(column, value).Error; err != nil {
		log.Printf("content by id update error = %v", err)
		return err
	}
	return nil
}

func getContentDetailsTable(contentID string) string {
	tableIndex := getContentTableIndex(contentID)
	table := fmt.Sprintf("cms_content.content_details_%d", tableIndex)
	log.Printf("content_id = %s, table = %s \n", contentID, table)
	return table
}

func getContentTableIndex(uuid string) int {
	// 计算UUID的哈希值
	hash := fnv.New32a()
	_, _ = hash.Write([]byte(uuid))
	hashValue := hash.Sum32()
	// 将哈希值映射到表的索引范围内
	// 创建一个大整数对象
	bigNum := big.NewInt(int64(hashValue))
	bigModulo := big.NewInt(contentNumTables)
	tableIndex := bigNum.Mod(bigNum, bigModulo).Int64()
	return int(tableIndex)
}
