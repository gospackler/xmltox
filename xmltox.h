// This is the header file for xmlrunner 
// xml runner is the starting point

// xml file 
// -> pdf
// -> png

typedef struct XMLRunner {
    char *XMLFileName;
    char *XMLData;
}XMLRunner;

/*
 * This is the basic point of entry where in the xmlData is passed on to 
 * get the pdf data as the response.
 */
char* GetPDFFile(char *, char *);

/*
 * This is the basic point of entry where in the xmlData is passed on to 
 * get the pdf data as the response.
 */
char* GetPngData(char *);

/*
 * When XML is passed, we get the Html out 
 * making use of the libxslt implementation.
 */
char *GetHTML(char *, char *);

/*
 * Pdf data is generated when the html is passed to it.
 * Use the web kit qt bindings for getting this done.
 */
char* GetPDFFromHTML(char *htmlData);
