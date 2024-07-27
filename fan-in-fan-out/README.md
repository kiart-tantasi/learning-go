# Description

fan-in fan-out is also about goroutine but I create a separate directory for it because it is a big topic in Go

# Fan-in

```
in
	\
		---> process
	/
in
```


# Fan-out

```
	    	---> out
        /
process
        \
    		---> out
```
