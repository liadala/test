const miesmuschel = require("./miesmuschel.js")

const tmi = require('tmi.js');

const client = new tmi.Client({
	options: { 
		debug: true,
		joinInterval: 2000,
		skipMembership: false,
		skipUpdatingEmotesets: false,
		updateEmotesetsTimer: 60000,
	},
	identity: {
		username: undefined,
		password: undefined
	},
	channels: [ 'liadala' ],
	connection:{
		reconnect:true,
		secure: true,
		reconnectInterval: 1000,
		maxReconnectInverval: 30000,
		reconnectDecay: 1.5,

		maxReconnectAttempts: Infinity
	}
});

client.connect();

client.on('message', (channel, userstate, message, self) => {
	var usersplit = userstate.username.split("")
	var userA = 0
	for (const iterator of usersplit) {
		userA += iterator.charCodeAt(0)
	}
});
