package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// TodoManager 구조체 정의
// 이 구조체는 Todo 항목을 관리하는 기능을 포함하고 있습니다.
type TodoManager struct {
	Todos    []Todo `json:"todos"`     // model 패키지의 Todo 항목 슬라이스
	FilePath string `json:"file_path"` // Todo 항목을 저장할 파일 경로
	NextID   int    `json:"next_id"`   // 다음 Todo 항목의 ID
}

// TodoManager 생성자 함수
func NewTodoManager(filePath string) *TodoManager {
	return &TodoManager{
		Todos:    make([]Todo, 0), // 빈 Todo 슬라이스로 초기화
		FilePath: filePath,        // 파일 경로 설정
		NextID:   1,               // 다음 ID를 1로 초기화
	}
}

// json 파일로 부터 Todo 항목을 로드하는 함수
func (tm *TodoManager) LoadFromFile() error {
	// 파일이 있는지 확인하고, 있다면 파일 내용을 바이트배열로 읽는다.
	if _, err := os.Stat(tm.FilePath); os.IsNotExist(err) {
		return nil // 파일이 없으면 에러 없이 종료
	}

	// 파일을 읽어와서 바이트 배열로 저장
	data, err := os.ReadFile(tm.FilePath)
	if err != nil {
		return fmt.Errorf("파일 읽기 실패: %v", err)
	}

	if len(data) == 0 {
		return nil // 파일이 비어있으면 에러 없이 종료
	}

	// JSON 데이터를 TodoManager 구조체로 언마샬링
	err = json.Unmarshal(data, tm)

	if err != nil {
		return fmt.Errorf("JSON 파싱 실패: %v", err)
	}

	return nil
}

// TodoManager 구조체를 json 파일로 저장하는 함수
func (tm *TodoManager) SaveToFile() error {
	data, err := json.MarshalIndent(tm, "", "  ")

	if err != nil {
		return fmt.Errorf("JSON 직렬화 실패: %v", err)
	}

	err = os.WriteFile(tm.FilePath, data, 0644)

	if err != nil {
		return fmt.Errorf("파일 쓰기 실패: %v", err)
	}

	return nil
}

// 신규 Todo 등록 함수
func (tm *TodoManager) AddTodo(title, description string) {
	todo := Todo{
		ID:          tm.NextID,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tm.Todos = append(tm.Todos, todo) // Todo 항목을 슬라이스에 추가
	tm.NextID++                       // 다음 ID를 증가시킴

	fmt.Printf("할일이 추가되었습니다. (ID: %d)\n", todo.ID)
}

// Todo 목록 출력 함수
func (tm *TodoManager) ListTodos(filter string) {
	var filteredTodos []Todo
	var count int

	// 필터링 로직
	for _, todo := range tm.Todos {
		switch filter {
		case "completed":
			if todo.Completed {
				filteredTodos = append(filteredTodos, todo)
			}
		case "pending":
			if !todo.Completed {
				filteredTodos = append(filteredTodos, todo)
			}
		default: // "all"
			filteredTodos = append(filteredTodos, todo)
		}
	}

	// 필터링된 Todo 항목 출력
	for _, todo := range filteredTodos {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}

		fmt.Printf("[%d] %s %s\n", todo.ID, status, todo.Title)
		if todo.Description != "" {
			fmt.Printf("    설명: %s\n", todo.Description)
		}
		fmt.Printf("    생성일: %s\n", todo.CreatedAt.Format(time.RFC3339))
		fmt.Println("---------------------------------------------------")
		count++
	}

	if count == 0 {
		fmt.Println("할 일이 없습니다.")
	} else {
		fmt.Printf("총 %d개의 할 일이 있습니다.\n", count)
	}
}

// Todo 항목의 완료 상태를 토글하는 함수
func (tm *TodoManager) Toggle(id int) error {
	idx, err := tm.findTodoIndex(id)
	if err != nil {
		return err // Todo 항목을 찾지 못한 경우 에러 리턴
	}

	todo := &tm.Todos[idx]           // Todo 항목의 포인터를 가져옴
	todo.Completed = !todo.Completed // 완료 상태를 토글
	todo.UpdatedAt = time.Now()      // 업데이트 시간 갱신

	return nil
}

// Todo 항목을 수정하는 함수
func (tm *TodoManager) UpdateTodo(id int, title, description string) error {
	idx, err := tm.findTodoIndex(id)
	if err != nil {
		return err // Todo 항목을 찾지 못한 경우 에러 리턴
	}

	todo := &tm.Todos[idx] // Todo 항목의 포인터를 가져옴

	if todo.Completed {
		return fmt.Errorf("완료된 할 일은 수정할 수 없습니다")
	}

	todo.Title = title
	todo.Description = description
	todo.UpdatedAt = time.Now() // 업데이트 시간 갱신

	return nil
}

// Todo 항목을 삭제하는 함수
func (tm *TodoManager) DeleteTodo(id int) error {
	idx, err := tm.findTodoIndex(id)
	if err != nil {
		return err // Todo 항목을 찾지 못한 경우 에러 리턴
	}

	// 해당 인덱스의 Todo 항목을 슬라이스에서 삭제
	tm.Todos = append(tm.Todos[:idx], tm.Todos[idx+1:]...)
	return nil
}

// Todo 항목의 인덱스를 ID로 찾아서 리턴 하는 함수
// Todo 객체를 리턴하지 않는 이유는 객체를 전달시점에서 슬라이스 내부의 배열 주소가 변경될 수 있기 때문입니다.
// 따라서 인덱스를 찾아서 해당 인덱스의 Todo 객체를 직접 수정하는 방식으로 구현합니다.
// ID가 존재하지 않는 경우 에러를 리턴합니다.
func (tm *TodoManager) findTodoIndex(id int) (int, error) {
	for i, todo := range tm.Todos {
		if todo.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("ID %d에 해당하는 할 일이 없습니다", id)
}
