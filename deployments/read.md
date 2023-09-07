Далее будет приведен пример установки docker и docker compose для убунуту:

Раздел установки из репозитория:

https://docs.docker.com/engine/install/ubuntu/

docker compose:

https://docs.docker.com/compose/install/linux/

Так же для запуска без прав sudo, необходимо:

https://stackoverflow.com/questions/48957195/how-to-fix-docker-got-permission-denied-issue

docker compose build

docker compose up

Если что то изменили в compose файле:

docker compose down

docker compose build (--no-cache флаг для полной пересборки например при обновлении git)

docker compose up