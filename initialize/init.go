package initialize

func StartInitSystemConfig(deps *Dependencies) {
	Migrate(deps)
	Site(deps)
	NodeSecret(deps)
	Node(deps)
	Email(deps)
	Device(deps)
	Invite(deps)
	Verify(deps)
	Subscribe(deps)
	Register(deps)
	Mobile(deps)
	Currency(deps)
	Telegram(deps)
}
