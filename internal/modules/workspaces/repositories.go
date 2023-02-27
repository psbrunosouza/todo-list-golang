package workspaces

import (
	"todo-list/internal/databases"

	"gorm.io/gorm"
)

func CreateWorkspace(workspace *Workspace) (result *gorm.DB) {
	return databases.PostgresDB.Create(workspace)
}

func ListWorkspace(workspaces *[]Workspace) (result *gorm.DB) {
	return databases.PostgresDB.Find(workspaces)
}

func FindWorkspace(id int, workspace *Workspace) (result *gorm.DB) {
	return databases.PostgresDB.First(workspace, id)
}

func UpdateWorkspace(workspace *Workspace) (result *gorm.DB) {
	return databases.PostgresDB.Save(workspace)
}

func DeleteWorkspace(workspace *Workspace) (result *gorm.DB) {
	return databases.PostgresDB.Unscoped().Delete(workspace)
}
