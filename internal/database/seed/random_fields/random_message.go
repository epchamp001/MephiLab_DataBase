package random_fields

import (
	"ORM_BD/models"
	"math/rand"
)

var courierMessages = []string{
	"Ваш заказ доставлен, спасибо за выбор нашей службы.",
	"Я прибыл на адрес, пожалуйста, выйдите.",
	"Не могу найти ваш адрес, уточните местоположение.",
	"Задерживаюсь из-за пробки, скоро буду.",
	"Ваш заказ оставлен на ресепшене, проверьте, пожалуйста.",
}

var supportMessages = []string{
	"Здравствуйте! Как я могу вам помочь?",
	"Ваш запрос принят в обработку.",
	"Извините за неудобства. Мы работаем над решением проблемы.",
	"Ваш возврат был успешно обработан.",
	"Пожалуйста, уточните детали вашего вопроса.",
	"Ваше обращение передано в соответствующий отдел.",
}

var clientMessages = []string{
	"Когда будет доставлен мой заказ?",
	"Я не получил чек на оплату.",
	"Курьер опоздал, прошу уточнить статус доставки.",
	"Можно ли изменить адрес доставки?",
	"Спасибо за быструю доставку!",
}

func GenerateMessage(senderType models.SenderTypeEnum) string {
	switch senderType {
	case models.ClientSender:
		return clientMessages[rand.Intn(len(clientMessages))]
	case models.CourierSender:
		return courierMessages[rand.Intn(len(courierMessages))]
	case models.SupportStaffSender:
		return supportMessages[rand.Intn(len(supportMessages))]
	default:
		return "Неопределенный отправитель."
	}
}
