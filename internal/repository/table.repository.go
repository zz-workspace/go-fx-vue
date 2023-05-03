package repository

import (
	"encoding/json"
	"strings"

	"fast-api.io/helpers"
	"fast-api.io/models"
	"github.com/fatih/structs"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Element struct {
	Elem datatypes.JSON `json:"elem"`
}

type TableRepository struct {
	Db *gorm.DB
}

func InitTableRepository(Db *gorm.DB) *TableRepository {
	repository := &TableRepository{
		Db: Db,
	}
	return repository
}

func (c TableRepository) CreateTable(table *models.Table) *models.Table {
	c.Db.Create(&table)
	return table
}

func (c TableRepository) FindTableById(id uint64) (models.Table, error) {
	var table models.Table
	err := c.Db.Where("id = ?", id).First(&table).Error
	return table, err
}

func (c TableRepository) FindAllTable(workspaceID uint64) ([]models.Table, helpers.Paginator) {
	var tables []models.Table
	var paginator helpers.Paginator
	// c.Db.Scopes(helpers.Paging(&paginator, &tables)).Find(&tables)
	c.Db.Where(&models.Table{WorkspaceID: workspaceID}).Find(&tables)
	return tables, paginator
}

func (c TableRepository) NewContent(id int, content interface{}) bool {
	print(id)
	var emptyStruct struct{}
	m := structs.Map(emptyStruct)
	m["data"] = content
	out, _ := json.Marshal(m["data"])
	print(string(out))
	tx := c.Db.Exec(`UPDATE tables SET contents = contents || ? WHERE id = ?`, string(out), id)
	if tx.Error != nil {
		return false
	} else {
		return true
	}
}

func (c TableRepository) FindContentByTable() []Element {
	var elements []Element
	// c.Db.Raw(`SELECT elem FROM tables CROSS JOIN jsonb_array_elements(contents) AS elem WHERE elem->>'id' = '2' LIMIT 100`).Scan(&elements)
	c.Db.Raw(`SELECT elem FROM tables CROSS JOIN jsonb_array_elements(contents) AS elem LIMIT 100`).Scan(&elements)
	return elements
}

func (c TableRepository) SearchContent(tableId int, searchConditions string) []Element {
	var elements []Element
	sql := []string{"SELECT elem FROM tables CROSS JOIN jsonb_array_elements(contents) AS elem WHERE id = ? AND ", searchConditions, " LIMIT 100"}
	c.Db.Raw(strings.Join(sql, ""), tableId).Scan(&elements)
	return elements
}

func (c TableRepository) UpdateContent(tableId int, searchConditions string, dataUpdate string) []Element {
	var elements []Element
	sql := []string{"UPDATE tables CROSS JOIN jsonb_array_elements(contents) AS elem SET ", dataUpdate, " WHERE id = ? AND ", searchConditions, " LIMIT 100"}
	c.Db.Raw(strings.Join(sql, ""), tableId).Scan(&elements)
	return elements
}

func (c TableRepository) Update(table *models.Table) *models.Table {
	c.Db.Save(&table)
	return table
}
