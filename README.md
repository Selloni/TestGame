http://localhost:8080/

register: регистрация по логину, паролю и роли

    {
        "login": "llog",
        "password": "sdsd11",
        "role": "customer"
    }

login: вход по логину, паролю и роли (-_-

    {
        "login": "llog",
        "password": "sdsd11",
        "role": "customer"
    }

start: выбрать одно задание и несколько грузчиков

    {
        "taskId": 4,
        "loaders": [5,2,4]
    }

tasks
- публичное: создать 5-2 задания 
- заказчик: список невыполненых задач
- грузчик: список выполненых задач

me: показать свои характеристики

Примечание: если у грузчика есть вредная привычка, он устает сильнее
