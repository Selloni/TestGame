Задачка такая: сделать онлайн игру - грузчики

Есть заказчик, есть грузчики. Заказчику необходимо переносить тяжелые грузы.
Заказчик обладает следующими свойствами:
- стартовый капитал (10 000р - 100 000р)
- умение нанимать грузчика
- набор заданий, которые нужно выполнить
  Каждый грузчик обладает следующими свойствами:
- максимально переносимый вес (5кг-30кг)
- "пьянство" (true,false)
- усталость (0-100%)
- зарплата (10 000р - 30 000р)

Суть следующая: генерируется N случайных заданий 
("название переносимых предметов", "вес"). 
Есть грузчики зарегистрировавшиеся на работу и получившие случайные свойства.
Задача заказчика - выбрать нужный набор грузчиков и выполнить задания.

Примечание: после кажной выполненной задачи грузчик устает на 20%,
если у грузчика есть вредная привычка его усталось повышаеться до 50%



    make docker - запускает контейнер с posqresql
    
    make run - запуск приложения

http://localhost:8080/

register: регистрация по логину, паролю и роли

    {
        "login": "llog",
        "password": "sdsd11",
        "role": "customer"
    }

login: вход по логину, паролю и роли (-_-

    {
        "login": "algin",
        "password": "qwer",
        "role": "loader"
    }

start: выбрать одно задание и несколько грузчиков(Доступно только заказику)

    {
        "taskId": 4, 
        "loaders": [5,2,4] // id грузчиков для найма
    }

tasks
- публичное: создание 5-2 заданий
- заказчик: список невыполненых задач
- грузчик: список выполненых задач

me: показать характеристики авторизованных пользователей




