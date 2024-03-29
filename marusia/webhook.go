package marusia

/**
Get code from github.com/SevereCloud/vksdk. And rewrite to gin server
*/

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/ThCompiler/go_game_constractor/pkg/language"
	"github.com/ThCompiler/go_game_constractor/pkg/logger"
	context2 "github.com/ThCompiler/go_game_constractor/pkg/logger/context"
)

// Version версия протокола.
const Version = "1.0"

// ApplicationType тип приложения.
type ApplicationType string

// Возможные значения.
const (
	Mobile  ApplicationType = "mobile"  // мобильное приложение
	Speaker ApplicationType = "speaker" // колонка
	VK      ApplicationType = "VK"      // страница вк
	Other   ApplicationType = "other"   // колонка
)

// RequestType тип ввода.
type RequestType string

// Возможные значения.
const (
	SimpleUtterance RequestType = "SimpleUtterance" // голосовой ввод
	ButtonPressed   RequestType = "ButtonPressed"   //  нажатие кнопки
)

// Типичные команды голосового ввода.
const (
	// OnStart команда запуска скилла. В скилл будет передана пустая строка
	// Command = "".
	OnStart = ""

	// OnInterrupt команда завершении скилла по команде "стоп", "выход" и т.д. в
	// скилл будет передано Command = "on_interrupt", чтобы у скилла была
	// возможность попрощаться с пользователем.
	OnInterrupt = "on_interrupt"
)

// RequestIn данные, полученные от пользователя.
type RequestIn struct {
	// Служебное поле: запрос пользователя, преобразованный для внутренней
	// обработки Марусей. В ходе преобразования текст, в частности, очищается
	// от знаков препинания, а числительные преобразуются в числа. При
	// завершении скилла по команде "стоп", "выход" и т.д. в скилл будет
	// передано "on_interrupt", чтобы у скилла была возможность попрощаться с
	// пользователем.
	Command string `json:"command"`

	// Полный текст пользовательского запроса, максимум 1024 символа.
	OriginalUtterance string `json:"original_utterance"`

	// Тип ввода.
	Type RequestType `json:"type"`

	// JSON, полученный с нажатой кнопкой от обработчика скилла (в ответе на
	// предыдущий запрос), максимум 4096 байт.
	Payload json.RawMessage `json:"payload,omitempty"`

	// Объект, содержащий слова и именованные сущности, которые Маруся
	// извлекла из запроса пользователя.
	NLU language.NLU `json:"nlu"`
}

// Screen структура для Interfaces.
type Screen struct{}

// Interfaces интерфейсы, доступные на устройстве пользователя.
type Interfaces struct {
	// Пользователь может видеть ответ скилла на экране и открывать ссылки
	// в браузере.
	Screen *Screen `json:"screen,omitempty"`
}

// IsScreen пользователь может видеть ответ скилла на экране и открывать
// ссылки в браузере.
func (i *Interfaces) IsScreen() bool {
	return i.Screen != nil
}

// Meta информация об устройстве, с помощью которого пользователь общается
// с Марусей.
type Meta struct {
	// Идентификатор клиентского приложения
	ClientID string `json:"client_id"`

	// Язык в POSIX-формате, максимум 64 символа.
	Locale string `json:"locale"`

	// Название часового пояса, включая алиасы, максимум 64 символа
	Timezone string `json:"timezone"`

	// Интерфейсы, доступные на устройстве пользователя.
	Interfaces Interfaces `json:"interfaces"`

	// Город пользователя на русском языке.
	CityRu string `json:"city_ru,omitempty"`
}

// Session данные о сессии.
type Session struct {
	// Уникальный идентификатор сессии, максимум 64 символа.
	SessionID string `json:"session_id"`

	// Идентификатор вызываемого скилла, присвоенный при создании.
	// Соответствует полю "Маруся ID" в настройках скилла.
	SkillID string `json:"skill_id"`

	// Признак новой сессии:
	//
	// true — пользователь начинает новый разговор с навыком,
	//
	// false — запрос отправлен в рамках уже начатого разговора.
	New bool `json:"new"`

	// Идентификатор сообщения в рамках сессии, максимум 8 символов.
	// Инкрементируется с каждым следующим запросом.
	MessageID int `json:"message_id"`

	// Данные об экземпляре приложения.
	Application Application `json:"application"`

	// Авторизационный токен Маруси.
	AuthToken string `json:"auth_token"`

	// Данные о пользователе. Передаётся, только если пользователь авторизован.
	User User `json:"user,omitempty"`
}

