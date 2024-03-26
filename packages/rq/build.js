const { build } = require("esbuild");


build({
    entryPoints: ["src/index.ts"],
    bundle: true,
    minify: true,
    platform: 'node', // for CJS
    outfile: "dist/index.js",
});

