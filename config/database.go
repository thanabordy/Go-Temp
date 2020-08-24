package config

func database() {
	config("DB_DRIVER", "pgsql")
	config("DB_HOST", env("DB_HOST", "127.0.0.1:5432"))
	config("DB_DATABASE", env("DB_DATABASE", "postgres"))
	config("DB_USER", env("DB_USER", "postgres"))
	config("DB_PASSWORD", env("DB_PASSWORD", ""))
}
