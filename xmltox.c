// #include<pdf.h>

// The webkit interaction lives here. 

// ----------------------------------


#include<libxml/HTMLtree.h>
#include<libxml/xmlIO.h>

#include<libxslt/xsltInternals.h>
#include<libxslt/xsltutils.h>
#include<libxslt/transform.h>

#include<wkhtmltox/pdf.h>

#include<stdbool.h>
#include<stdio.h>
#include<string.h>
#include<fcntl.h>

#define NOGRAPHICS 0
#define TMPHTML "tmp.html"
#define TMPPDF "tmp.pdf"

extern int xmlLoadExtDtdDefaultValue;

int getFileSize(FILE *fPtr) {
	int size = 0;
	// Calculate size and return
	fseek(fPtr, 0, SEEK_END);
	size = ftell(fPtr);
	fseek(fPtr, 0, 0);
	return size;
}


char* getFileInBuffer(FILE *fPtr) {

	int fileSize = getFileSize(fPtr);
	char *buffer = malloc(sizeof(char) * fileSize);
	/// Test code starts from here.
	int size = fread(buffer, sizeof(char), fileSize, fPtr);
	
	return buffer;
}

// Header files for 
// xmlChar 
// xmlNewDoc
char* GetHTML(char *xmlData, char *xslData) {

	xmlSubstituteEntitiesDefault(1);
	xmlLoadExtDtdDefaultValue = 1;
	xmlDocPtr xmlDoc = xmlParseDoc((xmlChar*) xmlData);
	xmlDocPtr xslDoc = xmlParseDoc((xmlChar*) xslData);

	xsltStylesheetPtr styleSheet = xsltParseStylesheetDoc(xslDoc);

	xmlDocPtr htmlDoc = xsltApplyStylesheet(styleSheet, xmlDoc, NULL);
	xmlChar *buf;
	int docLen;
	xsltSaveResultToString(&buf , &docLen, htmlDoc, styleSheet);

        // FIXME free buffers. 
	return buf;
}

bool createFile(char *htmlData) {
	FILE *htmlPtr = fopen(TMPHTML, "w");
	if(htmlPtr != NULL) {
		fputs(htmlData, htmlPtr);
	} else {
		return false;
	}
	fclose(htmlPtr);
	return true;
}


void error(wkhtmltopdf_converter * c, const char * msg) {
	fprintf(stderr, "Error: %s\n", msg);
}

void warning(wkhtmltopdf_converter * c, const char * msg) {
	fprintf(stderr, "Warning: %s\n", msg);
}


char* wkpdfCreate(char *fileName) {

	wkhtmltopdf_global_settings * gs;
	wkhtmltopdf_object_settings * os;
	wkhtmltopdf_converter * c;

	wkhtmltopdf_init(false);

	gs = wkhtmltopdf_create_global_settings();
	wkhtmltopdf_set_global_setting(gs, "out", fileName);
	os = wkhtmltopdf_create_object_settings();

	wkhtmltopdf_set_object_setting(os, "page", TMPHTML);
	c = wkhtmltopdf_create_converter(gs);

	wkhtmltopdf_set_error_callback(c, error);
	wkhtmltopdf_set_warning_callback(c, warning);

	wkhtmltopdf_add_object(c, os, NULL);
	if (! wkhtmltopdf_convert(c)) {
		return NULL;
	}

	//FIXME: Clear the created memory out of here.
	return fileName;
}

char* GetPDFFromHTML(char *htmlData) {

	char *pdfData; 
	bool status = createFile(htmlData);
	if(status) {
		pdfData = wkpdfCreate(TMPPDF);
	}
	return pdfData;
}

char* GetPDFFile(char *xmlData, char *xslData) {

    char *html = GetHTML(xmlData, xslData);
    return GetPDFFromHTML(html);
}

