var keythereum = require("keythereum");

var datadir = "/tmp/geth/";
var address = "0x433c0685d3bae6330052bc97a616ff36be6933cf";
var password = "";

var keyObject = keythereum.importFromFile(address, datadir);
var privateKey = keythereum.recover(password, keyObject);
console.log(privateKey.toString('hex'));
