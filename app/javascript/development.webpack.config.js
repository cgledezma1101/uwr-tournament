const path = require("path");

module.exports = {
    extends: path.resolve(__dirname, "base.webpack.config.js"),
    mode: "development",
    watch: true,
    watchOptions: {
        ignored: /node_modules/,
    },
};
