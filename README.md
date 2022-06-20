## About The Project
Technical task given to me by Constanta company to be accepted to their internship program.
<br>
(Golang Junior developer)

It's the emulator of the banking system built relying on Golang Clear Architecture (DI) with implemented CRUD operations upon Transactions entity with all required constraints and error handling.
<br>
Also JWT authentication has been added.

### Built With
* [Golang](https://go.dev/)
* [Gin Web Framework](https://gin-gonic.com/)
* [PostgreSQL](https://www.postgresql.org/)

## Open any CLI and run project using the following commands (make sure you are in the main project directory)
* `docker-compose up --build constanta-golang-emulator-task`
* `migrate -path ./migrations -database postgres://postgres:qwerty123@localhost:5437/postgres?sslmode=disable up`
* Open any browser and use this URL (or just send requests using Postman):
`localhost:8888/api/`
* Check the `Emulator API.postman_collection.json` file for available enpoints

## License
Distributed under the MIT License. See `LICENSE.txt` for more information.

## Contact
Zhaksylyk Ashim
<br>
Telegram messenger - [@duramash](https://t.me/duramash)
<br>
E-mails - duramash.02@gmail.com, z_ashim@kbtu.kz
<br>
Project Link: [https://github.com/DuramAsh/constanta-golang-emulator-task](https://github.com/DuramAsh/constanta-golang-emulator-task)
