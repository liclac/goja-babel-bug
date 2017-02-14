#!/usr/bin/env node
var fs = require('fs');
var Babel = require('./babel.js');

var src = fs.readFileSync('script.js', {encoding: 'utf-8'});
console.log(Babel.transform(src, {presets: ['latest']}).code);
