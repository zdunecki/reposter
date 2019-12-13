const rp = require("request-promise");
const variables = require("./variables");

//TODO: ?access_token - sometime we would like to pass query params in resource
module.exports = {
    get: async ({resource, token}) => await rp(`${variables.PINTEREST_API}${resource}?access_token=${token}`),
    post: async ({body, resource, token}) => {
        return await rp({
            uri: `${variables.PINTEREST_API}${resource}?access_token=${token}`,
            method: "POST",
            json: true,
            body
        });
    }
};