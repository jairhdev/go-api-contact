package config

import "fmt"

const (
	configValueDev    string = "dev"
	configValueTest   string = "test"
	configValueDocker string = "docker"
	configValueProd   string = "prod"
)

type cfg struct {
	environment string
}

// ***********************************************************************

// Inicia configurações do sistema baseado em variáveis de ambiente
func NewConfig(env string) error {
	ok := (env == configValueDev ||
		env == configValueTest ||
		env == configValueDocker ||
		env == configValueProd)
	if !ok {
		return fmt.Errorf("invalid environment")
	}

	newConfig := cfg{
		environment: env,
	}

	newConfig.startCommon()
	newConfig.startDev()
	newConfig.startTest()
	newConfig.startDocker()
	newConfig.startProd()

	return nil
}
