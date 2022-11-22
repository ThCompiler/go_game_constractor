---
comments: true
title: "Быстрый старт"
---
Перед началом создания игры убедитесь, что у Вас установлены `Go` и `Docker`.
??? info "Дополнительно"
    Инструкция по установке `Go` находится на сайте [go.dev](https://go.dev/learn/))

    Инструкция по установке `Docker` находится на сайте [docs.docker](https://docs.docker.com/engine/install/))

**Дальнейшая инструкция описывается для системы `Linux`**. а для остальных? им тоже надо, либо не писать об этом, сами догадаются

## Первый шаг: установка
```cmd
go install https://github.com/ThCompiler/go_game_constractor/scg@latest
```
нету, не будем путать
есть ли смысл дублировать одно и то же в двух местах? возможно, стоит оставить этот пункт только здесь

## Второй шаг: создание конфигурационного файла
Создать директорию проекта:
```cmd
mkdir dir_name
cd dir_name
```

Инициализировать **Go**-приложение:
```cmd
go mod init game_name
```

Игра описывается в `YAML` файле (поддерживается также `XML` и `JSON` формат).
Создать в директории проекта конфигурационный файл:
```cmd
touch skill.yml
```

Пример: конфигурационный файл для echo-игры, повторяющей реплику пользователя.

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

??? info "Дополнительно"
    Описание полей находится в разделе ["Генератор"](./gen_fields.md)

## Третий шаг: генерация скилла

Сгенерировать сервер с игрой:
```cmd
scg --output=./ --script=./skill.yml --http-server
```

В директории проекта появится папка `scg`:
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

## Четвёртый шаг: добавление логики с использованием библиотеки
Если Вам требуется реализация дополнительной логики, необходимо установить отсутствующие пакеты:
```cmd
go mod tidy
```
, открыть директорию со сценами `scg/internal/script/scenes` и добавить необходимую логику.

## Пятый шаг: настройка хранилища текстов
здесь (хотя лучше везде) надо выделять как-то названия, чтобы отличать их от команд и флагов
Тексты находятся в In-Memory хранилище `Redis`. Для его запуска следует воспользоваться Docker'ом:
```cmd
# Создать volume для данных из Redis'а, чтобы обеспечить их сохранение при перезапуске хранилища
sudo docker volume create tredis-vol

# Запустить Redis на 6380 порту
sudo docker run -v tredis-vol:/data -p 6380:6379 --name tredis -d redis redis-server  --save 60 1 --loglevel warning
```

## Шестой шаг: сборка и запуск сервиса с игрой

Собрать сервис в папку `/bin`:
```cmd
go build -o bin/skill scg/cmd
```

Приложение запускается на `8080` порту, убедитесь, что он свободен, или замените порт сервера в файле `scg/config/config.yml`:
```yaml
http:
  port: 'YOU_PORT'
```

Запустить хранилище `Redis` с указанием `url`:
```cmd
REDIS_URL=redis://localhost:6380 ./bin/skill
```

------------------------------------------------------------------
<h2 align="center">Поздравляем, Вы создали свою первую игру для голосового ассистента Маруся</h2>
------------------------------------------------------------------

!!! tip "hint"
    Для запуска игры Вы можете воспользоваться [отладчиком](https://skill-debugger.marusia.mail.ru/) от Маруси, указав в нём
    `url` для webhookа в формате `http://localhost:8080/v1/skill_name`.
