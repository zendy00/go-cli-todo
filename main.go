package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// flag 초기화
	var (
		add    = flag.Bool("add", false, "할 일을 추가합니다.")
		delete = flag.Bool("delete", false, "할 일을 삭제합니다.")
		update = flag.Bool("update", false, "할 일을 수정합니다.")
		toggle = flag.Bool("toggle", false, "완료 상태 토글")
		list   = flag.Bool("list", false, "할 일 목록을 출력합니다.")
		help   = flag.Bool("help", false, "사용법을 출력합니다.")

		id     = flag.Int("id", 0, "할 일 ID")
		title  = flag.String("title", "", "할 일 제목")
		desc   = flag.String("desc", "", "할 일 설명")
		filter = flag.String("filter", "all", "할 일 필터 (all|completed|pending)")
	)

	flag.Parse()

	// 도움말 출력
	if *help || len(os.Args) == 1 {
		printUsage()
		return
	}

	// TodoManager 인스턴스 생성
	todoManager := NewTodoManager("todos.json")

	// 저장된 데이터 읽기
	if err := todoManager.LoadFromFile(); err != nil {
		fmt.Printf("데이터 로드 중 오류 발생: %v\n", err)
		os.Exit(1)
	}

	// 명령어 처리
	var err error

	switch {
	case *add:
		if *title == "" {
			fmt.Println("[오류] 제목을 입력해야 합니다.(-title=\"<제목>\")")
			os.Exit(1)
		}
		todoManager.AddTodo(*title, *desc)

	case *list:
		todoManager.ListTodos(*filter)

	case *update:
		err = todoManager.UpdateTodo(*id, *title, *desc)
		fmt.Printf("할일이 수정되었습니다. (ID: %d)\n", *id)

	case *delete:
		err = todoManager.DeleteTodo(*id)
		fmt.Printf("할일이 삭제되었습니다. (ID: %d)\n", *id)

	case *toggle:
		if *id == 0 {
			fmt.Println("[오류] 완료할 할 일 ID를 입력해야 합니다.(-id)")
			os.Exit(1)
		}
		err = todoManager.Toggle(*id)
		fmt.Printf("할일이 완료 상태가 변경되었습니다. (ID: %d)\n", *id)

	default:
		printUsage()
		return
	}

	if err != nil {
		fmt.Printf("[오류] 작업 수행 중 오류 발생: %v\n", err)
		os.Exit(1)
	}

	// CUD 데이터 파일로 저장.
	if err := todoManager.SaveToFile(); err != nil {
		fmt.Printf("데이터 저장 중 오류 발생: %v\n", err)
		os.Exit(1)
	}

}

func printUsage() {
	println("TODO Manager - 명령행 할일 관리자 사용법:")
	println("  todo -add -title=\"<제목>\" [-desc=\"<설명>\"] - 할 일을 추가합니다.")
	println("  todo -list [-filter=all|completed|pending] 할 일 목록을 출력합니다.")
	println("  todo -done -id=<ID> - 완료된 할 일을 표시합니다.")
	println("  todo -delete -id=<ID> - 할 일을 삭제합니다.")
	println("  todo -update -id=<ID> -title=\"<제목>\" [-desc=\"<설명>\"] - 할 일을 수정합니다.")
	println("  todo -help - 사용법을 출력합니다.")
}
