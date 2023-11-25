const core = require('@actions/core'); // Importing the core module from the GitHub Actions toolkit
const axios = require('axios'); // Importing axios for making HTTP requests
const github = require('@actions/github'); // Importing GitHub toolkit to access repo and workflow information

async function run() {
  try {
    // Reading input parameters from the workflow file
    const userInput = core.getInput('input'); // User-provided JSON input
    const url = core.getInput('url'); // The URL to send the data to

    // Validate user input is valid JSON
    let inputJson;
    try {
      inputJson = JSON.parse(userInput);
    } catch (error) {
      throw new Error('Invalid JSON input');
    }

    // Constructing the git_info object
    const gitInfo = {
      workflow: github.context.workflow, // Name of the current workflow
      repo: github.context.repo.repo, // Name of the repository
    };

    // Constructing the body of the data to be sent
    const body = {
      git_info: gitInfo,
      input: inputJson,
    };

    // Sending data to the specified URL using a POST request
    const response = await axios.post(url, body);

    // Optionally, additional validation of the response can be done here

    // Setting the 'success' output parameter to true
    core.setOutput('success', true);
  } catch (error) {
    // If any error occurs, mark the action as failed and log the error message
    core.setFailed(`Action failed with error: ${error.message}`);
  }
}

run();
