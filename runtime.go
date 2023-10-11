package referralservice

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"referralservice/gen/models"

	"github.com/go-openapi/errors"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewRuntime() *Runtime {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := initDB()
	rt := &Runtime{
		errorLog: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		Logger:   log.New(os.Stdout, "[Referral-Service] ", log.Ldate|log.Ltime|log.Lshortfile),

		db: db,
	}

	rt.RunMigration()

	return rt
}

type Runtime struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	*log.Logger

	db *gorm.DB
}

func (r *Runtime) Info() *log.Logger {
	return r.infoLog
}

func (r *Runtime) Error() *log.Logger {
	return r.errorLog
}

func (rt *Runtime) Debugf(format string, args ...interface{}) {
	rt.Printf("[DEBUG] "+format, args...)
}

func (rt *Runtime) Infof(format string, args ...interface{}) {
	rt.Printf("[INFO] "+format, args...)
}

func (rt *Runtime) Warnf(format string, args ...interface{}) {
	rt.Printf("[WARN] "+format, args...)
}

func (rt *Runtime) Errorf(format string, args ...interface{}) {
	rt.Printf("[ERROR] "+format, args...)
}

func (r *Runtime) SetError(code int, msg string, args ...interface{}) error {
	return errors.New(int32(code), msg, args...)
}

func (r *Runtime) GetError(err error) errors.Error {
	if v, ok := err.(errors.Error); ok {
		return v
	}

	return errors.New(http.StatusInternalServerError, err.Error())
}

func (r *Runtime) DB() *gorm.DB {
	return r.db
}

// initDB DB initialization
func initDB() *gorm.DB {
	// DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panicf("[ERROR] db connection %v", err)
	}

	return db
}

func (r *Runtime) RunMigration() {
	r.Info().Println("Migrating DBs")
	r.db.AutoMigrate(&models.User{})
}
