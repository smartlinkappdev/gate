package app

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"cmd/gate/main.go/internal/action"
	"cmd/gate/main.go/internal/auth"
	"cmd/gate/main.go/internal/config"
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	log *slog.Logger
	srv *http.Server

	methods map[string]jsonrpc.Method
	conn    *gorm.DB

	auth auth.InterfaceAuth
}

func New(log *slog.Logger) InterfaceApp {
	return &App{
		log:     log,
		methods: map[string]jsonrpc.Method{},
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

	app.auth = auth.New()
	err = app.auth.Init(config.Auth.DSN)
	if err != nil {
		app.log.Error("Failed to init auth", "err", err)
		panic(err)
	}

	app.methods["user.get.info"] = action.UserGetInfo
	app.methods["user.get.users"] = action.UserGetUsers
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

	// auth
	router.Post("/login", app.login)
	router.Post("/signup", app.signup)

	// api
	router.Route("/v1", func(r chi.Router) {
		r.Use(app.authMiddleware)
		r.Post("/", app.handleRequest)
	})

	app.srv.Handler = router
}

func (app *App) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		err = json.Unmarshal(body, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		claims, err := app.auth.VerifyToken(request.Token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
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
		Username string `json:"username"`
		Password string `json:"password"`
	}

	lr := LoginRequest{}
	err = json.Unmarshal(body, &lr)
	if err != nil {
		app.log.Error("request", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	token, err := app.auth.Login(lr.Username, lr.Password)
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
		Username  string `json:"username"`
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

	token, id, err := app.auth.Signup(sr.Username, sr.Password)
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
		Username:  sr.Username,
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
		Conn:   app.conn,
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

	if method, ok := app.methods[request.Method]; ok {
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseByteSlice)
}
