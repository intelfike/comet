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

```
	cmt.Done(htmlgen.ChatItem(name, text))
```

### Request "/wait"
Execute when want wait.

```
	i := cmt.Wait(r)
```