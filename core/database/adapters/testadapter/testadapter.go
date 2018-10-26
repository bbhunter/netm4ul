package testadapter

import (
	"errors"
	"strconv"

	"github.com/netm4ul/netm4ul/core/config"
	"github.com/netm4ul/netm4ul/core/database/models"
	"github.com/netm4ul/netm4ul/tests"
)

//Test is the base struct for the "testadapter" adapter.
type Test struct {
	cfg *config.ConfigToml
}

//InitDatabase is only there to sets up the configuration of the database.
// This adapters doesn't require any connection or any other setup.
func InitDatabase(c *config.ConfigToml) *Test {
	test := Test{}
	test.cfg = c
	return &test
}

// General purpose functions

//Name return the name of the adapter. It is exported because it's called from other core packages.
func (test *Test) Name() string {
	return "TestAdapter"
}

//SetupDatabase TOFIX. It will probably just return nil because this adapters should not require any setup.
func (test *Test) SetupDatabase() error {
	return errors.New("Not implemented yet")
}

//DeleteDatabase TOFIX.
func (test *Test) DeleteDatabase() error {
	return errors.New("Not implemented yet")
}

//SetupAuth doesn't do anything. This adapter doesn't require any authentification setup.
func (test *Test) SetupAuth(username, password, dbname string) error {
	return nil
}

//Connect doesn't do anything. This adapter doesn't require any connection. It's there only to implement the interface.
func (test *Test) Connect(*config.ConfigToml) error {
	return nil
}

//GetUser return a models.User if the provided username correspond to the test username (stored in the /tests/values.go file).
func (test *Test) GetUser(username string) (models.User, error) {
	if tests.NormalUser.Name == username {
		return tests.NormalUser, nil
	}
	return models.User{}, models.ErrNotFound
}

//GetUserByToken return a models.User if the provided token correspond to the test token (stored in the /tests/values.go file).
func (test *Test) GetUserByToken(token string) (models.User, error) {
	if tests.NormalUser.Token == token {
		return tests.NormalUser, nil
	}
	return models.User{}, models.ErrNotFound
}

//CreateOrUpdateUser is a no-op.
func (test *Test) CreateOrUpdateUser(user models.User) error {
	return nil
}

//GenerateNewToken changes the in memory value of the user token
func (test *Test) GenerateNewToken(user models.User) error {
	tests.NormalUser.Token = "Changed token"
	return nil
}

//DeleteUser TOFIX
func (test *Test) DeleteUser(user models.User) error {
	return errors.New("Not implemented yet")
}

// Project

//CreateOrUpdateProject is a no-op
func (test *Test) CreateOrUpdateProject(projectName models.Project) error {
	return nil
}

//GetProjects returns the projects stored in the /tests/values.go file
func (test *Test) GetProjects() ([]models.Project, error) {
	return tests.NormalProjects, nil
}

//GetProject returns the project named "projectName" if it exist in the /tests/values.go file
//It returns an error if the project doesn't exist.
func (test *Test) GetProject(projectName string) (models.Project, error) {
	ps, err := test.GetProjects()
	var projects []models.Project
	projects = append([]models.Project{}, ps...)

	if err != nil {
		return models.Project{}, errors.New("Could not get projects" + err.Error())
	}

	for _, p := range projects {
		if p.Name == projectName {
			return p, nil
		}
	}
	return models.Project{}, models.ErrNotFound
}

//DeleteProject TOFIX
func (test *Test) DeleteProject(project models.Project) error {
	return errors.New("Not implemented yet")
}

// IP

//CreateOrUpdateIP is a no-op
func (test *Test) CreateOrUpdateIP(projectName string, ip models.IP) error {
	return nil
}

//CreateOrUpdateIPs is a no-op
func (test *Test) CreateOrUpdateIPs(projectName string, ip []models.IP) error {
	return nil
}

//GetIPs returns the IP addresses from the /tests/values.go file
func (test *Test) GetIPs(projectName string) ([]models.IP, error) {
	return tests.NormalIPs, nil
}

