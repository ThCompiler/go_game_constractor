---
title: "Быстрый старт"
---

Создание игры состоит минимум из двух этапов:

* Создание описания игры и генерация файлов
* Добавление не тривиальной логики в сгенерерованном сценарии
* **(Доп.)** Написания дополнительной логики связанной с вашим скиллом (Обращение на внейшние API, хранение данных пользователя и т.д.)

Ожидается, что у вас уже установлен `golang` и `docker`. 
> <h5>hint</h5>
> Инструкция по установке `golang` находится на сайте [go.dev](https://go.dev/learn/))
> 
> Инструкция по установке `docker` находится на сайте [docs.docker](https://docs.docker.com/engine/install/))

**Дальнейшая инструкция описывается для системы `Linux`**.

## Первый шаг: Установка
```cmd
mkdir tmp

# Скачиваем последний релиз
wget https://github.com/ThCompiler/go_game_constractor/releases/download/v0.1.3-alpha/scg.linux-amd64.tar.gz -P ./tmp

# Распаковываем архив
tar -xvf scg.linux-amd64.tar.gz -C ./tmp

# Добавляем генератор в утилиты пользователя
sudo cp ./tmp/scg.linux-amd64/scg /usr/loca/bin/

# Очищяем не нужные файлы
rm -r tmp
```

> <h5>hint</h5> 
> Полное описание установки есть в разделе ["Установка"](./install.md)


## Второй шаг: Создание описания скилла
Создадим директорию проекта:
```cmd
mkdir mar_skill
cd mar_skill
```

Инициализируем наше **Golang**-приложение:
```cmd
go mod init mar_skill
```

Игра описывается в `yml` файле (поддерживается также `xml` и `json` формат, но они в данном мануале не рассматриваются).
Поэтому создадим в проекте конфигурационный файл:
```cmd
toch skill.yml
```

Вы можете взять следующий пример за основу и изменить его под свои нужды. 
В нём описывается игра, которая повторяет за пользователем ввод.

```yaml
name: 'echo_game'
startScene: "hello"
goodByeCommand: "GoodeBye"
goodByeScene: "goodbye"
script:
  goodbye:
    text:
      string: "GoodyBye"
      tts: "GoodyBye"
    nextScenes:
      - 'goodbye'
  hello:
    text:
      string: "Hello boy. Is number {number}"
      tts: "Hello boy. Is number {number}"
      values:
        number: 'int64'
    nextScene: 'echo'
    isInfoScene: true
    error:
      scene: "goodbye"
  echo:
    text:
      string: "I will Repeat you word"
      tts: "I will Repeat you word"
    nextScenes:
      - 'echoRepeat'
    matchers:
      - 'any'
    error:
      base: "number"
  echoRepeat:
    text:
      string: "You say {userText}"
      tts: "You say {userText}"
      values:
        userText: 'string'
    nextScenes:
      - 'echoRepeat'
    matchers:
      - 'any'
    error:
      base: "number"
```

> <h5>hint</h5>
> Описание полей приводится в разделе ["Генератор"](./gen_fields.md)

## Третий шаг: Генерация скилла

После создания конфигурационного файла, сгенерируем нам сервер со скиллом.
```cmd
scg --output=./ --script=./skill.yml --http-server
```

В папке проекта появится папка `scg` с папкой нашего скилла:
```
- scg
    - cmd
    - internal
        - texts
            - store
                - redis
                - storesaver
            - consts
                - textsname
            - manager
                - usecase
        - skill_name
        - controller
            - http
                - v1
        - script
            - errors
            - matchers
            - payloads
            - scenes
    - config
        - resources
    - pkg
        - logger
        - str
        - httpserver
        - ginutilits
```

## Четвёртый шаг: Дополнительные изменения

Если вам требуется реализация дополнительной логики, то следует открыть директорию со сценами `scg/internal/script/scenes`
и добавить необходимую логику сценам.

> <h5>Важно</h5>
> Текущая версия генератора оставляет на разработчика задачу по установке правил перехода от выполненного набора правил к следующей сцене. 
> Т.е. после генерации вам обязательно надо зайти в каждую сцену и указать какая следующая сцена в параметр `nextScene` сцены.

## Пятый шаг: Настройка хранилища текстов

Тексты хранятся в in-memory хранилище `Redis`. Для его запуска воспользуемся докером:
```cmd
# Создадим volume для данных из Redisа, чтобы при перезапуски хранилища они сохранились
sudo docker volume create tredis-vol

# Запустим наше хранилище на 6380 порту
sudo docker run -v tredis-vol:/data -p 6380:6379 --name tredis -d redis redis-server  --save 60 1 --loglevel warning
```

## Шестой шаг: сборка и запуск сервиса со скиллом

Сначала собирём наш сервис в папку `/bin`:
```cmd
go build -o bin/skill skill_name/cmd
```

А теперь запустим с указанием `url` хранилища `Redis`:
```cmd
REDIS_URL=redis://localhost:6380 ./bin/skill
```

> <h5>hint</h5>
> Если вы хотите запуститься на отличном от `8080` порту, то в файле конфигураций сервиса `scg/skill_name/config/config.yml` 
> следует поменять поле `port` в разделе `http` с `8080` на желаемый вами порт.

------------------------------------------------------------------
<h2 align="center">Поздравляем вы создали свой скилл в марусе</h2>
------------------------------------------------------------------

> <h5>hint</h5>
> Для проверки своего скилла вы можете воспользоваться [отладчиком](https://skill-debugger.marusia.mail.ru/) от Маруси, указав в нём
> `url` для webhookа как `http://localhost:8080/v1/skill_name`.