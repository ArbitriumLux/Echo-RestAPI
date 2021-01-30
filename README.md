# Небольшое приложение предоставляющее REST API на Echo Framework 
##### Список запросов:
* GET localhost:1323/users - Вывести список Users
* GET localhost:1323/users/id - Вывести User по ID
* POST localhost:1323/users и тело запроса в JSON {"name":"Name"} - Добавить нового User
* PUT localhost:1323/users/id и тело запроса в JSON {"name":"Name"} - Редактировать User по ID
* DELETE localhost:1323/users/id - Удалить User по ID
---
По умолчанию имя JSON-файла с пользователями "Users.json"
---
