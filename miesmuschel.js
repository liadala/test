module.exports = {
    Ask:  function (username){
        let quote = {
            yes: [
                "Ja",
                "Ja definitiv",
                "Positiv",
                "Hättest du etwas anderes erwartet?",
                "Ja ich will 💍",
                "Nur mit guter bezahlung"
            ],
            no: [
                "Nein",
                "Weils du bist NEIN",
                "Nö",
                "Kener hat die Absicht hier Ja zu sagen.",
                "Gegenfrage würdest du nackt und mit Fleisch behängt vor einem hungrigen Tiger tanzen?",
                "Deswegen wird er auch nicht größer also nein!",
                "Ich musste dich jetzt einfach darauf Hinweisen. Du bist so hüpsch wie ein Badewannenstöpsel deswegen muss ich deine Anfrage leider ablehnen.",
                "Nein du stinkst geh dich erstmal waschen!",
                "Sprich mit meiner Hand.",
                "Ihre Bestellung wurde erfolgreich aufgenommen es werden 2502,35€ von ihrem Konto abgebucht. Danke",
                "Nein ich bin tot. Leg den Kranz hin und lass mich in Frieden ruhen",
                "Nein, das ist flüssiger Sonnenschein.",
                "Nein, ich lüge",
                "Nein. Ich bin gerade damit beschäftigt Menschen zu beobachten wie sie sich zum Affen machen.",
                "Diese Sache finde ich genauso positiv wie Durchfall!",
                "NEIN und wenn du nochmal so dämliches zeug frägst werfe ich dich ins Feuer und opfere dich der Göttin Brutzla"
            ]
        };
        let a = Math.floor(Math.random() * 2);
        let randomNumber;
        console.log(a);
    
        switch(a) {
            case 0:
                randomNumber = Math.floor(Math.random() * (quote.yes.length));
                return `${username} ${quote.yes[randomNumber]}`;
            case 1:
                randomNumber = Math.floor(Math.random() * (quote.no.length));
                return `${username} ${quote.no[randomNumber]}`;
            default:
                return `something went wrong`
        }
    }
}
