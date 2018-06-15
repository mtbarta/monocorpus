'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')
const defaults = require('./index.js')

module.exports = merge(defaults, prodEnv, {
  NODE_ENV: '"development"'
})
