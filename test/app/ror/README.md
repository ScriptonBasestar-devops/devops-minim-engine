# Ruby On Rails

## ruby 환경

```
rbenv 설치
rbenv install 2.5.3
rbenv global 2.5.3
gem install rails
gem env home

#rails 설치
sudo apt install libsqlite3-dev
gem install coffee-script
```

웹 실행
* unicorn https://github.com/unicorn-engine/unicorn
* puma
* passenger

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
### 도커 한번 실행시켜서 gem path 확인 해서 volume 확인
docker$ gem env
#/usr/local/gems

## start container
host$ docker run -it --rm \
-v ror:/usr/local/bundle/gems \
ruby:2.5.3-stretch bash

docker$ mkdir /app
docker$ cd /app


## copy_in project file
host$ docker cp . ${CONTAINER_ID}:/app

## 환경 설정
docker$ bindle install
docker$ apt update
docker$ apt install nodejs
docker$ rails server

## copy_out jar
host$ docker cp ${CONTAINER_ID}:/app ./${PROJECT_NAME}_app

## run
host$ rails server 
```

### 배포이미지 생성

```bash
## build
host$ docker build -f MinG.Deploy.Dockerfile -t test2 .

## run
host$ docker run --rm -d -p 8000:8000 test2
```


## 참고

https://www.digitalocean.com/community/tutorials/how-to-deploy-a-rails-app-with-unicorn-and-nginx-on-ubuntu-14-04