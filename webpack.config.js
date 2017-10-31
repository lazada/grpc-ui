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
            {
                test: /\.sass$/,
                use: [{
                    loader: "style-loader" // creates style nodes from JS strings
                }, {
                    loader: "css-loader" // translates CSS into CommonJS
                }, {
                    loader: "sass-loader", // compiles Sass to CSS
                    options: {
                        indented_syntax: 'on',
                        includePaths: [
                            path.resolve(__dirname + "/src/sass"),
                        ]
                    }
                }]
            }
        ],
    },
}