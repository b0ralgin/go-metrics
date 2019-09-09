### сборка

~~~~
make all
~~~~

### Запуск

~~~~
export METRICS_PORT=:8080 #порт HTTP сервера
export METRICS_FILE=log.txt #куда пишутся входящие метрики
export METRICS_PERIOD=5s #периодичность, с которой сохраняются метрики

./bin/metrics server
~~~~
