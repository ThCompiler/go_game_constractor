package runner

import (
	"encoding/json"
	"github.com/ThCompiler/go_game_constractor/marusia"
	"github.com/ThCompiler/go_game_constractor/pkg/language"
)

// RequestIn данные, полученные от пользователя.
type RequestIn struct {
	// Служебное поле: запрос пользователя, преобразованный для внутренней
	// обработки Марусей. В ходе преобразования текст, в частности, очищается
	// от знаков препинания, а числительные преобразуются в числа. При
	// завершении скилла по команде "стоп", "выход" и т.д. в скилл будет
	// передано "on_interrupt", чтобы у скилла была возможность попрощаться с
	// пользователем.
	Command string

	// Полный текст пользовательского запроса, максимум 1024 символа.
	OriginalUtterance string

	// Была ли нажата кнопка
	IsButton bool

	// JSON, полученный с нажатой кнопкой от обработчика скилла (в ответе на
	// предыдущий запрос), максимум 4096 байт.
	Payload json.RawMessage

	// Объект, содержащий слова и именованные сущности, которые Маруся
	// извлекла из запроса пользователя.
	NLU language.NLU
}

// Meta информация об устройстве, с помощью которого пользователь общается
// с Марусей.
type Meta struct {
	// Идентификатор клиентского приложения
	ClientID string

	// Язык в POSIX-формате, максимум 64 символа.
	Locale string

	// Название часового пояса, включая алиасы, максимум 64 символа
	Timezone string

	// Город пользователя на русском языке.
	CityRu string
}

// Session данные о сессии.
type Session struct {
	// Уникальный идентификатор сессии, максимум 64 символа.
	SessionID string

	// Идентификатор вызываемого скилла, присвоенный при создании.
	// Соответствует полю "Маруся ID" в настройках скилла.
	SkillID string

	// Признак новой сессии:
	//
	// true — пользователь начинает новый разговор с навыком,
	//
	// false — запрос отправлен в рамках уже начатого разговора.
	New bool

	// Идентификатор сообщения в рамках сессии, максимум 8 символов.
	// Инкрементируется с каждым следующим запросом.
	MessageID int

	// Данные об экземпляре приложения.
	Application Application

	// Данные о пользователе. Передаётся, только если пользователь авторизован.
	User User
}

// Application данные о приложении.
type Application struct {
	// Идентификатор экземпляра приложения, в котором пользователь общается с Марусей (максимум 64 символа).
	// Уникален в разрезе: «скилл + приложение (устройство)».
	ApplicationID string

	// Тип приложения (устройства). Возможные значения:
	//  • mobile;
	//  • speaker;
	//  • VK;
	//  • other.
	ApplicationType marusia.ApplicationType
}

type User struct {
	// Идентификатор аккаунта пользователя (максимум 64 символа). Уникален в разрезе: «скилл + аккаунт».
	UserId string

	// Идентификатор аккаунта пользователя в ВК, работает только если данное поле было включено разработчиками ВК навыков Маруси.
	// Не работает в отладки и локально.
	UserVKId string
}

// Request структура запроса jn голосового помощника.
type Request struct {
	// Информация об устройстве, с помощью которого пользователь общается с Марусей.
	Meta Meta

	// Данные, полученные от пользователя.
	Request RequestIn

	// Данные о сессии.
	Session Session
}

type Button struct {
	Title   string
	URL     string
	Payload interface{}
}

type Text struct {
	BaseText     string
	TextToSpeech string
}

type Result struct {
	Text          Text
	Buttons       []Button
	IsEndOfScript bool
}
