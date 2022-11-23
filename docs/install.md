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

??? example "Установка вручную"

    Для установки необходимо использовать последний [релиз](https://github.com/ThCompiler/go_game_constractor/releases/tag/v0.1.3-alpha) проекта.
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

    Убедитесь, что утилиты для распаковки и скачивания, используемые в примерах, установленны
    на Вашем устройстве.
    
    !!! tip "hint"
        Если Вам требуется другая архитектура, то вместо `amd64` укажите необходимую.
    
    !!! tip "hint"
        Если Вам требуется другая версия, то вместо `v0.1.4-alpha` укажите необходимую.

    === "Linux"

        ```cmd
          # Скачать последний релиз
          wget https://github.com/ThCompiler/go_game_constractor/releases/download/v0.1.4-alpha-dffd59b/scg.linux-amd64.tar.gz -P ./tmp
          
          # Распаковать архив
          tar -xvf ./tmp/scg.linux-amd64.tar.gz -C ./tmp
          
          # Добавить генератор в утилиты пользователя
          sudo cp ./tmp/scg.linux-amd64/bin/scg /usr/local/bin/
          
          # Очистить ненужные файлы
          rm -r tmp
        ```
    
    === "MacOS"
        ```cmd
          mkdir tmp
          
          # Скачать последний релиз
          cd tmp && curl -LO ./tmp https://github.com/ThCompiler/go_game_constractor/releases/download/v0.1.4-alpha-dffd59b/scg.darwin-amd64.tar.gz \
          && cd ..
        
          
          # Распаковать архив
          tar -xvf ./tmp/scg.darwin-amd64.tar.gz -C ./tmp
          
          # Добавить генератор в утилиты пользователя
          sudo cp ./tmp/scg.darwin-amd64/bin/scg /usr/local/bin/
          
          # Очистить ненужные файлы
          rm -r tmp
        ```

    === "Windows"
    
        Данные комманды прописаны для Powershell
        
        ```cmd
          mkdir tmp
          
          # Скачать последний релиз
          wget -Uri https://github.com/ThCompiler/go_game_constractor/releases/download/v0.1.4-alpha-dffd59b/scg.windows-amd64.zip -OutFile .\tmp\scg.windows-amd64.zip
          
          # Распаковать архив
          Expand-Archive -Path .\tmp\scg.windows-amd64.zip  -DestinationPath .\tmp\scg.windows-amd64 -Force
          
          # Добавить генератор в утилиты пользователя
          mkdir $env:USERPROFILE\scg
          copy .\tmp\scg.windows-amd64\bin\scg.exe $env:USERPROFILE\scg
          
          # Добавить папку в переменные среды
          $Env:Path += ";$env:USERPROFILE\scg"
          
          # Очистить ненужные файлы
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
          # Удалить папку из переменных среды
          $Env:Path = ( $Env:Path.Split(';') | Where-Object { $_ -ne "$env:USERPROFILE\scg" }) -join ';'
          
          # Удалить папку
          rd -r $Env:USERPROFILE\scg
        ```

    !!! warning "Важно"
        Перед началом работы с генератором обязательно инициализируйте Go-приложение в директории, где Вы хотите создать игру:
        `go init pkg_name`


------------------------------------------------------------

### Библиотека

#### Установка
```cmd
  go get github.com/ThCompiler/go_game_constractor@latest
```
