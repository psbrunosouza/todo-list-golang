# TODO LIST APP IN GO
Golang restAPI created to learn web develoment in GO using echo framework.

# 🚀 Features
- Create tasks
- Create subtasks
- Create workspaces
- Create colors
- Create ratings
- Create users

# 📋 Database Model
To understand project structure, I've created that database model

<div style="display: flex; justify-content: center; align-items: center; flex-direction: column">
  <img alt="database_der" src="https://i.imgur.com/tu2sKmu.png"/>
</div>

# 📑 Business Rules

- To create a new workspace we need to be loogged in the application
- To create a task with spaced repetition we need to create an iteration and add the task to it
- To mark as complete any task we need check if all the the subtasks were completed too 
  and request user confirmation to confirm

# 🏁 Get Started
### Install postgres or use docker

```docker-compose up```

### Install packages

```go mod download or make install```

### Run project main.go file

```go run cmd/main.go or make run```

# 📙 License
This project is under MIT LICENSE.