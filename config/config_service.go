package config

import "os"

// Parâmetros de configuração do sistema

// Sempre rodará
func (data cfg) startCommon() {
	os.Setenv("TEST_UNIT_COMMON", "common") // utilizado em teste unitário
	os.Setenv("FUSO", "America/Sao_Paulo")
}

func (data cfg) startDev() {
	if data.environment != configValueDev {
		return
	}
	os.Setenv("TEST_UNIT_DEV", "dev") // utilizado em teste unitário
	os.Setenv("HOST", "localhost")
	os.Setenv("HOST_PORT", "3001")

	os.Setenv("DB_POSTGRES_HOST", "localhost")
	os.Setenv("DB_POSTGRES_PORT", "5432")
	os.Setenv("DB_POSTGRES_NAME", "db-api-contact-dev")
	os.Setenv("DB_POSTGRES_USER", "postgres")
	os.Setenv("DB_POSTGRES_PWD", "12345")

	os.Setenv("MQ_RABBIT_HOST", "localhost")
	os.Setenv("MQ_RABBIT_PORT", "5672")
	os.Setenv("MQ_RABBIT_USER", "guest")
	os.Setenv("MQ_RABBIT_PWD", "guest")
}

func (data cfg) startTest() {
	if data.environment != configValueTest {
		return
	}
	os.Setenv("TEST_UNIT_TEST", "test") // utilizado em teste unitário
	os.Setenv("HOST", "localhost")
	os.Setenv("HOST_PORT", "3002")

	os.Setenv("DB_POSTGRES_HOST", "localhost")
	os.Setenv("DB_POSTGRES_PORT", "5432")
	os.Setenv("DB_POSTGRES_NAME", "db-api-contact-test")
	os.Setenv("DB_POSTGRES_USER", "postgres")
	os.Setenv("DB_POSTGRES_PWD", "12345")

	os.Setenv("MQ_RABBIT_HOST", "localhost")
	os.Setenv("MQ_RABBIT_PORT", "5672")
	os.Setenv("MQ_RABBIT_USER", "guest")
	os.Setenv("MQ_RABBIT_PWD", "guest")
}

func (data cfg) startDocker() {
	if data.environment != configValueDocker {
		return
	}
	os.Setenv("TEST_UNIT_DOCKER", "docker") // utilizado em teste unitário
	os.Setenv("HOST", "api-contact_nw")     //
	os.Setenv("HOST_PORT", "3003")

	os.Setenv("DB_POSTGRES_HOST", "dockerhost")
	os.Setenv("DB_POSTGRES_PORT", "5432")
	os.Setenv("DB_POSTGRES_NAME", "db-api-contact-docker")
	os.Setenv("DB_POSTGRES_USER", "postgres")
	os.Setenv("DB_POSTGRES_PWD", "12345")

	os.Setenv("MQ_RABBIT_HOST", "rabbitmq")
	os.Setenv("MQ_RABBIT_PORT", "5672")
	os.Setenv("MQ_RABBIT_USER", "guest")
	os.Setenv("MQ_RABBIT_PWD", "guest")
}

func (data cfg) startProd() {
	if data.environment != configValueProd {
		return
	}
	os.Setenv("TEST_UNIT_PROD", "prod") // utilizado em teste unitário
	os.Setenv("HOST", "localhost")
	os.Setenv("HOST_PORT", "3000")

	os.Setenv("DB_POSTGRES_HOST", "localhost")
	os.Setenv("DB_POSTGRES_PORT", "5432")
	os.Setenv("DB_POSTGRES_NAME", "db-api-contact")
	os.Setenv("DB_POSTGRES_USER", "postgres")
	os.Setenv("DB_POSTGRES_PWD", "12345")

	os.Setenv("MQ_RABBIT_HOST", "localhost")
	os.Setenv("MQ_RABBIT_PORT", "5672")
	os.Setenv("MQ_RABBIT_USER", "guest")
	os.Setenv("MQ_RABBIT_PWD", "guest")
}
