package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func RegisterSwaggerHandler(ctx context.Context, mux *runtime.ServeMux, swaggerPath, swaggerDir, defaultJsonFile string) {
	// Serve Swagger specification files (JSON or YAML)
	swaggerFS := http.FileServer(http.Dir(swaggerDir))
	mux.HandlePath("GET", swaggerPath+"/*/*", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		http.StripPrefix(swaggerPath, swaggerFS).ServeHTTP(w, r)
	})

	// // Redirect to Swagger UI index.html on the official CDN with a query parameter pointing to your local Swagger file
	// mux.HandlePath("GET", swaggerPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// 	swaggerFileURL := swaggerPath + "gobaseservice.swagger.json" // Adjust for JSON if needed
	// 	uiURL := fmt.Sprintf("https://petstore.swagger.io/?url=http://localhost:8080%s", swaggerFileURL)

	// 	http.Redirect(w, r, uiURL, http.StatusTemporaryRedirect)
	// })

	// Serve a simple HTML page that references the Swagger UI from the official CDN
	mux.HandlePath("GET", swaggerPath, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
            <!DOCTYPE html>
            <html>
            <head>
                <title>Swagger UI</title>
                <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/4.18.1/swagger-ui.css" />
            </head>
            <body>
                <div id="swagger-ui"></div>
                <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/4.18.1/swagger-ui-bundle.js"></script>
                <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/4.18.1/swagger-ui-standalone-preset.js"></script>
                <script>
                    window.onload = function() {
                        SwaggerUIBundle({
                            url: "`+swaggerPath+defaultJsonFile+`",  // Adjust this to point to your Swagger file
                            dom_id: '#swagger-ui',
                            presets: [
                                SwaggerUIBundle.presets.apis,
                                SwaggerUIStandalonePreset
                            ],
                            layout: "StandaloneLayout"
                        });
                    };
                </script>
            </body>
            </html>
        `)
	})
}
