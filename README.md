# XmltoX

`XmltoX` allows any link to be rendered to a `PNG` or `PDF`. It's a service that can be used for taking screenshots. The library uses the pool of goroutines to help with scaling. This is acheived through `bulldozer`.

* The library uses `firefox -marionete` backend to get the work done
* The marionette client - Handle the sessions and has the wrappers for screenshot and Navigate which will be sent out to the clinet. 

## Example

If `xvfb` is installed which is a linux virtual frame buffer, multiple instances of the renderer can run in the background to do the conversion. 

```
$ sh startup/firefoxrun.sh
```

An instance of firefox should work if marionette is enabled by default in the version of firefox used.

For converting a set of hits to goole to pdf try the test case. This uses the thread pool to get the task done. 

```
$ go test -v
```

Read through the code to get a better idea.

## Finalise 

```
$ killall -9 firefox
```

This will ensure that cpu is not hogged by the `firefox` zoombie processes. 


``` bash
$ xvfb-run firefox -marionette
```

## How it works ?

https://developer.mozilla.org/en-US/docs/Mozilla/QA/Marionette#How_does_it_work

### Worker queue using bulldozer.

If there are multiple requests and all of them need to be served, the library makes use of the backend

### Issues

* Please file issues if any to the library.

## Authors

- [George](www.github.com/georgethomas111)
