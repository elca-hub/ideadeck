package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"ideadeck/domain/repository"
	user_presenter "ideadeck/presenter/user"
	"ideadeck/usecase/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GinEngine struct {
	router     *gin.Engine
	port       Port
	ctxTimeout time.Duration
	sql        repository.SQL
	noSQL      repository.NoSQL
}

func NewGinServer(
	port Port,
	t time.Duration,
	db repository.SQL,
	session repository.NoSQL,
) *GinEngine {
	return &GinEngine{
		router:     gin.New(),
		port:       port,
		ctxTimeout: t,
		sql:        db,
		noSQL:      session,
	}
}

func (e *GinEngine) Listen() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	e.setupRouter(e.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf(":%d", e.port),
		Handler:      e.router,
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		// TODO: logの追加
		fmt.Println("web server running!")
		if err := server.ListenAndServe(); err != nil {
			// TODO: errorlog追加
			fmt.Println("web server stopped")
		}
	}()
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), e.ctxTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}

func (e *GinEngine) setupRouter(router *gin.Engine) {
	router.GET("/ping", e.healthCheckAction())

	router.POST("/api/v1/user", e.createUserAction())
}

func (e *GinEngine) healthCheckAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("hey")
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})
	}
}

func (e *GinEngine) createUserAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		uc := user.NewCreateUserInterator(e.sql.UserRepository(), e.noSQL.UserRepository(), user_presenter.NewCreatePresenter())

		_, err := uc.Execute(user.CreateUserInput{
			Name:     c.PostForm("name"),
			Email:    c.PostForm("email"),
			Password: c.PostForm("password"),
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"err":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
		})

		return
	}
}
