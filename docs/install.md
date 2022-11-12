---
title: "Установка генератора и подключения библиотеки"
---

## Установка генератора и подключения библиотеки

### Генератор

#### Установка

Для установки лучше использовать последний [релиз](https://github.com/ThCompiler/go_game_constractor/releases/tag/v0.1.3-alpha) проекта.
В релизе опубликованы генераторы для 3 различных ОС: *Linux*, *MacOS* и *Windows*.

![manual_images/release.png](static/release.png)

* *scg.darwin-amd64.tar.gz* -- генератор для Mac OS на базе amd64
* *scg.linux-amd64.tar.gz*  -- генератор для дистрибутивов Линукс на базе amd64
* *scg.windows-amd64.zip*   -- генератор для Windows на базе amd64

##### Пример установки

В примерах используются утилиты для распаковки и скачивания, которые могут быть не установленны
на вашем устройстве. Подразумевается, что вы сами их найдёте и установите.

###### Linux
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

###### MacOS
```cmd
  mkdir tmp
  
  # Скачиваем последний релиз
  curl -o ./tmp https://github.com/ThCompiler/go_game_constractor/releases/download/v0.1.3-alpha/scg.darwin-amd64.tar.gz
  
  # Распаковываем архив
  tar -xvf scg.darwin-amd64.tar.gz -C ./tmp
  
  # Добавляем генератор в утилиты пользователя
  sudo cp ./tmp/scg.darwin-amd64/scg /usr/loca/bin/
  
  # Очищяем не нужные файлы
  rm -r tmp
```

###### MacOS
```cmd
  mkdir tmp
  
  # Скачиваем последний релиз
  wget https://github.com/ThCompiler/go_game_constractor/releases/download/v0.1.3-alpha/scg.windows-amd64.zip -P .\tmp
   
  # Скачиваем распаковшик zip-архива
  wget https://www.7-zip.org/a/7z2201-x64.exe -P .\tmp
  
  # Распаковываем архив
  .\tmp\7z2201-x64.exe x scg.windows-amd64.zip -o .\tmp
  
  # Добавляем генератор в утилиты пользователя
  mkdir %USERPROFILE%/scg
  copy .\tmp\scg.windows-amd64\scg %USERPROFILE%\scg
  
  # Очищяем не нужные файлы
  rd -r tmp
```

### Библиотека

#### Установка
```cmd
  go get github.com/ThCompiler/go_game_constractor@latest
```