---
title: "Быстрый старт"
---

Создание игры состоит минимум из двух этапов:

* Создание описания игры и генерация файлов
* Добавление не тривиальной логики в сгенерерованном сценарии
* **(Доп.)** Написания дополнительной логики связанной с вашим скиллом (Обращение на внейшние API, хранение данных пользователя и т.д.)

## Создание описания игры
Игра описывается в `yml` файле (поддерживается также `xml` и `json` формат, но они в данном мануале не рассматриваются)

#### Пример
Данный файл описывает игру которая повторяет за пользователем ввод
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

## Генерация кода по описанию

### Пример использования
```(cmd)
scg --output=./scg/ --script=./scg/example/echo_game.yml
```

### Описание

#### Аргументы:
- `-o` `--output=file` - путь к директории куда надо сохранить сгенерированные файлы
- `-s` `--script` - путь к файлу с описанием
- `-v` `--version` - показывает версию исполняемого файла
- `-u` `--update` - выполнить генерацию с сохранением пользовательских файлов

#### Уточнение:
Флаг `--update` сохраняет изменения пользователя
Новые сгенерированные строки, отличные от строк в текущей версии файлов, будут записаны как комментарии. этих строк в  будут вставлены в текущий код
Эти комментарии ограничены строками ```// >>>>>>> Generated```.
Решение о применении изменений остается за вами, так же как и решение об удалении ненужного функциональна.