# Protey_microservice
**Проект для практики в НТЦ Протей**

## 1 Спринт
Реализован самый базовый каркас программы.
 Добавлен пакет Config с конфиг-файлом и go файлом, который инициализирует работу с конфигом
> **Важно, что на GitHub выложен .env файл. В реальном проекте я бы этого не делал, но в учебном задании для удобства проверяющих я его не добавлял в .gitignore**  

Для работы с конфигурационными файлами использую (Viper)[https://github.com/spf13/viper]  

Так же добавлен в pkg мой самописный логгер, который может работать с двумя видами логгов:
Info - обычная информация о том, как выполняется программа
Err - Ошибки, которые возникают во время работы программы  

Выглядит примерно вот так:  
![image](https://github.com/ykkssyaa/Protey_microservice/assets/64478650/e348fef7-1619-4327-bd2a-708277f50eb4)

---
