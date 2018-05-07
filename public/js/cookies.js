let date = new Date();

if (!document.cookie) {
    console.log("Writing at:", date);
    document.cookie = "First seen=" + date.toUTCString();
    document.cookie ="inceptionID=abc123"
}

document.cookie = "last seen=" + date.toUTCString();


console.log("Reading at", date);
console.log(document.cookie)