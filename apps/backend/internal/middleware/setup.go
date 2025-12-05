package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/mondayy1/llm-games/internal/config"
)

func Setup(cfg *config.Config, e *echo.Echo) {

	// ✅ RequestID: 각 요청에 고유한 ID 부여 (추적 및 디버깅 목적)
	e.Use(middleware.RequestID())

	// ✅ Logger: 요청 및 응답 로깅 설정
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${method} ${uri} ${status} request_id=${id}\n`,
	}))

	// ✅ Recover: 패닉 발생 시 복구 및 로그 출력
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 스택 크기: 1KB
		LogLevel:  log.ERROR,
	}))

	// ✅ Gzip: 응답 압축 (성능 최적화)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level:     5,   // 압축 레벨 (1-9)
		MinLength: 256, // 최소 압축 크기 (256바이트 이상만 압축)
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "metrics") // 특정 경로는 압축 제외
		},
	}))
}