// Application данные о приложении.
type Application struct {
	// Идентификатор экземпляра приложения, в котором пользователь общается с Марусей (максимум 64 символа).
	// Уникален в разрезе: «скилл + приложение (устройство)».
	ApplicationID string `json:"application_id"`

	// Тип приложения (устройства). Возможные значения:
	//  • mobile;
	//  • speaker;
	//  • VK;
	//  • other.
	ApplicationType ApplicationType `json:"application_type"`
}

// User данные о пользователе.
type User struct {
	// Идентификатор аккаунта пользователя (максимум 64 символа). Уникален в разрезе: «скилл + аккаунт».
	UserID string `json:"user_id"`

	// Идентификатор аккаунта пользователя в ВК, работает только если данное поле было
	// включено разработчиками ВК навыков Маруси.
	// Не работает в отладки и локально.
	UserVKID string `json:"user_vk_id,omitempty"`
}

// Request структура запроса.
type Request struct {
	// Информация об устройстве, с помощью которого пользователь общается с Марусей.
	Meta Meta `json:"meta"`

	// Данные, полученные от пользователя.
	Request RequestIn `json:"request"`

	// Данные о сессии.
	Session Session `json:"session"`

	// Версия протокола.
	Version string `json:"version"`

	// Служебное поле. Позволяет передать какие-то данные о запросе в обработчик сцен
	Context context.Context
}

// BindingType тип для DefaultPayload.
type BindingType string

// Возможные значения.
const (
	BindingTypeSuggest BindingType = "suggest"
)

// DefaultPayload дефолтная нагрузка.
type DefaultPayload struct {
	BindingType    BindingType `json:"binding_type"`
	Index          int         `json:"index"`
	TargetPhraseID string      `json:"target_phrase_id"`
}

// Button кнопка.
type Button struct {
	Title   string      `json:"title"`
	Payload interface{} `json:"payload,omitempty"`
	URL     string      `json:"url,omitempty"`
}

// CardType тип карточки.
type CardType string

// Возможные значения.
const (
	// Одно изображение.
	BigImage CardType = "BigImage"

	// Набор изображений.
	ItemsList CardType = "ItemsList"
)

// CardItem элемент карточки.
type CardItem struct {
	// ID изображения из раздела "Медиа-файлы" настроек в VKApps.
	ImageID int `json:"image_id"`
}

// Card описание карточки — сообщения с поддержкой изображений.
type Card struct {
	// Тип карточки.
	Type CardType `json:"type"`

	// Заголовок изображения.
	Title string `json:"title"`

	// Описание изображения.
	Description string `json:"description"`

	// ID изображения из раздела "Медиа-файлы" настроек в VKApps
	// (игнорируется для типа ItemsList).
	ImageID int `json:"image_id,omitempty"`

	// Список изображений, каждый элемент является объектом формата BigImage.
	Items []CardItem `json:"items,omitempty"`
}

// NewBigImage возвращает карточку с картинкой.
func NewBigImage(title, description string, imageID int) *Card {
	return &Card{
		Type:        BigImage,
		Title:       title,
		Description: description,
		ImageID:     imageID,
	}
}

// NewItemsList возвращает карточку с набором картинок.
func NewItemsList(title, description string, items []CardItem) *Card {
	return &Card{
		Type:        ItemsList,
		Title:       title,
		Description: description,
		Items:       items,
	}
}

// NewImageList возвращает карточку с набором картинок.
func NewImageList(title, description string, imageIDs ...int) *Card {
	items := make([]CardItem, len(imageIDs))

	for i := 0; i < len(imageIDs); i++ {
		items[i].ImageID = imageIDs[i]
	}

	return NewItemsList(title, description, items)
}

// Response данные для ответа пользователю.
type Response struct {
	// Текст, который следует показать и сказать пользователю. Максимум 1024
	// символа. Не должен быть пустым. В тексте ответа можно указать переводы
	// строк последовательностью «\n».
	Text string `json:"text"`

	// Ответ в формате TTS (text-to-speech), максимум 1024 символа.
	// Поддерживается расстановка ударений с помощью '+'.
	TTS string `json:"tts,omitempty"`

	// Кнопки (suggest'ы), которые следует показать пользователю. Кнопки можно
	// использовать как релевантные ответу ссылки или подсказки для
	// продолжения разговора.
	Buttons []Button `json:"buttons,omitempty"`

	// Признак конца разговора:
	//
	// true — сессию следует завершить,
	//
	// false — сессию следует продолжить.
	EndSession bool `json:"end_session"`

	// Описание карточки — сообщения с поддержкой изображений.
	// Важно! Если указано данное поле, то поле text игнорируется.
	Card *Card `json:"card,omitempty"`
}

