## xlog

**Debug log output to console**

## How to use

```go
    var xlog xlog.Log
    
    func init() {
        xlog = xlog.NewLog(true)
    }
    
	// Set whether to output logs to console 
	xlog.SetDebug(false)

    // Set log message prefix
	xlog.SetPrefix("debug: ")

    // show log
	xlog.Println("hello world")

    xlog.Printf("hello %s\n","world")
    
    xlog.Printfn("hello %s","world")
```