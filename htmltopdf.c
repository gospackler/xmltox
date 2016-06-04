#include<wkhtmltox/pdf.h>

#include<string.h>
#include<stdio.h>
#include<stdbool.h>

#include "xmltox.h"

void error_pdf(wkhtmltopdf_converter * c, const char * msg) {
	fprintf(stderr, "Error: %s\n", msg);
}

void warning_pdf(wkhtmltopdf_converter * c, const char * msg) {
	fprintf(stderr, "Warning: %s\n", msg);
}


wkhtmltopdf_global_settings* wkpdfInit(char *pdfFileName) {

	wkhtmltopdf_init(false);
	wkhtmltopdf_global_settings * gs;
	gs = wkhtmltopdf_create_global_settings();
	wkhtmltopdf_set_global_setting(gs, "out", pdfFileName);
	return gs;
}

bool wkpdfCreate(char *htmlFileName, wkhtmltopdf_global_settings* gs) {

	wkhtmltopdf_object_settings * os;
	wkhtmltopdf_converter * c;

	os = wkhtmltopdf_create_object_settings();
	wkhtmltopdf_set_object_setting(os, "page", htmlFileName);
	c = wkhtmltopdf_create_converter(gs);

	wkhtmltopdf_set_error_callback(c, error_pdf);
	wkhtmltopdf_set_warning_callback(c, warning_pdf);

	wkhtmltopdf_add_object(c, os, NULL);
	if (! wkhtmltopdf_convert(c)) {
		return false;
	}
	

	wkhtmltopdf_destroy_converter(c);
	return true;
}

void wkpdfDestroy() {

	wkhtmltopdf_deinit();
}

bool  GetPDFFile(Status *status) {
	// FIXME Create the appropriate hooks for the rest of the library to use.
	wkhtmltopdf_global_settings *gs = wkpdfInit(status->pdfFileName);
	wkpdfCreate(status->htmlFileName, gs);
}
