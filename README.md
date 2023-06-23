# Chat-client for gRPC course project

[![build-and-test](https://github.com/satanaroom/chat_client/actions/workflows/build-and-test.yml/badge.svg?branch=main)](https://github.com/satanaroom/chat_client/actions/workflows/build-and-test.yml)
[![golangci-lint](https://github.com/satanaroom/chat_client/actions/workflows/golangci-lint.yml/badge.svg?branch=main)](https://github.com/satanaroom/chat_client/actions/workflows/golangci-lint.yml)

Chat-client - это CLI-приложение для создание чатов, подключения и отправку сообщений.
Клиент проходит аутентификацию в сервисе [auth](https://github.com/satanaroom/auth), 
получая **refresh** и **access** токены на основе `jwt`. Запросы на создание чата, 
отправку сообщение и подключения отправляются на сервис [chat_server](https://github.com/satanaroom/chat_server).

## Quick start
Для работы необходимо установить [Docker](https://docs.docker.com/engine/install/) и [Docker Compose](https://docs.docker.com/compose/install/).

1. Склонировать репозиторий
``` bash
git clone https://github.com/satanaroom/chat_client.git
```

2. Запустить контейнер Redis
``` bash
docker-compose up -d
```

3. Сбилдить проект
``` bash
make build
```

## CLI commands
Приложение поддерживает следующие команды:

**login** - авторизация на сервере чата при помощи логина и пароля.

Флаги: 
  * username, u - имя пользователя
  * password, p - пароль пользователя

Пример:
```bash
chat-client login -u john -p qwerty
```

**create** - создание комнаты чата.

Флаги:
* usernames, u - участники чата, разделённые запятой

Пример:
```bash
chat-client create -u john,semen,petr
```

**connect** - подключение к существующему чату. После успешного подключения становится доступен терминал для воода сообщений.

Флаги:
* chat-id, c - id существующего чата.

Пример:
```bash
chat-client connect -c 1234-AAAA-1234-AAAA
```

