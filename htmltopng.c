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
int wkpngCreate(char *htmlFile, char **pngData) {

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

	if (!wkhtmltoimage_convert(c)) {
		fprintf(stderr, "Conversion failed!");
		return -1;
	}

	len = wkhtmltoimage_get_output(c, &data);
	printf("%ld len\n", len);
	wkhtmltoimage_destroy_converter(c);
	*pngData = (char *) data; 
	return len;
}

int GetPNGData(Status *status, char **buffer) {

	// At the end of the function. 
	int len = wkpngCreate(status->htmlFileName, buffer);
	if(len > 0) {
		status->imageConverted = true;
	}
	return len;
}
