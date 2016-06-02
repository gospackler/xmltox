#include "xmltox.h"

#include<stdlib.h>
#include<stdio.h>
#include<string.h>

int getFileSize(FILE *fPtr) {
	int size = 0;
	// Calculate size and return
	fseek(fPtr, 0, SEEK_END);
	size = ftell(fPtr);
	fseek(fPtr, 0, 0);
	return size;
}

char* getFileInBuffer(char *fileName) {
	FILE *fPtr;
	fPtr = fopen(fileName, "r");
	if(fPtr == NULL) {
		return "";
	}
	int fileSize = getFileSize(fPtr);
	char *buffer = malloc(sizeof(char) * fileSize);
	int size = fread(buffer, sizeof(char), fileSize, fPtr);
	
	return buffer;
}

Status* InitStatus(char *uid, char *xmlData, char *xslData) {
	Status *status = (Status*) malloc(sizeof(Status));
	status->XMLData = xmlData;
	status->XSLData = xslData;
	status->tmpFileName = uid;
	int len = strlen(uid);
	// 6 is .html + \0
	status->htmlFileName = (char*)malloc(len + 6);
	status->htmlFileName = strcpy(status->htmlFileName, status->tmpFileName);
	status->htmlFileName = strcat(status->htmlFileName, ".html");


	// 5 is .pdf + \0
	status->pdfFileName = (char*) malloc(len + 5);
	status->pdfFileName = strcpy(status->pdfFileName, status->tmpFileName);
	status->pdfFileName = strcat(status->pdfFileName, ".pdf");
	status->imageConverted = false; 
	status->pngData = NULL;
	return status;
}

Status* InitStatusFromFile(char *uid, char *xml, char *xsl) {
	char *xmlData = getFileInBuffer(xml);
	char *xslData = getFileInBuffer(xsl);
	return InitStatus(uid, xmlData, xslData);
}

bool FinishStatus(Status *status) {
	// Xml and XSL data comes from the go side so no need to de allocate
	if(status->imageConverted) {
		wkhtmltoimage_deinit();
	}

	// Finally
	free(status);
	return true;
}
