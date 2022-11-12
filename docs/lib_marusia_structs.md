---
title: "Структуры Маруси"
---

Данный раздел предназначен для понимания работы библиотеки.
В проекте используются следующие json структуры для общения с голосовым помощником.


## Запрос от WebHook Маруси к скилу
-------------------------------------------------------------------------------------------
Запрос содержит четыре поля:

### Структура запроса

| Поле      | Тип                           | Описание                                                                      |
|-----------|-------------------------------|-------------------------------------------------------------------------------|
| `meta`    | [`meta`](#meta)               | Информация об устройстве, с помощью которого пользователь общается с Марусей. |
| `request` | [`request`](#request-marusia) | Данные, полученные от пользователя.                                           |
| `session` | [`session`](#session)         | Данные о сессии.                                                              |
| `version` | `string`                      | Версия протокола. Текущая — `1.0`.                                            |

### `meta`

| Поле         | Тип      | Описание                                                                                                                                                               |
|--------------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `locale`     | `string` | Язык в POSIX-формате, максимум 64 символа.                                                                                                                             |
| `timezone`   | `string` | Название часового пояса, включая алиасы, максимум 64 символа.                                                                                                          |
| `interfaces` | `array`  | Интерфейсы, доступные на устройстве пользователя. Сейчас всегда присылается `screen` — пользователь может видеть ответ скилла на экране и открывать ссылки в браузере. |

### `session`

| Поле                   | Тип                           | Описание                                                                                                                                                                                                           |
|------------------------|-------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `session_id`           | `string`                      | Уникальный идентификатор сессии, максимум 64 символа.                                                                                                                                                              |
| `user_id` (deprecated) | `string`                      | Идентификатор экземпляра приложения, в котором пользователь общается с Марусей, максимум 64 символа. **Важно!** Это поле устарело, вместо него стоит использовать `session.application.application_id` (см. ниже). |
| `skill_id`             | `string`                      | Идентификатор вызываемого скилла, присвоенный при создании. Соответствует полю «Маруся ID» в настройках скилла.                                                                                                    |
| `new`                  | `boolean`                     | Признак новой сессии: • `true` — пользователь начинает новый разговор с вызова навыка, • `false` — запрос отправлен в рамках уже начатого разговора.                                                               |
| `message_id`           | `integer`                     | Идентификатор сообщения в рамках сессии, максимум 8 символов. Инкрементируется с каждым следующим запросом.                                                                                                        |
| `user`                 | [`user`](#user)               | Данные о пользователе. Передаётся, только если пользователь авторизован. (см. ниже).                                                                                                                               |
| `application`          | [`application`](#application) | Данные об экземпляре приложения (см. ниже).                                                                                                                                                                        |
| `auth_token`           | `object`                      | Авторизационный токен Маруси.                                                                                                                                                                                      |

### `user`

| Поле      | Тип      | Описание                                                                                          |
|-----------|----------|---------------------------------------------------------------------------------------------------|
| `user_id` | `string` | Идентификатор аккаунта пользователя (максимум 64 символа). Уникален в разрезе: «скилл + аккаунт». |

### `application`

| Поле               | Тип      | Описание                                                                                                                                                     |
|--------------------|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `application_id`   | `string` | Идентификатор экземпляра приложения, в котором пользователь общается с Марусей (максимум 64 символа). Уникален в разрезе: «скилл + приложение (устройство)». |
| `application_type` | `string` | Тип приложения (устройства). Возможные значения: • `mobile`; • `speaker`; • `VK`; • `other`.                                                                 |

### <a id="request-marusia"></a> `request`

| Поле                 | Тип      | Описание                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
|----------------------|----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `command`            | `string` | Пользовательский текст, очищенный от слов, не влияющих на смысл предложения. В ходе преобразования текст, в частности, очищается от знаков препинания, удаляются обращения к Марусе и слова, выражающие просьбу, например «пожалуйста», «слушай» и т. д., числительные преобразуются в числа. При завершении скилла (по команде «стоп», «выход» и так далее) в него будет передана команда `on_interrupt`, чтобы у скилла была возможность попрощаться с пользователем. |
| `original_utterance` | `string` | Полный текст пользовательского запроса, максимум 1024 символа.                                                                                                                                                                                                                                                                                                                                                                                                          |
| `type`               | `string` | Тип ввода, обязательное свойство. Возможные значения: • `SimpleUtterance` — голосовой ввод; • `ButtonPressed` — нажатие кнопки.                                                                                                                                                                                                                                                                                                                                         |
| `payload`            | `object` | JSON, полученный с нажатой кнопкой от обработчика скилла (в ответе на предыдущий запрос), максимум 4 096 байт. Передаётся, только если была нажата кнопка с payload.                                                                                                                                                                                                                                                                                                    |
| `nlu`                | `object` | Объект, содержащий слова и именованные сущности, которые Маруся извлекла из запроса пользователя, в поле `tokens` (`array`).                                                                                                                                                                                                                                                                                                                                            |

### Пример запроса

```json
{
  "meta": {
    "client_id": "MailRu-VC/1.0",
    "locale": "ru_RU",
    "timezone": "Europe/Moscow",
    "interfaces": {
      "screen": {}
    }
  },
  "request": {
    "command": "какая очередь в столовой",
    "original_utterance": "какая очередь в столовой",
    "type": "SimpleUtterance",
    "payload": {},
    "nlu": {
      "tokens": [
        "какая",
        "очередь",
        "в",
        "столовой"
      ]
    }
  },
  "session": {
    "session_id": "01bfd28fe3a326-c-2-fea35db06d4-a8930",
    "user_id": "f63bc4d9e9c89abe10fbe874b5400b67c0df41f86143ec22629b00be606a1dac",
    "skill_id": "5b23aa28b9cbd41ad25-21-2-60c7-121d4b",
    "new": false,
    "message_id": 1,
    "user": {
      "user_id": "c825511e862f23f3728a58cd3b15896cd243c7460237c651944b7499c7c9a425"
    },
    "application": {
      "application_id": "f63bc4d9e9c89abe10fbe874b5400b67c0df41f86143ec22629b00be606a1dac",
      "application_type": "mobile"
    }
  },
  "version": "1.0"
}
```
## Формат ответа обработчика скилла Марусе
-------------------------------------------------------------------------------------------

### Структура ответа

| Поле       | Тип                     | Описание                                           |
|------------|-------------------------|----------------------------------------------------|
| `response` | [`response`](#response) | _Обязательное_. Данные для ответа пользователю.    |
| `session`  | [`session`](#session)   | _Обязательное_. Данные о сессии.                   |
| `version`  | `string`                | _Обязательное_. Версия протокола. Текущая — `1.0`. |

### `response`

| Поле          | Тип                                        | Описание                                                                                                                                                                                                                                                    |
|---------------|--------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `text`        | `string` или `array`                       | _Обязательное_. Текст, который следует показать и сказать пользователю, максимум 1 024 символа. Не должен быть пустым. В тексте ответа можно указать переводы строк последовательностью `\n`. Если передать массив строк, то сообщения разобьются на баблы. |
| `tts`         | `object`                                   | Ответ в [формате TTS](https://dev.vk.com/marusia/tts) (text-to-speech), максимум 1 024 символа. Поддерживается расстановка ударений с помощью '+'.                                                                                                          |
| `buttons`     | `array`                                    | Кнопки (suggest), которые следует показать пользователю. Кнопки можно использовать как релевантные ответу ссылки или подсказки для продолжения разговора.                                                                                                   |
| `end_session` | `boolean`                                  | _Обязательное_. Признак конца разговора: • `true` — сессию следует завершить, • `false` — сессию следует продолжить.                                                                                                                                        |
| `card`        | [`card`](https://dev.vk.com/marusia/cards) | Описание карточки — сообщения с различным контентом. Подробнее о типах карточек и описание структур в [специальном разделе](https://dev.vk.com/marusia/cards).                                                                                              |
| `commands`    | `array`                                    | Команды. Поле позволяет передать несколько сообщений в нужном порядке. На данный момент поддерживаются только [карточки](https://dev.vk.com/marusia/cards).                                                                                                 |

### `buttons`

| Поле      | Тип      | Описание                                                                                                                                                                                                         |
|-----------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `title`   | `string` | _Обязательное_. Текст кнопки, максимум 64 символа.                                                                                                                                                               |
| `url`     | `string` | URL, который откроется при нажатии на кнопку, максимум 1 024 байта. Если свойство URL не указано, по нажатию на кнопку навыку будет отправлен текст кнопки. Пока кнопки с URL не поддерживаются в приложении VK. |
| `payload` | `object` | Любой JSON, который нужно отправить скиллу, если эта кнопка будет нажата, максимум 4 096 байт.                                                                                                                   |

### `session`

| Поле         | Тип      | Описание                                                                                                                    |
|--------------|----------|-----------------------------------------------------------------------------------------------------------------------------|
| `session_id` | `string` | _Обязательное_. Уникальный идентификатор сессии, максимум 64 символа.                                                       |
| `user_id`    | `string` | _Обязательное_. Идентификатор экземпляра приложения, в котором пользователь общается с Марусей, максимум 64 символа.        |
| `message_id` | `string` | _Обязательное_. Идентификатор сообщения в рамках сессии, максимум 8 символов. Инкрементируется с каждым следующим запросом. |

### Пример ответа

```json
{
  "response": {
    "text": "Сейчас очередь в столовой 5 человек.",
    "tts": "Сейчас очередь в столовой пять человек.",
    "buttons": [
      {
        "title": "Надпись на кнопке",
        "payload": {},
        "url": "https://example.com/"
      }
    ],
    "end_session": true
  },
  "session": {
    "session_id": "01bfd28fe3a326-c-2-fea35db06d4-a8930",
    "user_id": "f63bc4d9e9c89abe10fbe874b5400b67c0df41f86143ec22629b00be606a1dac",
    "skill_id": "5b23aa28b9cbd41ad25-21-2-60c7-121d4b",
    "new": false,
    "message_id": 1,
    "user": {
      "user_id": "c825511e862f23f3728a58cd3b15896cd243c7460237c651944b7499c7c9a425"
    },
    "application": {
      "application_id": "f63bc4d9e9c89abe10fbe874b5400b67c0df41f86143ec22629b00be606a1dac",
      "application_type": "mobile"
    }
  },
  "version": "1.0"
}
```
