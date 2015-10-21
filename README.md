```
import "github.com/ondi/go-log"

log_rotate := log.NewRotateLogWriter(LogFile, LogSize, LogBackup, DuplicateOnStderr)
logger := log.NewLogger(LogLevel)
logger.SetOutput(log_rotate)
log.SetLogger(logger)
log.SetOutput(log_rotate)

log.Info("%v", "test")
```

