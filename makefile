BINARY_NAME = todo
BUILD_DIR = bin
MODULE_PATH = .

.PHONY: all build run clean

all: build

build:
	@echo "TODO 앱 빌드중"
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MODULE_PATH)
	@echo "빌드 완료: $(BUILD_DIR)/$(BINARY_NAME)"

run: build
	@echo "TODO 앱 실행중"
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "빌드 디렉토리 정리중"
	rm -rf $(BUILD_DIR)
	@echo "정리 완료"	