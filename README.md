# comet lib for golang

## Usage

### init
Execute once.

```
var cmt = comet.NewComet("realtimesession")
```

### Request "/"
Execute when first request.

```
	cmt.Start(w, r)
```

### Request "/post"
Execute when want end of "Wait()".
Using Argument "i" of "DoneAll()" for send data to "Wait()".

```
	cmt.DoneAll(i)
```

### Request "/wait"
Execute when want wait.

```
	i := cmt.Wait(r)
```

### Request "/exit"
Execute when exit.

```
	cmt.End(r)
```