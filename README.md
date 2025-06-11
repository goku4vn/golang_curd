# Golang CRUD with postgresql

## Start from scratch

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/google/uuid

# dev only
sudo go install github.com/air-verse/air@latest
```

### Run the project to develop

```bash
make dev
```

### Client to test the API
1. Download Insomnia from `https://insomnia.rest`

2. Import the `golang_crud_insomnia` file in the root directory of this project.