// https://full-stack.blend.com/how-we-write-github-actions-in-go.html
const os = require('os');
// const fetch = require('node-fetch');
import fetch from 'node-fetch';
const fs = require('fs');
// import fs from 'fs';
const { exec } = require('node:child_process');
// import { exec } from 'node:child_process';
const core = require('@actions/core');
const github = require('@actions/github');

async function getReleaseBinaryURL(os, arch) {
  if (arch === 'x64') {
    arch = 'amd64';
  }

  try {
    const response = await fetch(
      'https://api.github.com/repos/zchryr/build-notifier-action/releases/latest',
      {
        method: 'get',
        headers: {
          Accept: 'application/vnd.github+json',
          'X-GitHub-Api-Version': '2022-11-28',
        },
      }
    );

    const data = await response.json();

    if (os === 'darwin') {
      for (var element of data['assets']) {
        if (element['name'].endsWith(`darwin-${arch}`)) {
          return element['browser_download_url'];
        }
      }
    } else if (os === 'linux') {
      for (var element of data['assets']) {
        if (element['name'].endsWith('linux-amd64')) {
          return element['browser_download_url'];
        }
      }
    } else if (os === 'win32') {
      for (var element of data['assets']) {
        if (element['name'].endsWith('windows-amd64')) {
          return element['browser_download_url'];
        }
      }
    }
  } catch (error) {
    console.log(error);
  }
}

async function downloadFile(url, fileName) {
  const res = await fetch(url);
  const fileStream = fs.createWriteStream(process.cwd() + '/' + fileName);
  await new Promise((resolve, reject) => {
    res.body.pipe(fileStream);
    res.body.on('error', reject);
    fileStream.on('finish', resolve);
    // console.log(`Download completed. File: ${process.cwd()}/${fileName}`);
  }).then();

  return process.cwd() + '/' + fileName;
}

async function chooseBinary() {
  const platform = os.platform();
  const arch = os.arch();

  const url = await getReleaseBinaryURL(platform, arch);
  const binary = await downloadFile(url, 'build-notifier-action');

  // Make executable.
  if (platform === 'win32') {
    fs.rename('build-notifier-action', 'build-notifier-action.exe');
  } else if (platform === 'linux' || platform === 'darwin') {
    exec(`chmod +x ${binary}`);
  }

  return binary;
}

function runBinary(binary, args) {
  exec(binary + ' ' + args, (error, stdout, stderr) => {
    if (error) {
      console.error(`exec error: ${error}`);
      return;
    }

    console.log(stdout);
  });
}

try {
  const body = core.getInput('body');
  const url = core.getInput('url');
  const response = core.getInput('response');

  console.log(`body: ${body}`);
  console.log(`url: ${url}`);
  console.log(`response: ${response}`);

  // const binary = await chooseBinary();
  // runBinary(binary, 'send -h');
} catch (error) {
  core.setFailed(error.message);
}
