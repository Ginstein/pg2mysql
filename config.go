package pg2mysql

type Config struct {
	MySQL struct {
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"mysql"`

	PostgreSQL struct {
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"postgresql"`

	MigrateTables []string `yaml:"migrate_tables"`
}

func MigrateTableContains(migrateTables []string, migrateTable string) bool {
	if len(migrateTables) == 0 {
		return true
	}
	for _, table := range migrateTables {
		if migrateTable == table {
			return true
		}
	}
	return false
}
