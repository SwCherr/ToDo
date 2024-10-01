all: clean mocks

mocks: clean
	mockgen -source=pkg/service/service.go -destination=mocks/service/mock_service.go
	mockgen -source=pkg/repository/repository.go -destination=mocks/repository/mock_repository.go
	
test: mocks
	go test --coverprofile=c.out ./pkg/handler/

gcov: test
	go tool cover -html=c.out

clean: 
	rm -rf mocks/
	rm -rf c.out

.PHONY: clean

# -source определяет исходный файл, из которого будет сгенерирован mock-объект.
# -destination определяет файл, в который будет записан сгенерированный mock-объект.
# -package определяет имя пакета, в котором будет находиться сгенерированный mock-объект.