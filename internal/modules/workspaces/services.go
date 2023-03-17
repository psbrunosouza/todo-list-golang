package workspaces

import "todo-list/internal/entities"

func CreateWorkspaceService(workspace *entities.Workspace) error {
	if result := CreateWorkspace(workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func DeleteWorkspaceService(workspace *entities.Workspace) error {
	if result := DeleteWorkspace(workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func FindWorkspaceService(id int, workspace *entities.Workspace) error {
	if result := FindWorkspace(id, workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func ListWorkspacesService(workspaces *[]entities.Workspace) error {
	if result := ListWorkspace(workspaces); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func UpdateWorkspaceService(id int, workspace *entities.Workspace) error {
	if result := UpdateWorkspace(id, workspace); result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
