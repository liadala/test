const miesmuschel = require("./miesmuschel.js")

const tmi = require('tmi.js');
const { channel } = require("tmi.js/lib/utils");

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
	if (self){ return };

	(function(){
		if (message.startsWith("!love")) {
			var param = message.split(" ")
			if (param.length == 0){
				bot.say(channel, "no parameter")
				return;
			}
			
			var usersplit = userstate.username.split("")
			var userA = 0
			for (const iterator of usersplit) {
				userA += iterator.charCodeAt(0)
			}

			usersplit = param[1].split(" ")
			var userB = 0
			for (const iterator of usersplit) {
				userB += iterator.charCodeAt(0)
			}

			var min = Math.min(userA,userB)
			var max = Math.max(userA,userB)

			let matching = Math.round((min/max)*100)

			bot.say(channel, `@${userstate.username} and ${param[1]} love value is ${matching}%`)
		}
	})()
});

// ignore this its just to get output from anonymous login
var bot = {
	say: function(channel, message){
		console.log(channel, message)
	}
}