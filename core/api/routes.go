package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

//Routes is responsible for seting up all the handler function for the API
func (api *API) Routes() {
	api.Router = mux.NewRouter()
	api.Router.UseEncodedPath()
	// Add content-type json header !
	api.Router.Use(api.jsonMiddleware)
	api.Router.Use(api.authMiddleware)

	// GET
	api.Router.HandleFunc(api.Prefix+"/", api.GetIndex).Methods("GET")

	api.Router.HandleFunc(api.Prefix+"/users/{name}", api.GetUser).Methods("GET")

	api.Router.HandleFunc(api.Prefix+"/nodes", api.GetNodes).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/nodes/{id}", api.GetNode).Methods("GET")

	api.Router.HandleFunc(api.Prefix+"/projects", api.GetProjects).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}", api.GetProject).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/domains", api.GetDomains).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/domains/{domain}", api.GetDomain).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/algorithm", api.GetAlgorithm).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips", api.GetIPsByProjectName).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips/{ip}/ports", api.GetPortsByIP).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips/{ip}/ports", api.GetPortsByIP).Queries("protocol", "{protocol}").Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips/{ip}/ports/{port}", api.GetPortByIP).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips/{ip}/ports/{port}", api.GetPortByIP).Queries("protocol", "{protocol}").Methods("GET") // register optionnal protocol (tcp/udp...)
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips/{ip}/ports/{port}/uris", api.GetURIsByPort).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips/{ip}/ports/{port}/uris/{uri}", api.GetURIByPort).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips/{ip}/routes", api.GetRoutesByIP).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/raws", api.GetRawsByProject).Methods("GET")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/raws/{module}", api.GetRawsByModule).Methods("GET")

	// POST

	//	user
	api.Router.HandleFunc(api.Prefix+"/users/create", api.CreateUser).Methods("POST")
	api.Router.HandleFunc(api.Prefix+"/users/login", api.UserLogin).Methods("POST")
	api.Router.HandleFunc(api.Prefix+"/users/logout", api.UserLogout).Methods("POST")
	//	objects
	api.Router.HandleFunc(api.Prefix+"/projects", api.PostProject).Methods("POST")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/algorithm", api.PostAlgorithm).Methods("POST")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/run", api.RunModules).Methods("POST")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/run/{module}", api.RunModule).Methods("POST")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips", api.PostIP).Methods("POST")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/ips/{ip}/ports", api.PostPortsByIP).Methods("POST")
	api.Router.HandleFunc(api.Prefix+"/projects/{name}/domains", api.PostDomain).Methods("POST")

	//PUT
	//... updates ...

	// DELETE
	api.Router.HandleFunc(api.Prefix+"/projects/{name}", api.DeleteProject).Methods("DELETE")

	api.Router.NotFoundHandler = http.HandlerFunc(api.NotFound)
}
