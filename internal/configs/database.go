package configs

type DBConfig struct {
	Dialect      string
	Dsn          string
	MaxOpenConn  int    // if <=0, then there is no limit on the number of open connections
	MaxIdleConn  int    // if <=0, no idle connections are retained
	ConnLifetime string // if d =0, connections are reused forever
	DisableLog   bool   // disable sql queries logging
}
