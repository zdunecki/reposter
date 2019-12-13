const rp = require("request-promise");
const redis = require("./redis");

const variables = require("./variables");

const getTokenUrl = code => `${variables.PINTEREST_API}/oauth/token?grant_type=authorization_code&client_id=${variables.CLIENT_ID}&client_secret=${variables.CLIENT_SECRET}&code=${code}`;

module.exports = async (req, res) => {
    try {
        const response = await rp({
            uri: getTokenUrl(req.query.code),
            method: "POST",
        });

        await redis.set(redis.TOKEN_KEY, JSON.parse(response).access_token);

        res.send(response);
    } catch (e) {
        console.log(e);

        res.status(500);
        res.send("Something went wrong")
    }
};