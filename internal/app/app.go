package app

import (
	"cmd/gate/main.go/internal/action2"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"gorm.io/driver/clickhouse"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cmd/gate/main.go/internal/action"
	"cmd/gate/main.go/internal/auth"
	"cmd/gate/main.go/internal/config"
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	log *slog.Logger
	srv *http.Server

	methods  map[string]jsonrpc.Method
	download map[string]jsonrpc.Method

	conn    *gorm.DB
	conn2   *gorm.DB
	conn3ch *gorm.DB

	auth auth.InterfaceAuth
}

func New(log *slog.Logger) InterfaceApp {
	return &App{
		log:      log,
		methods:  map[string]jsonrpc.Method{},
		download: map[string]jsonrpc.Method{},
	}
}

func (app *App) Init(config *config.Config) {
	app.srv = &http.Server{
		Addr: config.Port,
	}

	app.initRouter()

	conn, err := gorm.Open(postgres.Open(config.Storage.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	app.conn = conn

	conn2, err := gorm.Open(postgres.Open("host=45.9.27.162 user=root password=LfhTG7T7dzwmkXh dbname=sldb_ym port=5432"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	app.conn2 = conn2

	host := "ch.platina360.ru"
	port := 9443
	username := "zosimenko%40adventum.ru"
	password := "hi%2BRCVE%23LR9qxACE"
	database := "sberbank"

	dsn := fmt.Sprintf(
		"https://%s:%d?username=%s&password=%s&database=%s&secure=true",
		host, port, username, password, database,
	)

	conn3ch, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	app.conn3ch = conn3ch

	app.auth = auth.New()
	err = app.auth.Init(config.Auth.DSN)
	if err != nil {
		app.log.Error("Failed to init auth", "err", err)
		panic(err)
	}

	app.methods["bio.get.bio"] = action.BioGetBio

	app.methods["user.get.user"] = action.UserGetUser
	app.methods["user.get.users"] = action.UserGetUsers

	app.methods["group.create.group"] = action.GroupCreateGroup
	app.methods["group.get.group"] = action.GroupGetGroup
	app.methods["group.get.groups"] = action.GroupGetGroups
	app.methods["group.join.group"] = action.GroupJoinGroup
	app.methods["group.leave.group"] = action.GroupLeaveGroup

	app.methods["group.invite.users"] = action.GroupInviteUsers
	app.methods["group.add.links"] = action.GroupAddLinks
	app.methods["group.delete.user"] = action.GroupDeleteUser

	app.methods["metric.get.metric"] = action.MetricGetMetric
	app.methods["metric.get.table"] = action.MetricGetTable
	app.methods["metric.get.chart"] = action.MetricGetChart

	app.methods["link.create.link"] = action.LinkCreateLink
	app.methods["link.get.links"] = action.LinkGetLinks
	app.methods["link.get.link"] = action.LinkGetLink

	app.download["metric.download.chart"] = action.MetricDownloadChart
	app.download["metric.download.bar"] = action.MetricDownloadBar

	app.methods["ch.get.chart"] = action2.CHGetChart

}

func (app *App) Start() {
	app.log.Info("Listening on " + app.srv.Addr)
	err := app.srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (app *App) initRouter() {
	router := chi.NewRouter()
	router.Use(middleware.SetHeader("Access-Control-Allow-Origin", "*"))
	router.Use(middleware.SetHeader("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"))
	router.Use(middleware.SetHeader("Access-Control-Allow-Methods", " GET, PUT, POST, DELETE, OPTIONS"))
	// auth
	router.Post("/login", app.login)
	router.Post("/signup", app.signup)

	// api
	router.Route("/v1", func(r chi.Router) {
		r.Use(middleware.SetHeader("Access-Control-Allow-Methods", " GET, PUT, POST, DELETE, OPTIONS"))
		r.Use(app.authMiddleware)
		r.Post("/", app.handleRequest)
		r.Options("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("OPTIONS")
			w.WriteHeader(http.StatusOK)
		})
	})

	router.Route("/download", func(r chi.Router) {
		r.Use(middleware.SetHeader("Access-Control-Allow-Methods", " GET, PUT, POST, DELETE, OPTIONS"))
		r.Use(app.authMiddleware)
		//r.Post("/", app.handleDownloadRequest)
		r.Post("/", app.downloadCSVHandler)
		r.Post("/bar", app.downloadCSVBarHandler)
		r.Options("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("OPTIONS")
			w.WriteHeader(http.StatusOK)
		})
	})

	app.srv.Handler = router
}

func (app *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		type Response struct {
			Auth  bool  `json:"auth"`
			Error error `json:"error"`
		}

		response := Response{Auth: false, Error: nil}

		body, err := io.ReadAll(r.Body)
		defer func() {
			err = r.Body.Close()
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				response.Auth = false
				response.Error = err
				res, _ := json.Marshal(response)
				_, _ = w.Write(res)
				return
			}
		}()

		request := jsonrpc.Request{}
		err = json.Unmarshal(body, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.Auth = false
			response.Error = err
			res, _ := json.Marshal(response)
			_, _ = w.Write(res)
			return
		}

		claims, err := app.auth.VerifyToken(request.Token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.Auth = false
			response.Error = err
			res, _ := json.Marshal(response)
			_, _ = w.Write(res)
			return
		}

		r.Body = io.NopCloser(strings.NewReader(string(body)))
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "id", claims.ID)))
	})
}

