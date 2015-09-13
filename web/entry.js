"use strict";

require("./node_modules/bootstrap/dist/css/bootstrap.css");
require("./style.css");

var content = require("./content.js");
var moreContent = require("./moreContent.js");

document.write(content);
document.write("<br>");
document.write(moreContent.ones);
document.write("<br>");
document.write(moreContent.tens);
document.write("<br>");
document.write(moreContent.hundreds);
document.write("<br>");
