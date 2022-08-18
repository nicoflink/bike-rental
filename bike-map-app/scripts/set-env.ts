// Script provided by Sara Lara: https://medium.com/@natchiketa/angular-cli-and-os-environment-variables-4cfa3b849659

import { writeFile } from 'fs';

// This is good for local dev environments, when it's better to
// store a projects environment variables in a .gitignore'd file
require('dotenv').config();

// args provided 
var argv = require('yargs/yargs')(process.argv.slice(2)).argv;

// Would be passed to script like this:
// `ts-node set-env.ts --environment=dev`
// we get it from yargs's argv object
const environment = argv.environment;
const isProd = environment === 'prod';

const targetPath = `./src/environments/environment.${environment}.ts`;
const envConfigFile = `
export const environment = {
  production: ${isProd},
  apiKey: "${process.env["GMAPS_API_KEY"]}",

};
`
writeFile(targetPath, envConfigFile, function (err) {
  if (err) {
    console.log(err);
  }

  console.log(`Output generated at ${targetPath}`);
});