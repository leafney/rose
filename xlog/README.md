## xlog

**Simple log output to console**

## How to use

### Example

```go
    var xlog xlog.*Log
    
    func init() {
        // simple
        xlog = NewXLog()
        
        // more
        xlog = NewXLog(WithDebug(true)).
            SetDebug(true).
            //SetPrefix("").
            //SetPrefix("[hello]").
            SetEnable(true)
    }
	
    // show log
    xlog.Debugln("hello","world")
    xlog.Info("hello")
    xlog.Infof("hello %v", "world")
    xlog.Errorln("hello", "world")
    xlog.Fatal("fatal")

    // example
    // 2023/12/19 20:15:00 [XLog] [INFO]: hello world
```

### Log Level

- debug
- info (default)
- error
- fatal

### Config

- `WithDebug()` (default `False`)
- `WithLevel()` (default `Info`)
- `WithEnable()` (default `True`)
- `WithPrefix()` (default `[XLog]`)
- `WithFlags()`

---

- `SetDebug()`
- `SetLevel()`
- `SetEnable()`
- `SetPrefix()`
- `SetFlags()`

---