let date = new Date();

if (!document.cookie) {
    date.setTime(date.getTime()+(365*24*60*60*1000));
    document.cookie = "fID=" + sha1(window + date) + ";expires="+ date.toUTCString();
}