// AddURL добавляет к ответу кнопку с ссылкой.
func (r *Response) AddURL(title, url string) {
	if r.Buttons == nil {
		r.Buttons = make([]Button, 0)
	}

	r.Buttons = append(r.Buttons, Button{
		Title: title,
		URL:   url,
	})
}

// AddButton добавляет к ответу кнопку с полезной нагрузкой.
//
// Если полезная нагрузка не нужна, можно передать nil.
func (r *Response) AddButton(title string, payload interface{}) {
	if r.Buttons == nil {
		r.Buttons = make([]Button, 0)
	}

	r.Buttons = append(r.Buttons, Button{
		Title:   title,
		Payload: payload,
	})
}

// responseSession данные о сессии.
type responseSession struct {
	SessionID string `json:"session_id"`
	MessageID int    `json:"message_id"`
	UserID    string `json:"user_id"`
}

// response структура ответа серверу.
type response struct {
	Response Response        `json:"response"` // Данные для ответа.
	Session  responseSession `json:"session"`  // Данные о сессии.
	Version  string          `json:"version"`  // Версия протокола.
}

type eventFunc func(r Request) (Response, error)

// Webhook структура.
type Webhook struct {
	context2.LogObject
	event eventFunc
}

// NewWebhook возвращает новый Webhook.
func NewWebhook(l logger.Interface) *Webhook {
	return &Webhook{
		LogObject: context2.NewLogObject(l),
	}
}

// OnEvent обработчик скилла.
//
// Таймаут ожидания ответа — 5 секунд, после чего сервер Маруси завершит
// сессию.
func (wh *Webhook) OnEvent(f eventFunc) {
	wh.event = f
}

// HandleFunc обработчик http запросов.
func (wh *Webhook) HandleFunc(c HTTPContext) {
	mediatype, _, err := mime.ParseMediaType(c.GetHeader("Content-Type"))
	if err != nil {
		wh.Log(c.GetContext()).Error(
			"%s + "+http.StatusText(http.StatusBadRequest),
			errors.Wrap(err, "http - marusia server, bad content-type"),
		)
		c.SendErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))

		return
	}

	if mediatype != "application/json" {
		wh.Log(c.GetContext()).Error(
			"%s + "+http.StatusText(http.StatusBadRequest),
			"http - marusia server, bad body type",
		)
		c.SendErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))

		return
	}

	var req Request

	if errParse := c.ParseRequest(&req); errParse != nil {
		wh.Log(c.GetContext()).Error(errors.Wrap(errParse, "http - marusia server"))
		c.SendErrorResponse(http.StatusBadRequest, "invalid request body")

		return
	}

	req.Context = c.GetContext()

	resp, errWebhook := wh.event(req)

	if errWebhook != nil {
		wh.Log(c.GetContext()).Error(errors.Wrap(errWebhook, "webhook - error server"))
		c.SendErrorResponse(http.StatusInternalServerError, errWebhook.Error())

		return
	}

	fullResponse := response{
		Response: resp,
		Session: responseSession{
			SessionID: req.Session.SessionID,
			MessageID: req.Session.MessageID,
			UserID:    req.Session.Application.ApplicationID,
		},
		Version: Version,
	}

	// Возвращаем данные
	c.SetHeader("Content-Type", "application/json; encoding=utf-8")
	err = c.SendResponse(http.StatusOK, fullResponse)

	if err != nil {
		wh.Log(c.GetContext()).Error(errors.Wrap(errWebhook, "http - error send response"))
	}
}

// GinHandleFunc обработчик http запросов для gin.Context.
func (wh *Webhook) GinHandleFunc(c *gin.Context) {
	wh.HandleFunc(&GinHTTPContext{Context: c})
}

// BaseHTTPHandleFunc обработчик http запросов для gin.Context.
func (wh *Webhook) BaseHTTPHandleFunc(w http.ResponseWriter, r *http.Request) {
	wh.HandleFunc(&BaseHTTPContext{Req: r, Resp: w})
}
