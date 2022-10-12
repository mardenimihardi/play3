run:
	@echo "RUNNING..."
	@go run ./app.go

remockgen:
	@echo "===Remockgen==="
	@mockgen -source=./internal/repository/repository.go -destination=./testfile/repository/mock_repository.go
	@mockgen -source=./internal/usecase/usecase.go -destination=./testfile/usecase/mock_usecase.go
	@echo "====Done==="