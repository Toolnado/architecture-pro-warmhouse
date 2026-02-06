# Задание 1. Анализ и планирование

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Система поддерживает мануальное подключение новых устройств, с помощью выезда сотрудника на место.
- Пользователи могут управлять отоплением (включить/выключить).

**Мониторинг температуры:**

- Система собирает данные о температуре с помощью запросов от сервера к датчикам.
- Пользователи могут проверять температуру в домах через веб-интерфейс.

### 2. Анализ архитектуры монолитного приложения

- Язык программирования: Go
- База данных: PostgreSQL
- Архитектура: Монолитная, все компоненты системы (обработка запросов, бизнес-логика, работа с данными) находятся в рамках одного приложения.
- Взаимодействие: Синхронное, запросы обрабатываются последовательно.
- Масштабируемость: Ограничена, так как монолит сложно масштабировать по частям.
- Развертывание: Требует остановки всего приложения.

### 3. Определение доменов и границы контекстов

Управление пользователями:
- Авторизация в системе

Управление устройствами:
- Регистрация устройств в системе
- Управление отоплением

Телеметрия:
- Получение данных с датчиков
- Отправка данных на веб-интерфейс
- Отслеживание состояния устройств

### 4. Проблемы монолитного решения

- Сложность разработки и тестирования: Высокая связанность компонентов затрудняет разработку и повышает шанс возникновения ошибок.
- Сложность масштабирования: Невозможно масштабировать отдельные компоненты.
- Недоступность приложения: Требуется полная остановка для деплоя.

### 5. Визуализация контекста системы — диаграмма С4

