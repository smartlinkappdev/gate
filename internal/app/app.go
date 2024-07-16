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
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	log *slog.Logger
	srv *http.Server

	methods map[string]jsonrpc.Method
	conn    *gorm.DB
	conn2   *gorm.DB

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

	conn2, err := gorm.Open(postgres.Open("host=45.9.27.162 user=root password=LfhTG7T7dzwmkXh dbname=sldb_ym port=5432"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	app.conn2 = conn2

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

	app.methods["metric.get.metric"] = action.MetricGetMetric

	app.methods["link.create.link"] = action.LinkCreateLink
	app.methods["link.get.links"] = action.LinkGetLinks
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
