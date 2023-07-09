
```
f, e := os.OpenFile("/home/fanyx/mine/FeMPC/pprof", os.O_RDWR|os.O_CREATE, 0666)
			always.Nil(e)
			e = pprof.Lookup("heap").WriteTo(f, 0)
			always.Nil(e)
			f.Close()

go tool pprof "file"            
```