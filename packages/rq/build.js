const { build } = require("esbuild");


build({
    entryPoints: ["src/runtime.ts"],
    bundle: true,
    minify: true,
    platform: 'node', // for CJS
    outfile: "dist/runtime.js",
});

