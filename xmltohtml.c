#include<libxml/HTMLtree.h>
#include<libxml/xmlIO.h>

#include<libxslt/xsltInternals.h>
#include<libxslt/xsltutils.h>
#include<libxslt/transform.h>

#include<string.h>
#include<stdbool.h>

#include "xmltox.h"

extern int xmlLoadExtDtdDefaultValue;

bool createFile(char *fileName, char *data) {
	FILE *fPtr = fopen(fileName, "w");
	if(fPtr != NULL) {
		fputs(data, fPtr);
	} else {
		return false;
	}
	fclose(fPtr);
	return true;
}

char* getHTML(char *xmlData, char *xslData) {

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

bool GetHTML(Status *status) {

	char *htmlData = getHTML(status->XMLData, status->XSLData);
	return createFile(status->htmlFileName, htmlData);
}
