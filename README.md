# ⚡ music-service 
Реализация онлайн библиотеки песен 🎶🎶


- Получение данных библиотеки с фильтрацией по всем полям и
пагинацией
- Получение текста песни с пагинацией по куплетам
- Удаление песни
- Изменение данных песни
- Добавление новой песни

swagger


### Запуск из корня проекта

```make
docker compose -f ./deployment/docker-compose.yml up
```
После этого сервис будер доступен на `localhost:8080`

Документация swagger ui после docker compose доступна на localhost:8082  
[Документация openapi](./docs/openapi.yaml)



### Визуальная схема таблиц БД:

```sql

+------------------+            +----------------------+
|     groups       |            |        songs         |
+------------------+            +----------------------+
| id   (PK)        |<-----------| group_id (FK)        |
| name             |            | id   (PK)            |
+------------------+            | name                 |
                                | release_dat e        |
                                | lyrics               |
                                | link                 |
                                | group_id             |
                                +----------------------+

```

Для удобства база данных инициализируется с тестовыми данными 
- Миграции при старте приложения
- Индексы по полям фильтрации
- Используются транзакции при добавлении песни / группы

