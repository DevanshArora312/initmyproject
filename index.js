#!/usr/bin/env node

const os = require('os');
const { spawn } = require('child_process');
const path = require('path');

// Determine the binary based on the OS
let binaryName;
switch (os.platform()) {
    case 'win32':
        binaryName = 'initmyproject-windows-amd64.exe';
        break;
    case 'darwin':
        binaryName = 'initmyproject-macos-amd64';
        break;
    case 'linux':
        binaryName = 'initmyproject-linux-amd64';
        break;
    default:
        console.error(`Unsupported platform: ${os.platform()}`);
        process.exit(1);
}

// Path to the binary
const binaryPath = path.join(__dirname, 'bin', binaryName);

// Spawn the binary with arguments from the command line
const child = spawn(binaryPath, process.argv.slice(2), { stdio: 'inherit' });

child.on('error', (err) => {
    console.error(`Failed to start subprocess: ${err}`);
    process.exit(1);
});

child.on('exit', (code) => {
    process.exit(code);
});
