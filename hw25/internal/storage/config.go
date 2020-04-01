package storage


type Config struct {
	Server	       	string `yaml:"Server"`
	User           	string `yaml:"User"`
	Pass           	string `yaml:"Pass"`
	Database       	string `yaml:"Database"`
	TimeoutConnect 	int    `yaml:"TimeoutConnect"`
	TimeoutExecute 	int    `yaml:"TimeoutExecute"`
	Type 			string `yaml:"Type"`
}

