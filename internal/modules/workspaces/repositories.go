package workspaces

import (
	"todo-list/internal/entities"

	"gorm.io/gorm"
)

type WorkspaceRepository interface {
	Create(workspace *entities.Workspace) *gorm.DB
	List(workspaces *[]entities.Workspace) *gorm.DB
	Find(id int, workspace *entities.Workspace) *gorm.DB
	Update(id int, workspace *entities.Workspace) *gorm.DB
	Delete(workspace *entities.Workspace) *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(workspace *entities.Workspace) (result *gorm.DB) {
	return r.db.Create(workspace)
}

func (r *repository) List(workspaces *[]entities.Workspace) (result *gorm.DB) {
	return r.db.Preload("Tasks").Preload("Tasks.Workspace").Find(workspaces)
}

func (r *repository) Find(id int, workspace *entities.Workspace) (result *gorm.DB) {
	return r.db.Preload("Tasks").Preload("Tasks.Workspace").First(workspace, id)
}

func (r *repository) Update(id int, workspace *entities.Workspace) (result *gorm.DB) {
	r.db.Model(workspace).Where("id = ?", id).Updates(workspace)
	return r.db.Preload("Tasks").Preload("Tasks.Workspace").First(workspace, id)
}

func (r *repository) Delete(workspace *entities.Workspace) (result *gorm.DB) {
	return r.db.Unscoped().Delete(workspace)
}
