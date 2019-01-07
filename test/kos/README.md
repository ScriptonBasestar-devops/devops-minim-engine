# Kotlin on Springboot

KoS

## 소스 받기
```bash
## 처음이면
git clone http://github.com/source/repo.git

## 기존소스면
git clean -dfx
git fetch origin
git reset --hard origin/master
```

## 이미지 사용

### 빌드

```bash
## start container
host$ docker run -it --rm \
-v gradle_cache:/home/gradle/.gradle \
gradle:5.0.0-jdk8 bash

## copy_in project file
host$ docker cp . ${CONTAINER_ID}:/home/gradle/app

## build gradle
docker$ cd app
docker$ gradle build

## copy_out jar
host$ docker cp ${CONTAINER_ID}:/home/gradle/app/build/libs/kos-0.0.1-SNAPSHOT.jar app.jar

## run
host$ java -Djava.security.egd=file:/dev/./urandom -jar /app.jar
```

### 배포이미지 생성

```bash
## build
host$ docker build -f MinG.Deploy.Dockerfile -t test2 .

## run
host$ docker run --rm -d -p 8080:8080 test2
```


## 참고

https://github.com/sdeleuze/spring-kotlin-functional.git 
