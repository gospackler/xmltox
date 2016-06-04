#ifndef XMLTOX_H
#define XMLTOX_H

#include<stdbool.h>

/**
 * Status acts as a placeholder for information flow.
 */
typedef struct Status {
    char *XMLData;
    char *XSLData;
    char *tmpFileName;
    char *htmlFileName;
    char *pdfFileName;
    bool imageConverted;
    char *pngData;
}Status;

/*
 * This is the basic point of entry where in the xmlData is passed on to 
 * get the pdf data as the response.
 */
bool GetPDFFile(Status *);

/*
 * This is the basic point of entry where in the xmlData is passed on to 
 * get the pdf data as the response.
 * @return - length of the png data.
 */
int GetPNGData(Status *, char **);

// Splitting into two functions for easy use in go.
// Generates PNG and returns the length of PNG created. 
int GenPNG(Status *);

// Returns the pointer to the first PNG instance.
char* GetPNG(Status *);

Status* InitStatus(char *, char *, char *);

// C programs use this. de alllocate yourself the allocated one or FIXME later.
Status* InitStatusFromFile(char *, char *, char *);
bool FinishStatus(Status *);

/*
 * Create HTML based on the contents of the status object passed.
 */
bool GetHTML(Status *);

/*
 * Pdf data is generated when the html is passed to it.
 * Use the web kit qt bindings for getting this done.
 */
char* GetPDFFromHTML(char *htmlData);

#endif /* Header file ends */
