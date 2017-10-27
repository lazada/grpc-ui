const path = require('path');

module.exports = {
    entry: './src/index.js',
    output: {
        path: path.resolve(__dirname) + "/static",
        filename: "bundle.js"
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                include: [
                    path.resolve(__dirname + "/src"),
                ],
                loader: "babel-loader"
            },
        ],
    },
}