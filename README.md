## RU Описание ProcessorService

ProcessorService — это сервис который читает сообщения из топика Handler брокера Kafka, изменяет JSON и отправляет его в другой топик.  Вот основные детали:

### Настройки

Вы можете настроить ProcessorService с использованием следующих флагов командной строки или переменных окружения:

- `SERVER_ADDRESS` (`-a`):**Адрес сервера**: По умолчанию — `localhost:9090`.
- `LOG_LEVEL` (`-l`): **Уровень логирования**: По умолчанию установлен на `info`.
- `BROKER_ADDRESS` (`-b`): **Адрес брокера Kafka**: По умолчанию — `localhost:29092`.

### Middleware

2. **Middleware логирования**:
    - Позволяет осуществлять сквозное логирование ы помощью библиотеки logrus.


### Service
- Создает woorker pool
- Запускает consumer
- С помощью воркеров обрабатывает данные поступающие от consumer
- Отправляет данные через producer


### Consumer

- Запускается для чтения сообщений из топика Handler.
- Читает сообщение с JSON, topic_name и requestID.
- Передает данные в слой service для их обработки.


### Producer

- Запускается для отправки сообщений с новыми данными в топики Kafka с названиями полученными из topic_name.


## Начало работы

Чтобы запустить ProcessorService, используйте следующую команду:

```bash
go run main.go -a <адрес_сервера> -l <уровень_логирования> -b <адрес_брокера>
```

Замените `<адрес_сервера>`, `<уровень_логирования>` и `<адрес_брокера>` на необходимые значения.

## Зависимости

ProcessorService зависит от следующих внешних библиотек:

- [GitHub - Sarama](https://github.com/IBM/sarama): Библиотека для работы с Kafka
- [Github.com/sirupsen/logrus](https://github.com/sirupsen/logrus): Библиотека для логирования
- [Github.com/natefinch/lumberjack](https://github.com/natefinch/lumberjack): Библиотека для конфигурирования записи в файл logrus
- [Github.com/caarlos0/env](https://github.com/caarlos0/env): Библиотека для получения данных из переменных окружения

Не стесняйтесь вносить свой вклад, сообщать об ошибках или давать обратную связь!

____________________________________________________________



## EN  ProcessorService

ProcessorService is a service that reads messages from the Handler topic of the Kafka broker, modifies JSON, and sends it to another topic. Here are the main details:

### Settings

You can configure ProcessorService using the following command-line flags or environment variables:

- `SERVER_ADDRESS` (`-a`): **Logging Level**: Defaults to `info`.
- `LOG_LEVEL` (`-l`): **Server Address**: Defaults to `localhost:9090`.
- `BROKER_ADDRESS` (`-b`): **Kafka Broker Address**: Defaults to `localhost:29092`.

### Middleware

2. **Logging Middleware**:
    - Enables cross-cutting logging using the logrus library.

### Service

- Creates a worker pool.
- Launches a consumer.
- Processes incoming data from the consumer using workers.
- Sends processed data through a producer.

### Consumer

- Starts to read messages from the Handler topic.
- Reads messages with JSON, topic_name, and requestID.
- Passes the data to the service layer for processing.

### Producer

- Starts for sending messages with new data to Kafka topics with names obtained from topic_name.

## Getting Started

To run SecondService, use the following command:

```bash
go run main.go -a <server_address> -l <logging_level> -b <kafka_broker_address>
```

Replace `<server_address>`, `<logging_level>`, and `<kafka_broker_address>` with the desired values.

## Dependencies

SecondService depends on the following external libraries:

- [GitHub - Sarama](https://github.com/IBM/sarama): Library for working with Kafka.
- [GitHub.com/sirupsen/logrus](https://github.com/sirupsen/logrus): Logging library.
- [GitHub.com/natefinch/lumberjack](https://github.com/natefinch/lumberjack): Library for configuring logrus file logging.
- [GitHub.com/caarlos0/env](https://github.com/caarlos0/env): Library for retrieving data from environment variables.

Feel free to contribute, report bugs, or provide feedback!
