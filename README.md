# XmltoX

What is the motive ?

* Render xmlfiles with the xsl file in stock. 
* Split into packages
    * package which takes in any url and renders returns the png.
        * This uses the firefox -marionete backend to get the work done
    *  the marionette client - Handle the sessions and has the wrappers for screenshot and Navigate which will be sent out to the clinet. 

xmltox - Two functions
 ```
 New() {
    // Initializes
    // xslFile
    // marionateClientObject
 }
 
 x.GetPNG(xmlContent) {
    // Returns the png as bytes. 
 }
 x.GetPDF(xmlContent) {
    // Converts the png to pdf.
 }
 ```
 
## Marionette

Marionette is a project from mozilla which allows access to the renderer of firefox called gecko. Run `firefox -marionette` behind `xvfb-run` as a headless browser.

``` bash
$ xvfb-run firefox -marionette
```

https://developer.mozilla.org/en-US/docs/Mozilla/QA/Marionette

## TODO
* Make xmltox use `github.com/gospacker/bulldozer` to make it possible to use process pools. 
	* Create a type task converter of type `task interface in bulldozer`
	* Create TaskConverter object and wrap in the taskInput struct.
		* inputData []byte 
		* GetConverter() - This has backing queue with `sync` locks to ensure that it could be accessed concurrently.
			- github.com/gospackler/bulldozer/queue
        * Response 
		* outPutData
