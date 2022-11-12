---
title: "Установка генератора и подключения библиотеки"
---


## Порядок установки каждого из компонентов

- Генератор
  - Открыть последний [релиз](https://github.com/ThCompiler/go_game_constractor/releases/tag/v0.1.3-alpha) проекта
  - Выбрать версию генератора под вашу архитектуру
    ![manual_images/release.png](static/release.png)
      * *scg.darwin-amd64.tar.gz* -- генератор для Mac OS на базе amd64
      * *scg.linux-amd64.tar.gz*  -- генератор для дистрибутивов Линукс на базе amd64
      * *scg.windows-amd64.zip*   -- генератор для Windows на базе amd64
- Библиотека
  - Установка библиотеки для работы с игрой
      ```cmd
      go get github.com/ThCompiler/go_game_constractor@latest
      ```