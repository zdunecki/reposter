const pinterest = require("./pinterest");
const redis = require("./redis");

//TODO: secure this route from outside world
module.exports = async (req, res) => {
    const token = await redis.get(redis.TOKEN_KEY);

    const response = await pinterest.post({
        resource: "/boards/",
        body: req.body,
        token: token
    });

    res.json(response);
};