# Шахматы
## Общая инофрмация
Приложение использует Postgres в качестве БД для хранения информации о юзерах, играх, истории ходов и состояний досок во время игры. 

## Запуск приложения
Приложение поднимается в докере.

Файл ```env.example``` содержит

1) POSTGRES_URL - его при желании можно поменять на свой в своем файле ```.env```
2) JWT_SECRET - приложение использует jwt для авторизации
3) APP_URL - необязательное поле для локального тестирования

Для начала работы достаточно выполнить команду:

```docker-compose up -d```

## Работа с приложением

API описано сваггером, который доступен по ```localhost:8080```. На данный момент поддерживается следующий функционал:

### User

Позволяет создать пользоваетля (игрока). Для регистрации нового пользователя необходимо отправть новый пароль.

Возвращает id нового пользователя

### Session

Создаёт jwt для пользователя. Необходимо отправить id и password пользователя. 

Возвращает jwt.

Все дальнейшие действия потребуют авторизации по полученному jwt.

### Game

Созаёт игроку и/или позволяет присоединиться игроку к уже имеющейся игре.

Необходимо указать желаемый цвет, за которой будете играть. true - белый, false - чёрный. Если будет найдена игра с одним игроком противоположного цвета - вас добавят в неё. Если нет, то будет создана новая игра с вами. Игра начинатется после присоединения второго игрока.

Возвращает id игры.

### Move

Позволяет сделать ход в уже созданой игре. 

Необходимо указать id игры и ход в фотмате 

```"from": "string", "to": "string", "newFigure": 0,```

Поля ```from``` и ```to``` это поля на доске. Узавать нужно в формате ```E2```.  ```newFigure``` - на случай, если ваша пешка достигла последней горизонтали, необходимо указать новую фигуру, в которую она превращатеся. Указывается rune первой буквы названия фигуры (например queen - 'q' = 113).

Возвращает запись move.

### Give-up

Возможность сдаться до окончания партии. Необходимо отправть id игры.

### History

Получить историю ходов по игре. Необходимо отправть id игры.

### Board

Получить информацию о текущем состоянии доски. Необходимо отправть id игры.