//GetIP returns the IP address from the /tests/values.go file given the project name and an ip string
func (test *Test) GetIP(projectName string, ip string) (models.IP, error) {
	return tests.NormalIPs[0], nil
}

//DeleteIP TOFIX
func (test *Test) DeleteIP(ip models.IP) error {
	return errors.New("Not implemented yet")
}

// Domain

//CreateOrUpdateDomain is a no-op
func (test *Test) CreateOrUpdateDomain(projectName string, domain models.Domain) error {
	return errors.New("Not implemented yet")
}

//CreateOrUpdateDomains is a no-op
func (test *Test) CreateOrUpdateDomains(projectName string, domain []models.Domain) error {
	return errors.New("Not implemented yet")
}

//GetDomains return all the domains
func (test *Test) GetDomains(projectName string) ([]models.Domain, error) {
	return tests.NormalDomains, nil
}

//GetDomain TOFIX
func (test *Test) GetDomain(projectName string, domain string) (models.Domain, error) {
	return models.Domain{}, errors.New("Not implemented yet")
}

//DeleteDomain TOFIX
func (test *Test) DeleteDomain(projectName string, domain models.Domain) error {
	return errors.New("Not implemented yet")
}

// Port

//CreateOrUpdatePort is a no-op
func (test *Test) CreateOrUpdatePort(projectName string, ip string, port models.Port) error {
	return nil
}

//CreateOrUpdatePorts is a no-op
func (test *Test) CreateOrUpdatePorts(projectName string, ip string, port []models.Port) error {
	return nil
}

//GetPorts return multiple ports
func (test *Test) GetPorts(projectName string, ip string) ([]models.Port, error) {
	return tests.NormalPorts, nil
}

//GetPort return a port
func (test *Test) GetPort(projectName string, ip string, port string) (models.Port, error) {
	ports, err := test.GetPorts(projectName, ip)
	if err != nil {
		return models.Port{}, err
	}

	for _, p := range ports {
		if strconv.Itoa(int(p.Number)) == port {
			return p, nil
		}
	}
	return models.Port{}, models.ErrNotFound
}

//DeletePort TOFIX
func (test *Test) DeletePort(projectName string, ip string, port models.Port) error {
	return errors.New("Not implemented yet")
}

// URI (directory and files)

//CreateOrUpdateURI is a no-op
func (test *Test) CreateOrUpdateURI(projectName string, ip string, port string, uri models.URI) error {
	return nil
}

//CreateOrUpdateURIs is a no-op
func (test *Test) CreateOrUpdateURIs(projectName string, ip string, port string, uris []models.URI) error {
	return nil
}

//GetURIs return all the URIs
func (test *Test) GetURIs(projectName string, ip string, port string) ([]models.URI, error) {
	return tests.NormalURIs, nil
}

//GetURI return one URI
func (test *Test) GetURI(projectName string, ip string, port string, uri string) (models.URI, error) {
	uris, err := test.GetURIs(projectName, ip, port)
	if err != nil {
		return models.URI{}, err
	}
	for _, u := range uris {
		if u.Name == uri {
			return u, nil
		}
	}
	return models.URI{}, models.ErrNotFound
}

// DeleteURI TOFIX
func (test *Test) DeleteURI(projectName string, ip string, port string, dir models.URI) error {
	return errors.New("Not implemented yet")
}

// Raw data

//AppendRawData is a no-op
func (test *Test) AppendRawData(projectName string, raw models.Raw) error {
	return nil
}

//GetRaws return all the raws for one project name
func (test *Test) GetRaws(projectName string) ([]models.Raw, error) {
	raws, ok := tests.NormalRaws[projectName]
	if !ok {
		return []models.Raw{}, models.ErrNotFound
	}
	return raws, nil
}

//GetRawModule return all the raws for the provided module
func (test *Test) GetRawModule(projectName string, moduleName string) (map[string][]models.Raw, error) {
	raws, err := test.GetRaws(projectName)
	if err != nil {
		return nil, err
	}

	if len(raws) == 0 {
		return nil, models.ErrNotFound
	}
	//TOFIX : return actual raw data...
	return map[string][]models.Raw{}, nil
}
