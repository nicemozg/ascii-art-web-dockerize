FROM golang:latest

# Установите рабочую директорию внутри контейнера
WORKDIR /app

# Скопируйте содержимое текущей директории (включая файлы .go) внутрь контейнера
COPY . .

# Соберите приложение
RUN go build -o main .

# Установите переменную окружения PORT для указания порта, на котором будет работать веб-сервер
ENV PORT=8080

# Metadata with multiple authors
LABEL maintainers="nice.mozg@gmail.com, nurgisaserikkali@gmail.com" \
      version="1.0" \
      description="ASCII Art Web Application"

# Команда для запуска веб-сервера при запуске контейнера
CMD ["./main"]