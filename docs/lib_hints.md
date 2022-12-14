---
comments: true
title: "Дополнительно"
---

В данном разделе рассказывается о работе некоторых частей библиотеки.

## <a id="stash_scene"></a> Команды StashScene и ApplyStashScene

Данные команды позволяют переходит на отдельную цепочку сцен, без потери основной. 
Если вводится команда `StashScene`, то текущая сцена сохраняется [`directorом`](./lib_structs.md#director)  в **стэк** и заменяется следующей.
Сохранённая сцена находится в стэке до тех пор, пока не будет вызвана команда `ApplyStashScene`. Данная команда, заменяет
текущую сцена на сцену из вершины стэка сохранённых сцен. 

**Примером** использования данного функционала может являться сцена справки. В момент её вызова текущая сцена сохраняется. Далее пользователь изучает справку.
И затем пользователь снова оказывается на сцене, из которой вызывал справку.

## <a id="matchers"></a> Работа наборов правил

[`Director`](./lib_structs.md#director) проверяет введённый текст пользователя последовательно по всем указанным наборам правил для текущей сцены, до первого совпадения.
Для каждого набора правил генерируются константы с названием **`MatcherName`MatchedString**, где `MatcherName` - название набора правил.

При проверке внутри сцены какой из наборов сработал у [`Request`](./lib_structs.md#request) следует проверять поле `NameMatched`.

```go
// React function of actions after scene has been played
func (sc *Echo) React(ctx *scene.Context) scene.Command {
	switch {
		// Matcher select
	case ctx.Request.NameMatched == base_matchers.AnyMatchedString:
        sc.NextScene = GoodbyeScene
	}
    
	return scene.NoCommand
}
```


### Стандартные наборы правил

Библиотека предоставляет два вида стандартных наборов правил:

* Правила на основе регулярных выражений
* Правила на основе соответствия элементу списка

#### Правила на основе регулярных выражений

Для создания правила на регулярных выражений, требуется вызывать функцию `matchers.NewRegexMather` и
указать регулярное выражение и строку, которая будет уникальным идентификатором данного набора правил. 

После выполнения если ввод пользователя подошёл под правило, то набор правил вернёт ту подстроку, которая **соответствовала** регулярному выражению.

Генератор создаёт наборы правил этого типа подобным образом:
```go
// name matched string for RegexMatchers
const (
    CheckedMatchedString = "da"
)

// RegexMatchers
var (
    CheckedMatcher = matchers.NewRegexMather("*", CheckedMatchedString)
)

```

#### Правила на основе соответствия элементу списка

Для создания правила на основе соответствия элементу списка, требуется вызывать функцию `matchers.NewSelectorMatcher` и 
указать массив элементов, одному из которых должен соответствовать ввод пользователя, и строку, которая будет уникальным 
идентификатором данного набора правил.

Генератор создаёт наборы правил этого типа подобным образом:
```go
// replace string for SelectsMatchers
const (
	AgreedMatchedString = "Da"
)

// SelectsMatchers
var (
	AgreedMatcher = matchers.NewSelectorMatcher(
		[]string{
			"da",
			"no",
		},
		AgreedMatchedString,
	)
)
```

!!! tip "hint"
    Также вы можете дописать свой набор правил в соответствии с интерфейсом [`MessageMatcher`](./lib_structs.md#message_matcher).

## <a id="inf_scene"></a> Информационные сцены

В нашей библиотеке сцены могут быть с обработкой реакции пользователя и без неё. Иногда удобно разбить большую реплику голосового помощника на несколько сцен.
Например, когда часть реплики надо сказать только один раз, а остальную множество раз. 

Данная характеристика [сцены](./lib_structs.md#scene) регулируется флагом `withReact` в коде и флагом [`isInfoScene`](./gen_fields.md#isInfoScene) в конфигурационном файле.

Если сцена является информационной, то [`director`](./lib_structs.md#director) при формировании ответа для голосового помощника 
получает от такой сцены только текст реплики. Далее заращивает следующие сцены пока не доберётся до не информационной сцены.
После чего все считанные реплики объединяет в еденную реплику, **где реплики с разных сцен разделены переносом строки**.
От не информационной сцены [`director`](./lib_structs.md#director) берёт остальную информацию. **И только сейчас отправляет ответ**.

#### Пример

Рассмотрим следующие две последовательные сцен:

```yaml
  hello:
    text:
      string: "Hello boy."
      tts: "Hello boy."
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
    buttons:
      Dore:
        text: "Привет"
    matchers:
      - 'any'
    error:
      base: "number"
```

После выполнения сцены `hello`, будет сохранены её реплики и сразу будет запрос информации о следующей сцене `echo`. 
По итогу в ответ голосовому помощнику будет отправлена следующая реплика:
```yaml
BaseText: "Hello boy. \n I will Repeat you word"
TextToSpeech: "Hello boy. \n I will Repeat you word"
```

А также кнопка `Dore`, набор правил `any` и ошибка `number` **только** из сцены `echo`.