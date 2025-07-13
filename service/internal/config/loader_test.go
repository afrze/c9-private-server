package config

import "testing"

func TestLoadEnv(t *testing.T) {
	t.Setenv("MONGODB_URI", "mongodb://test")
	t.Setenv("MSSQL_CONN", "sqlserver://test")
	t.Setenv("JWT_SECRET", "secret")

	c := Load()

	if c.MongoURI != "mongodb://test" {
		t.Fatalf("expected MongoURI to be set")
	}
	if c.RateLimit != 60 {
		t.Fatalf("default RateLimit not applied")
	}
}
