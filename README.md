# Boilerplate for nginx with Let’s Encrypt on docker-compose


## Инсталляция

Конфигурация для локальной разработки.

Для переноса на домен:
1. Раскомментировать директивы для Let`s Encrypt

2. Закомментировать подключение самоподписанных сертификатов 

3. Исправить конфигурации:
        - Add domains and email addresses to init-letsencrypt.sh
        - Replace all occurrences of example.org with primary domain (the first one you added to init-letsencrypt.sh) in data/nginx/app.conf

4. Выполнить скрипт:
        ./init-letsencrypt.sh
        `init-letsencrypt.sh` извлекает и обеспечивает продление сертификата Let’s

5. Run the server:
        docker-compose --compatibility up -d

6. Stop the server:
        docker-compose stop

PHPMyAdmin
https://localhost/pma
        login: root
        password: irol_root_pw


## License
All code in this repository is licensed under the terms of the `MIT License`. For further information please refer to the `LICENSE` file.

