const date = new Date();

if (!document.cookie) {
    document.cookie = "fID=" + sha1(window + date.toUTCString());
}