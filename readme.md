# 📝 CLI TODO Manager

<div align="center">

![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macOS%20%7C%20windows-lightgrey.svg)

**간단하고 효율적인 명령행 기반 할일 관리자**

[설치하기](#-설치) •
[사용법](#-사용법) •
[기능](#-주요-기능) •
[예시](#-사용-예시)

</div>

---

## 🚀 소개

CLI TODO Manager는 Go로 작성된 가볍고 빠른 명령행 할일 관리 도구입니다. 복잡한 GUI 없이 터미널에서 바로 할일을 관리할 수 있어 개발자들에게 특히 유용합니다.

### ✨ 왜 CLI TODO Manager인가?

- 🚀 **빠른 속도**: 가벼운 바이너리로 즉시 실행
- 💾 **로컬 저장**: JSON 파일로 데이터를 안전하게 보관
- 🔧 **간단한 사용법**: 직관적인 명령어 체계
- 🎯 **개발자 친화**: 터미널 환경에 최적화
- 📱 **크로스 플랫폼**: Windows, macOS, Linux 지원

---

## 📦 설치

### Go가 설치된 경우

```bash
# 레포지토리 클론
git clone https://github.com/zendy00/go-cli-todo.git
cd go-cli-todo

# Makefile 기반 빌드
make build

# 시스템 PATH에 추가 (선택사항)
sudo mv bin/todo /usr/local/bin/
```

---

## 🎯 주요 기능

| 기능            | 설명                | 명령어                               |
| --------------- | ------------------- | ------------------------------------ |
| ➕ **할일 추가** | 새로운 할일 생성    | `todo -add -title="..." -desc="..."` |
| 📋 **목록 조회** | 할일 목록 확인      | `todo -list`                         |
| ✅ **상태 변경** | 완료/미완료 토글    | `todo -toggle -id=1`                 |
| ✏️ **내용 수정** | 할일 제목/설명 변경 | `todo -update -id=1 -title="..."`    |
| 🗑️ **할일 삭제** | 불필요한 할일 제거  | `todo -delete -id=1`                 |
| 🔍 **필터링**    | 상태별 할일 조회    | `todo -list -filter=completed`       |

---

## 🛠️ 사용법

### 기본 명령어

```bash
# 도움말 보기
todo -help

# 할일 추가
todo -add -title="Go 프로젝트 완성하기" -desc="CLI TODO 앱 개발"

# 전체 목록 보기
todo -list

# 완료된 할일만 보기
todo -list -filter=completed

# 미완료 할일만 보기
todo -list -filter=pending

# 할일 완료 표시
todo -toggle -id=1

# 할일 수정
todo -update -id=1 -title="새로운 제목" -desc="새로운 설명"

# 할일 삭제
todo -delete -id=1
```

### 필터 옵션

- `all` (기본값): 모든 할일 표시
- `completed`: 완료된 할일만 표시
- `pending`: 미완료 할일만 표시

---

## 📸 사용 예시

### 할일 추가하기
```bash
$ todo -add -title="Go 언어 학습" -desc="기본 문법부터 고급 기능까지"
할일이 추가되었습니다. (ID: 1)
```

### 목록 확인하기
```bash
$ todo -list
==================== 할일 목록 ====================
[1] ❌ Go 언어 학습
    설명: 기본 문법부터 고급 기능까지
    생성일: 2025-06-27 14:30:15
---------------------------------------------------
[2] ✅ 프로젝트 설계
    설명: TODO 앱 구조 설계 완료
    생성일: 2025-06-27 14:25:10
---------------------------------------------------
총 2개의 할일이 있습니다.
```

### 상태 변경하기
```bash
$ todo -toggle -id=1
할일 상태가 변경되었습니다. (ID: 1, 상태: 완료)
```

---

## 🏗️ 프로젝트 구조

```
go-cli-todo/
├── .gitignore            # Git 버전관리 제외 파일 목록
├── main.go               # 메인 엔트리포인트
├── todo.go               # Todo 구조체 및 관련 함수
├── todo_manager.go       # TodoManager 구조체 및 메서드
├── todos.json            # 할일 데이터 저장 파일 (자동 생성)
├── go.mod                # Go 모듈 설정 파일
├── README.md             # 프로젝트 설명서
├── LICENSE               # 라이선스 파일
└── Makefile              # 빌드 및 관리용 Makefile
```

---

## 🧠 기술적 특징

### 사용된 Go 개념들

- **구조체(Struct)**: Todo와 TodoManager 데이터 모델링
- **메서드**: 구조체에 연결된 동작 정의
- **슬라이스(Slice)**: 동적 할일 목록 관리
- **JSON 마샬링**: 데이터 직렬화/역직렬화
- **Flag 패키지**: 명령행 인자 파싱
- **에러 처리**: Go의 관용적 에러 처리 패턴
- **파일 I/O**: 데이터 영속성 구현

### 데이터 구조

```go
type Todo struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

---

## 📄 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.

---

## 🙋‍♂️ 지원

문제가 있으시거나 질문이 있으시면:

- 📧 **이메일**: zendy00@gmail.com
- 💬 **이슈 트래커**: [GitHub Issues](https://github.com/zendy00/go-cli-todo/issues)
- 💡 **토론**: [GitHub Discussions](https://github.com/zendy00/go-cli-todo/discussions)

---
