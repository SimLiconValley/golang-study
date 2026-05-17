# Go Study Repository

이 저장소는 [roadmap.sh/golang](https://roadmap.sh/golang) 커리큘럼을 따라 Go 언어를 체계적으로 학습하는 공간입니다.

## 학습 목표

- Go 언어의 기본 문법 및 핵심 개념 완벽 숙달
- 로드맵의 각 노드별 예제 코드 작성 및 동작 원리 이해
- 학습한 개념을 정형화된 데이터(CSV)로 관리하여 장기 기억화 (Anki 활용 등)
- 효율적이고 유지보수 가능한 Go 프로젝트 구조 학습

## 학습 가이드 (Hybrid Workflow)

학습 생산성을 높이기 위해 큰 주제별 폴더 구조와 통합 단어장을 사용합니다.

### 1. 디렉토리 구조

- `/src/[주제]/`: 로드맵의 큰 섹션 단위로 폴더를 구성합니다.
  - 예: `/src/introduction-to-go`
- `concepts.csv`: 각 주제 폴더 내에 위치하며, 해당 주제에서 배운 모든 개념과 용어를 기록합니다.
- `[node-name].go`: 로드맵의 각 세부 노드별 예제 코드를 파일 단위로 작성합니다.

### 2. CSV 파일 형식

`concepts.csv` 파일은 VS Code의 CSV Editor 등을 활용하여 편집하며, 다음과 같은 형식을 권장합니다:
`개념, 설명, 관련 코드/비고`

## 진행 상황 확인

세부 학습 진도는 [GitHub Issues](https://github.com/SimLiconValley/golang-study/issues)에서 관리합니다. 현재 진행 중인 로드맵 체크리스트는 관련 이슈를 참조하세요.

---

## 기존 학습 (A Tour of Go)

이전 학습 내역은 `src/` 디렉토리의 새로운 구조에 맞춰 재배치되었습니다.

- `/docs`: 학습 관련 문서 및 PlantUML 다이어그램
- `/src`: 로드맵 단계별 학습 코드 및 개념 데이터
