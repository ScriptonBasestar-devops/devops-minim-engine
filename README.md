# Minimal imaGe

## 목적

* 빌드 배포 이미지 분리해서 최소화된 배포 이미지 사용
* 도커만 잘 알면 오류없이 사용가능한 배포툴

## 특징

**알아서 해 주지 않고 빌드 및 배포환경을 수동셋팅**

대신 sample repository에서 자주 사용하는 DockerImage의 예시나 추천설정을 제시 해 준다.
 
자동화를 안하는 이유:

- 복사 붙여넣기해서 쓰는게 오히려 열받는 시간낭비를 줄여준다.
- 자동으로 하거나 셀렉트박스로 하는경우 세부적인 설정을 못해서 더 난감한 상황


## 설정

Docker 이미지가 만들어진 후 아래 설정을 이용해서 이미지를 추가 생성
```yaml
vcs:
  name: git
  branch: master
  command:
    - 'git reset --hard HEAD'
    - 'git pull'

build:
  # 아래 명령이 실행될 root directory
  root_path: /application
  dockerbuild: MinG.Build.Dockerfile
# 이미지에 소스코드 및 파일 추가
# copy와 volume이 순서대로 둘 다 동작하기 때문에 충돌발생가능. 한가지만 사용하는게 좋지만
# 그런건 알아서 처리  
  copy_to:
    #- (file/dir):name:path
    - file:name1:/etc/filepath
  volume_to:
    - dir:name3:/application/target/resources/
  test:
    - gradle test
  script:
    - 'gradle clean'
    - 'gradle build'
  extract_from:
#    - 'name:path'
    - 'file:appjar:./app/build/app.jar'

deploy:
  root_path: /application
#별다른 설정이 없는 경우 아래처럼 이미지명을 그대로 입력해서 사용 가능 
#  dockerbuild: MinG.Deploy.Dockerfile
  dockerimage: openjdk:8-jre-alpine
  inject_to:
#    - name:path
    - appjar:/application/app.jar
  script:
    - 'java -jar app.jar'
```

## 실행

go 실행환경. 빌드된 ./bin 시스템 path에 추가

(왠지겹치는 이름이 많을 것 같아 걱정되는 이름)

```bash
ming
ming all
ming build
ming vcs
ming deploy
```
