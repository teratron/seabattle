package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

type Logger struct {
	File string
	Info,
	Warning,
	Error,
	Debug *log.Logger
}

func New() *Logger {
	/*file, err := os.Create("./configs/warn.log")
	defer func() {
		err = file.Close()
	}()*/
	return &Logger{
		Info:    log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		Warning: log.New(os.Stdout, "WARNING\t", log.Ldate|log.Ltime|log.Lshortfile),
		Error:   log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		Debug:   log.New(os.Stderr, "DEBUG\t", log.Ldate|log.Ltime|log.Llongfile),
	}
}

// Помощник serverError записывает сообщение об ошибке в errorLog и
// затем отправляет пользователю ответ 500 "Внутренняя ошибка сервера".
func (l *Logger) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	err = l.Error.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Помощник clientError отправляет определенный код состояния и соответствующее описание
// пользователю. Мы будем использовать это в следующий уроках, чтобы отправлять ответы вроде 400 "Bad
// Request", когда есть проблема с пользовательским запросом.
func (l *Logger) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Мы также реализуем помощник notFound. Это просто
// удобная оболочка вокруг clientError, которая отправляет пользователю ответ "404 Страница не найдена".
func (l *Logger) notFound(w http.ResponseWriter) {
	l.clientError(w, http.StatusNotFound)
}

/*func New() *log.Logger {
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { err = f.Close() }()

	_ = log.New(f, "INFO\t", log.Ldate|log.Ltime)
	//
	return nil
}*/
