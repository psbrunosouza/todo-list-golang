package workspaces

import "todo-list/internal/models"

func CreateWorkspaceService(workspace *models.Workspace) error {
	if result := CreateWorkspace(workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteWorkspaceService(workspace *models.Workspace) error {
	if result := DeleteWorkspace(workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindWorkspaceService(id int, workspace *models.Workspace) error {
	if result := FindWorkspace(id, workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListWorkspacesService(workspaces *[]models.Workspace) error {
	if result := ListWorkspace(workspaces); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateWorkspaceService(workspace *models.Workspace) error {
	if result := UpdateWorkspace(workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
