# Python On Django

PoD

버전 관리는 virtualenv 사용. 다른 것 사용하려면 별도 설정

* anaconda https://conda.io/docs/user-guide/tasks/manage-environments.html
* pienv https://github.com/pypa/pipenv


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
### 도커 한번 실행시켜서 python path 확인 해서 volume 확인
docker$ python -m site

## start container
host$ docker run -it --rm \
-v py37:/usr/local/lib/python3.7 \
python:3.7-stretch bash

docker$ mkdir /app
docker$ cd /app


## copy_in project file
host$ docker cp . ${CONTAINER_ID}:/app

## 환경 설정
docker$ pip install virtualenv
docker$ virtualenv env
dockder$ source env/bin/activate
docker$ pip install -r requirements.txt

## copy_out jar
host$ docker cp ${CONTAINER_ID}:/app ./${PROJECT_NAME}_app

## run
host$ gunicorn ${PROJECT_NAME}.wsgi:application --bind 0.0.0.0:8000 
```

### 배포이미지 생성

```bash
## build
host$ docker build -f MinG.Deploy.Dockerfile -t test2 .

## run
host$ docker run --rm -d -p 8000:8000 test2
```
