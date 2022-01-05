const miesmuschel = require("./miesmuschel.js")

const tmi = require('tmi.js');

const client = new tmi.Client({
	channels: [ 'liadala' ]
});

client.connect();

client.on('message', (channel, userstate, message, self) => {
	console.log(miesmuschel.Ask(tags['display-name']));
});
