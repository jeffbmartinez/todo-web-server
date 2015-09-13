"use strict";

const ones = [1, 2, 3];
const tens = ones.map(i => i * 10);
const hundreds = ones.map(i => i * 100);

module.exports = {
    ones: ones,
    tens: tens,
    hundreds: hundreds
};
