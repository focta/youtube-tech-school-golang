package util

import "github.com/spf13/viper"

// Lesson12 3 ファイルの作成後にファイルからの読み込みを行う値の構造体を設定する
type Config struct {
	// Lesson12 4 Unmarshalingを使ってタグを設定から読み出す値との紐づけを記載する
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADRESS"`
}


// Lesson12 5 viperが設定を読み込むための関数を作る
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app") // app.env を読み取るため、ファイル名の部分のappを指定
	viper.SetConfigType("env") // app.env を読み取るため、ファイル拡張子の.envを指定

	viper.AutomaticEnv() // 環境変数からも読取りを行ってくれる

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return

}