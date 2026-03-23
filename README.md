# animal face server

## introduction
- 멋쟁이 사자처럼 2~3주차 세션 자료입니다
- 사람의 얼굴 사진을 넣으면 그에 맞는 동물상이 나오는 프로젝트에 대한 서버입니다

## skills
- Gin # Framework
- Go  # Language
- GORM # ORM
- net/http or Resty # HTTP Client
- swaggo/swag # API Doc

## directory structure
~~~
animal-face-go/
    - cmd/
        - server/
            - main.go # EntryPoint

    - internal # 외부에서 참조 불가능한 비지니스 로직
        - controller/ # Controller (Gin Handlers)
        - service/ # Service: 비지니스 로직
        - domain/ # Entity / Model (GORM)
        - repository/ # Data Acess
        - dto/ # Request/Response 객체
        - client # AI server Client
    - config/ # 설정 로드
    - go.mod # dependency 관리
    - docs/ # swagger 문서
~~~

## roadmap
- [x] go mod init animal-face-go 프로젝트 초기화
- [x] gin을 이용한 기본 서버 띄우기 및 heathcheck api
- [x] GORM 설정 및 DB연동
- [] 이미지 업로드를 위한 Multipart Form 핸들러 구현
- [] Resty를 이용한 AI 서버 연동 로직 작성
- [] 분석 결과 저장 및 RESTful API 완성
- [] swag를 이용한 API 문서 자동화