[c4-context-diagram](https://www.planttext.com?text=VLDDSzem4BtxLsovb320Bpbn2aEQXXaeD4wQ8oFiZQs9FbnfNSV_lYjZWqDdUhAiNRttzkdfpfDmL6a5gzl3oXh23pei_Tf7NwH9LPKX54IbNyIn4_KybrHKnygZIqEleNbgTLmgOGB6p5Ecv35UtSxsOUjbkucry8Gklg3FI-wC8fAa49A6gngw4jnBaJkXOGQlmcbujCm4usKoMoUmjSOgffSfB5rQICAKAeStKRH7vwqP1rLJ63rQZROsw4RyimiBi6z01GyIu4OtVJy354QcKXXkN1isGYb1qXfuFbS2Hb8YBvNq17LXGPuwFGYJ0Rlm9dEMuQ50Fcnoq3Qh5FyBWhJon6hSLwxobwcZqy1ZtGt3y8D9LrQNmZI1TAcKhJs4cJoG1OVQ4gecC9ApwamNL32Q9Ek0aKjq8jWolv0TldtGk6lVjkA2IztnLGTtAWJcjNs63F_85FjHCUFLnS4M-WbLP_XWb5yIMVD0nRINUxOdCBpY4LxMZ190BiUl9IPeCW-x7_7kuO5njEPH_PIT-5sXfqyyxJKr9YqSerM-Wxu270XzGZvbF5-GSnpY_qErMCEmV0BAPSYt1uDrV_NJOInELlQIXZ3Ny-wBnikK_Mx-oPZXhBSnC4YrTUz1v5QOIgY-jGrk46iu9k1aNb28VIyr0fXCsbXRPVC64XBf-s8ouVqPR8J9Av7Z0hxRTkDUa3WApnjxwobtcFpSj8NXmriCNcHRvFJgYeThfInWO6B4Dmb1M1JTyJAy_xy0)


# Задание 2. Проектирование микросервисной архитектуры

**Диаграмма контейнеров (Containers)**

[c4-container-to-be](https://www.planttext.com?text=ZLRBajem4Bpp5GedjWhMbptbj3psGMN9aWMI8oNRWwr2jbnwu91K_ZqZMyR6DgbmeF2WwHxrJ0yyAasbDWadyuVzNAIQiXJaVYjcCFZ4qe2R44YiTQO-Uvwa-Nt4T6nyeq06U1fIVH-8nCiuJItAHEcGCz-RFqpMDhHRlNa9LHgaryI_p_3iOA2Pva0kKR9WD98q8HFoWyg4l0ga8K-1K6T4ICXe-pwPFP4L2wH0_XCBG96f36Ac8T16mjrWi0QfH3gozOt9y5Ka87A47kB33cCaZmL9Q4eZJ8s1QfP6XAOXEJ785K6M32IrM4C4snIy-vamQKZbUPHZLR4j2j7wAXpUaTy3Wg_BbKOv-8RPu-2JQPPn5Y2dI6qz7q03RJ-ifJYn40lQBSb1oB8-MrW8vHsh0crC0uS4j3nZWTTKD6FxY6h8wHb1f-ibUIcVBCUBiEyRH435iMloJ3B1KZqcKXYhmfXGesFyfgeGxt-DRpNTEzqhEJVbOyrGH2sIaEnN0JGccHG7XhIbxWboT6elRzf0NnI1BhwBIuYOqj8HA4X3gmrE8dA4s4THQTTXkkZN8DjMiIxVN72EmwKiXyBUX5JJCK7cuAWS1q4BQQEAUvPzj8Ti3RjbHyM5N-eO-fLO2ppiKrLmheNIaOJDjxTQoo30yTETYo2G4ws0ShhS1hiM3YS4xNjWaIarx8ElbIeOQf7wIRmjIq05bCDYLinQKx26KaLFMat1H8MZiVUbE1R6NPK1CYi25kuBFHof_J1LvpI8fKZHUqN7JUf6bkIucyXIR1lD_tDnyTDFFNAEtoSYDDns_jMjW5KHiF1rUWdNM0dL35WxLhl3Ba4fkBtI1t2tTohJhwuNpRXdozWQjjlrfW9mgTTUNj6ZtNJeQ9INg-hirQ6MEwTjCtSIRbckTcbfh8kh-a3wtFJD00v9KxVE31ZL5ApEhQs1op5KNYvnjQiYN82u4jgbDurmbUWeor6lBj1JSslKhp0Qng-ARlDtndTjVCvK3AfUIQFthzxxy_ETAwATqfdC3rpBZNoNrZxs3-Ps0suGztLm0PoTy8VqV-xUchllbBEGNfg65IHCARTXRgZVcSGNq9VUuLbQyJ_YEDe_3dy1)

**Диаграмма компонентов (Components)**

[c4-component-device-service](https://www.planttext.com?text=XLLBSzem4BxxLsmvaPaclkJKKuZ9WsbeQ83sw95jOJIn9LUF4ARJ_zwLRFa1peG3ICi-lls-bNIZD97Qb0L4rt4aoafmodKyfsyifMigxK_mXV6qC1c5lTQL-XQ6aXokSgRt9Z6AobHmZK5NgIZ3gY3SfhjKEYjO4aRNbojhsYwVmf8eJMNOBtHLeMyGQAOB2gqLvepaafHm2JKGy4Y226iHngcSaeh5ET7qG8vVOJ9RBU2Xtarm-o3izuimcYgGz8-XIYlOIL72Rvh0TZ6vwAMoJSIgBe5XMrpwYbsoNwHW6R7vh3jKL9PCAIQu6kGggL8aft4YnIkL6BUi3N3h33RJTx9x9NOnKqUUxgNWmYX0zah3MKeqveG3iWkPQtwGVvvCQskS9PYXeMTED4c8ScXNGkbSqlNF9xjRQo4HSXq3Y9zWjYvTV2iCpuWy-gmT2qFc9nVmDm3yj28v0QpqKXI5wzHg4BK69y3TUWEFX6VDlbugUDniLYUoz0L4XCDI3JX9SwQqF7ReNXfBvslBt1g5pIa5Jo9dgRKi2KSDMXwQH8xghyWytx7Sj5iagH0vY9rtryVAgJ0Aokg6RNMOejhm6QGrvLcjEh4BBNfodn1jz3ukMORK7OYqsipG0ilMuZeSx3uUrdgo1YF-CTsLK0m7Q42ujzaoTkvWbY9ePRVHotOE0eiv9b5O8NrJmBYRFZUbfmqs9CILIGgczcwOF8ihRvlqpjB0w3T3Cfjp2FKvn9B_Wc1zn0keZE_UzNHF4IFFur9afh2KFjO6M3g3pRiGcmQzPUY55iChvspy6uPADsgblSRyXFVZ-o7dmdhTV2X6ki1JnvCRAzhJz7NiNXejUtwQwcE4TxC_2hJht6OuYnYRa3LveoSNpI2sS-tVO0kEg7LeAjjsY4tBuoUukmDdZv0N7QRFFyBd-_kBOTdMUGnnotQ3chwpMgsn16D3seoU_w-97F4yUnNS6m0hYTEMkl6qWQTUPqFg7hHUjU066R0lw7y0)

**Диаграмма кода (Code)**

[c4-code-device-registry](https://www.planttext.com?text=tLNTJjim5BxtKnnnitPhyW0MGWmYqKfeGgLSJRimyLbcBObTszc82ky-9qvQ1rnGRvUBo3hV-J_7do-DPTesLGaP_X4vhh4GnkhsMd8an0fR8dHdo0GhDAiW6HHXr7HmfMIDjONPZPmJebZ-cnK8PqSLps172BWlBvanS8-wLzprikvBGAEIr9agQOILkZr0oVxuPN2DsaUQvVPf0Or1lU8Bvs4xfp3pq09GQwddWUqDsj2Cz-SLVzVcraNpe0c3TJRVVnorUb2SMOm5-pZB34kCspgpNhjNVWdxEVHz5xXPPHJS34HT1C9lhCA8UDEgcFYXApNcPYahZ0Btbjd6H81hZQuX_AkbO4M5wSRztlGh0a-g76U_ncs3naxAFQsoI06JI3wV4-942XySYjfb_9Frjs9SxOb7tvfmqfCvxrVtgbrbNakylxAzhf_ToIlxGHf-DuCiZkpZdJIsq6YY9IV07obyDjiopQxYj_o_wOdhoYTORZPtSCjQrCItPybgNg8UUfA0yNH8HruaGKFUKDvVHo5piYyh8WqvAvHFM2a49fJZWJxnDPRC2bcRNq8PCcQT91TZehJheo4ZdNV8OUG36Ep0-NEQHdQ3WgXKYPLx9mw-qlHYe3NgADCskfvYhyY00ijpLAySJExnGIModc4uuoGkSNA9DNVl9Fa7)

# Задание 3. Разработка ER-диаграммы

Добавьте сюда ER-диаграмму. Она должна отражать ключевые сущности системы, их атрибуты и тип связей между ними.

# Задание 4. Создание и документирование API

### 1. Тип API

Укажите, какой тип API вы будете использовать для взаимодействия микросервисов. Объясните своё решение.

### 2. Документация API

Здесь приложите ссылки на документацию API для микросервисов, которые вы спроектировали в первой части проектной работы. Для документирования используйте Swagger/OpenAPI или AsyncAPI.

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1) сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```

2) Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

3) Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.


# **Задание 6. Разработка MVP**

Необходимо создать новые микросервисы и обеспечить их интеграции с существующим монолитом для плавного перехода к микросервисной архитектуре. 

### **Что нужно сделать**

1. Создайте новые микросервисы для управления телеметрией и устройствами (с простейшей логикой), которые будут интегрированы с существующим монолитным приложением. Каждый микросервис на своем ООП языке.
2. Обеспечьте взаимодействие между микросервисами и монолитом (при желании с помощью брокера сообщений), чтобы постепенно перенести функциональность из монолита в микросервисы. 

В результате у вас должны быть созданы Dockerfiles и docker-compose для запуска микросервисов.
