module github.com/theovidal/onyxcord

go 1.16

require (
	github.com/bwmarrin/discordgo v0.23.3-0.20210627161652-421e14965030
	github.com/fatih/color v1.10.0
	github.com/go-redis/redis/v8 v8.4.4
	github.com/joho/godotenv v1.3.0
	go.mongodb.org/mongo-driver v1.4.0
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/bwmarrin/discordgo => github.com/FedorLap2006/discordgo v0.22.1-0.20210810220050-f6231aba9904
