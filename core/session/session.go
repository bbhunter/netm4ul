package session

import (
	"crypto/tls"
	"net"
	"strconv"
	"strings"

	"github.com/netm4ul/netm4ul/core/config"
	"github.com/netm4ul/netm4ul/modules"
	"github.com/netm4ul/netm4ul/modules/recon/dns"
	"github.com/netm4ul/netm4ul/modules/recon/nmap"
	"github.com/netm4ul/netm4ul/modules/recon/shodan"
	"github.com/netm4ul/netm4ul/modules/recon/traceroute"
	mgo "gopkg.in/mgo.v2"
)

// Connection type, to handle either use of TLS or not
type Connector struct {
	TLSConn *tls.Conn
	Conn    net.Conn
}

type Session struct {
	Modules      map[string]modules.Module
	Config       config.ConfigToml
	ConnectionDB *mgo.Session
	Connector    Connector
}

func NewSession(c config.ConfigToml) *Session {
	s := Session{
		Modules: make(map[string]modules.Module, 0),
	}
	// populate all modules
	s.Config = c
	s.loadModule()
	return &s
}

func (s *Session) Register(m modules.Module) {
	s.Modules[strings.ToLower(m.Name())] = m
}

func (s *Session) loadModule() {
	s.Register(traceroute.NewTraceroute())
	s.Register(dns.NewDns())
	s.Register(nmap.NewNmap())
}

func (s *Session) GetServerIPPort() string {
	return s.Config.Server.IP + ":" + strconv.FormatUint(uint64(s.Config.Server.Port), 10)
}

func (p *Session) loadModule() {
	p.Register(traceroute.NewTraceroute())
	p.Register(shodan.NewShodan())
	p.Register(dns.NewDns())
	p.Register(nmap.NewNmap())
}

func (s *Session) GetAPIIPPort() string {
	return s.Config.Server.IP + ":" + strconv.FormatUint(uint64(s.Config.API.Port), 10)
}
