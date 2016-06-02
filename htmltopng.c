#include<wkhtmltox/image.h>

#include<string.h>
#include<stdio.h>
#include<stdbool.h>

#include "xmltox.h"

void error_png(wkhtmltoimage_converter * c, const char * msg) {
	fprintf(stderr, "Error: %s\n", msg);
}

void warning_png(wkhtmltoimage_converter * c, const char * msg) {
	fprintf(stderr, "Warning: %s\n", msg);
}

// Pretty much copy pasted from wkhtmltopdf/examples/
int WkpngCreate(char *htmlFile, char **pngData) {

	wkhtmltoimage_global_settings * gs;
	wkhtmltoimage_converter * c;
	const unsigned char * data;
	long len;

	wkhtmltoimage_init(false);

	gs = wkhtmltoimage_create_global_settings();

	wkhtmltoimage_set_global_setting(gs, "in", htmlFile);
	wkhtmltoimage_set_global_setting(gs, "fmt", "png");

	c = wkhtmltoimage_create_converter(gs, NULL);
	wkhtmltoimage_set_error_callback(c, error_png);
	wkhtmltoimage_set_warning_callback(c, warning_png);

	printf("\n\n pngData(address passed) before allocation %p \n", *pngData);
	if (!wkhtmltoimage_convert(c)) {
		fprintf(stderr, "Conversion failed!");
		return -1;
	}

	printf("\n\n Data before allocation %p \n", data);
	len = wkhtmltoimage_get_output(c, &data);
	printf("\n\n Data after allocation %p \n", data);
	printf("%ld len\n", len);
	wkhtmltoimage_destroy_converter(c);
	*pngData = (char *) data; 
	printf("\n\n pngData(address passed) after allocation %p \n", *pngData);
	return len;
}

int GetPNGData(Status *status, char **buffer) {

	// At the end of the function. 
	printf("\n\n Address before allocation %p \n", *buffer);
	int len = WkpngCreate(status->htmlFileName, buffer);

	printf("\n\n Address after allocation %p \n", *buffer);
	if(len > 0) {
		status->imageConverted = true;
	}
	return len;
}

// Geerates the png and returns the length of the png generated.
int GenPNG(Status *status) {

	int len = WkpngCreate(status->htmlFileName, &status->pngData);

	printf("\n\n Address after allocation %p \n", status->pngData);
	if(len > 0) {
		status->imageConverted = true;
	}
	return len;
}

// Returns the pointer to the starting address of the generated PNG.
char* GetPNG(Status *status) {
	return status->pngData;
}
