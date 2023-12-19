## xlog

**Simple log output to console**

## How to use

```go
    var xlog xlog.*Log
    
    func init() {
        // simple
        xlog = NewXLog(true)
        
        // more
        xlog = NewXLog(false).
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
