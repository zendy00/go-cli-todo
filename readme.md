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
[예시](#-사용-예시) •
[기여하기](#-기여하기)

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
git clone https://github.com/zendy00/cli-todo-manager.git
cd cli-todo-manager

# 빌드
go build -o todo main.go

# 시스템 PATH에 추가 (선택사항)
sudo mv todo /usr/local/bin/
```

### 릴리즈 다운로드

최신 릴리즈에서 운영체제에 맞는 바이너리를 다운로드하세요:

[![Download](https://img.shields.io/badge/Download-Latest%20Release-blue.svg)](https://github.com/zendy00/cli-todo-manager/releases)

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
cli-todo-manager/
├── main.go           # 메인 애플리케이션 코드
├── todos.json        # 데이터 저장 파일 (자동 생성)
├── README.md         # 프로젝트 문서
├── LICENSE          # 라이선스 파일
└── examples/        # 사용 예시들
    └── demo.gif     # 데모 GIF
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

## 🤝 기여하기

프로젝트에 기여해주셔서 감사합니다! 다음과 같은 방법으로 참여할 수 있습니다:

### 🐛 버그 리포트
버그를 발견하셨나요? [이슈를 생성](https://github.com/zendy00/cli-todo-manager/issues/new)해 주세요.

### 💡 기능 제안
새로운 기능 아이디어가 있으시면 [토론](https://github.com/zendy00/cli-todo-manager/discussions)에서 공유해 주세요.

### 🔧 코드 기여

1. 레포지토리를 포크합니다
2. 새로운 브랜치를 생성합니다 (`git checkout -b feature/amazing-feature`)
3. 변경사항을 커밋합니다 (`git commit -m 'Add amazing feature'`)
4. 브랜치에 푸시합니다 (`git push origin feature/amazing-feature`)
5. Pull Request를 생성합니다

---

## 🗺️ 로드맵

### v1.1 (예정)
- [ ] 할일 우선순위 기능
- [ ] 마감일 설정 및 알림
- [ ] 태그 시스템
- [ ] 검색 기능

### v1.2 (예정)
- [ ] 할일 카테고리 분류
- [ ] 통계 기능 (완료율, 생산성 지표)
- [ ] 데이터 내보내기/가져오기
- [ ] 컬러 테마 지원

### v2.0 (미래)
- [ ] 팀 협업 기능
- [ ] 웹 대시보드
- [ ] 모바일 동기화

---

## 📄 라이선스

이 프로젝트는 MIT 라이선스 하에 배포됩니다. 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.

---

## 🙋‍♂️ 지원

문제가 있으시거나 질문이 있으시면:

- 📧 **이메일**: your.email@example.com
- 💬 **이슈 트래커**: [GitHub Issues](https://github.com/zendy00/cli-todo-manager/issues)
- 💡 **토론**: [GitHub Discussions](https://github.com/zendy00/cli-todo-manager/discussions)

---

<div align="center">

**CLI TODO Manager를 사용해주셔서 감사합니다!** ⭐

스타를 눌러주시면 프로젝트 개발에 큰 도움이 됩니다.

[![GitHub stars](https://img.shields.io/github/stars/zendy00/cli-todo-manager.svg?style=social&label=Star)](https://github.com/zendy00/cli-todo-manager)
[![GitHub forks](https://img.shields.io/github/forks/zendy00/cli-todo-manager.svg?style=social&label=Fork)](https://github.com/zendy00/cli-todo-manager/fork)

</div>