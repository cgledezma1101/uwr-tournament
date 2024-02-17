const path = require("path");
const ESLintPlugin = require("eslint-webpack-plugin");

module.exports = {
    entry: "./src/index.tsx",
    devtool: "inline-source-map",
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: "ts-loader",
                exclude: /node_modules/,
            },
        ],
    },
    resolve: {
        extensions: [".tsx", ".ts", ".js"],
    },
    output: {
        filename: "bundle.js",
        path: path.resolve(__dirname, "..", "assets", "builds"),
    },
    plugins: [
        new ESLintPlugin({
            extensions: [".ts", ".tsx", ".js", ".jsx"],
        }),
    ],
};
