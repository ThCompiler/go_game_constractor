---
comments: true
title: "Установка генератора и подключения библиотеки"
---

## Установка генератора и подключения библиотеки

### Генератор

#### Установка

Чтобы установить генератор с помощью `Go` достаточно выполнить команду:
```cmd
go install github.com/ThCompiler/go_game_constractor/scg@latest
```
предлагаю полностью удалить отсюда установку ручками

??? example "Установка ручками из релизов"

    Для установки лучше использовать последний [релиз](https://github.com/ThCompiler/go_game_constractor/releases/tag/v0.1.3-alpha) проекта.
    В релизе опубликованы генераторы для 3 различных ОС: *Linux*, *MacOS* и *Windows*.

    ![manual_images/release.png](static/release.png)

    * *scg.darwin-amd64.tar.gz* - генератор для Mac OS на базе amd64
    * *scg.darwin-arm64.tar.gz* - генератор для Mac OS на базе arm64
    * *scg.linux-amd64.tar.gz*  - генератор для дистрибутивов Линукс на базе amd64
    * *scg.linux-arm64.tar.gz*  - генератор для дистрибутивов Линукс на базе arm64
    * *scg.linux-rm.tar.gz*     - генератор для дистрибутивов Линукс на базе arm
    * *scg.linux-i386.tar.gz*   - генератор для дистрибутивов Линукс на базе i386
    * *scg.windows-amd64.zip*   - генератор для Windows на базе amd64
    * *scg.windows-arm64.zip*   - генератор для Windows на базе arm64
    * *scg.windows-arm.zip*     - генератор для Windows на базе arm
    * *scg.windows-i386.zip*    - генератор для Windows на базе i386

    ## Пример

    В примерах используются утилиты для распаковки и скачивания, которые могут быть не установленны
    на вашем устройстве. Подразумевается, что вы сами их найдёте и установите. 
    
    !!! tip "hint"
        Если вам требуется другая архитектура, то в строке со скачиванием вместо `amd64` укажите вашу архитектуру.
    
    !!! tip "hint"
        Если вам требуется другая версия, то в строке со скачиванием вместо `v0.1.4-alpha` укажите нужную версию.

    === "Linux"

        ```cmd
          # Скачиваем последний релиз
          wget https://github.com/ThCompiler/go_game_constractor/releases/download/v0.1.4-alpha-dffd59b/scg.linux-amd64.tar.gz -P ./tmp
          
          # Распаковываем архив
          tar -xvf ./tmp/scg.linux-amd64.tar.gz -C ./tmp
          
          # Добавляем генератор в утилиты пользователя
          sudo cp ./tmp/scg.linux-amd64/bin/scg /usr/local/bin/
          
          # Очищяем не нужные файлы
          rm -r tmp
        ```
    
    === "MacOS"
        ```cmd
          mkdir tmp
          
          # Скачиваем последний релиз
          cd tmp && curl -LO ./tmp https://github.com/ThCompiler/go_game_constractor/releases/download/v0.1.4-alpha-dffd59b/scg.darwin-amd64.tar.gz \
          && cd ..
        
          
          # Распаковываем архив
          tar -xvf ./tmp/scg.darwin-amd64.tar.gz -C ./tmp
          
          # Добавляем генератор в утилиты пользователя
          sudo cp ./tmp/scg.darwin-amd64/bin/scg /usr/local/bin/
          
          # Очищяем не нужные файлы
          rm -r tmp
        ```

    === "Windows"
    
        Данные комманды прописаны для использования Powershell
        
        ```cmd
          mkdir tmp
          
          # Скачиваем последний релиз
          wget -Uri https://github.com/ThCompiler/go_game_constractor/releases/download/v0.1.4-alpha-dffd59b/scg.windows-amd64.zip -OutFile .\tmp\scg.windows-amd64.zip
          
          # Распаковываем архив
          Expand-Archive -Path .\tmp\scg.windows-amd64.zip  -DestinationPath .\tmp\scg.windows-amd64 -Force
          
          # Добавляем генератор в утилиты пользователя
          mkdir $env:USERPROFILE\scg
          copy .\tmp\scg.windows-amd64\bin\scg.exe $env:USERPROFILE\scg
          
          # Добавим папку в переменные среды
          $Env:Path += ";$env:USERPROFILE\scg"
          
          # Очищяем не нужные файлы
          rd -r tmp
        ```
    
    ## Удаление
    
    === "Linux"
        ```cmd  
          rm -r /usr/local/bin/scg
        ```
    
    === "MacOS"
        ```cmd
          rm -r /usr/local/bin/scg
        ```
    
    === "Windows"
        ```cmd
          # Удалим папку из переменных среды
          $Env:Path = ( $Env:Path.Split(';') | Where-Object { $_ -ne "$env:USERPROFILE\scg" }) -join ';'
          
          # Удалить папку
          rd -r $Env:USERPROFILE\scg
        ```

    !!! warning "Важно"
        Перед началом работы с генератором обязательно инициализируйте Go приложение в директории, где Вы хотите создать игру:
        `go init pkg_name`


------------------------------------------------------------

### Библиотека

#### Установка
```cmd
  go get github.com/ThCompiler/go_game_constractor@latest
```
