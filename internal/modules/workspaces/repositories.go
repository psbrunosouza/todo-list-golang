package workspaces

import (
	"todo-list/internal/databases"
	"todo-list/internal/models"

	"gorm.io/gorm"
)

func CreateWorkspace(workspace *models.Workspace) (result *gorm.DB) {
	return databases.PostgresDB.Create(workspace)
}

func ListWorkspace(workspaces *[]models.Workspace) (result *gorm.DB) {
	return databases.PostgresDB.Preload("Tasks").Preload("Tasks.Workspace").Find(workspaces)
}

func FindWorkspace(id int, workspace *models.Workspace) (result *gorm.DB) {
	return databases.PostgresDB.Preload("Tasks").Preload("Tasks.Workspace").First(workspace, id)
}

func UpdateWorkspace(id int, workspace *models.Workspace) (result *gorm.DB) {
	databases.PostgresDB.Model(workspace).Where("id = ?", id).Updates(workspace)
	return FindWorkspace(id, workspace)
}

func DeleteWorkspace(workspace *models.Workspace) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(workspace)
}
