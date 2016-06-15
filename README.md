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
 
 Get the POC of the PDF with gofpdf from the PNG.
* Complete test.go to use `gofpdf`
* Move on to the marionette implementation. 

