const redis = require("redis");
const {promisify} = require('util');

const redisClient = redis.createClient();
const set = promisify(redisClient.set).bind(redisClient);
const get = promisify(redisClient.get).bind(redisClient);

module.exports = {
    TOKEN_KEY: "poster:pinterest-token",
    get,
    set
};