func (app *App) login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		app.log.Error("request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type LoginResponse struct {
		Token string       `json:"token"`
		User  *entity.User `json:"user"`
	}

	lr := LoginRequest{}
	err = json.Unmarshal(body, &lr)
	if err != nil {
		app.log.Error("request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	token, err := app.auth.Login(lr.Email, lr.Password)
	if err != nil {
		app.log.Error("request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, err = w.Write([]byte(token))
	if err != nil {
		app.log.Error("request", "error", err)
		return
	}
}
func (app *App) signup(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		app.log.Error("request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	type SignupRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	sr := SignupRequest{}
	err = json.Unmarshal(body, &sr)
	if err != nil {
		app.log.Error("request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	token, id, err := app.auth.Signup(sr.Email, sr.Password)
	if err != nil {
		app.log.Error("request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	user := entity.User{
		ID:        id,
		FirstName: sr.FirstName,
		LastName:  sr.LastName,
		Email:     sr.Email,
		Role:      "USER",
	}

	err = app.conn.Create(&user).Error
	if err != nil {
		app.log.Error("request", "error", err)
		return
	}

	_, err = w.Write([]byte(token))
	if err != nil {
		app.log.Error("request", "error", err)
		return
	}
}

func (app *App) handleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer func() {
		err = r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}()

	request := jsonrpc.Request{}
	response := jsonrpc.Response{}

	id := r.Context().Value("id").(int)

	options := jsonrpc.Options{
		Log:     app.log,
		Conn:    app.conn,
		Conn2:   app.conn2,
		Conn3ch: app.conn3ch,
		UserID:  id,
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	response.Auth = true
	options.Params = request.Params

	if method, ok := app.methods[request.Method]; ok {
		response.Result, err = method(options)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			response.Error = err.Error()
		}
	} else if method, ok := app.download[request.Method]; ok {
		response.Result, err = method(options)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			response.Error = err.Error()
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "Method does not exist"
	}

	responseByteSlice, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if request.Method == "metric.get.file" {
		w.Header().Set("Content-Type", "application/octet-stream")

	} else {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseByteSlice)
}

func (app *App) handleDownloadRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer func() {
		err = r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}()

	request := jsonrpc.Request{}
	response := jsonrpc.Response{}

	id := r.Context().Value("id").(int)

	options := jsonrpc.Options{
		Log:    app.log,
		Conn:   app.conn,
		Conn2:  app.conn2,
		UserID: id,
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	response.Auth = true
	options.Params = request.Params

	if method, ok := app.download[request.Method]; ok {
		response.Result, err = method(options)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			response.Error = err.Error()
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "Method does not exist"
	}

	responseByteSlice, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
	}

	if request.Method == "metric.download.chart" || request.Method == "metric.download.bar" {
		w.Header().Set("Content-Type", "application/octet-stream")

	} else {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseByteSlice)
}

func (app *App) downloadCSVHandler(w http.ResponseWriter, r *http.Request) {
	// Установка заголовков ответа
	w.Header().Set("Content-Disposition", "attachment;filename=data.csv")
	w.Header().Set("Content-Type", "text/csv")

	body, err := io.ReadAll(r.Body)
	defer func() {
		err = r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}()

	request := jsonrpc.Request{}
	response := jsonrpc.Response{}

	id := r.Context().Value("id").(int)

	options := jsonrpc.Options{
		Log:    app.log,
		Conn:   app.conn,
		Conn2:  app.conn2,
		UserID: id,
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	response.Auth = true
	options.Params = request.Params

	if method, ok := app.download[request.Method]; ok {
		response.Result, err = method(options)

		// Создание CSV writer
		writer := csv.NewWriter(w)
		defer writer.Flush()

		type Result struct {
			Date      time.Time
			Dimension string
			PageViews int
			Users     int
		}

		var params []Result
		err = json.Unmarshal(response.Result, &params)
		if err != nil {
			fmt.Println(err)
		}

		if err := writer.Write([]string{"Date", "Dimension", "PageViews", "Users"}); err != nil {
			http.Error(w, "Error writing record to CSV", http.StatusInternalServerError)
			return
		}

		for _, record := range params {
			str := []string{record.Date.Format("02.01.2006 15:04"), record.Dimension, strconv.Itoa(record.PageViews), strconv.Itoa(record.Users)}

			if err := writer.Write(str); err != nil {
				http.Error(w, "Error writing record to CSV", http.StatusInternalServerError)
				return
			}
		}

		if err != nil {
			w.WriteHeader(http.StatusOK)
			response.Error = err.Error()
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "Method does not exist"
	}

}

func (app *App) downloadCSVBarHandler(w http.ResponseWriter, r *http.Request) {
	// Установка заголовков ответа
	w.Header().Set("Content-Disposition", "attachment;filename=data.csv")
	w.Header().Set("Content-Type", "text/csv")

	body, err := io.ReadAll(r.Body)
	defer func() {
		err = r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}()

	request := jsonrpc.Request{}
	response := jsonrpc.Response{}

	id := r.Context().Value("id").(int)

	options := jsonrpc.Options{
		Log:     app.log,
		Conn:    app.conn,
		Conn2:   app.conn2,
		Conn3ch: app.conn3ch,
		UserID:  id,
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	response.Auth = true
	options.Params = request.Params

	if method, ok := app.download[request.Method]; ok {
		response.Result, err = method(options)

		// Создание CSV writer
		writer := csv.NewWriter(w)
		defer writer.Flush()

		type Result struct {
			Dimension string
			PageViews int
			Users     int
		}

		var params []Result
		err = json.Unmarshal(response.Result, &params)
		if err != nil {
			fmt.Println(err)
		}

		if err := writer.Write([]string{"Dimension", "PageViews", "Users"}); err != nil {
			http.Error(w, "Error writing record to CSV", http.StatusInternalServerError)
			return
		}

		for _, record := range params {
			str := []string{record.Dimension, strconv.Itoa(record.PageViews), strconv.Itoa(record.Users)}

			if err := writer.Write(str); err != nil {
				http.Error(w, "Error writing record to CSV", http.StatusInternalServerError)
				return
			}
		}

		if err != nil {
			w.WriteHeader(http.StatusOK)
			response.Error = err.Error()
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		response.Error = "Method does not exist"
	}

}
