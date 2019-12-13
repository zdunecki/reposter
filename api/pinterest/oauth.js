const variables = require("./variables");

const oauthUrl = `${variables.PINTEREST_URL}/oauth/?response_type=code&redirect_uri=${variables.REDIRECT_URI}&client_id=${variables.CLIENT_ID}&scope=read_public,write_public`;

module.exports = async (req, res) => {
    res.writeHead(302, {"Location": oauthUrl});
    return res.end()
};